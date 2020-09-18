package application

import (
	"go-projection/infrastructure"
	"go-projection/shared/actor"

	. "github.com/AsynkronIT/protoactor-go/actor"
)

type stockActor struct {
	*actor.ReciveActor
	persister infrastructure.Persister
}

func NewStockActor(p infrastructure.Persister) Actor {
	a := &stockActor{
		ReciveActor: actor.NewReciveActor(),
		persister:   p,
	}

	a.When(func(ctx Context, msg *Started) {
		recoverState(a, ctx)
	})

	a.When(func(_ Context, msg StockCreatedEvent) {
		a.State = infrastructure.Stock{
			Id:         msg.ID,
			LocationID: msg.LocationId,
			Quantity:   msg.Amount,
		}
	})

	a.When(func(_ Context, msg ShippedToLocationEvent) {
		update(a, func(s *infrastructure.Stock) {
			s.Quantity += msg.Amount
		})
	})

	a.When(func(_ Context, msg ShippedFromLocationEvent) {
		update(a, func(s *infrastructure.Stock) {
			s.Quantity -= msg.Amount
		})
	})

	return a
}

func recoverState(s *stockActor, ctx Context) {

	actorID := ctx.Self().Id

	stock := s.persister.Get(actorID)

	s.State = stock
}

func update(s *stockActor, mutator func(s *infrastructure.Stock)) {
	stock := s.State.(infrastructure.Stock)

	mutator(&stock)

	s.persister.Update(stock)

	s.State = stock
}
