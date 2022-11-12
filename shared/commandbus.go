package shared

type CommandBus interface {
	Subscribe(commandType string, handler CommandHandler) Unsubscribe
	Send(commandType string, query Command)
}

type CommandHandler func(payload Command) error

type Unsubscribe func()

type Command any
