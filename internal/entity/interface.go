package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	OrderList() ([]Order, error)
}
