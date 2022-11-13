package annotations

import (
	"homesensor/shared"

	"go.uber.org/fx"
)

func AsQueryHandlerRegistration(qhreg any) any {
	return fx.Annotate(
		qhreg,
		fx.ResultTags(`group:"qhregs"`),
		fx.As(new(shared.QueryHandlerRegistration)),
	)
}

func AsQueryBus(qb any) any {
	return fx.Annotate(
		qb,
		fx.ParamTags(``, `group:"qhregs"`),
		fx.As(new(shared.QueryBus)),
	)
}
