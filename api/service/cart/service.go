package cart

import (
	"fmt"

	"github.com/Tutuacs/learn_golang/api.git/types"
)

func getCartItemsIDs(items []types.CartItem) ([]int, error) {
	productIDs := make([]int, len(items))
	for i, item := range items {
		if item.Quantity <= 0 {
			return nil, fmt.Errorf("invalid quantity for the product %v", item.ProductID)
		}

		productIDs[i] = item.ProductID
	}

	return productIDs, nil
}

func (h *Handler) createOrder(products []types.Product, items []types.CartItem, userID int) (int, float64, error) {
	productMap := make(map[int]types.Product)
	for _, product := range products {
		productMap[product.ID] = product
	}

	if err := checkIfCartIsInStock(items, productMap); err != nil {
		return 0, 0, nil
	}

	totalPrice := calculateTotalPrice(items, productMap)

	for _, item := range items {
		product := productMap[item.ProductID]
		product.Quantity -= item.Quantity

		h.productStore.UpdateProduct(product)
	}

}

func checkIfCartIsInStock(items []types.CartItem, products map[int]types.Product) error {
	if len(items) == 0 {
		return fmt.Errorf("cart is empty")
	}

	for _, item := range items {
		product, ok := products[item.ProductID]
		if !ok {
			return fmt.Errorf("product %d is not available in the store, please refresh your cart", item.ProductID)
		}

		if product.Quantity < int(item.Quantity) {
			return fmt.Errorf("product %s is not available in the quantity request", product.Name)
		}
	}

	return nil
}

func calculateTotalPrice(items []types.CartItem, products map[int]types.Product) float64 {
	var total float64

	for _, item := range items {
		product := products[item.ProductID]
		total += product.Price * float64(item.Quantity)
	}

	return total
}
