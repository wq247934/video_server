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
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/user_login/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户接口"
                ],
                "summary": "用户登录",
                "operationId": "/user_login/login",
                "parameters": [
                    {
                        "description": "body",
                        "name": "polygon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.UserLoginInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/dto.UserLoginOutput"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/user_login/logout": {
            "get": {
                "description": "用户退出登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户接口"
                ],
                "summary": "用户退出",
                "operationId": "/user_login/logout",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/video/upload": {
            "post": {
                "description": "用户上传视频",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "视频接口"
                ],
                "summary": "上传视频",
                "operationId": "/video/upload",
                "parameters": [
                    {
                        "description": "body",
                        "name": "polygon",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.VideoUploadInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/video/videos": {
            "get": {
                "description": "视频列表",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "视频接口"
                ],
                "summary": "视频列表",
                "operationId": "/video/videos",
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/middleware.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "array",
                                            "items": {
                                                "$ref": "#/definitions/dto.VideoIndexOutput"
                                            }
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.UserLoginInput": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "密码",
                    "type": "string",
                    "example": "123456"
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "dto.UserLoginOutput": {
            "type": "object",
            "properties": {
                "token": {
                    "description": "token",
                    "type": "string",
                    "example": "admin"
                },
                "username": {
                    "description": "用户名",
                    "type": "string",
                    "example": "admin"
                }
            }
        },
        "dto.VideoIndexOutput": {
            "type": "object",
            "properties": {
                "author_name": {
                    "description": "视频作者名称",
                    "type": "string",
                    "example": "张三"
                },
                "title": {
                    "description": "视频标题",
                    "type": "string",
                    "example": "标题1"
                },
                "video_id": {
                    "description": "视频ID",
                    "type": "integer",
                    "example": 1
                }
            }
        },
        "dto.VideoUploadInput": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "introduction": {
                    "description": "视频简介",
                    "type": "string",
                    "example": "视频简介..."
                },
                "public": {
                    "type": "boolean",
                    "example": true
                },
                "title": {
                    "description": "视频标题",
                    "type": "string",
                    "example": "标题1"
                }
            }
        },
        "middleware.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "errmsg": {
                    "type": "string"
                },
                "errno": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}