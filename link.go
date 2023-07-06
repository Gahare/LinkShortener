package LinkShortener

type Link struct {
	Id        int     `json:"id"`
	LongLink  *string `json:"longLink"`
	ShortLink *string `json:"shortLink"`
}
