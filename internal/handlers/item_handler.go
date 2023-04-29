package handlers

import (
	"github.com/labstack/echo"
	"github.com/salom12/memcache/internal/repositories"
)

type ItemHandler struct {
	r *repositories.ItemRepository
}

func NewItemHandler(r *repositories.ItemRepository) *ItemHandler {
	return &ItemHandler{r: r}
}

func (h *ItemHandler) GetItems(c echo.Context) error {
	return nil
}

func (h *ItemHandler) GetItem(c echo.Context) error {
	return nil
}

func (h *ItemHandler) CreateItem(c echo.Context) error {
	return nil
}

func (h *ItemHandler) UpdateItem(c echo.Context) error {
	return nil
}

func (h *ItemHandler) DeleteItem(c echo.Context) error {
	return nil
}
