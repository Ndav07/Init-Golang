package calculate_price

type OrderOutputDTO struct {
	ID string
	Price float64
	Tax float64
	FinalPrice float64
}