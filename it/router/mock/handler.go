package mock

import (
	"net/http"

	"github.com/tayalone/go-ess-package/router"
)

/*GetGlobalFromCtx get "global" from Ctx and return HTTP Response */
func GetGlobalFromCtx(c router.Context) {
	value, isExist := c.Get("global")
	if !isExist {
		c.JSON(http.StatusInternalServerError, map[string]string{
			"message": "can't get 'global' which setting from 'UseGlobal'",
		})
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
		"global":  value,
	})
}
