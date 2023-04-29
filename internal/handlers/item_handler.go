package handlers

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo"
	"github.com/salom12/memcache/internal/cache"
	"github.com/salom12/memcache/internal/repositories"
)

type ItemHandler struct {
	r *repositories.ItemRepository
}

func NewItemHandler(r *repositories.ItemRepository) *ItemHandler {
	return &ItemHandler{r: r}
}

func (h *ItemHandler) GetItems(c echo.Context) error {
	items, err := h.r.Get()
	if err != nil {
		return fmt.Errorf("couldn't get items")
	}
	return c.JSON(200, items)
}

func (h *ItemHandler) GetItem(c echo.Context) error {
	id := c.Param("id")

	mc := cache.GetMemcache()

	// check if item is already in memory
	data, err := mc.Get(id)
	if err == nil {
		fmt.Println("get from cache")
		return c.JSON(200, data)
	}

	ui64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	fmt.Println("get from db")
	item, err := h.r.GetByID(ui64)
	if err != nil {
		return err
	}

	// cache item
	mc.Set(id, item)

	return c.JSON(200, item)
}

func (h *ItemHandler) CreateItem(c echo.Context) error {
	return nil
}

func (h *ItemHandler) UpdateItem(c echo.Context) error {
	return nil
}

func (h *ItemHandler) DeleteItem(c echo.Context) error {
	id := c.Param("id")

	ui64, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return err
	}

	return c.JSON(200, h.r.Delete(ui64))
}
