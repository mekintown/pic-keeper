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
        "license": {
            "name": "Apache 2.0"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/authen/v1/login": {
            "post": {
                "description": "Login",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "authen"
                ],
                "summary": "Login via email and password",
                "parameters": [
                    {
                        "description": "email and password of the user",
                        "name": "Credentials",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginCredentials"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The token will be returned inside the data field",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONSuccessResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "400": {
                        "description": "Incorrect input",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "404": {
                        "description": "User does not exist",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Unhandled internal server error",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/model.JSONErrorResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "error": {
                                            "type": "object"
                                        },
                                        "status": {
                                            "type": "string"
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
        "model.JSONErrorResult": {
            "type": "object",
            "properties": {
                "error": {},
                "status": {
                    "type": "string",
                    "example": "failed"
                }
            }
        },
        "model.JSONSuccessResult": {
            "type": "object",
            "properties": {
                "data": {},
                "status": {
                    "type": "string",
                    "example": "success"
                }
            }
        },
        "model.LoginCredentials": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@mail.com"
                },
                "password": {
                    "type": "string",
                    "example": "abc123"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Pic-keeper APIs",
	Description:      "This is the back-end documentation of the pic-keeper project",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}