package app

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"dolphin/salesManager/pkg/e"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, data interface{}) (int, int) {
	err := c.ShouldBind(&data)
	if err != nil {
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
