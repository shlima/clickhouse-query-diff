package cmd

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/static"
	"github.com/shlima/clickhouse-query-diff/internal/pkg/parser"
	"github.com/shlima/clickhouse-query-diff/internal/pkg/service/select_diff"
)

type PostSelectRequest struct {
	Left  string `json:"left"`
	Right string `json:"right"`
}

func Api() error {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	}))

	app.Post("/api/select", postSelect)
	app.Get("/*", static.New("./frontend/.output/public"))

	return app.Listen(":3333")
}

func postSelect(c fiber.Ctx) error {
	service := select_diff.New()

	req := new(PostSelectRequest)
	err := c.Bind().Body(req)
	if err != nil {
		return err
	}

	service.SetLeftSQL(req.Left)
	service.SetRightSQL(req.Right)

	diff, err := service.ColumnsDiffHTML()
	switch {
	case errors.Is(err, parser.ErrNoSelect):
		return c.Send(nil)
	case err != nil:
		return fmt.Errorf("failed to diff html: %w", err)
	default:
		return c.Send([]byte(diff))
	}
}
