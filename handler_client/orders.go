package handler_client

import (
	"fmt"
	"log"
	"net/http"

	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent"
	"github.com/billc-dev/tuango-go/ent/order"
	"github.com/billc-dev/tuango-go/ent/orderitem"
	"github.com/billc-dev/tuango-go/ent/post"
	"github.com/billc-dev/tuango-go/ent/postitem"
	"github.com/gofiber/fiber/v2"
)

type orderForm struct {
	PostID  string             `json:"postId"`
	Order   map[string]float64 `json:"order"`
	Comment string             `json:"comment"`
}

func CreateOrder(c *fiber.Ctx) error {
	u, ok := c.Locals("user").(*ent.User)

	if !ok {
		return fiber.NewError(http.StatusNotFound, "User not found")
	}

	orderForm := new(orderForm)

	if err := c.BodyParser(orderForm); err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusBadRequest, "Could not parse order form")
	}

	postID, err := database.DBConn.Post.Query().
		Where(post.ID(orderForm.PostID), post.StatusEQ(post.StatusOpen)).
		FirstID(c.Context())

	if err != nil {
		return fiber.NewError(http.StatusBadRequest, "Could not find post")
	}

	postItems, err := database.DBConn.PostItem.Query().Where(postitem.PostID(postID)).All(c.Context())

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query post items")
	}

	if len(postItems) == 0 {
		return fiber.NewError(http.StatusInternalServerError, "No post items were found")
	}

	errors := fiber.Map{}

	for itemId, qty := range orderForm.Order {
		if itemId == "" || qty == 0 {
			continue
		}

		found := false
		for _, postItem := range postItems {
			if postItem.ID == itemId {
				found = true
				if *postItem.Stock-qty < 0 {
					errors[itemId] = "剩餘數量已更新! 請重新下單!"
				}
			}
		}
		if !found {
			return fiber.NewError(http.StatusBadRequest, fmt.Sprintf(`Order contains invalid post item id "%s"`, itemId))
		}
	}

	if len(errors) > 0 {
		return c.JSON(fiber.Map{
			"errors":    errors,
			"postItems": postItems,
		})
	}

	previousOrder, err := database.DBConn.Order.Query().
		Where(order.PostID(postID), order.StatusNEQ(order.StatusCanceled)).
		Order(order.ByOrderNum(sql.OrderDesc())).
		Limit(1).
		All(c.Context())

	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Could not query previousOrder")
	}

	orderNum := 1

	if len(previousOrder) == 1 {
		orderNum = *previousOrder[0].OrderNum + 1
	}

	tx, err := database.DBConn.Debug().Tx(c.Context())
	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not start transaction")
	}

	newOrder, err := tx.Order.Create().
		SetUserID(u.ID).
		SetPostID(postID).
		SetOrderNum(orderNum).
		SetComment(orderForm.Comment).
		Save(c.Context())

	if err != nil {
		log.Print(err)
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "Could not create order")
	}

	orderItems := []*ent.OrderItemCreate{}

	for itemId, qty := range orderForm.Order {
		if itemId == "" || qty == 0 {
			continue
		}

		for _, postItem := range postItems {
			if postItem.ID == itemId {
				orderItems = append(orderItems,
					tx.OrderItem.Create().
						SetOrderID(newOrder.ID).
						SetPostItemID(postItem.ID).
						SetIdentifier(*postItem.Identifier).
						SetName(*postItem.Name).
						SetPrice(*postItem.Price).
						SetQty(qty),
				)
				pi, err := tx.PostItem.Query().Where(postitem.ID(postItem.ID)).First(c.Context())
				if err != nil {
					log.Print(err)
					tx.Rollback()
					return fiber.NewError(http.StatusInternalServerError, "Could not query post item")
				}
				err = pi.Update().AddStock(-qty).Exec(c.Context())
				// TODO: check if stock is negative here
				if err != nil {
					log.Print(err)
					tx.Rollback()
					return fiber.NewError(http.StatusInternalServerError, "Could not update post item stock")
				}
			}
		}
	}

	if len(orderItems) == 0 {
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "No order to create")
	}

	err = tx.OrderItem.CreateBulk(orderItems...).Exec(c.Context())

	if err != nil {
		log.Print(err)
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "Could not create order")
	}

	err = tx.Post.UpdateOneID(postID).AddOrderCount(1).Select(post.FieldID).Exec(c.Context())
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Could not increment order count")
	}

	err = tx.Commit()
	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not commit order transaction")
	}

	newOrder, err = database.DBConn.Order.Query().
		Where(order.ID(newOrder.ID)).
		Select(
			order.FieldID, order.FieldOrderNum, order.FieldComment,
			order.FieldHasName, order.FieldCreatedAt,
		).
		WithOrderItems(func(oiq *ent.OrderItemQuery) {
			oiq.Select(
				orderitem.FieldID, orderitem.FieldIdentifier, orderitem.FieldName,
				orderitem.FieldPrice, orderitem.FieldQty,
			).Order(orderitem.ByIdentifier())
		}).
		First(c.Context())

	if err != nil {
		// TODO: this should not throw an error since the order is created
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not query order")
	}

	return c.JSON(fiber.Map{
		"data": newOrder,
	})
}

func CancelOrder(c *fiber.Ctx) error {
	orderID := c.Params("id")

	u, ok := c.Locals("user").(*ent.User)

	if !ok {
		return fiber.NewError(http.StatusNotFound, "User not found")
	}

	tx, err := database.DBConn.Debug().Tx(c.Context())
	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not start transaction")
	}

	existingOrder, err := tx.Order.
		UpdateOneID(orderID).
		Where(
			order.UserID(u.ID),
			order.StatusEQ(order.StatusOrdered),
			order.HasPostWith(post.StatusEQ(post.StatusOpen)),
		).
		SetStatus(order.StatusCanceled).
		Save(c.Context())

	if err != nil {
		log.Print(err)
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "Could not cancel order")
	}

	orderItems, err := tx.OrderItem.
		Query().
		Where(orderitem.OrderID(orderID)).
		All(c.Context())

	if err != nil {
		log.Print(err)
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "Could not query order items")
	}

	for _, item := range orderItems {
		err := tx.PostItem.
			UpdateOneID(*item.PostItemID).
			AddStock(*item.Qty).
			Exec(c.Context())

		if err != nil {
			log.Print(err)
			tx.Rollback()
			return fiber.NewError(http.StatusInternalServerError, "Could not increment post item qty")
		}
	}

	err = tx.OrderItem.
		Update().
		Where(orderitem.OrderID(orderID)).
		SetStatus(orderitem.StatusCanceled).
		Exec(c.Context())

	if err != nil {
		log.Print(err)
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "Could not cancel order items")
	}

	err = tx.Post.
		UpdateOneID(*existingOrder.PostID).
		AddOrderCount(-1).
		Exec(c.Context())

	if err != nil {
		log.Print(err)
		tx.Rollback()
		return fiber.NewError(http.StatusInternalServerError, "Could not decrement post order count")
	}

	err = tx.Commit()
	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not commit order")
	}

	return c.JSON(fiber.Map{
		"success": true,
	})
}
