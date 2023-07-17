package status

import "github.com/gin-gonic/gin"

type (
	response struct {
		Text string `json:"text"`
	}

	Handler struct {
		resp response
	}
)

// NewHandler defines a handler constructor.
func NewHandler() *Handler {
	return &Handler{
		resp: response{
			Text: "OK",
		},
	}
}

// CheckStatus -  HTTP GET handler for status endpoint.
func (h Handler) CheckStatus() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.JSON(200, h.resp)
	}
}
