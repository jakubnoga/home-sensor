package core

import (
	"fmt"
	"homesensor/shared"
)

type queryBus struct {
	handlers map[string]shared.QueryHandler
}

func NewQueryBus() shared.QueryBus {
	return &queryBus{
		handlers: make(map[string]shared.QueryHandler),
	}
}

func (b *queryBus) SetHandler(queryType string, handler shared.QueryHandler) shared.Unset {
	b.handlers[queryType] = handler

	return func() {
		delete(b.handlers, queryType)
	}
}

func (b *queryBus) Send(queryType string, query shared.Query) (shared.Result, error) {
	handler := b.handlers[queryType]

	if handler == nil {
		return nil, fmt.Errorf("unhandled query: %s", queryType)
	}

	return handler.Handle(query)
}