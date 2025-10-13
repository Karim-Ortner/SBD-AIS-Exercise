package repository

import (
	"ordersystem/model"
	"time"
)

type DatabaseHandler struct {
	// drinks represent all available drinks
	drinks []model.Drink
	// orders serves as order history
	orders []model.Order
}

// todo
func NewDatabaseHandler() *DatabaseHandler {
	drinks := []model.Drink{
		{ID: 1, Name: "Coca Cola", Price: 2.5, Description: "Classic Coke"},
		{ID: 2, Name: "Pepsi", Price: 2.0, Description: "Pepsi Cola"},
		{ID: 3, Name: "Orange Juice", Price: 3.0, Description: "Freshly squeezed orange juice"},
	}
	orders := []model.Order{
		{DrinkID: 1, CreatedAt: time.Now(), Amount: 2},
		{DrinkID: 2, CreatedAt: time.Now(), Amount: 1},
		{DrinkID: 1, CreatedAt: time.Now(), Amount: 3},
	}
	// Init the drinks slice with some test data -DONE
	// drinks := ...

	// Init orders slice with some test data -DONE

	return &DatabaseHandler{
		drinks: drinks,
		orders: orders,
	}
}

func (db *DatabaseHandler) GetDrinks() []model.Drink {
	return db.drinks
}

func (db *DatabaseHandler) GetOrders() []model.Order {
	return db.orders
}

// todo
func (db *DatabaseHandler) GetTotalledOrders() map[uint64]uint64 {
	totalledOrders := make(map[uint64]uint64)
	for _, order := range db.orders {
		totalledOrders[order.DrinkID] += order.Amount
	}
	// calculate total orders
	// key = DrinkID, value = Amount of orders
	// totalledOrders map[uint64]uint64 -DONE
	return totalledOrders
}

func (db *DatabaseHandler) AddOrder(order *model.Order) {
	db.orders = append(db.orders, *order)
	// todo
	// add order to db.orders slice -DONE
}
