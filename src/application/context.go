package application

import (
	"go-projection/infrastructure"

	"github.com/AsynkronIT/protoactor-go/actor"
)

var emptyRootContext *actor.RootContext = actor.EmptyRootContext

var props *actor.Props

func init() {
	persister := infrastructure.NewPersister()

	props = actor.PropsFromProducer(func() actor.Actor { return NewStockActor(persister) })

}

func Send(id string, msg interface{}) {
	pid, _ := emptyRootContext.SpawnNamed(props, id)

	emptyRootContext.Send(pid, msg)
}
