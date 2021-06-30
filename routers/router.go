package routers

import (
	"dolphin/salesManager/controller"
	"dolphin/salesManager/controller/StatementController"
	"dolphin/salesManager/middleware/cors"
	"net/http"

	"github.com/gin-gonic/gin"

	_ "dolphin/salesManager/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	middlewarejwt "dolphin/salesManager/middleware/jwt"
	"dolphin/salesManager/pkg/export"
	"dolphin/salesManager/pkg/qrcode"
	"dolphin/salesManager/pkg/upload"
	"dolphin/salesManager/routers/api"
	"dolphin/salesManager/service/clientService"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	url := ginSwagger.URL("http://192.168.59.65:8030/swagger/doc.json")
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Cors())

	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.POST("/upload", api.UploadImage)
	r.POST("/api/auth/addClient", clientService.AddClient)
	//r.GET("/GetTSRGrade", controller.TSRGradeController)

	projectRoute := r.Group("/projectData")
	projectRoute.POST("/GetTSRGrade", controller.TSRGradeController)
	projectRoute.POST("/GetTSRTop", controller.GetTSRTopController)
	projectRoute.GET("/GetDataOverviewBaseData", StatementController.DataOverviewBaseDataController)
	projectRoute.GET("/GetDataOverviewProductData", StatementController.DataOverviewProductController)
	projectRoute.GET("/GetNetPowerPolicyData", StatementController.NetPowerPolicyDataController)
	projectRoute.GET("/AchievementMoon", StatementController.AchievementMoonController)
	projectRoute.GET("/GetTrafficMeasurementData", StatementController.TrafficMeasurementDataController)
	projectRoute.GET("/GetTSRMonthlyPerformanceData", StatementController.TSRMonthlyPerformanceDataController)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(middlewarejwt.JWT())

	return r
}
