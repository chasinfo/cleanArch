package database

import (
	"database/sql"

	"github.com/chasinfo/cleanArch/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) OrderList() ([]entity.Order, error) {
	rows, err := r.Db.Query("SELECT id, price, tax, final_price FROM orders ORDER BY id")

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	orders := []entity.Order{}

	for rows.Next() {
		var o entity.Order
		err = rows.Scan(&o.ID, &o.Price, &o.Tax, &o.FinalPrice)

		if err != nil {
			panic(err)
		}
		orders = append(orders, o)
	}

	return orders, nil
}
