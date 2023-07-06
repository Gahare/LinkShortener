package service

import (
	"LinkShortener"
	"LinkShortener/pkg/repository"
)

type LinkService struct {
	repo repository.Linker
}

func NewLinkService(repo repository.Linker) *LinkService {
	return &LinkService{repo: repo}
}

func (s *LinkService) ShortenLink(link LinkShortener.Link) (string, error) {
	return s.repo.ShortenLink(link)
}

func (s *LinkService) LengthenLink(link LinkShortener.Link) (string, error) {
	return s.repo.LengthenLink(link)
}
