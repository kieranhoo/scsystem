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
        "/auth/signin": {
            "post": {
                "description": "Sign in account for admin.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "Sign In",
                        "name": "sign_in",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SignInResponse"
                        }
                    }
                }
            }
        },
        "/auth/signup": {
            "post": {
                "description": "Register account for admin.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "parameters": [
                    {
                        "description": "Sign Up",
                        "name": "sign_up",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.SignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SignUpResponse"
                        }
                    }
                }
            }
        },
        "/common/healthcheck": {
            "get": {
                "description": "health check api server.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "common"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/room": {
            "get": {
                "description": "Get Room information from database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.DataResponse"
                        }
                    }
                }
            }
        },
        "/room/activity": {
            "get": {
                "description": "Get User Activity for checkin.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "student id",
                        "name": "sid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "room id",
                        "name": "room",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.DataResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Save activity type in/out.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "parameters": [
                    {
                        "description": "CheckInRequest",
                        "name": "activity_param",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CheckInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    }
                }
            }
        },
        "/room/history": {
            "get": {
                "description": "Get History.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "room id",
                        "name": "room_id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "date only",
                        "name": "date",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.HistoryDataResponse"
                        }
                    }
                }
            }
        },
        "/room/register": {
            "post": {
                "description": "Register Room.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "room"
                ],
                "parameters": [
                    {
                        "description": "RegisterRoom",
                        "name": "sign_in",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.RegistrationRoomRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    }
                }
            }
        },
        "/stat/chart": {
            "get": {
                "description": "Get Chart Data IN/OUT.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stat"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "room id",
                        "name": "room_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.ChartMetadata"
                        }
                    }
                }
            }
        },
        "/stat/rooms": {
            "get": {
                "description": "Get Room Data IN/OUT.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "stat"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.RoomStat"
                        }
                    }
                }
            }
        },
        "/user": {
            "get": {
                "description": "Get information by student id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "student id",
                        "name": "sid",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.SignInResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.ChartData": {
            "type": "object",
            "properties": {
                "label": {
                    "description": "day",
                    "type": "string"
                },
                "value": {
                    "description": "counter in/out",
                    "type": "integer"
                }
            }
        },
        "domain.ChartDataValue": {
            "type": "object",
            "properties": {
                "in": {
                    "$ref": "#/definitions/domain.ChartData"
                },
                "out": {
                    "$ref": "#/definitions/domain.ChartData"
                }
            }
        },
        "domain.ChartMetadata": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.ChartDataValue"
                    }
                },
                "month": {
                    "type": "string"
                },
                "room_id": {
                    "type": "string"
                },
                "room_name": {
                    "type": "string"
                },
                "year": {
                    "type": "string"
                }
            }
        },
        "domain.CheckInRequest": {
            "type": "object",
            "required": [
                "activity_type",
                "admin_id",
                "registration_id"
            ],
            "properties": {
                "activity_type": {
                    "type": "string"
                },
                "admin_id": {
                    "type": "string"
                },
                "registration_id": {
                    "type": "string"
                }
            }
        },
        "domain.DataResponse": {
            "type": "object",
            "required": [
                "data",
                "success"
            ],
            "properties": {
                "data": {},
                "success": {
                    "type": "boolean"
                }
            }
        },
        "domain.HistoryData": {
            "type": "object",
            "properties": {
                "activity_type": {
                    "type": "string"
                },
                "admin_id": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "end_day": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "type": "string"
                },
                "office_name": {
                    "type": "string"
                },
                "org_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "registration_time": {
                    "type": "string"
                },
                "room_name": {
                    "type": "string"
                },
                "start_day": {
                    "type": "string"
                },
                "supervisor": {
                    "type": "string"
                },
                "time": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "domain.HistoryDataResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.HistoryData"
                    }
                },
                "room_name": {
                    "type": "string"
                },
                "total_in": {
                    "type": "integer"
                },
                "total_out": {
                    "type": "integer"
                }
            }
        },
        "domain.RegistrationRoomRequest": {
            "type": "object",
            "required": [
                "email",
                "end_day",
                "first_name",
                "last_name",
                "phone_number",
                "registration_time",
                "room_id",
                "start_day",
                "student_id",
                "supervisor"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "end_day": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                },
                "registration_time": {
                    "type": "string"
                },
                "room_id": {
                    "type": "string"
                },
                "start_day": {
                    "type": "string"
                },
                "student_id": {
                    "type": "string"
                },
                "supervisor": {
                    "type": "string"
                }
            }
        },
        "domain.Response": {
            "type": "object",
            "required": [
                "msg",
                "success"
            ],
            "properties": {
                "msg": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        },
        "domain.RoomStat": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.RoomStatData"
                    }
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "domain.RoomStatData": {
            "type": "object",
            "properties": {
                "in": {
                    "type": "integer"
                },
                "out": {
                    "type": "integer"
                },
                "room_id": {
                    "type": "string"
                },
                "room_name": {
                    "type": "string"
                },
                "total": {
                    "type": "integer"
                }
            }
        },
        "domain.SignInRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "domain.SignInResponse": {
            "type": "object",
            "required": [
                "email",
                "success",
                "token"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "domain.SignUpRequest": {
            "type": "object",
            "required": [
                "email",
                "first_name",
                "id",
                "last_name",
                "password",
                "phone_number"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "last_name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "phone_number": {
                    "type": "string"
                }
            }
        },
        "domain.SignUpResponse": {
            "type": "object",
            "required": [
                "msg",
                "success"
            ],
            "properties": {
                "msg": {
                    "type": "string"
                },
                "success": {
                    "type": "boolean"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Student Checkin System",
	Description:      "This is a documentation for the Student Checkin System API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
