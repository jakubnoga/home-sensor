package shared

import (
	"fmt"
)

type QueryBus interface {
	SetHandler(queryType string, handler QueryHandler) Unset
	Send(queryType string, query Query) (Result, error)
}

type QueryHandler interface {
	Handle(payload Query) (Result, error)
}

type QueryHandlerRegistration interface {
	QueryHandler
	QueryType() string
}

type QueryHandlerFunc func(payload Query) (Result, error)

func (f QueryHandlerFunc) Handle(payload Query) (Result, error) {
	return f(payload)
}

type Unset func()

type Query any
type Result any

type hsQueryBus struct {
	handlers map[string]QueryHandler
}

func NewHsQueryBus() QueryBus {
	return &hsQueryBus{
		handlers: make(map[string]QueryHandler),
	}
}

func (b *hsQueryBus) SetHandler(queryType string, handler QueryHandler) Unset {
	b.handlers[queryType] = handler

	return func() {
		delete(b.handlers, queryType)
	}
}

func (b *hsQueryBus) Send(queryType string, query Query) (Result, error) {
	handler := b.handlers[queryType]

	if handler == nil {
		return nil, fmt.Errorf("unhandled query: %s", queryType)
	}

	return handler.Handle(query)
}
