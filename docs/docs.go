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
        "/data": {
            "get": {
                "description": "Get a list of data",
                "produces": [
                    "application/json"
                ],
                "summary": "Get data",
                "operationId": "get-data",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.Data"
                            }
                        }
                    }
                }
            }
        },
        "/thirtyconsumer": {
            "get": {
                "description": "Get a list of consumers and their usage that is almost 30% of total",
                "produces": [
                    "application/json"
                ],
                "summary": "Get consumers that adds up to around 30% to total consumption",
                "operationId": "get-tirthy-consumer",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.TopThirtyConsumer"
                            }
                        }
                    }
                }
            }
        },
        "/topconsumer": {
            "get": {
                "description": "Get a list of consumers and their usage for the last 10 minutes",
                "produces": [
                    "application/json"
                ],
                "summary": "Get top consumers",
                "operationId": "get-top-consumer",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.TopConsumer"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.Data": {
            "description": "Data structure for API",
            "type": "object",
            "properties": {
                "consumerId": {
                    "description": "id of consumer",
                    "type": "string"
                },
                "meterId": {
                    "description": "id of meter",
                    "type": "string"
                },
                "meterReading": {
                    "description": "the reading from the meter",
                    "type": "integer"
                },
                "timestamp": {
                    "description": "timestamp of data",
                    "type": "string"
                }
            }
        },
        "models.TopConsumer": {
            "description": "TopConsumer structure for API",
            "type": "object",
            "properties": {
                "consumerId": {
                    "type": "string"
                },
                "totalReading": {
                    "type": "integer"
                }
            }
        },
        "models.TopThirtyConsumer": {
            "description": "TopThirtyConsumer structure for API",
            "type": "object",
            "properties": {
                "consumers": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.TopConsumer"
                    }
                },
                "totalConsumption": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
