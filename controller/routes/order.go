package routes

import (
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/models"
	"azura-lab-intern/study-case-1/repository"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func UpdateOrder(orderRepo *repository.OrderRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var orderBody models.Order

		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&orderBody)

		if err != nil {
			log.Println("Error Decoding json : ", err.Error())
			helpers.ErrorResponseJSON(w, "Json Body Is Invalid", http.StatusBadRequest)
			return
		}
		log.Println(orderBody)
		log.Println("HELLO")
		updatedOrder, err := orderRepo.ChangeOrderAmount(orderBody.ID, orderBody.Amount)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {

				log.Println("Order Not found: ", err.Error())
				helpers.ErrorResponseJSON(w, "Order Not Found", http.StatusOK)
				return
			}
			log.Println("Error while updating order: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success Updating cart", updatedOrder)
	}
}

func DeleteOrderByID(orderRepo *repository.OrderRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if len(id) < 1 {
			log.Println("Error product by id : id query not found")
			helpers.ErrorResponseJSON(w, "id params required", http.StatusBadRequest)
			return
		}
		id_int, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Error product by id : ", err.Error())
			helpers.ErrorResponseJSON(w, "Invalid id params", http.StatusBadRequest)
			return
		}

		deletedOrder, err := orderRepo.DeleteOrderByID(id_int)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				helpers.ErrorResponseJSON(w, "Order Not Found", http.StatusOK)
				return
			}
			log.Println("Error when deleting order : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success to delete order", deletedOrder)
	}
}
func CreateOrder(cartRepo *repository.CartRepository, orderRepo *repository.OrderRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var OrderBody models.Order

		err := json.NewDecoder(r.Body).Decode(&OrderBody)
		if err != nil {
			log.Println("Error while decoding json: ", err.Error())
			helpers.ErrorResponseJSON(w, "Body is invalid", http.StatusBadRequest)
			return
		}

		_, err = cartRepo.GetCartByID(OrderBody.CartID)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {

				helpers.ErrorResponseJSON(w, "Cart is not found, please create the cart first", http.StatusInternalServerError)
				return
			}
			log.Println("Error while searching cart: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		newOrder, err := orderRepo.CreateOrder(OrderBody)

		if err != nil {
			log.Println("Error creating order: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		helpers.SuccessResponseJSON(w, "Success Creating Order", newOrder)
	}
}
