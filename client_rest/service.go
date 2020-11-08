package main

type Service struct {
	repo Repository
}

func NewstringService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}
