package database

import "ql/first-steps/internal/order/entity"

type OrderRepositoryInterface interface {
	Save(order *entity.Order) error
	GetTotal() (int, error)
}