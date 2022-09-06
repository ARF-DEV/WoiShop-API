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

type OrderResponseSchema struct {
	models.Order
	ProductName  string `json:"product_name"`
	ProductPrice int    `json:"product_price"`
}

type CartResponseSchema struct {
	models.Cart
	Orders []OrderResponseSchema `json:"orders"`
}

func GetCartByID(cartRepo *repository.CartRepository,
	orderRepo *repository.OrderRepository,
	productRepo *repository.ProductRepository) http.HandlerFunc {

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
		var orderList []OrderResponseSchema
		for _, order := range orders {
			var or OrderResponseSchema
			product, err := productRepo.GetProductByID(order.ProductID)

			if err != nil {
				log.Println("Error While getting products : ", err.Error())
				helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			or = OrderResponseSchema{
				Order:        order,
				ProductName:  product.Name,
				ProductPrice: or.ProductPrice,
			}

			orderList = append(orderList, or)
		}

		res := CartResponseSchema{
			Cart:   *cart,
			Orders: orderList,
		}

		helpers.SuccessResponseJSON(w, "Success Getting Cart By ID", res)

	})
}

func GetAllCart(cartRepo *repository.CartRepository,
	orderRepo *repository.OrderRepository,
	productRepo *repository.ProductRepository) http.HandlerFunc {
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

			var orderList []OrderResponseSchema
			for _, order := range orders {
				var or OrderResponseSchema
				product, err := productRepo.GetProductByID(order.ProductID)

				if err != nil {
					log.Println("Error While getting products : ", err.Error())
					helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
					return
				}

				or = OrderResponseSchema{
					Order:        order,
					ProductName:  product.Name,
					ProductPrice: product.Price,
				}

				orderList = append(orderList, or)
			}
			cartRes = CartResponseSchema{
				Cart:   carts[i],
				Orders: orderList,
			}

			res = append(res, cartRes)

		}

		helpers.SuccessResponseJSON(w, "Succes Getting All Carts", res)

	})
}
