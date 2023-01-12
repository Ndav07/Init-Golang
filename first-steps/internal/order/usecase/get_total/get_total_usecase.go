package get_total

import database "ql/first-steps/internal/order/infra/database/order_repository"

type GetTotalUseCase struct {
	OrderRepository database.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository database.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {
	total, err := c.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}
	return &GetTotalOutputDTO{
		Total: total,
	}, nil 
}