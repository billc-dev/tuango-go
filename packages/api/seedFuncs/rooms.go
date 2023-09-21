package seedfuncs

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/billc-dev/tuango-go/ent"
	"github.com/billc-dev/tuango-go/ent/room"
	"github.com/gofiber/fiber/v2"
)

type RoomUser struct {
	ID   string `json:"_id"`
	User string `json:"user"`
}

type JSONRoom struct {
	ID        string     `json:"_id"`
	UpdatedAt string     `json:"lastCreatedAt"`
	Users     []RoomUser `json:"users"`
}

func SeedRooms(client *ent.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jsonFile, err := os.Open("seeds/rooms.json")

		if err != nil {
			return err
		}
		defer jsonFile.Close()
		byteValue, _ := io.ReadAll(jsonFile)
		var rooms []JSONRoom
		json.Unmarshal(byteValue, &rooms)

		// delete
		client.RoomUser.Delete().Where().ExecX(c.Context())
		client.Room.Delete().Where().ExecX(c.Context())

		var bulkRooms []*ent.RoomCreate

		for _, r := range rooms {
			t, err := time.Parse("2006-01-02T15:04:05.000Z", r.UpdatedAt)
			if err != nil {
				return err
			}

			bulkRooms = append(bulkRooms,
				client.Room.Create().
					SetID(r.ID).
					SetName("").
					SetType(room.TypePrivate).
					SetUpdatedAt(t),
			)
		}

		batch := 3000

		for i := 0; i < len(bulkRooms); i += batch {
			j := i + batch
			if j > len(bulkRooms) {
				j = len(bulkRooms)
			}
			_, err := client.Room.CreateBulk(bulkRooms[i:j]...).Save(c.Context())

			if err != nil {
				return err
			}
		}

		var bulkRoomUsers []*ent.RoomUserCreate

		for _, r := range rooms {
			t, err := time.Parse("2006-01-02T15:04:05.000Z", r.UpdatedAt)
			if err != nil {
				return err
			}

			bulkRooms = append(bulkRooms,
				client.Room.Create().
					SetID(r.ID).
					SetName("").
					SetType(room.TypePrivate).
					SetUpdatedAt(t),
			)

			for _, user := range r.Users {
				bulkRoomUsers = append(bulkRoomUsers,
					client.RoomUser.Create().
						SetRoomID(r.ID).
						SetUserID(user.User).
						SetLastReadMessageID(""),
				)
			}

		}

		for i := 0; i < len(bulkRoomUsers); i += batch {
			j := i + batch
			if j > len(bulkRoomUsers) {
				j = len(bulkRoomUsers)
			}
			_, err := client.RoomUser.CreateBulk(bulkRoomUsers[i:j]...).Save(c.Context())

			if err != nil {
				return err
			}
		}

		return c.JSON(fiber.Map{})
	}

}
