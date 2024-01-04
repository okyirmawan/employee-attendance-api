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
        "/attendance/checkin": {
            "post": {
                "description": "Record check-in for attendance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Check-in for attendance",
                "parameters": [
                    {
                        "description": "Check In",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto_request.CheckInDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    }
                }
            }
        },
        "/attendance/checkout/{id}": {
            "put": {
                "description": "Record check-out for attendance",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Check-out for attendance",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Attendance ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Check Out",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto_request.CheckOutDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    }
                }
            }
        },
        "/attendance/history/{employee_id}": {
            "put": {
                "description": "Retrieve a list of all attendance.",
                "produces": [
                    "application/json"
                ],
                "summary": "Find Attendance by Employee ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Employee ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success response",
                        "schema": {
                            "$ref": "#/definitions/handler.SuccessResp"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    }
                }
            }
        },
        "/employees": {
            "get": {
                "description": "Retrieve a list of all employees.",
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieve all employees",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.EmployeeDTO"
                            }
                        }
                    },
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new employee with the provided data.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new employee",
                "parameters": [
                    {
                        "description": "Employee data",
                        "name": "employee",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.EmployeeDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.EmployeeDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    }
                }
            }
        },
        "/employees/{id}": {
            "put": {
                "description": "Update employee details by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Update Employee",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the employee",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Employee details to be updated",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.EmployeeDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.EmployeeDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete employee by ID",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete Employee",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the employee",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.EmployeeDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    }
                }
            }
        },
        "/employees/{nip}": {
            "get": {
                "description": "Get employee details by NIP",
                "produces": [
                    "application/json"
                ],
                "summary": "Find Employee by NIP",
                "parameters": [
                    {
                        "type": "string",
                        "description": "NIP of the employee",
                        "name": "nip",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.EmployeeDTO"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ErrorResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.EmployeeDTO": {
            "type": "object",
            "required": [
                "date_of_birth",
                "gender",
                "name",
                "nip"
            ],
            "properties": {
                "date_of_birth": {
                    "type": "string"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "nip": {
                    "type": "string"
                }
            }
        },
        "dto_request.CheckInDTO": {
            "type": "object",
            "required": [
                "employee_id",
                "location_in"
            ],
            "properties": {
                "employee_id": {
                    "type": "integer"
                },
                "location_in": {
                    "type": "string"
                }
            }
        },
        "dto_request.CheckOutDTO": {
            "type": "object",
            "required": [
                "employee_id",
                "location_out"
            ],
            "properties": {
                "employee_id": {
                    "type": "integer"
                },
                "location_out": {
                    "type": "string"
                }
            }
        },
        "handler.ErrorResp": {
            "type": "object",
            "properties": {
                "errorCode": {
                    "type": "integer"
                },
                "messages": {}
            }
        },
        "handler.SuccessResp": {
            "type": "object",
            "properties": {
                "data": {},
                "success": {
                    "type": "boolean"
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
