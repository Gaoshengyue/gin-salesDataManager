package clientService

import (
	"dolphin/salesManager/pkg/app"
	"dolphin/salesManager/pkg/e"
	"dolphin/salesManager/schema/ServiceSchema/clientSchema"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Add Client
// @Produce  json
// @Param app_id body string true "app_id"
// @Param app_secret body string true "app_secret"
// @Param desc body string true "desc"
// @Param app_name body string true "app_name"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/auth/addClient [post]
func AddClient(c *gin.Context) {
	var (
		appG          = app.Gin{C: c}
		addClientForm clientSchema.AddClientForm
	)

	httpCode, errCode := app.BindAndValid(c, addClientForm)
	if errCode != e.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}
	clientSchemaObj := clientSchema.Client{AppId: addClientForm.AppId}
	exists, err := clientSchemaObj.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_CALCULATION_TSR_GRADE, nil)
		return
	}
	fmt.Println(exists)
	if exists {
		appG.Response(http.StatusOK, e.ERROR_EXIST_CALCULATION_TSR_GRADE, nil)
		return
	}

	clientService := clientSchema.Client{
		AppId:     addClientForm.AppId,
		AppSecret: addClientForm.AppSecret,
		AppName:   addClientForm.AppName,
		Desc:      addClientForm.Desc,
	}
	if err := clientService.Add(); err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_CALCULATION_TSR_GRADE, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}
