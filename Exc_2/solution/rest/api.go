package rest

import (
	"encoding/json"
	"net/http"
	"ordersystem/model"
	"ordersystem/repository"

	"github.com/go-chi/render"
)

// GetMenu 			godoc
// @tags 			Menu
// @Description 	Returns the menu of all drinks
// @Produce  		json
// @Success 		200 {array} model.Drink
// @Router 			/api/menu [get]
func GetMenu(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		drinks := db.GetDrinks()
		render.Status(r, http.StatusOK)
		render.JSON(w, r, drinks)
		// todo
		// get slice from db
		// render.Status(r, http.StatusOK)
		// render.JSON(w, r, <your-slice>) - DONE
	}
}

// GetOrders 		godoc
// @tags 			Order
// @Description 	Returns the orders of all drinks
// @Produce  		json
// @Success 		200 {array} model.Order
// @Router 			/api/order/all [get]
func GetOrders(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		orders := db.GetOrders()
		render.Status(r, http.StatusOK)
		render.JSON(w, r, orders)
	}
}

// todo create GetOrders /api/order/all -DONE

// GetMenu 			godoc
// @tags 			TotalOrder
// @Description 	Returns the total orders of all drinks
// @Produce  		json
// @Success 		200 {object} map[uint64]uint64
// @Router 			/api/order/totalled [get]
func GetOrdersTotal(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		totalledOrders := db.GetTotalledOrders()
		render.Status(r, http.StatusOK)
		render.JSON(w, r, totalledOrders)
	}
}

// todo create GetOrdersTotal /api/order/total -DONE

// PostOrder 		godoc
// @tags 			Order
// @Description 	Adds an order to the db
// @Accept 			json
// @Param 			b body model.Order true "Order"
// @Produce  		json
// @Success 		200
// @Failure     	400
// @Router 			/api/order [post]
func PostOrder(db *repository.DatabaseHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		order := model.Order{}

		err := json.NewDecoder(r.Body).Decode(&order)
		if err != nil {
			render.Status(r, http.StatusBadRequest)
			render.JSON(w, r, map[string]string{"error": "invalid request body"})
			return
		}

		db.AddOrder(&order)
		// todo
		// declare empty order struct
		// err := json.NewDecoder(r.Body).Decode(&<your-order-struct>)
		// handle error and render Status 400
		// add to db
		render.Status(r, http.StatusOK)
		render.JSON(w, r, "ok")
	}
}
