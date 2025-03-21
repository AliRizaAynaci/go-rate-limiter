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
        "/logs": {
            "get": {
                "description": "Get the last 100 log entries with optional filtering",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "logs"
                ],
                "summary": "Get logs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Filter by log level (DEBUG, INFO, WARN, ERROR)",
                        "name": "level",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by endpoint",
                        "name": "endpoint",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search in message content",
                        "name": "search",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.LogEntry"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.LogEntry": {
            "type": "object",
            "properties": {
                "endpoint": {
                    "type": "string",
                    "example": "/api/protected-endpoint"
                },
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "level": {
                    "type": "string",
                    "example": "INFO"
                },
                "message": {
                    "type": "string",
                    "example": "Request processed successfully"
                },
                "timestamp": {
                    "type": "string",
                    "example": "2024-03-20T15:04:05Z"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Rate Limiter API",
	Description:      "A REST API for rate limiting with Redis and SQLite",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
