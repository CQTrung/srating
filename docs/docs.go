// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/auth/login": {
            "post": {
                "tags": [
                    "auth"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.LoginResponse"
                        }
                    }
                }
            }
        },
        "/auth/refresh": {
            "post": {
                "tags": [
                    "auth"
                ],
                "summary": "Refresh token",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.RefreshTokenRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.RefreshTokenResponse"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "tags": [
                    "auth"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/dashboard": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "dashboard"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/feedbacks": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "feedback"
                ],
                "summary": "Get all feedback",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "page",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "user_id",
                        "name": "user_id",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "level",
                        "name": "level",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "start_date",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "end_date",
                        "name": "end_date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "feedback"
                ],
                "summary": "Create feedback",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.Feedback"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/feedbacks/:id": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "feedback"
                ],
                "summary": "Get feedback by detail",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/feedbacks/search": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "feedback"
                ],
                "summary": "Search feedback",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/media": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "media"
                ],
                "summary": "Get all media",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "media"
                ],
                "summary": "Upload media",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get user detail",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "user"
                ],
                "summary": "Create user by admin",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CreateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/employees": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "user"
                ],
                "summary": "Get all employee",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "user"
                ],
                "summary": "Update employee",
                "parameters": [
                    {
                        "description": "payload",
                        "name": "payload",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UpdateUserRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/employees/:id": {
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "user"
                ],
                "summary": "Delete employee",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/users/status": {
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "user"
                ],
                "summary": "Change status",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.CreateUserRequest": {
            "type": "object",
            "properties": {
                "department_id": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "field": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/domain.Role"
                },
                "short_name": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/domain.Status"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.Department": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.User"
                    }
                }
            }
        },
        "domain.Feedback": {
            "type": "object",
            "required": [
                "level",
                "user_id"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "level": {
                    "$ref": "#/definitions/domain.Level"
                },
                "note": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "domain.Level": {
            "type": "integer",
            "enum": [
                1,
                2,
                3,
                4
            ],
            "x-enum-varnames": [
                "VeryGood",
                "Good",
                "Normal",
                "Bad"
            ]
        },
        "domain.LoginRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "is_remember_me": {
                    "type": "boolean"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.LoginResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                },
                "department": {
                    "$ref": "#/definitions/domain.Department"
                },
                "email": {
                    "type": "string"
                },
                "field": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "media": {
                    "$ref": "#/definitions/domain.Media"
                },
                "phone": {
                    "type": "string"
                },
                "refresh_token": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/domain.Role"
                },
                "short_name": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/domain.Status"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.Media": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "filename": {
                    "description": "Type     string ` + "`" + `gorm:\"column:type\" json:\"type\"` + "`" + `",
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "updated_at": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "domain.RefreshTokenRequest": {
            "type": "object",
            "properties": {
                "refresh_token": {
                    "type": "string"
                }
            }
        },
        "domain.RefreshTokenResponse": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "domain.Role": {
            "type": "string",
            "enum": [
                "employee",
                "admin",
                "manager"
            ],
            "x-enum-varnames": [
                "EmployeeRole",
                "AdminRole",
                "ManagerRole"
            ]
        },
        "domain.Status": {
            "type": "integer",
            "enum": [
                1,
                0
            ],
            "x-enum-varnames": [
                "ActiveStatus",
                "InActiveStatus"
            ]
        },
        "domain.UpdateUserRequest": {
            "type": "object",
            "properties": {
                "department_id": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "field": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/domain.Role"
                },
                "short_name": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/domain.Status"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "domain.User": {
            "type": "object",
            "required": [
                "department_id",
                "email",
                "field",
                "full_name",
                "password",
                "phone",
                "short_name",
                "username"
            ],
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "department": {
                    "$ref": "#/definitions/domain.Department"
                },
                "department_id": {
                    "type": "integer"
                },
                "email": {
                    "type": "string"
                },
                "feedbacks": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Feedback"
                    }
                },
                "field": {
                    "type": "string"
                },
                "full_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "media": {
                    "$ref": "#/definitions/domain.Media"
                },
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "$ref": "#/definitions/domain.Role"
                },
                "short_name": {
                    "type": "string"
                },
                "status": {
                    "$ref": "#/definitions/domain.Status"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "basic"
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
