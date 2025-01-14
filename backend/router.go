package main

import (
	"fmt"
	"log"

	"github.com/efydb/config"
	"github.com/efydb/handlers"
	"github.com/efydb/util"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

const portStr = ":8080"

func CreateRouter() {
	router := fiber.New(
		fiber.Config{
			BodyLimit:   20 * 1024 * 1024,
			JSONEncoder: json.Marshal,
			JSONDecoder: json.Unmarshal,
		},
	)

	router.Use(cors.New(cors.Config{
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	router.Get("/", func(c *fiber.Ctx) error {
		return util.OkResponse(c)
	})

	router.Static("/files", fmt.Sprintf("%s/files/", config.RootDir()))

	user := router.Group("/users")
	user.Get("/", handlers.GetUsers)
	user.Get("/:name", handlers.GetUserInfo)
	user.Get("/account", handlers.GetUser)
	user.Post("/register", handlers.CreateUser)
	user.Post("/login", handlers.LoginUser)
	user.Patch("/update", handlers.UpdateUser)
	user.Post("/promote", handlers.PromoteUser)
	user.Delete("/delete", handlers.DeleteUser)

	themes := router.Group("/themes")
	themes.Get("/", handlers.GetThemes)
	themes.Get("/:id", handlers.GetTheme)
	themes.Post("/create", handlers.CreateTheme)
	themes.Patch("/edit", handlers.EditTheme)
	themes.Delete("/delete", handlers.DeleteTheme)
	themes.Post("/approve", handlers.ApproveTheme)

	fmt.Printf("Listening on http://localhost%s/", portStr)
	log.Fatal(router.Listen(portStr))
}
