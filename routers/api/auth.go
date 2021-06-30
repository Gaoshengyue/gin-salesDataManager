package api

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"dolphin/salesManager/pkg/app"
	"dolphin/salesManager/pkg/e"
	"dolphin/salesManager/pkg/util"
	"dolphin/salesManager/schema/ServiceSchema/authSchema"
)

// @Summary Get Auth
// @Produce  json
// @Param app_id query string true "app_id"
// @Param app_secret query string true "app_secret"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [get]
func GetAuth(c *gin.Context) {
	appG := app.Gin{C: c}
	valid := validation.Validation{}

	appId := c.PostForm("AppId")
	appSecret := c.PostForm("AppSecret")

	a := authSchema.Auth{AppId: appId, AppSecret: appSecret}
	ok, _ := valid.Valid(&a)

	if !ok {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	authServiceObj := authSchema.Auth{AppId: appId, AppSecret: appSecret}
	isExist, err := authServiceObj.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, e.ERROR_AUTH, nil)
		return
	}

	token, err := util.GenerateToken(appId, appSecret)
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"token": token,
	})
}
