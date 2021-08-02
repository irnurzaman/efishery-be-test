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
        "/auth/login": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "summary": "Login user",
                "operationId": "LoginUser",
                "parameters": [
                    {
                        "description": "all fields are mandatory",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqLoginUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"token\": \"{token}\"}",
                        "schema": {
                            "$ref": "#/definitions/models.RespLoginUser"
                        }
                    },
                    "400": {
                        "description": "{\"remark\": \"Invalid parse body request to JSON\"}",
                        "schema": {
                            "$ref": "#/definitions/models.RespError"
                        }
                    },
                    "401": {
                        "description": "{\"remark\": \"Invalid authentication for phone {phone}\"}",
                        "schema": {
                            "$ref": "#/definitions/models.RespError"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "tags": [
                    "Auth"
                ],
                "summary": "Register new user",
                "operationId": "RegisterUser",
                "parameters": [
                    {
                        "description": "all fields are mandatory",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ReqRegisterUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\"password\": \"{password}\"}",
                        "schema": {
                            "$ref": "#/definitions/models.RespRegisterUser"
                        }
                    },
                    "400": {
                        "description": "{\"remark\": \"Invalid parse body request to JSON\"}",
                        "schema": {
                            "$ref": "#/definitions/models.RespError"
                        }
                    },
                    "422": {
                        "description": "{\"remark\": \"Phone number has been registered\"}",
                        "schema": {
                            "$ref": "#/definitions/models.RespError"
                        }
                    }
                }
            }
        },
        "/auth/verify": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Verify and extract JWT",
                "operationId": "VerifyToken",
                "responses": {
                    "200": {
                        "description": "{\"claims\": Model}",
                        "schema": {
                            "$ref": "#/definitions/models.RespVerifyToken"
                        }
                    },
                    "400": {
                        "description": "{\"remark\": \"Invalid parse body request to JSON\"}",
                        "schema": {
                            "$ref": "#/definitions/models.RespError"
                        }
                    },
                    "401": {
                        "description": "{\"remark\": \"Invalid token verification\"}",
                        "schema": {
                            "$ref": "#/definitions/models.RespError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ReqLoginUser": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "models.ReqRegisterUser": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                }
            }
        },
        "models.RespError": {
            "type": "object",
            "properties": {
                "remark": {
                    "type": "string"
                }
            }
        },
        "models.RespLoginUser": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "models.RespRegisterUser": {
            "type": "object",
            "properties": {
                "password": {
                    "type": "string"
                }
            }
        },
        "models.RespVerifyToken": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "timestamp": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
	Title:       "Auth Service API",
	Description: "This is a simple auth service for generating and verifying JWT.",
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
