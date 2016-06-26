package controllers

import (
	"github.com/revel/revel"
	"revel-starter/app/models"
)

type OrdersController struct {
	*revel.Controller
}

var orders = []models.Order{
	models.Order{1, "Event 1", 1},
	models.Order{2, "Event 2", 2},
	models.Order{3, "Event 3", 3},
}

// Return all Orders
func (c OrdersController) List() revel.Result {
	return c.RenderJson(orders)
}

// Return Orders by :id
func (c OrdersController) Show(orderID int) revel.Result {
	var res models.Order

	for _, order := range orders {
		if order.ID == orderID {
			res = order
			break
		}
	}

	if res.ID == 0 {
		return c.NotFound("Could not find beer")
	}

	return c.RenderJson(res)
}
