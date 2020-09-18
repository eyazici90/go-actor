package application

type (
	StockCreatedEvent struct {
		ID         string
		LocationId string
		Amount     int
	}
	ShippedToLocationEvent struct {
		ID         string
		LocationId string
		Amount     int
	}

	ShippedFromLocationEvent struct {
		ID         string
		LocationId string
		Amount     int
	}
)
