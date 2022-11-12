package application

import "homesensor/shared"

const (
	queryType = "CheckInternetSpeed"
)

type InternetSpeedQueryHandler struct {
	service InternetSpeedApplicationService
}

func NewHsInternetSpeedQueryHandler(service InternetSpeedApplicationService) shared.QueryHandlerRegistration {
	return &InternetSpeedQueryHandler{
		service,
	}
}

func (qh *InternetSpeedQueryHandler) Handle(payload shared.Query) (shared.Result, error) {
	return qh.service.CheckInternetSpeed()
}

func (qh *InternetSpeedQueryHandler) QueryType() string {
	return queryType
}