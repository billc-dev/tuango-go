package seedfuncs

import (
	"github.com/billc-dev/tuango-go/ent"
	"github.com/gofiber/fiber/v2"
)

func SeedAll(client *ent.Client) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// delete
		client.Message.Delete().Where().Exec(c.Context())
		client.Room.Delete().Where().Exec(c.Context())
		client.Like.Delete().Where().Exec(c.Context())
		client.Deliver.Delete().Where().Exec(c.Context())
		client.Notify.Delete().Where().Exec(c.Context())
		client.Comment.Delete().Where().ExecX(c.Context())
		client.Complete.Delete().Where().Exec(c.Context())
		client.OrderHistory.Delete().Where().ExecX(c.Context())
		client.OrderItem.Delete().Where().ExecX(c.Context())
		client.Order.Delete().Where().ExecX(c.Context())
		client.PostItem.Delete().Where().ExecX(c.Context())
		client.Post.Delete().Where().ExecX(c.Context())
		client.User.Delete().Where().Exec(c.Context())

		SeedUsers(client)(c)
		SeedPosts(client)(c)
		SeedOrders(client)(c)
		SeedCompletes(client)(c)
		SeedComments(client)(c)
		SeedNotifies(client)(c)
		SeedDelivers(client)(c)
		SeedLikes(client)(c)
		SeedRooms(client)(c)
		SeedMessages(client)(c)

		return c.JSON(fiber.Map{})
	}
}
