package usecase

func (c *CreateOrderUseCase) GetOrders() ([]OrderOutputDTO, error) {

	orders, err := c.OrderRepository.OrderList()

	if err != nil {
		return []OrderOutputDTO{}, err
	}

	var dtos []OrderOutputDTO

	for _, order := range orders {

		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}

		dtos = append(dtos, dto)
	}

	return dtos, nil
}
