package service

import (
	"LinkShortener"
	"LinkShortener/pkg/repository"
)

type Linker interface {
	ShortenLink(longLink LinkShortener.Link) (string, error)
	LengthenLink(longLink LinkShortener.Link) (string, error)
}

type Service struct {
	Linker
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Linker: NewLinkService(repos.Linker),
	}
}
