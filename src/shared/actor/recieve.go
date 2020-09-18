package actor

import (
	"reflect"

	"github.com/AsynkronIT/protoactor-go/actor"
)

type ReciveActor struct {
	State    interface{}
	handlers map[reflect.Type]reflect.Value
}

func NewReciveActor() *ReciveActor {
	return &ReciveActor{}
}

func (a *ReciveActor) When(handler interface{}) {
	hVal := reflect.ValueOf(handler)

	if hVal.Kind() != reflect.Func {
		panic("handler type should be func(ctx actor.Context, msg interface{})")
	}

	hType := hVal.Type()

	msgType := hType.In(1)

	if a.handlers == nil {
		a.handlers = make(map[reflect.Type]reflect.Value)
	}

	a.handlers[msgType] = hVal
}

func (a *ReciveActor) Receive(ctx actor.Context) {
	msg := ctx.Message()
	msgVal := reflect.ValueOf(msg)

	msgType := msgVal.Type()

	handler := a.handlers[msgType]

	in := []reflect.Value{reflect.ValueOf(ctx), msgVal}

	handler.Call(in)
}
