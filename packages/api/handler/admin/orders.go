package admin

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent"
	"github.com/billc-dev/tuango-go/ent/order"
	"github.com/billc-dev/tuango-go/ent/orderitem"
	"github.com/billc-dev/tuango-go/ent/post"
	"github.com/billc-dev/tuango-go/ent/postitem"
	"github.com/billc-dev/tuango-go/ent/predicate"
	"github.com/billc-dev/tuango-go/ent/user"
	"github.com/gofiber/fiber/v2"
)

func GetOrders(c *fiber.Ctx) error {
	m := c.Queries()

	orderWhere := []predicate.Order{}

	if pickupNum := m["pickup_num"]; pickupNum != "" {
		if pickupNum, err := strconv.Atoi(pickupNum); err == nil {
			orderWhere = append(orderWhere,
				order.HasUserWith(
					user.PickupNum(float64(pickupNum)),
				))
		}
	}

	if userID := m["user_id"]; userID != "" {
		orderWhere = append(orderWhere, order.UserID(userID))
	}

	if isInStock := m["is_in_stock"]; isInStock != "" {
		orderWhere = append(orderWhere, order.IsInStock(isInStock == "true"))
	}

	if arrivedYesterday := m["arrived_yesterday"]; arrivedYesterday == "true" {
		orderWhere = append(orderWhere, order.CreatedAtLTE(time.Now().AddDate(0, 0, -1)))
	}

	postWhere := []predicate.Post{}

	if postNum := m["post_num"]; postNum != "" {
		postNum, err := strconv.Atoi(postNum)
		if err == nil {
			postWhere = append(postWhere, post.PostNumEQ(postNum))
		}
	}

	if text := m["text"]; text != "" {
		postWhere = append(postWhere,
			post.Or(
				post.TitleContains(text),
				post.BodyContains(text),
				post.HasPostItemsWith(postitem.NameContains(text)),
			),
		)
	}

	if deadline := m["deadline"]; deadline != "" {
		postWhere = append(postWhere, post.DeadlineEQ(deadline))
	}

	if deliveryDate := m["delivery_date"]; deliveryDate != "" {
		postWhere = append(postWhere, post.DeliveryDateEQ(deliveryDate))
	}

	if storageType := m["storage_type"]; storageType != "" {
		storageType := post.StorageType(storageType)
		err := post.StorageTypeValidator(storageType)
		if err != nil {
			return fiber.NewError(
				http.StatusInternalServerError,
				fmt.Sprintf(`Post storage type "%v" is not valid`, storageType),
			)
		}
		postWhere = append(postWhere, post.StorageTypeEQ(storageType))
	}

	if sellerID := m["seller_id"]; sellerID != "" {
		postWhere = append(postWhere, post.SellerID(sellerID))
	}

	status := post.Status(m["status"])
	if err := post.StatusValidator(status); err == nil {
		postWhere = append(postWhere, post.StatusEQ(status))
	}

	orderWhere = append(orderWhere, order.HasPostWith(postWhere...))

	limit := 20
	page := 0

	if int, err := strconv.Atoi(m["page"]); err == nil {
		page = int
	}

	offset := page * limit

	orders, err := database.DBConn.Debug().Order.
		Query().
		Where(orderWhere...).
		WithPost(func(pq *ent.PostQuery) {
			pq.
				Select(post.FieldPostNum, post.FieldTitle).
				WithSeller(func(uq *ent.UserQuery) {
					uq.Select(user.FieldDisplayName)
				})
		}).
		WithUser(func(uq *ent.UserQuery) {
			uq.Select(user.FieldDisplayName, user.FieldPictureURL)
		}).
		WithOrderItems(func(oiq *ent.OrderItemQuery) {
			oiq.Order(orderitem.ByIdentifier())
		}).
		Order(order.ByCreatedAt(sql.OrderDesc())).
		Limit(limit).
		Offset(offset).
		All(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query orders")
	}

	count, err := database.DBConn.Order.Query().Where(orderWhere...).Count(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query order count")
	}

	return c.JSON(fiber.Map{
		"data": fiber.Map{
			"orders":      orders,
			"count":       count,
			"hasNextPage": len(orders) == limit,
		},
	})
}

func GetOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")

	o, err := database.DBConn.Debug().Order.
		Query().
		Where(order.ID(orderID)).
		WithPost(func(pq *ent.PostQuery) {
			pq.
				Select(post.FieldPostNum, post.FieldTitle).
				WithSeller(func(uq *ent.UserQuery) {
					uq.Select(user.FieldDisplayName)
				})
		}).
		WithUser(func(uq *ent.UserQuery) {
			uq.Select(user.FieldDisplayName, user.FieldPictureURL)
		}).
		WithOrderItems(func(oiq *ent.OrderItemQuery) {
			oiq.Order(orderitem.ByIdentifier())
		}).
		First(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Post not found")
	}

	return c.JSON(fiber.Map{
		"data": o,
	})
}

type createOrderItem struct {
	PostItemID string  `json:"post_item_id"`
	Identifier string  `json:"identifier"`
	Name       string  `json:"name"`
	HasName    bool    `json:"has_name"`
	Qty        float64 `json:"qty"`
	Price      float64 `json:"price"`
	Location   string  `json:"location"`
	Status     string  `json:"status"`
}

type createOrderForm struct {
	UserID  string            `json:"user_id"`
	PostID  string            `json:"post_id"`
	IsExtra bool              `json:"is_extra"`
	Order   []createOrderItem `json:"order"`
	Comment string            `json:"comment"`
}

func CreateOrder(c *fiber.Ctx) error {
	orderForm := new(createOrderForm)
	if err := c.BodyParser(orderForm); err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusBadRequest, "Could not parse order form")
	}

	ordersToCreate := map[order.Status][]createOrderItem{}
	for _, orderItem := range orderForm.Order {
		status := order.Status(orderItem.Status)
		err := order.StatusValidator(status)
		if err != nil {
			log.Print(err)
			return fiber.NewError(http.StatusBadRequest, "Order item has invalid order status")
		}
		ordersToCreate[status] = append(ordersToCreate[status], orderItem)
	}

	if len(ordersToCreate) == 0 {
		return fiber.NewError(http.StatusBadRequest, "No orders to create")
	}

	orderNum := 0
	if !orderForm.IsExtra {
		previousOrder, err := database.DBConn.Order.Query().
			Where(order.PostID(orderForm.PostID), order.StatusNEQ(order.StatusCanceled)).
			Order(order.ByOrderNum(sql.OrderDesc())).
			Limit(1).
			All(c.Context())
		if err != nil {
			log.Print(err)
			return fiber.NewError(http.StatusInternalServerError, "Could not query previous order")
		}
		if len(previousOrder) == 1 {
			orderNum = *previousOrder[0].OrderNum + 1
		} else {
			orderNum = 1
		}
	}

	tx, err := database.DBConn.Debug().Tx(c.Context())
	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not start transaction")
	}

	for status, orderItems := range ordersToCreate {
		newOrder, err := tx.Order.Create().
			SetUserID(orderForm.UserID).
			SetPostID(orderForm.PostID).
			SetStatus(status).
			SetOrderNum(orderNum).
			SetComment(orderForm.Comment).
			Save(c.Context())
		if err != nil {
			log.Print(err)
			tx.Rollback()
			return fiber.NewError(http.StatusInternalServerError, "Could not create order")
		}

		orderItemsCreate := []*ent.OrderItemCreate{}

		for _, orderItem := range orderItems {
			orderItemsCreate = append(orderItemsCreate,
				tx.OrderItem.Create().
					SetOrderID(newOrder.ID).
					SetPostItemID(orderItem.PostItemID).
					SetIdentifier(orderItem.Identifier).
					SetName(orderItem.Name).
					SetHasName(orderItem.HasName).
					SetQty(orderItem.Qty).
					SetPrice(orderItem.Price).
					SetLocation(orderItem.Location).
					SetStatus(orderitem.Status(status)),
			)
		}

		if len(orderItems) == 0 {
			tx.Rollback()
			return fiber.NewError(http.StatusBadRequest, "No order items to create")
		}

		err = tx.OrderItem.CreateBulk(orderItemsCreate...).Exec(c.Context())
		if err != nil {
			log.Print(err)
			tx.Rollback()
			return fiber.NewError(http.StatusInternalServerError, "Could not create order items")
		}
	}

	// TODO: increment post order count if orderNum != 0

	err = tx.Commit()
	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not commit order transaction")
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
