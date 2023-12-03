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
        "/confirm": {
            "post": {
                "description": "Confirmation of the registered user, for further use in the system",
                "produces": [
                    "application/json"
                ],
                "summary": "Confirm a registered user",
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/login": {
            "post": {
                "description": "Get login and password to get jwt token for further work in the system",
                "produces": [
                    "application/json"
                ],
                "summary": "Gets data to login in to the system",
                "parameters": [
                    {
                        "description": "Login",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.ReqLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.RespLogin"
                        }
                    }
                }
            }
        },
        "/logout": {
            "delete": {
                "description": "deletes the token from the database, thus ceasing to serve it",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete JWT token from database",
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
        "/refresh": {
            "put": {
                "description": "request with refreshToken, obtain new token for continued system access and operation",
                "produces": [
                    "application/json"
                ],
                "summary": "Gets refreshToken and gives new Token in to work in system",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/http.RespLogin"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Registration a new User with Confirmation",
                "produces": [
                    "application/json"
                ],
                "summary": "UserRegister",
                "parameters": [
                    {
                        "description": "Register",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.UserRegister"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "http.ReqLogin": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "http.RespLogin": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string"
                },
                "refreshToken": {
                    "type": "string"
                }
            }
        },
        "http.UserRegister": {
            "type": "object",
            "properties": {
                "FirstName": {
                    "type": "string"
                },
                "LastName": {
                    "type": "string"
                },
                "Login": {
                    "type": "string"
                },
                "Password": {
                    "type": "string"
                },
                "Phone": {
                    "type": "string"
                },
                "isConfirmed": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/auth",
	Schemes:          []string{},
	Title:            "Auth Service API",
	Description:      "Auth service API in Go using Gin Framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}