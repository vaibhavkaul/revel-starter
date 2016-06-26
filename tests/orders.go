package tests

import (
	"github.com/revel/revel/testing"
)

type OrdersControllerTest struct {
	testing.TestSuite
}

func (t *OrdersControllerTest) Before() {
	println("Set up")
}

func (t *OrdersControllerTest) After() {
	println("Tear down")
}

func (t *OrdersControllerTest) TestList() {
	t.Get("/orders")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}

func (t *OrdersControllerTest) TestShow() {
	t.Get("/orders/1")
	t.AssertOk()
	t.AssertContentType("application/json; charset=utf-8")
}