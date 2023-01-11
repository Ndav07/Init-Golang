package usecase

import (
	"ql/first-steps/internal/order/entity"
	"ql/first-steps/internal/order/infra/database/order_repository"
)

type CalculateFinalPriceUseCase struct {
	OrderRepository database.OrderRepositoryInterface
}

func NewCalculateFinalPriceUseCase(orderRepository database.OrderRepository) *CalculateFinalPriceUseCase {
	return &CalculateFinalPriceUseCase{
		OrderRepository: &orderRepository,
	}
}

func (c *CalculateFinalPriceUseCase) Execute(input OrderInputDTO) (*OrderOutputDTO, error) {
	order, err := entity.NewOder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}
	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}
	return &OrderOutputDTO{
		ID: order.ID,
		Price: order.Price,
		Tax: order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}