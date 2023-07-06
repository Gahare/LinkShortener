package repository

import (
	"LinkShortener"
	"errors"
	"github.com/jmoiron/sqlx"
)

type Linker interface {
	ShortenLink(longLink LinkShortener.Link) (string, error)
	LengthenLink(longLink LinkShortener.Link) (string, error)
}

type Repository struct {
	Linker
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Linker: NewLinkPostgres(db),
	}
}

func advanceToken(linkToken string) (string, error) {
	if len(linkToken) == 0 {
		return "", errors.New("overflow error")
	}
	letterRunes := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	endLetter := linkToken[len(linkToken)-1:]
	slicedToken := linkToken[0 : len(linkToken)-1]
	var err error
	if endLetter == "_" {
		endLetter = "a"
		slicedToken, err = advanceToken(slicedToken)
		if err != nil {
			return "", err
		}
	} else {
		for i := 0; i < len(letterRunes)-1; i++ {
			if endLetter == string(letterRunes[i]) {
				endLetter = string(letterRunes[i+1])
				break
			}
		}
	}
	return slicedToken + endLetter, nil
}
