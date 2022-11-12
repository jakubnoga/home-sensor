package core

import (
	"homesensor/shared"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

type commandBus struct {
	handlers map[string]map[string]shared.CommandHandler
}

func NewCommandBus() shared.CommandBus {
	return &commandBus{
		handlers: make(map[string]map[string]shared.CommandHandler),
	}
}

func (b *commandBus) Subscribe(commandType string, handler shared.CommandHandler) shared.Unsubscribe {
	typeHandlers := b.handlers[commandType]
	if typeHandlers == nil {
		typeHandlers = make(map[string]shared.CommandHandler)
		b.handlers[commandType] = typeHandlers
	}
	key := gonanoid.Must()
	typeHandlers[key] = handler
	return func() {
		delete(typeHandlers, key)
	}
}

func (b *commandBus) Send(commandType string, command shared.Command) {
	typeHandlers := b.handlers[commandType]

	if typeHandlers == nil {
		return
	}

	for _, handler := range typeHandlers {
		handler(command)
	}
}
