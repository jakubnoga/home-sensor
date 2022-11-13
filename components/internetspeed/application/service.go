package application

import (
	"homesensor/components/internetspeed/domain"
	"homesensor/components/internetspeed/ports"
)

type InternetSpeedApplicationService interface {
	CheckInternetSpeed() (domain.InternetSpeed, error)
}

type HsInternetSpeedApplicationService struct {
	adapter ports.CheckInternetSpeedPort
}

func NewInternetSpeedApplicationService(adapter ports.CheckInternetSpeedPort) *HsInternetSpeedApplicationService {
	return &HsInternetSpeedApplicationService{
		adapter,
	}
}

func (s *HsInternetSpeedApplicationService) CheckInternetSpeed() (domain.InternetSpeed, error) {
	return s.adapter.CheckInternetSpeed()
}
