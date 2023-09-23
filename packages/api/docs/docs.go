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
        "/admin/v1/posts": {
            "get": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Paginate posts",
                "parameters": [
                    {
                        "type": "number",
                        "description": "Post number",
                        "name": "post_num",
                        "in": "query"
                    },
                    {
                        "enum": [
                            "open",
                            "open",
                            "closed",
                            "completed",
                            "canceled"
                        ],
                        "type": "string",
                        "description": "Post status",
                        "name": "status",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Text",
                        "name": "text",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Deadline",
                        "name": "deadline",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Delivery date",
                        "name": "delivery_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Seller ID",
                        "name": "seller_id",
                        "in": "query"
                    },
                    {
                        "type": "number",
                        "default": 0,
                        "description": "Page (0-based)",
                        "name": "page",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/admin.PaginationResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/admin.Object"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "posts": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/ent.Post"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "BearerToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "posts"
                ],
                "summary": "Create post",
                "parameters": [
                    {
                        "description": "Post body",
                        "name": "post",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.createPostForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/admin.PaginationResult"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "allOf": [
                                                {
                                                    "$ref": "#/definitions/admin.Object"
                                                },
                                                {
                                                    "type": "object",
                                                    "properties": {
                                                        "posts": {
                                                            "type": "array",
                                                            "items": {
                                                                "$ref": "#/definitions/ent.Post"
                                                            }
                                                        }
                                                    }
                                                }
                                            ]
                                        }
                                    }
                                }
                            ]
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/utils.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "admin.Object": {
            "type": "object"
        },
        "admin.PaginationResult": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object",
                    "properties": {
                        "count": {
                            "type": "integer"
                        },
                        "has_next_page": {
                            "type": "boolean"
                        }
                    }
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "admin.createPostForm": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "deadline": {
                    "type": "string"
                },
                "delivery_date": {
                    "type": "string"
                },
                "images": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Image"
                    }
                },
                "is_in_stock": {
                    "type": "boolean"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "properties": {
                            "name": {
                                "type": "string"
                            },
                            "price": {
                                "type": "number"
                            },
                            "stock": {
                                "type": "number"
                            }
                        }
                    }
                },
                "post_num": {
                    "type": "integer"
                },
                "seller_id": {
                    "type": "string"
                },
                "storage_type": {
                    "$ref": "#/definitions/post.StorageType"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "ent.Post": {
            "type": "object",
            "properties": {
                "body": {
                    "description": "Body holds the value of the \"body\" field.",
                    "type": "string"
                },
                "comment": {
                    "description": "Comment holds the value of the \"comment\" field.",
                    "type": "string"
                },
                "comment_count": {
                    "description": "CommentCount holds the value of the \"comment_count\" field.",
                    "type": "integer"
                },
                "created_at": {
                    "description": "CreatedAt holds the value of the \"created_at\" field.",
                    "type": "string"
                },
                "deadline": {
                    "description": "Deadline holds the value of the \"deadline\" field.",
                    "type": "string"
                },
                "delivered": {
                    "description": "Delivered holds the value of the \"delivered\" field.",
                    "type": "boolean"
                },
                "delivery_date": {
                    "description": "DeliveryDate holds the value of the \"delivery_date\" field.",
                    "type": "string"
                },
                "extra_fee": {
                    "description": "ExtraFee holds the value of the \"extra_fee\" field.",
                    "type": "number"
                },
                "extra_total": {
                    "description": "ExtraTotal holds the value of the \"extra_total\" field.",
                    "type": "number"
                },
                "id": {
                    "description": "ID of the ent.",
                    "type": "string"
                },
                "images": {
                    "description": "Images holds the value of the \"images\" field.",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/schema.Image"
                    }
                },
                "is_in_stock": {
                    "description": "IsInStock holds the value of the \"is_in_stock\" field.",
                    "type": "boolean"
                },
                "like_count": {
                    "description": "LikeCount holds the value of the \"like_count\" field.",
                    "type": "integer"
                },
                "normal_fee": {
                    "description": "NormalFee holds the value of the \"normal_fee\" field.",
                    "type": "number"
                },
                "normal_total": {
                    "description": "NormalTotal holds the value of the \"normal_total\" field.",
                    "type": "number"
                },
                "order_count": {
                    "description": "OrderCount holds the value of the \"order_count\" field.",
                    "type": "integer"
                },
                "post_num": {
                    "description": "PostNum holds the value of the \"post_num\" field.",
                    "type": "integer"
                },
                "seller_id": {
                    "description": "SellerID holds the value of the \"seller_id\" field.",
                    "type": "string"
                },
                "status": {
                    "description": "Status holds the value of the \"status\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/post.Status"
                        }
                    ]
                },
                "storage_type": {
                    "description": "StorageType holds the value of the \"storage_type\" field.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/post.StorageType"
                        }
                    ]
                },
                "title": {
                    "description": "Title holds the value of the \"title\" field.",
                    "type": "string"
                },
                "updated_at": {
                    "description": "UpdatedAt holds the value of the \"updated_at\" field.",
                    "type": "string"
                }
            }
        },
        "post.Status": {
            "type": "string",
            "enum": [
                "open",
                "open",
                "closed",
                "completed",
                "canceled"
            ],
            "x-enum-varnames": [
                "DefaultStatus",
                "StatusOpen",
                "StatusClosed",
                "StatusCompleted",
                "StatusCanceled"
            ]
        },
        "post.StorageType": {
            "type": "string",
            "enum": [
                "roomTemp",
                "farmGoods",
                "refrigerated",
                "frozen"
            ],
            "x-enum-varnames": [
                "StorageTypeRoomTemp",
                "StorageTypeFarmGoods",
                "StorageTypeRefrigerated",
                "StorageTypeFrozen"
            ]
        },
        "schema.Image": {
            "type": "object",
            "properties": {
                "lg": {
                    "type": "string"
                },
                "md": {
                    "type": "string"
                },
                "sm": {
                    "type": "string"
                }
            }
        },
        "utils.HTTPError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer",
                    "example": 400
                },
                "message": {
                    "type": "string",
                    "example": "status bad request"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerToken": {
            "description": "Bearer token",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "5.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Tuango API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}