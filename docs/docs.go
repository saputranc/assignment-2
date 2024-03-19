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
        "/orders": {
            "get": {
                "description": "Get Order with Item Data",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "operationId": "get-orders-with-items",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/GetOrdersResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create Order Data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "operationId": "create-new-order",
                "parameters": [
                    {
                        "description": "request body json",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewOrderRequestDto"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/dto.NewOrderResponseDto"
                        }
                    }
                }
            }
        },
        "/orders/{orderId}": {
            "put": {
                "description": "Update Order Data By Id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "operationId": "update-order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order's id",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "request body json",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.NewOrderRequestDto"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewOrderResponseDto"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete Order Data By Id (Tidak berjalan hanya tampilan)",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "operationId": "delete-order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "order's id",
                        "name": "orderId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.NewOrderResponseDto"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "GetOrdersResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.OrderWithItems"
                    }
                },
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "dto.GetItemResponseDto": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "itemCode": {
                    "type": "string"
                },
                "itemId": {
                    "type": "integer"
                },
                "orderId": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "dto.NewItemRequestDto": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string",
                    "example": "BMW"
                },
                "itemCode": {
                    "type": "string",
                    "example": "889"
                },
                "quantity": {
                    "type": "integer",
                    "example": 13
                }
            }
        },
        "dto.NewOrderRequestDto": {
            "type": "object",
            "properties": {
                "customerName": {
                    "type": "string",
                    "example": "John Doe"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.NewItemRequestDto"
                    }
                },
                "orderedAt": {
                    "type": "string",
                    "example": "2023-07-10T21:21:46+00:00"
                }
            }
        },
        "dto.NewOrderResponseDto": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "success"
                },
                "statusCode": {
                    "type": "integer",
                    "example": 200
                }
            }
        },
        "dto.OrderWithItems": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "customerName": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.GetItemResponseDto"
                    }
                },
                "orderId": {
                    "type": "integer"
                },
                "orderedAt": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
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
