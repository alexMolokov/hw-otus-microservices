// Package api GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package api

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
        "/api/v1/user": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Создание пользователя.",
                "parameters": [
                    {
                        "description": "Запрос на создание",
                        "name": "_",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserCreateRequest"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.UserCreateResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/internalhttp.ResponseErrors"
                        }
                    },
                    "500": {
                        "description": "Some error",
                        "schema": {
                            "$ref": "#/definitions/internalhttp.ResponseError"
                        }
                    }
                }
            }
        },
        "/api/v1/user/{id}": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Получение пользователя по ID",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Пользователь",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/internalhttp.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/internalhttp.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Some error",
                        "schema": {
                            "$ref": "#/definitions/internalhttp.ResponseError"
                        }
                    }
                }
            },
            "put": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Изменение пользователя.",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Запрос на изменение данных",
                        "name": "_",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UserUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/internalhttp.ResponseOk"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/internalhttp.ResponseErrors"
                        }
                    },
                    "500": {
                        "description": "Some error",
                        "schema": {
                            "$ref": "#/definitions/internalhttp.ResponseError"
                        }
                    }
                }
            },
            "delete": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "v1"
                ],
                "summary": "Удаление пользователя по ID",
                "parameters": [
                    {
                        "type": "integer",
                        "example": 1,
                        "description": "ID пользователя",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Ok",
                        "schema": {
                            "$ref": "#/definitions/internalhttp.ResponseOk"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/internalhttp.ResponseError"
                        }
                    },
                    "404": {
                        "description": "Not found",
                        "schema": {
                            "$ref": "#/definitions/internalhttp.ResponseError"
                        }
                    },
                    "500": {
                        "description": "Some error",
                        "schema": {
                            "$ref": "#/definitions/internalhttp.ResponseError"
                        }
                    }
                }
            }
        },
        "/health": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "system"
                ],
                "summary": "Проверка здоровья сервиса",
                "responses": {
                    "200": {
                        "description": "Сервис работает корректно",
                        "schema": {
                            "$ref": "#/definitions/model.StatusResponse"
                        }
                    }
                }
            }
        },
        "/ready": {
            "get": {
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "system"
                ],
                "summary": "Проверка сервиса принимать трафик",
                "responses": {
                    "200": {
                        "description": "Сервис может принимать трафик",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "503": {
                        "description": "Сервис не может принимать трафик",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "internalhttp.ResponseError": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean",
                    "example": true
                },
                "message": {
                    "type": "string",
                    "example": "Some error message"
                }
            }
        },
        "internalhttp.ResponseErrors": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean",
                    "example": true
                },
                "errors": {
                    "type": "array",
                    "items": {}
                }
            }
        },
        "internalhttp.ResponseOk": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean",
                    "example": false
                },
                "message": {
                    "type": "string",
                    "example": "Ok"
                }
            }
        },
        "model.StatusResponse": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "string",
                    "example": "OK"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "alex.molokov@yandex.ru"
                },
                "firstName": {
                    "type": "string",
                    "example": "Молоков"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "lastName": {
                    "type": "string",
                    "example": "Алексей"
                },
                "phone": {
                    "type": "string",
                    "example": "+79035431754"
                },
                "username": {
                    "type": "string",
                    "example": "alex.molokov"
                }
            }
        },
        "model.UserCreateRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "alex.molokov@yandex.ru"
                },
                "firstName": {
                    "type": "string",
                    "example": "Молоков"
                },
                "lastName": {
                    "type": "string",
                    "example": "Алексей"
                },
                "phone": {
                    "type": "string",
                    "example": "+79035431754"
                },
                "username": {
                    "type": "string",
                    "example": "alex.molokov"
                }
            }
        },
        "model.UserCreateResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "boolean",
                    "example": false
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "message": {
                    "type": "string",
                    "example": "OK"
                }
            }
        },
        "model.UserUpdateRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "alex.molokov@yandex.ru"
                },
                "firstName": {
                    "type": "string",
                    "example": "Молоков"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "lastName": {
                    "type": "string",
                    "example": "Алексей"
                },
                "phone": {
                    "type": "string",
                    "example": "+79035431754"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Microservices Otus Service (Курс микросервисы Otus)",
	Description:      "Описание API методов",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
