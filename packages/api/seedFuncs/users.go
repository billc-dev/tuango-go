package seedfuncs

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/billc-dev/tuango-go/ent"
	"github.com/billc-dev/tuango-go/ent/user"
	"github.com/gofiber/fiber/v2"
)

type JSONUser struct {
	ID                       string      `json:"_id"`
	Username                 string      `json:"username"`
	DisplayName              string      `json:"displayName"`
	PictureUrl               string      `json:"pictureUrl"`
	PickupNum                float64     `json:"pickupNum"`
	Role                     user.Role   `json:"role"`
	Status                   user.Status `json:"status"`
	Notified                 bool        `json:"notified"`
	LinePay                  bool        `json:"linepay"`
	Fb                       bool        `json:"fb"`
	Comment                  string      `json:"comment"`
	DeliveredOrderCountLimit int         `json:"deliveredOrderCountLimit"`
	CreatedAt                string      `json:"createdAt"`
	Message                  struct {
		Notified   bool `json:"notified"`
		NotifiedAt bool `json:"notifiedAt"`
	} `json:"message"`
}

func SeedUsers(client *ent.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("seeds/users.json")
		if err != nil {
			return err
		}
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		var users []JSONUser
		json.Unmarshal(byteValue, &users)

		bulk := make([]*ent.UserCreate, len(users))

		client.User.Delete().Where().Exec(c.Context())

		for i, user := range users {
			t, err := time.Parse("2006-01-02T15:04:05.000Z", user.CreatedAt)
			if err != nil {
				return err
			}

			bulk[i] = client.User.
				Create().
				SetID(user.ID).
				SetUsername(user.Username).
				SetDisplayName(user.DisplayName).
				SetPictureURL(user.PictureUrl).
				SetPickupNum(user.PickupNum).
				SetRole(user.Role).
				SetStatus(user.Status).
				SetNotified(user.Notified).
				SetLinePay(user.LinePay).
				SetFb(user.Fb).
				SetComment(user.Comment).
				SetDeliveredOrderCountLimit(user.DeliveredOrderCountLimit).
				SetCreatedAt(t).
				SetUpdatedAt(t)
		}

		bulkUsers, err := client.User.CreateBulk(bulk...).Save(c.Context())

		if err != nil {
			return err
		}

		log.Print(bulkUsers[0].ID)

		return c.JSON(fiber.Map{})
	}
}
