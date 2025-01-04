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
        "/api/v1/flyer/{id}": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Generates new flyer background images using AI with specified color palette",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "flyer"
                ],
                "summary": "Generate new flyer backgrounds",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Generation parameters",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.CreateRequest"
                        }
                    }
                ],
                "responses": {
                    "202": {
                        "description": "Successfully queued generation",
                        "schema": {
                            "$ref": "#/definitions/model.CreateResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request parameters",
                        "schema": {
                            "$ref": "#/definitions/model.GenerationErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized access",
                        "schema": {
                            "$ref": "#/definitions/model.GenerationErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/model.GenerationErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/flyer/{id}/{path}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieves a previously generated flyer background image by file path",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "image/png"
                ],
                "tags": [
                    "flyer"
                ],
                "summary": "Get a generated flyer background",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Image path",
                        "name": "path",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Image data",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "401": {
                        "description": "Unauthorized access",
                        "schema": {
                            "$ref": "#/definitions/model.GenerationErrorResponse"
                        }
                    },
                    "404": {
                        "description": "Image not found",
                        "schema": {
                            "$ref": "#/definitions/model.GenerationErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "$ref": "#/definitions/model.GenerationErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/login": {
            "post": {
                "description": "Authenticate user and return JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "Login credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successfully logged in",
                        "schema": {
                            "$ref": "#/definitions/model.LoginResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/model.UserErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Invalid credentials",
                        "schema": {
                            "$ref": "#/definitions/model.UserErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/register": {
            "post": {
                "description": "Register a new user with username, password and email",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "User registration details",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully registered",
                        "schema": {
                            "$ref": "#/definitions/model.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body or username taken",
                        "schema": {
                            "$ref": "#/definitions/model.UserErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get detailed information about a user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get user details",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User details",
                        "schema": {
                            "$ref": "#/definitions/model.UserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized access",
                        "schema": {
                            "$ref": "#/definitions/model.UserErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/model.UserErrorResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/model.UserErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Delete a user account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Delete user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User deleted",
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUserResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized access",
                        "schema": {
                            "$ref": "#/definitions/model.UserErrorResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/model.UserErrorResponse"
                        }
                    },
                    "404": {
                        "description": "User not found",
                        "schema": {
                            "$ref": "#/definitions/model.UserErrorResponse"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}/api-keys": {
            "put": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Update user's API keys for text and image generation services",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Add/update API keys (idempotent)",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "API keys",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateAPIKeysRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "API keys updated",
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUserResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/model.UserErrorResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized access",
                        "schema": {
                            "$ref": "#/definitions/model.UserErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "imagegen.ProviderType": {
            "type": "string",
            "enum": [
                "openai"
            ],
            "x-enum-varnames": [
                "OpenAI"
            ]
        },
        "model.CreateRequest": {
            "description": "Parameters for flyer background generation",
            "type": "object",
            "required": [
                "image_model_provider",
                "text_model_provider"
            ],
            "properties": {
                "color_palette": {
                    "description": "Color palette to use for the background\nExample: \"metalic gray and emerald green\"",
                    "type": "string"
                },
                "image_model_provider": {
                    "description": "Image generation model provider (e.g. \"openai\")\nRequired: true\nExample: \"openai\"",
                    "allOf": [
                        {
                            "$ref": "#/definitions/imagegen.ProviderType"
                        }
                    ]
                },
                "text_model_provider": {
                    "description": "Text generation model provider (e.g. \"openai\", \"anthropic\", \"google-vertex\")\nRequired: true\nExample: \"openai\"",
                    "allOf": [
                        {
                            "$ref": "#/definitions/textgen.ProviderType"
                        }
                    ]
                }
            }
        },
        "model.CreateResponse": {
            "description": "Response containing paths to generated backgrounds",
            "type": "object",
            "properties": {
                "background_paths": {
                    "description": "Array of file paths to the generated background images\nExample: [\"/images/1234567890/image.png\"]",
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.GenerationErrorResponse": {
            "description": "Error response from the API",
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "model.LoginRequest": {
            "description": "Login request payload",
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "Password hash\nExample: \"$2a$10$...\"",
                    "type": "string"
                },
                "username": {
                    "description": "Username for authentication\nExample: \"john_doe\"",
                    "type": "string"
                }
            }
        },
        "model.LoginResponse": {
            "description": "Login response",
            "type": "object",
            "properties": {
                "id": {
                    "description": "User's unique identifier\nExample: \"123e4567-e89b-12d3-a456-426614174000\"",
                    "type": "string"
                },
                "token": {
                    "description": "JWT token for authentication\nExample: \"eyJhbGciOiJIUzI1NiIs...\"",
                    "type": "string"
                }
            }
        },
        "model.RegisterRequest": {
            "description": "Registration request payload",
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "description": "Email address\nExample: \"john@example.com\"",
                    "type": "string"
                },
                "password": {
                    "description": "Password hash\nExample: \"$2a$10$...\"",
                    "type": "string",
                    "minLength": 8
                },
                "username": {
                    "description": "Desired username\nExample: \"john_doe\"",
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3
                }
            }
        },
        "model.RegisterResponse": {
            "description": "Registration request payload",
            "type": "object",
            "properties": {
                "token": {
                    "description": "Token for all endpoints required auth",
                    "type": "string"
                },
                "user": {
                    "description": "Basic Info regarding User",
                    "allOf": [
                        {
                            "$ref": "#/definitions/model.UserRegisterationInfo"
                        }
                    ]
                }
            }
        },
        "model.UpdateAPIKeysRequest": {
            "description": "Request to add/update API keys",
            "type": "object",
            "required": [
                "image_api_key",
                "image_provider",
                "text_api_key",
                "text_provider"
            ],
            "properties": {
                "image_api_key": {
                    "description": "API key for image generation service\nExample: \"sk-...\"",
                    "type": "string"
                },
                "image_provider": {
                    "description": "Image generation service provider\nExample: \"openai\"",
                    "type": "string"
                },
                "text_api_key": {
                    "description": "API key for text generation service\nExample: \"sk-...\"",
                    "type": "string"
                },
                "text_provider": {
                    "description": "Text generation service provider\nExample: \"anthropic\"",
                    "type": "string"
                }
            }
        },
        "model.UpdateUserResponse": {
            "description": "Response for updating user, such as deleting a user",
            "type": "object",
            "properties": {
                "message": {
                    "description": "Message regarding successful operation",
                    "type": "string"
                }
            }
        },
        "model.UserErrorResponse": {
            "description": "Error response from the API",
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "model.UserRegisterationInfo": {
            "description": "Basic Info regarding User",
            "type": "object",
            "properties": {
                "email": {
                    "description": "Email address\nExample: \"john@example.com\"",
                    "type": "string"
                },
                "id": {
                    "description": "user's UUID\nExample: \"john_doe\"",
                    "type": "string"
                },
                "username": {
                    "description": "Desired username\nExample: \"john_doe\"",
                    "type": "string"
                }
            }
        },
        "model.UserResponse": {
            "description": "Response for getting user, everything but password and api keys",
            "type": "object",
            "properties": {
                "active_status": {
                    "description": "Whether the account is active\nExample: true",
                    "type": "boolean"
                },
                "created_at": {
                    "description": "Account creation timestamp\nExample: \"2024-01-04T12:00:00Z\"",
                    "type": "string"
                },
                "email": {
                    "description": "User's email address\nExample: \"john@example.com\"",
                    "type": "string"
                },
                "id": {
                    "description": "Unique identifier for the user\nExample: \"123e4567-e89b-12d3-a456-426614174000\"",
                    "type": "string"
                },
                "is_admin": {
                    "description": "Whether the user has admin privileges\nExample: false",
                    "type": "boolean"
                },
                "last_login": {
                    "description": "Last login timestamp\nExample: \"2024-01-04T12:00:00Z\"",
                    "type": "string"
                },
                "updated_at": {
                    "description": "Last update timestamp\nExample: \"2024-01-04T12:00:00Z\"",
                    "type": "string"
                },
                "username": {
                    "description": "Username for login\nExample: \"john_doe\"",
                    "type": "string"
                }
            }
        },
        "textgen.ProviderType": {
            "type": "string",
            "enum": [
                "openai",
                "anthropic",
                "google-vertex"
            ],
            "x-enum-varnames": [
                "OpenAI",
                "Anthropic",
                "GoogleVertex"
            ]
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
