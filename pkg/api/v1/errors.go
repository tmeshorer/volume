package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	unsuccessful = Reply{Success: false}
	notFound     = Reply{Success: false, Error: "resource not found"}
	notAllowed   = Reply{Success: false, Error: "method not allowed"}
)

// ErrResp contain new error response
func ErrResp(err string) Reply {
	rep := Reply{Success: false}
	rep.Error = err
	return rep
}

// NotFound returns 404 response for the API.
func NotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, notFound)
}

// NotAllowed returns 405 response for the API.
func NotAllowed(c *gin.Context) {
	c.JSON(http.StatusMethodNotAllowed, notAllowed)
}
