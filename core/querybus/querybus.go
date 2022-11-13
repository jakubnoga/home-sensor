package querybus

import (
	"fmt"
	"homesensor/shared"

	"go.uber.org/zap"
)

type QueryBus struct {
	log      *zap.Logger
	handlers map[string]shared.QueryHandler
}

func NewQueryBus(log *zap.Logger, registrations ...shared.QueryHandlerRegistration) *QueryBus {
	log.Sugar().Infof("Creating QueryBus with %d handlers", len(registrations))
	handlers := make(map[string]shared.QueryHandler)
	bus := &QueryBus{log, handlers}
	for _, reg := range registrations {
		bus.SetHandler(reg.QueryType(), reg)
	}

	return bus
}

func (b *QueryBus) SetHandler(queryType string, handler shared.QueryHandler) shared.Unset {
	b.log.Sugar().Infof("Setting QueryHandler for type: %s", queryType)
	b.handlers[queryType] = handler

	return func() {
		delete(b.handlers, queryType)
	}
}

func (b *QueryBus) Send(queryType string, query shared.Query) (shared.Result, error) {
	handler := b.handlers[queryType]

	if handler == nil {
		return nil, fmt.Errorf("unhandled query: %s", queryType)
	}

	return handler.Handle(query)
}
