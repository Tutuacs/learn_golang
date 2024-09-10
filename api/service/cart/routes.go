package cart

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/gorilla/mux"

	"github.com/Tutuacs/learn_golang/api.git/types"
	"github.com/Tutuacs/learn_golang/api.git/utils"
)

type Handler struct {
	store        types.OrderStore
	productStore types.ProductStore
}

func NewHandler(store types.OrderStore, productStore types.ProductStore) *Handler {
	return &Handler{store: store, productStore: productStore}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/cart/checkout", h.handleCheckout).Methods(http.MethodGet)
}

func (h *Handler) handleCheckout(w http.ResponseWriter, r *http.Request) {
	var cart types.CartCheckoutDTO
	if err := utils.ParseJSON(r, &cart); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(cart); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	productIDs, err := getCartItemsIDs(cart.Items)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	products, err := h.productStore.GetProductsByIDs(productIDs)
}
