package main

import (
	"log"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/salom12/memcache/internal/cache"
	"github.com/salom12/memcache/internal/database"
	"github.com/salom12/memcache/internal/handlers"
	"github.com/salom12/memcache/internal/repositories"
	"github.com/salom12/memcache/pkg/memcache"
)

func main() {
	// init database
	if err := database.InitDB(); err != nil {
		log.Fatal(err)
	}

	// init memcache
	cache.InitMemcache(10, &memcache.LFUEviction{})

	// create echo instance
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// setting up repositories and handlers
	itemRepo := repositories.NewItemRepository(database.GetDB())
	itemHandler := handlers.NewItemHandler(itemRepo)

	api := e.Group("/api/v1")

	// register items routes
	items := api.Group("/items")
	items.GET("", itemHandler.GetItems)
	items.GET("/:id", itemHandler.GetItem)
	items.POST("", itemHandler.CreateItem)
	items.PUT("/:id", itemHandler.UpdateItem)
	items.DELETE("/:id", itemHandler.DeleteItem)

	e.Logger.Fatal(e.Start(":8080"))
}
