package handler

import (
	"LinkShortener"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) send(c *gin.Context) {
	var input LinkShortener.Link

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	short, err := h.services.Linker.ShortenLink(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"shortLink": short,
		})
	}
}

func (h *Handler) receive(c *gin.Context) {
	input := c.Param("shortLink")
	link := LinkShortener.Link{ShortLink: &input}
	long, err := h.services.Linker.LengthenLink(link)
	if err != nil {
		if err != sql.ErrNoRows {
			newErrorResponse(c, http.StatusInternalServerError, err.Error())
		} else {
			newErrorResponse(c, http.StatusNotFound, "No link with suck credentials exist")
		}
	} else {
		c.JSON(http.StatusOK, map[string]interface{}{
			"longLink": long,
		})
	}
}
