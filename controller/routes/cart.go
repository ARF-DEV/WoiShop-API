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
	"gopkg.in/guregu/null.v3"
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

func GetCartByUserID(cartRepo *repository.CartRepository,
	orderRepo *repository.OrderRepository,
	productRepo *repository.ProductRepository) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		userData, ok := r.Context().Value("user_data").(models.User)
		if !ok {
			log.Println("userData not found")
			helpers.ErrorResponseJSON(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		cart, err := cartRepo.GetCartByUserID(userData.ID)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			log.Println("Error While getting Cart By ID : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if errors.Is(err, sql.ErrNoRows) {
			cart, _ = cartRepo.AddCart(models.Cart{
				UserID:           userData.ID,
				DeliveryMethodID: 1,
				Note:             null.NewString("", false),
			})
		}

		orders, err := orderRepo.GetOrderByCartID(cart.ID)

		if err != nil {
			log.Println("Error While getting Orders By Cart ID : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		var orderList []OrderResponseSchema = []OrderResponseSchema{}
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

		helpers.SuccessResponseJSON(w, "Success Getting All Carts", res)

	})
}

func DeleteCartByID(cartRepo *repository.CartRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		err = cartRepo.DeleteCartByID(id_int)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				helpers.ErrorResponseJSON(w, "Cart Not Found", http.StatusOK)
				return
			}
			log.Println("Error when deleting cart : ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success to cart", nil)
	}
}

func AddOrderToCart(cartRepo *repository.CartRepository, orderRepo *repository.OrderRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var OrderBody models.Order

		err := json.NewDecoder(r.Body).Decode(&OrderBody)
		if err != nil {
			log.Println("Error while decoding json: ", err.Error())
			helpers.ErrorResponseJSON(w, "Body is invalid", http.StatusBadRequest)
			return
		}

		_, err = cartRepo.GetCartByID(OrderBody.CartID)

		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			log.Println("Error while searching cart: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		if errors.Is(err, sql.ErrNoRows) {
			user, _ := r.Context().Value("user_data").(models.User)
			cartRepo.AddCart(models.Cart{
				UserID:           user.ID,
				DeliveryMethodID: 1,
			})
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

func UpdateCart(cartRepo *repository.CartRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var cartBody models.Cart

		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(&cartBody)

		if err != nil {
			log.Println("Error Decoding json : ", err.Error())
			helpers.ErrorResponseJSON(w, "Json Body Is Invalid", http.StatusBadRequest)
			return
		}

		updatedCart, err := cartRepo.UpdateCart(cartBody)
		if err != nil {
			log.Println("Error while updating cart: ", err.Error())
			helpers.ErrorResponseJSON(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		helpers.SuccessResponseJSON(w, "Success Updating cart", updatedCart)
	}
}

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
			helpers.ErrorResponseJSON(w, "id query required", http.StatusBadRequest)
			return
		}
		id_int, err := strconv.Atoi(id)

		if err != nil {
			log.Println("Error product by id : ", err.Error())
			helpers.ErrorResponseJSON(w, "Invalid id query", http.StatusBadRequest)
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
