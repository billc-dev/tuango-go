package handler_admin

import (
	"log"
	"net/http"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/billc-dev/tuango-go/database"
	"github.com/billc-dev/tuango-go/ent/deliver"
	"github.com/billc-dev/tuango-go/ent/predicate"
	"github.com/gofiber/fiber/v2"
)

func GetFinance(c *fiber.Ctx) error {
	m := c.Queries()

	deliverWhere := []predicate.Deliver{}

	if from := m["from"]; from != "" {
		t, err := time.Parse("2006-01-02", from) // TODO: double check time zone
		if err != nil {
			log.Print(err)
			return fiber.NewError(http.StatusBadRequest, "Could not parse from date")
		}
		deliverWhere = append(deliverWhere, deliver.CreatedAtGTE(t))
	}

	if to := m["to"]; to != "" {
		t, err := time.Parse("2006-01-02", to)
		if err != nil {
			log.Print(err)
			return fiber.NewError(http.StatusBadRequest, "Could not parse to date")
		}
		deliverWhere = append(deliverWhere, deliver.CreatedAtGTE(t))
	}

	var results []struct {
		Year  int
		Month int
		Day   int
		Total float64
		Fee   float64
	}

	err := database.DBConn.Debug().Deliver.Query().
		Where(deliverWhere...).
		Modify(func(s *sql.Selector) {
			// s.Where(sql.GTE(deliver.FieldCreatedAt, "2022-01-04T16:28:18.65236Z"))
			s.Select(
				sql.As("extract(year from created_at)", "year"),
				sql.As("extract(month from created_at)", "month"),
				sql.As("extract(day from created_at)", "day"),
				sql.As("SUM(normal_total + extra_total)", "total"),
				sql.As("SUM(normal_fee + extra_fee)", "fee"),
			).GroupBy(
				"extract(year from created_at)",
				"extract(month from created_at)",
				"extract(day from created_at)",
			).OrderBy("year", "month", "day")
		}).
		Scan(c.Context(), &results)

	if err != nil {
		log.Print(err)
		return fiber.NewError(http.StatusInternalServerError, "Could not start transaction")
	}

	return c.JSON(fiber.Map{
		"data":  results,
		"count": len(results),
	})
}
