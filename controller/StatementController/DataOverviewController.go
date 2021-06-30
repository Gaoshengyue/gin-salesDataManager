package StatementController

import (
	"dolphin/salesManager/pkg/app"
	"dolphin/salesManager/pkg/e"
	"dolphin/salesManager/service/StatementService"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary GetDataOverviewBaseData
// @Success 200 {object} StatementControllerResponseSchema.BaseDataResponseSchema
// @Failure 500 {object} app.Response
// @Router /projectData/GetDataOverviewBaseData [Get]
func DataOverviewBaseDataController(c *gin.Context) {
	appG := app.Gin{C: c}
	//　获取数据概览基础数据
	DataOverviewBaseData, err := StatementService.GetBaseData()
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_EXIST_CALCULATION_TSR_GRADE, nil)
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, DataOverviewBaseData)
	}
}

// @Summary GetDataOverviewProductData
// @Success 200 {object} StatementControllerResponseSchema.ProductDataResponseSchema
// @Failure 500 {object} app.Response
// @Router /projectData/GetDataOverviewProductData [Get]
func DataOverviewProductController(c *gin.Context) {
	appG := app.Gin{C: c}
	//　获取数据概览各个产品数据分布
	ProductDistributionArray, err := StatementService.GetProductDistribution()
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_EXIST_CALCULATION_TSR_GRADE, nil)
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, ProductDistributionArray)
	}
}
