package order

import (
	"database/sql"

	"github.com/Tutuacs/learn_golang/api.git/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(order types.Order) (int, error) {

	sql := "INSERT INTO orders (userId, total, status, address) VALUES (?,?,?,?)"

	res, err := s.db.Exec(sql, order.UserId, order.Total, order.Status, order.Address)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil

}

func (s *Store) CreateOrderItem(orderItem types.OrderItem) error {
	sql := "INSERT INTO order_items (orderId, productId, quantity, price) VALUES (?, ?, ?, ?)"
	_, err := s.db.Exec(sql, orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}
