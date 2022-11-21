package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html"
)

func main() {

	rooturl := "http://localhost:8080/"

	app := fiber.New(fiber.Config{
		Views: html.New("./templates", ".html"),
	})
	app.Static("/assets", "./web_res")
	app.Use(recover.New())

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect(rooturl + "home")
	})
	app.Get("/home", func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{
			"Title":     "example->home",
			"Version":   "v1.0",
			"Webpath":   rooturl,
			"Ram":       "0",
			"IsRunning": true,
			"Dbstats":   map[string]string{"a": "b", "c": "d"},
		})
	})
	app.Get("/list/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		return c.Render("id", fiber.Map{
			"Title":   "example->id",
			"Webpath": rooturl,
			"ID":      id,
		})
	})

	app.Get("/list", func(c *fiber.Ctx) error {
		return c.Render("list", fiber.Map{
			"Title":   "example->list",
			"Webpath": rooturl,
			"List":    []string{"a", "b", "c", "d"},
		})
	})
	api := app.Group("/api")
	//the following is accessible via /api/return, or: http://localhost:8080/api/return
	api.Get("/return", func(c *fiber.Ctx) error {
		return c.Redirect(rooturl + "home")
	})

	app.Listen(":8080")
}
