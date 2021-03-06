// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/auth/addClient": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Add Client",
                "parameters": [
                    {
                        "description": "app_id",
                        "name": "app_id",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "app_secret",
                        "name": "app_secret",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "desc",
                        "name": "desc",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "app_name",
                        "name": "app_name",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/tags/import": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Import Image",
                "parameters": [
                    {
                        "type": "file",
                        "description": "Image File",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/auth": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "summary": "Get Auth",
                "parameters": [
                    {
                        "type": "string",
                        "description": "app_id",
                        "name": "app_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "app_secret",
                        "name": "app_secret",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/projectData/AchievementMoon": {
            "get": {
                "summary": "AchievementMoon  ???????????????",
                "responses": {
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/projectData/AchievementSummary": {
            "get": {
                "summary": "AchievementSummary  ???????????????",
                "responses": {
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/projectData/GetDataOverviewBaseData": {
            "get": {
                "summary": "GetDataOverviewBaseData",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/StatementControllerResponseSchema.BaseDataResponseSchema"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/projectData/GetDataOverviewProductData": {
            "get": {
                "summary": "GetDataOverviewProductData",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/StatementControllerResponseSchema.ProductDataResponseSchema"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/projectData/GetNetPowerPolicyData": {
            "get": {
                "summary": "GetNetPowerPolicyData",
                "parameters": [
                    {
                        "type": "number",
                        "description": "??????????????????",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "????????????",
                        "name": "current",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/StatementControllerResponseSchema.NetPowerPolicyDataResponseSchema"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/projectData/GetTSRGrade": {
            "post": {
                "summary": "GetTSRGrade",
                "parameters": [
                    {
                        "description": "??????????????????????????????",
                        "name": "gradeBody",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/TSRPortraitControllerRequestSchema.TSRGradeControllerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/TSRPortraitControllerResponseSchema.TSRGradeControllerResponseSchema"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/projectData/GetTSRMonthlyPerformanceData": {
            "get": {
                "summary": "GetTSRMonthlyPerformanceData",
                "parameters": [
                    {
                        "type": "number",
                        "description": "??????????????????",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "????????????",
                        "name": "current",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/StatementControllerResponseSchema.TSRMonthlyPerformanceResponseSchema"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/projectData/GetTSRTop": {
            "post": {
                "summary": "GetTSRTop",
                "parameters": [
                    {
                        "description": "????????????",
                        "name": "topBody",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/TSRPortraitControllerRequestSchema.TSRTopControllerRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/TSRPortraitControllerResponseSchema.TSRTopControllerResponseSchema"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        },
        "/projectData/GetTrafficMeasurementData": {
            "get": {
                "summary": "GetTrafficMeasurementData",
                "parameters": [
                    {
                        "type": "number",
                        "description": "??????????????????",
                        "name": "page_size",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "description": "????????????",
                        "name": "current",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/StatementControllerResponseSchema.NetPowerPolicyDataResponseSchema"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/app.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "StatementControllerResponseSchema.BaseDataResponseSchema": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "$ref": "#/definitions/StatementSchema.DataOverviewBaseDataSchema"
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "StatementControllerResponseSchema.NetPowerPolicyDataResponseSchema": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/StatementSchema.NetPowerPolicySchema"
                    }
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "StatementControllerResponseSchema.ProductDataResponseSchema": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/StatementSchema.DataOverviewProductDistribution"
                    }
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "StatementControllerResponseSchema.TSRMonthlyPerformanceResponseSchema": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/StatementSchema.TSRMonthlyPerformance"
                    }
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "StatementSchema.DataOverviewBaseDataSchema": {
            "type": "object",
            "properties": {
                "call_connect_rate": {
                    "description": "?????????",
                    "type": "number"
                },
                "call_time_mean": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "total_call_count": {
                    "description": "????????????",
                    "type": "integer"
                },
                "total_customer_count": {
                    "description": "????????????",
                    "type": "integer"
                }
            }
        },
        "StatementSchema.DataOverviewProductDistribution": {
            "type": "object",
            "properties": {
                "product_count": {
                    "description": "????????????",
                    "type": "integer"
                },
                "product_name": {
                    "description": "????????????",
                    "type": "string"
                }
            }
        },
        "StatementSchema.NetPowerPolicySchema": {
            "type": "object",
            "properties": {
                "activeMonth": {
                    "description": "?????????",
                    "type": "integer"
                },
                "annualizedPremium": {
                    "description": "????????????",
                    "type": "number"
                },
                "batchName": {
                    "description": "????????????",
                    "type": "string"
                },
                "insuredAmount": {
                    "description": "??????",
                    "type": "number"
                },
                "policyStatus": {
                    "description": "????????????",
                    "type": "string"
                },
                "productName": {
                    "description": "????????????",
                    "type": "string"
                },
                "tsrid": {
                    "description": "??????ID",
                    "type": "integer"
                },
                "underwritingTime": {
                    "description": "????????????",
                    "type": "string"
                }
            }
        },
        "StatementSchema.TSRMonthlyPerformance": {
            "type": "object",
            "properties": {
                "activeMonth": {
                    "description": "?????????",
                    "type": "number"
                },
                "callCount": {
                    "description": "????????????",
                    "type": "integer"
                },
                "connectCount": {
                    "description": "????????????",
                    "type": "integer"
                },
                "contactListCount": {
                    "description": "???????????????",
                    "type": "integer"
                },
                "firstConnectCount": {
                    "description": "???????????????",
                    "type": "integer"
                },
                "secondConnectCount": {
                    "description": "???????????????",
                    "type": "integer"
                },
                "tsrid": {
                    "description": "TSR????????????",
                    "type": "integer"
                },
                "tsrname": {
                    "description": "TSR??????",
                    "type": "string"
                }
            }
        },
        "TSRPortraitControllerRequestSchema.TSRGradeControllerRequest": {
            "type": "object",
            "properties": {
                "call_count_rate": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "call_count_standard": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "call_count_time_rate": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "call_count_time_standard": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "call_time_rate": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "call_time_standard": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "end_time": {
                    "description": "????????????",
                    "type": "string"
                },
                "phone_list": {
                    "description": "????????????",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "policy_premium_mean_rate": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "policy_premium_mean_standard": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "premium_grade_rate": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "premium_standard": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "start_time": {
                    "description": "????????????",
                    "type": "string"
                },
                "tsr_monthly_call_grade_rate": {
                    "description": "???????????????????????????",
                    "type": "number"
                },
                "tsr_monthly_call_grade_standard": {
                    "description": "???????????????????????????",
                    "type": "number"
                }
            }
        },
        "TSRPortraitControllerRequestSchema.TSRTopControllerRequest": {
            "type": "object",
            "properties": {
                "current": {
                    "description": "????????????",
                    "type": "integer"
                },
                "customer_age_bottom": {
                    "description": "??????????????????",
                    "type": "integer"
                },
                "customer_age_top": {
                    "description": "??????????????????",
                    "type": "integer"
                },
                "end_time": {
                    "description": "????????????",
                    "type": "string"
                },
                "name_list_type": {
                    "description": "????????????",
                    "type": "string"
                },
                "order_place": {
                    "description": "????????????",
                    "type": "string"
                },
                "page_size": {
                    "description": "????????????",
                    "type": "integer"
                },
                "phone_list": {
                    "description": "????????????",
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                },
                "premiums_five_star": {
                    "description": "?????????????????????????????????",
                    "type": "number"
                },
                "premiums_four_star": {
                    "description": "?????????????????????????????????",
                    "type": "number"
                },
                "premiums_one_star": {
                    "description": "?????????????????????????????????",
                    "type": "number"
                },
                "premiums_three_star": {
                    "description": "?????????????????????????????????",
                    "type": "number"
                },
                "premiums_two_star": {
                    "description": "?????????????????????????????????",
                    "type": "number"
                },
                "product_type": {
                    "description": "????????????",
                    "type": "string"
                },
                "quality_five_star": {
                    "description": "??????????????????????????????",
                    "type": "number"
                },
                "quality_four_star": {
                    "description": "??????????????????????????????",
                    "type": "number"
                },
                "quality_one_star": {
                    "description": "??????????????????????????????",
                    "type": "number"
                },
                "quality_three_star": {
                    "description": "??????????????????????????????",
                    "type": "number"
                },
                "quality_two_star": {
                    "description": "??????????????????????????????",
                    "type": "number"
                },
                "renewal_five_star": {
                    "description": "??????????????????????????????",
                    "type": "number"
                },
                "renewal_four_star": {
                    "description": "??????????????????????????????",
                    "type": "number"
                },
                "renewal_one_star": {
                    "description": "??????????????????????????????",
                    "type": "number"
                },
                "renewal_three_star": {
                    "description": "??????????????????????????????",
                    "type": "number"
                },
                "renewal_two_star": {
                    "description": "??????????????????????????????",
                    "type": "number"
                },
                "start_time": {
                    "description": "????????????",
                    "type": "string"
                }
            }
        },
        "TSRPortraitControllerResponseSchema.TSRGradeControllerResponseSchema": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/TSRPortraitSchema.TSRGradeDetail"
                    }
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "TSRPortraitControllerResponseSchema.TSRTopControllerResponseSchema": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/TSRPortraitSchema.TsrStarPageResponseSchema"
                    }
                },
                "msg": {
                    "type": "string"
                }
            }
        },
        "TSRPortraitSchema.TSRGradeDetail": {
            "type": "object",
            "properties": {
                "call_connect_count": {
                    "description": "??????????????????",
                    "type": "integer"
                },
                "call_count_day_mean": {
                    "description": "????????????",
                    "type": "number"
                },
                "call_count_grade": {
                    "description": "????????????",
                    "type": "number"
                },
                "call_count_time_grade": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "call_count_time_mean": {
                    "description": "????????????",
                    "type": "number"
                },
                "call_time_day_mean": {
                    "description": "????????????",
                    "type": "number"
                },
                "call_time_grade": {
                    "description": "????????????",
                    "type": "number"
                },
                "call_time_total": {
                    "description": "???????????????",
                    "type": "integer"
                },
                "name_list_call_count": {
                    "description": "???????????????",
                    "type": "integer"
                },
                "name_list_call_count_mean": {
                    "description": "??????????????????????????????",
                    "type": "number"
                },
                "name_list_call_grade": {
                    "description": "????????????????????????",
                    "type": "number"
                },
                "name_list_customer_count": {
                    "description": "????????????????????????",
                    "type": "integer"
                },
                "policy_count": {
                    "description": "?????????",
                    "type": "integer"
                },
                "policy_total_premium": {
                    "description": "????????????",
                    "type": "number"
                },
                "record_count": {
                    "description": "????????????(????????????)",
                    "type": "integer"
                },
                "tsr_grade": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "tsr_id": {
                    "description": "??????ID",
                    "type": "integer"
                },
                "tsr_name": {
                    "description": "????????????",
                    "type": "string"
                },
                "tsr_policy_grade": {
                    "description": "????????????????????????",
                    "type": "number"
                },
                "tsr_policy_premium_mean": {
                    "description": "????????????????????????",
                    "type": "number"
                },
                "tsr_policy_premium_mean_grade": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "tsr_skill_grade": {
                    "description": "??????????????????",
                    "type": "number"
                },
                "tsr_status_grade": {
                    "description": "??????????????????",
                    "type": "number"
                }
            }
        },
        "TSRPortraitSchema.TsrStarPageResponseSchema": {
            "type": "object",
            "properties": {
                "current": {
                    "description": "????????????",
                    "type": "integer"
                },
                "pageSize": {
                    "description": "????????????",
                    "type": "integer"
                },
                "total": {
                    "description": "????????????",
                    "type": "integer"
                },
                "tsrList": {
                    "description": "????????????",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/TSRPortraitSchema.TsrStarResponseSchema"
                    }
                }
            }
        },
        "TSRPortraitSchema.TsrStarResponseSchema": {
            "type": "object",
            "properties": {
                "Top": {
                    "description": "??????",
                    "type": "integer"
                },
                "policyStar": {
                    "description": "????????????",
                    "type": "integer"
                },
                "premiumsStar": {
                    "description": "????????????",
                    "type": "integer"
                },
                "qualifiedStar": {
                    "description": "????????????",
                    "type": "integer"
                },
                "totalStar": {
                    "description": "?????????",
                    "type": "integer"
                },
                "tsrId": {
                    "description": "??????ID",
                    "type": "integer"
                },
                "tsrName": {
                    "description": "????????????",
                    "type": "string"
                }
            }
        },
        "app.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {
                    "type": "object"
                },
                "msg": {
                    "type": "string"
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "Golang Gin API",
	Description: "An example of gin",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
