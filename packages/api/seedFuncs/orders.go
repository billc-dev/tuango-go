package seedfuncs

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/billc-dev/tuango-go/ent"
	"github.com/billc-dev/tuango-go/ent/order"
	"github.com/billc-dev/tuango-go/ent/orderitem"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JSONOrder struct {
	ID                string       `json:"_id"`
	UserId            string       `json:"userId"`
	PostId            string       `json:"postId"`
	Title             string       `json:"title"`
	DisplayName       string       `json:"displayName"`
	SellerDisplayName string       `json:"sellerDisplayName"`
	OrderNum          int          `json:"orderNum"`
	PostNum           int          `json:"postNum"`
	Comment           string       `json:"comment"`
	SellerComment     string       `json:"sellerComment"`
	HasName           bool         `json:"hasName"`
	IsExtra           bool         `json:"isExtra"`
	Fb                bool         `json:"fb"`
	IsInStock         bool         `json:"isInStock"`
	Status            order.Status `json:"status"`
	CreatedAt         string       `json:"createdAt"`
	Order             []struct {
		OID      string           `json:"_id"`
		ID       string           `json:"id"`
		Item     string           `json:"item"`
		Qty      float64          `json:"qty"`
		Price    float64          `json:"price"`
		Status   orderitem.Status `json:"status"`
		Location string           `json:"location"`
		HasName  bool             `json:"hasName"`
	} `json:"order"`
}

func SeedOrders(client *ent.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("seeds/orders.json")

		if err != nil {
			return err
		}
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		var orders []JSONOrder
		json.Unmarshal(byteValue, &orders)

		// delete
		client.OrderHistory.Delete().Where().ExecX(c.Context())
		client.OrderItem.Delete().Where().ExecX(c.Context())
		client.Order.Delete().Where().ExecX(c.Context())

		users, err := client.User.Query().All(c.Context())
		if err != nil {
			return err
		}

		var bulkOrders []*ent.OrderCreate

		posts, err := client.Post.Query().WithPostItems().All(c.Context())
		if err != nil {
			return err
		}

		for _, order := range orders {
			var found bool
			for _, p := range posts {
				if p.ID == order.PostId {
					found = true
					break
				}
			}
			if !found {
				continue
			}
			t, err := time.Parse("2006-01-02T15:04:05.000Z", order.CreatedAt)
			if err != nil {
				return err
			}
			u, err := getUserById(users, order.UserId)
			if err != nil {
				return err
			}

			bulkOrders = append(bulkOrders,
				client.Order.
					Create().
					SetID(order.ID).
					SetUserID(u.ID).
					SetPostID(order.PostId).
					SetOrderNum(order.OrderNum).
					SetComment(order.Comment).
					SetSellerComment(order.SellerComment).
					SetHasName(order.HasName).
					SetIsExtra(order.IsExtra).
					SetStatus(order.Status).
					SetFb(order.Fb).
					SetIsInStock(order.IsInStock).
					SetCreatedAt(t).
					SetUpdatedAt(t),
			)

		}
		batch := 2500

		for i := 0; i < len(bulkOrders); i += batch {
			j := i + batch
			if j > len(bulkOrders) {
				j = len(bulkOrders)
			}
			_, err := client.Order.CreateBulk(bulkOrders[i:j]...).Save(c.Context())

			if err != nil {
				return err
			}
		}

		var orderItems []*ent.OrderItemCreate

		itemsIdCount := make(map[string]int)

		for _, order := range orders {
			var found bool
			for _, p := range posts {
				if p.ID == order.PostId {
					found = true
					break
				}
			}
			if !found {
				// log.Print(order.PostNum, " ", order.Title)
				continue
			}
			for _, item := range order.Order {
				itemId := item.OID

				if itemsIdCount[item.OID] > 0 {
					itemId = primitive.NewObjectID().Hex()
				} else {
					itemsIdCount[item.OID]++
				}
				var postItemId string
				for _, post := range posts {
					if post.ID == order.PostId {
						for _, postItem := range post.Edges.PostItems {
							if *postItem.Identifier == item.ID {
								postItemId = postItem.ID
								break
							}
						}
						break
					}
				}

				if postItemId == "" {
					panic("could not find post item id")
				}

				orderItems = append(orderItems,
					client.OrderItem.
						Create().
						SetOrderID(order.ID).
						SetID(itemId).
						SetIdentifier(item.ID).
						SetName(item.Item).
						SetPrice(item.Price).
						SetQty(item.Qty).
						SetStatus(item.Status).
						SetLocation(item.Location).
						SetPostItemID(postItemId).
						SetHasName(item.HasName),
				)

			}
		}

		batch = 5000

		for i := 0; i < len(orderItems); i += batch {
			j := i + batch
			if j > len(orderItems) {
				j = len(orderItems)
			}
			_, err := client.OrderItem.CreateBulk(orderItems[i:j]...).Save(c.Context())

			if err != nil {
				return err
			}
		}

		return c.JSON(fiber.Map{})
	}
}
