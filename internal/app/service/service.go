package service

import "time"

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) DaysLeft() int64 {
	data := time.Date(2024, time.January, 1, 0, 0, 0, 0, time.UTC)

	remains := time.Until(data)
	
	return int64(remains.Hours()) / 24
}
