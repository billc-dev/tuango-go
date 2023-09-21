package seedfuncs

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/billc-dev/tuango-go/ent"
	"github.com/billc-dev/tuango-go/ent/message"
	"github.com/billc-dev/tuango-go/ent/schema"
	"github.com/gofiber/fiber/v2"
)

type Read struct {
	UserId string `json:"userId"`
	Read   bool   `json:"read"`
}

type JSONMessage struct {
	ID        string        `json:"_id"`
	RoomId    string        `json:"roomId"`
	UserId    string        `json:"userId"`
	PostId    *string       `json:"post"`
	OrderId   *string       `json:"order"`
	Text      *string       `json:"text"`
	CreatedAt string        `json:"createdAt"`
	UpdatedAt string        `json:"updatedAt"`
	Type      message.Type  `json:"type"`
	Image     *schema.Image `json:"imageUrl"`
	Read      []Read        `json:"read"`
}

func SeedMessages(client *ent.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("seeds/chats.json")

		if err != nil {
			return err
		}
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		var messages []JSONMessage
		json.Unmarshal(byteValue, &messages)

		// delete
		client.Message.Delete().Where().ExecX(c.Context())

		var bulkMessages []*ent.MessageCreate

		for _, m := range messages {
			t, err := time.Parse("2006-01-02T15:04:05.000Z", m.UpdatedAt)
			if err != nil {
				return err
			}

			bulkMessages = append(bulkMessages,
				client.Message.Create().
					SetID(m.ID).
					SetRoomID(m.RoomId).
					SetUserID(m.UserId).
					SetType(m.Type).
					SetNillablePostID(m.PostId).
					SetNillableOrderID(m.OrderId).
					SetNillableImage(m.Image).
					SetNillableText(m.Text).
					SetUpdatedAt(t),
			)
		}

		batch := 3000

		for i := 0; i < len(bulkMessages); i += batch {
			j := i + batch
			if j > len(bulkMessages) {
				j = len(bulkMessages)
			}
			_, err := client.Message.CreateBulk(bulkMessages[i:j]...).Save(c.Context())

			if err != nil {
				return err
			}
		}

		// var bulkRoomUsers []*ent.RoomUserCreate

		// for _, r := range messages {
		// 	t, err := time.Parse("2006-01-02T15:04:05.000Z", r.UpdatedAt)
		// 	if err != nil {
		// 		return err
		// 	}

		// 	bulkMessages = append(bulkMessages,
		// 		client.Room.Create().
		// 			SetID(r.ID).
		// 			SetName("").
		// 			SetType(room.TypePrivate).
		// 			SetUpdatedAt(t),
		// 	)

		// 	for _, user := range r.Users {
		// 		bulkRoomUsers = append(bulkRoomUsers,
		// 			client.RoomUser.Create().
		// 				SetRoomID(r.ID).
		// 				SetUserID(user.User),
		// 		)
		// 	}

		// }

		// for i := 0; i < len(bulkRoomUsers); i += batch {
		// 	j := i + batch
		// 	if j > len(bulkRoomUsers) {
		// 		j = len(bulkRoomUsers)
		// 	}
		// 	_, err := client.RoomUser.CreateBulk(bulkRoomUsers[i:j]...).Save(c.Context())

		// 	if err != nil {
		// 		return err
		// 	}
		// }

		return c.JSON(fiber.Map{})
	}

}
