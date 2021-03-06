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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Grupo GIT",
            "url": "http://www.grupogit.com/",
            "email": "contacto@grupogit.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/authenticate/": {
            "post": {
                "description": "get datos del usuario",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Authenticación de usuarios",
                "parameters": [
                    {
                        "description": "Necesita un objeto",
                        "name": "objeto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Login"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Authenticate"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/authenticatemovil/": {
            "post": {
                "description": "get datos del usuario en la aplicación movil",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Authenticación de usuarios en app movil",
                "parameters": [
                    {
                        "description": "Necesita un objeto",
                        "name": "objeto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginMovil"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.AuthenticateMovil"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/change+password/": {
            "post": {
                "description": "Recibe el id hasheado del solicitante de verificación de cuenta y valida si el usario ya esta verificado o no.s",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Verificación de Usaurios nuevos",
                "parameters": [
                    {
                        "description": "Necesita un objeto",
                        "name": "objeto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ChangePassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retorna un true",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/auth/recovery+password/": {
            "post": {
                "description": "Envia un email al correo del usuario solicitante",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Recuperacion de contraseñas",
                "parameters": [
                    {
                        "description": "Necesita un objeto",
                        "name": "objeto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RecoveryPassword"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retorna un true",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/language/create/": {
            "post": {
                "description": "Agrega un nuevo idioma.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "language"
                ],
                "summary": "Agregar Idioma.",
                "parameters": [
                    {
                        "description": "Necesita un objeto",
                        "name": "objeto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Language"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retorna el id del idioma creado",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/language/delete/:id": {
            "delete": {
                "description": "Elimina algun idioma Existente.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "language"
                ],
                "summary": "Elimina idioma existente.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Language ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retorna un true",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/language/languages/": {
            "get": {
                "description": "Retorna un Array de Idiomas.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "language"
                ],
                "summary": "Lista de los idiomas existentes.",
                "responses": {
                    "200": {
                        "description": "Retorna un array de idiomas",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Language"
                            }
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/language/update/:id": {
            "put": {
                "description": "Modifica el dato del idioma registrado.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "language"
                ],
                "summary": "Modificar datos del idioma.",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Language ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Necesita un objeto",
                        "name": "objeto",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Language"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Retorna un true",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "default": {
                        "description": "error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Authenticate": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "resourcedataid": {
                    "type": "integer",
                    "example": 1
                },
                "roleid": {
                    "type": "integer",
                    "example": 2
                },
                "rolename": {
                    "type": "string"
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "model.AuthenticateMovil": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "resourcedataid": {
                    "type": "integer",
                    "example": 1
                },
                "roleid": {
                    "type": "integer",
                    "example": 2
                },
                "status": {
                    "type": "boolean"
                }
            }
        },
        "model.ChangePassword": {
            "type": "object",
            "properties": {
                "hashuserid": {
                    "type": "string",
                    "example": "8b459f91bacc7157222c4c1894c7229ae2d58eec745084bfdf1aa87b65"
                },
                "newpassword": {
                    "type": "string",
                    "example": "ex@ample16"
                }
            }
        },
        "model.Language": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 0
                },
                "languagename": {
                    "type": "string",
                    "example": "Español"
                }
            }
        },
        "model.Login": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "example@grupogit.com"
                },
                "password": {
                    "type": "string",
                    "example": "example15"
                }
            }
        },
        "model.LoginMovil": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "example@grupogit.com"
                },
                "password": {
                    "type": "string",
                    "example": "example15"
                },
                "tokenmovil": {
                    "type": "string"
                }
            }
        },
        "model.RecoveryPassword": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "example@grupogit.com"
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
	Host:        "localhost:8080",
	BasePath:    "/v2",
	Schemes:     []string{},
	Title:       "Grupo GIT API",
	Description: "This is a api rest full.",
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
