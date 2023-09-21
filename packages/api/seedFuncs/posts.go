package seedfuncs

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/billc-dev/tuango-go/ent"
	"github.com/billc-dev/tuango-go/ent/post"
	"github.com/billc-dev/tuango-go/ent/schema"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JSONPost struct {
	ID           string           `json:"_id"`
	UserId       string           `json:"userId"`
	PostNum      int              `json:"postNum"`
	Title        string           `json:"title"`
	Body         string           `json:"body"`
	Deadline     string           `json:"deadline"`
	DeliveryDate string           `json:"deliveryDate"`
	LikeCount    int              `json:"likeCount"`
	CommentCount int              `json:"commentCount"`
	OrderCount   int              `json:"orderCount"`
	NormalTotal  float64          `json:"normalTotal"`
	NormalFee    float64          `json:"normalFee"`
	ExtraTotal   float64          `json:"extraTotal"`
	ExtraFee     float64          `json:"extraFee"`
	StorageType  post.StorageType `json:"storageType"`
	Status       post.Status      `json:"status"`
	Comment      string           `json:"comment"`
	Delivered    bool             `json:"delivered"`
	IsInStock    bool             `json:"fb"`
	CreatedAt    string           `json:"createdAt"`
	UpdatedAt    string           `json:"updatedAt"`
	ImageUrls    []schema.Image   `json:"imageUrls"`
	Items        []struct {
		OID     string  `json:"_id"`
		ID      string  `json:"id"`
		Item    string  `json:"item"`
		Price   float64 `json:"price"`
		ItemQty float64 `json:"itemQty"`
	} `json:"items"`
}

func SeedPosts(client *ent.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("seeds/posts.json")

		if err != nil {
			return err
		}
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		var posts []JSONPost
		json.Unmarshal(byteValue, &posts)

		// delete
		client.PostItem.Delete().Where().ExecX(c.Context())
		client.Post.Delete().Where().ExecX(c.Context())

		users, err := client.User.Query().All(c.Context())
		if err != nil {
			return err
		}

		bulkPosts := make([]*ent.PostCreate, len(posts))

		for i, post := range posts {
			t, err := time.Parse("2006-01-02T15:04:05.000Z", post.CreatedAt)
			if err != nil {
				return err
			}
			u, err := getUserById(users, post.UserId)
			if err != nil {
				return err
			}

			bulkPosts[i] = client.Post.
				Create().
				SetSellerID(u.ID).
				SetID(post.ID).
				SetPostNum(post.PostNum).
				SetTitle(post.Title).
				SetBody(post.Body).
				SetDeadline(post.Deadline).
				SetDeliveryDate(post.DeliveryDate).
				SetLikeCount(post.LikeCount).
				SetCommentCount(post.CommentCount).
				SetOrderCount(post.OrderCount).
				SetImages(post.ImageUrls).
				SetNormalTotal(post.NormalTotal).
				SetNormalFee(post.NormalFee).
				SetExtraTotal(post.ExtraTotal).
				SetExtraFee(post.ExtraFee).
				// SetPostFinance(&ent.PostFinance{
				// 	ID:          post.ID,
				// 	PostID:      post.ID,
				// 	NormalTotal: post.NormalTotal,
				// 	NormalFee:   post.NormalFee,
				// 	ExtraTotal:  post.ExtraTotal,
				// 	ExtraFee:    post.ExtraFee,
				// }).
				SetStorageType(post.StorageType).
				SetStatus(post.Status).
				SetComment(post.Comment).
				SetDelivered(post.Delivered).
				SetIsInStock(post.IsInStock).
				SetCreatedAt(t).
				SetUpdatedAt(t)
		}

		batch := 2500

		for i := 0; i < len(bulkPosts); i += batch {
			j := i + batch
			if j > len(bulkPosts) {
				j = len(bulkPosts)
			}
			_, err := client.Post.CreateBulk(bulkPosts[i:j]...).Save(c.Context())

			if err != nil {
				return err
			}
		}

		// var bulkPostFinances []*ent.PostFinanceCreate

		// for _, post := range posts {
		// 	bulkPostFinances = append(bulkPostFinances,
		// 		client.PostFinance.
		// 			Create().
		// 			SetID(post.ID).
		// 			SetPostID(post.ID).
		// 			SetNormalTotal(post.NormalTotal).
		// 			SetNormalFee(post.NormalFee).
		// 			SetExtraTotal(post.ExtraTotal).
		// 			SetExtraFee(post.ExtraFee),
		// 	)
		// }

		// batch = 5000

		// for i := 0; i < len(bulkPostFinances); i += batch {
		// 	j := i + batch
		// 	if j > len(bulkPostFinances) {
		// 		j = len(bulkPostFinances)
		// 	}
		// 	_, err := client.PostFinance.CreateBulk(bulkPostFinances[i:j]...).Save(c.Context())

		// 	if err != nil {
		// 		return err
		// 	}
		// }

		var bulkPostItems []*ent.PostItemCreate

		postItemsIdCount := make(map[string]int)

		for _, post := range posts {
			for _, item := range post.Items {
				itemId := item.OID

				if postItemsIdCount[item.OID] > 0 {
					itemId = primitive.NewObjectID().Hex()
				} else {
					postItemsIdCount[item.OID]++
				}

				bulkPostItems = append(bulkPostItems, client.PostItem.
					Create().
					SetPostID(post.ID).
					SetID(itemId).
					SetIdentifier(item.ID).
					SetName(item.Item).
					SetPrice(item.Price).
					SetStock(item.ItemQty),
				)
			}
		}

		for key, element := range postItemsIdCount {
			if element <= 1 {
				delete(postItemsIdCount, key)
			}
		}

		batch = 10000

		for i := 0; i < len(bulkPostItems); i += batch {
			j := i + batch
			if j > len(bulkPostItems) {
				j = len(bulkPostItems)
			}
			_, err := client.PostItem.CreateBulk(bulkPostItems[i:j]...).Save(c.Context())

			if err != nil {
				return err
			}
		}

		return c.JSON(fiber.Map{})
	}

}
