package repository

import (
	"LinkShortener"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type LinkPostgres struct {
	db *sqlx.DB
}

func NewLinkPostgres(db *sqlx.DB) *LinkPostgres {
	return &LinkPostgres{db: db}
}

func (r *LinkPostgres) ShortenLink(link LinkShortener.Link) (string, error) {
	regQuery := fmt.Sprintf("SELECT short_link FROM links ORDER BY id DESC LIMIT 1")
	regular := r.db.QueryRow(regQuery)
	var possibleShortLink, shortLink string
	var err error
	errReg := regular.Scan(&possibleShortLink)
	if errReg != nil {
		if errReg != sql.ErrNoRows {
			return "", errReg
		} else {
			shortLink = "aaaaaaaaaa"
		}
	} else {
		duplicate, dupErr := r.findDuplicates(*link.LongLink)
		if dupErr == nil {
			return duplicate, nil
		} else if dupErr != sql.ErrNoRows {
			return "", dupErr
		}
		shortLink, err = advanceToken(possibleShortLink)
		if err != nil {
			return "", err
		}
	}
	query := fmt.Sprintf("INSERT INTO links (long_link, short_link) values ($1, $2)")
	r.db.QueryRow(query, link.LongLink, shortLink)
	return shortLink, nil
}

func (r *LinkPostgres) findDuplicates(testLink string) (string, error) {
	dupQuery := fmt.Sprintf("SELECT short_link from links WHERE long_link = $1")
	duplicate := r.db.QueryRow(dupQuery, testLink)
	var shortLink string
	err := duplicate.Scan(&shortLink)
	if err != nil {
		return "", err
	}
	return shortLink, err
}

func (r *LinkPostgres) LengthenLink(link LinkShortener.Link) (string, error) {
	regQuery := fmt.Sprintf("SELECT long_link FROM links WHERE short_link = $1")
	regular := r.db.QueryRow(regQuery, link.ShortLink)
	var longLink string
	errReg := regular.Scan(&longLink)
	if errReg != nil {
		if errReg == sql.ErrNoRows {
			return "", errReg
		}
	}
	return longLink, nil
}
