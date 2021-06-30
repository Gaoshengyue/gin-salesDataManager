package controller

import (
	"dolphin/salesManager/pkg/app"
	"dolphin/salesManager/pkg/e"
	"dolphin/salesManager/schema/ControllerSchema/TSRPortraitControllerSchema/TSRPortraitControllerRequestSchema"
	"dolphin/salesManager/schema/ServiceSchema/TSRPortraitSchema"
	"dolphin/salesManager/service/TSRPortraitService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary GetTSRGrade
// @Param gradeBody body TSRPortraitControllerRequestSchema.TSRGradeControllerRequest false "保单平均保费评分标准"
// @Success 200 {object} TSRPortraitControllerResponseSchema.TSRGradeControllerResponseSchema
// @Failure 500 {object} app.Response
// @Router /projectData/GetTSRGrade [Post]
func TSRGradeController(c *gin.Context) {
	appG := app.Gin{C: c}
	//例子，获取地区与省份的对照对象
	var TSRGradeRequest TSRPortraitControllerRequestSchema.TSRGradeControllerRequest
	//初始化默认参数
	TSRGradeRequest.InitDefaultRequest()
	//获取Query参数，根据地址复制到对象属性中　对应form
	if c.BindJSON(&TSRGradeRequest) != nil {
		appG.Response(http.StatusUnprocessableEntity, e.INVALID_PARAMS, nil)
		return
	}
	TSRGradeList, err := TSRPortraitService.CalculationTSRGrade(TSRGradeRequest)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_EXIST_CALCULATION_TSR_GRADE, nil)
	} else {
		appG.Response(http.StatusOK, e.SUCCESS, TSRGradeList)
	}

	return
}

var f TSRPortraitSchema.TsrStarPageResponseSchema

// @Summary GetTSRTop
// @Param topBody body TSRPortraitControllerRequestSchema.TSRTopControllerRequest false "坐席排名"
// @Success 200 {object} TSRPortraitControllerResponseSchema.TSRTopControllerResponseSchema
// @Failure 500 {object} app.Response
// @Router /projectData/GetTSRTop [Post]
func GetTSRTopController(c *gin.Context) {
	appG := app.Gin{C: c}
	var TSRTopRequest TSRPortraitControllerRequestSchema.TSRTopControllerRequest
	TSRTopRequest.InitDefaultRequest()
	// //获取Query参数，根据地址复制到对象属性中　对应form

	if c.BindJSON(&TSRTopRequest) != nil {
		appG.Response(http.StatusUnprocessableEntity, e.INVALID_PARAMS, nil)
		return
	}
	TSRStarsPage, err := TSRPortraitService.GetTSRTop(TSRTopRequest)
	if err != nil {
		appG.Response(http.StatusBadRequest, e.ERROR_EXIST_CALCULATION_TSR_TOP, nil)
	}
	// if TSRStars == nil {
	// 	appG.Response(http.StatusBadRequest, e.ERROR_EXIST_CALCULATION_TSR_TOP, nil)
	// }
	appG.Response(http.StatusOK, e.SUCCESS, TSRStarsPage)

	return
}
