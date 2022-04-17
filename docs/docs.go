// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://localhost:8000/",
        "contact": {
            "name": "yangguojie"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/save_user_info": {
            "post": {
                "description": "更新用户信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "description": "用户登录参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/table.UserInfo"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/table.UserInfo"
                        }
                    }
                }
            }
        },
        "/api/v1/search_user_info": {
            "get": {
                "description": "查询用户信息!!",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "查询用户信息!",
                "parameters": [
                    {
                        "type": "string",
                        "name": "uid",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Bearer 用户令牌",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/table.UserInfo"
                        }
                    }
                }
            }
        },
        "/user/login": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "用户登录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户登录参数",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/table.UserLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/table.loginInfo"
                        }
                    }
                }
            }
        },
        "/user/refresh_token": {
            "get": {
                "description": "刷新用户token!!",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "刷新用户token!",
                "parameters": [
                    {
                        "type": "string",
                        "name": "username",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/table.RefreshTokenRes"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "table.RefreshTokenRes": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "table.UserInfo": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "table.UserLogin": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "table.loginInfo": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                },
                "u_id": {
                    "type": "integer"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "毕设后端接口文档",
	Description:      "毕设后端接口文档 部分接口",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
