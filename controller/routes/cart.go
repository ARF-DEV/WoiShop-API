package routes

import (
	"azura-lab-intern/study-case-1/helpers"
	"azura-lab-intern/study-case-1/models"
	"azura-lab-intern/study-case-1/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type CartResponseSchema struct {
	models.Cart
	Orders []models.Order `json:"orders"`
}

func GetCartByID(cartRepo *repository.CartRepository, orderRepo *repository.OrderRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		if len(id) < 1 {
			log.Println("Error product by id : id query not found")
			helpers.ErrorResponseJSON(w, "id query required", http.StatusBadRequest)
			return
		}
		id_int, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Error product by id : ", err.Error())
			helpers.ErrorResponseJSON(w, "Invalid id query", http.StatusBadRequest)
			return
		}

		cart, err := cartRepo.GetCartByID(id_int)

		if err != nil {
			log.Println("Error While getting Cart By ID : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		orders, err := orderRepo.GetOrderByCartID(id_int)

		if err != nil {
			log.Println("Error While getting Orders By Cart ID : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		res := CartResponseSchema{
			Cart:   *cart,
			Orders: orders,
		}

		helpers.SuccessResponseJSON(w, "Success Getting Cart By ID", res)

	})
}

func GetAllCart(cartRepo *repository.CartRepository, orderRepo *repository.OrderRepository) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		carts, err := cartRepo.GetAllCart()

		if err != nil {
			log.Println("Error While getting Carts : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		var res []CartResponseSchema

		for i := range carts {
			var cartRes CartResponseSchema
			orders, err := orderRepo.GetOrderByCartID(carts[i].ID)

			if err != nil {
				log.Println("Error While getting orders : ", err.Error())
				helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			cartRes = CartResponseSchema{
				Cart:   carts[i],
				Orders: orders,
			}

			res = append(res, cartRes)

		}

		helpers.SuccessResponseJSON(w, "Succes Getting All Carts", res)

	})
}
