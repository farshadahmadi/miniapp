// Code generated by go-swagger; DO NOT EDIT.

package miniapp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "api for mini application to investigate split sdk synchronization problem",
    "title": "API",
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/v1",
  "paths": {
    "/feature/{feature}": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "feature"
        ],
        "summary": "For requesting feature flag data.",
        "operationId": "featureFlag",
        "parameters": [
          {
            "type": "string",
            "description": "feature id.",
            "name": "feature",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/FeatureFlagRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "History request success",
            "schema": {
              "$ref": "#/definitions/FeatureFlagResponse"
            }
          },
          "401": {
            "description": "401 Unauthorized",
            "schema": {
              "$ref": "#/definitions/UnauthorizedResponse"
            }
          },
          "500": {
            "description": "500 Internal server error",
            "schema": {
              "$ref": "#/definitions/InternalServerErrorResponse"
            }
          },
          "503": {
            "description": "503 Service Unavailable",
            "schema": {
              "$ref": "#/definitions/ServiceUnavailableResponse"
            }
          },
          "504": {
            "description": "504 Gateway Timeout",
            "schema": {
              "$ref": "#/definitions/GatewayTimeoutResponse"
            }
          },
          "default": {
            "description": "Default response",
            "schema": {
              "$ref": "#/definitions/DefaultResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "DefaultResponse": {
      "description": "Default response",
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "FeatureFlagRequest": {
      "type": "object",
      "required": [
        "trafficId"
      ],
      "properties": {
        "attributes": {
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        },
        "trafficId": {
          "$ref": "#/definitions/TrafficId"
        }
      }
    },
    "FeatureFlagResponse": {
      "type": "object",
      "properties": {
        "flag": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "GatewayTimeoutResponse": {
      "description": "504 Gateway Timeout - The server cannot handle the request. Generally, this is a temporary state",
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "enum": [
            504
          ]
        },
        "message": {
          "type": "string"
        }
      }
    },
    "InternalServerErrorResponse": {
      "description": "500 Internal server error",
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "enum": [
            500
          ]
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ServiceUnavailableResponse": {
      "description": "503 Service unavailable - The server cannot handle the request (because it is overloaded or down for maintenance). Generally, this is a temporary state",
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "enum": [
            503
          ]
        },
        "message": {
          "type": "string"
        }
      }
    },
    "TrafficId": {
      "type": "string",
      "pattern": "^[\\S]{1,}$"
    },
    "UnauthorizedResponse": {
      "description": "401 Unauthorized",
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "enum": [
            401
          ]
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Feature-flag service related endpoints.",
      "name": "feature"
    }
  ]
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "api for mini application to investigate split sdk synchronization problem",
    "title": "API",
    "version": "1.0.0"
  },
  "host": "localhost",
  "basePath": "/v1",
  "paths": {
    "/feature/{feature}": {
      "post": {
        "consumes": [
          "application/json"
        ],
        "produces": [
          "application/json"
        ],
        "tags": [
          "feature"
        ],
        "summary": "For requesting feature flag data.",
        "operationId": "featureFlag",
        "parameters": [
          {
            "type": "string",
            "description": "feature id.",
            "name": "feature",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/FeatureFlagRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "History request success",
            "schema": {
              "$ref": "#/definitions/FeatureFlagResponse"
            }
          },
          "401": {
            "description": "401 Unauthorized",
            "schema": {
              "$ref": "#/definitions/UnauthorizedResponse"
            }
          },
          "500": {
            "description": "500 Internal server error",
            "schema": {
              "$ref": "#/definitions/InternalServerErrorResponse"
            }
          },
          "503": {
            "description": "503 Service Unavailable",
            "schema": {
              "$ref": "#/definitions/ServiceUnavailableResponse"
            }
          },
          "504": {
            "description": "504 Gateway Timeout",
            "schema": {
              "$ref": "#/definitions/GatewayTimeoutResponse"
            }
          },
          "default": {
            "description": "Default response",
            "schema": {
              "$ref": "#/definitions/DefaultResponse"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "DefaultResponse": {
      "description": "Default response",
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "FeatureFlagRequest": {
      "type": "object",
      "required": [
        "trafficId"
      ],
      "properties": {
        "attributes": {
          "type": "object",
          "additionalProperties": {
            "type": "object"
          }
        },
        "trafficId": {
          "$ref": "#/definitions/TrafficId"
        }
      }
    },
    "FeatureFlagResponse": {
      "type": "object",
      "properties": {
        "flag": {
          "type": "string"
        },
        "value": {
          "type": "string"
        }
      }
    },
    "GatewayTimeoutResponse": {
      "description": "504 Gateway Timeout - The server cannot handle the request. Generally, this is a temporary state",
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "enum": [
            504
          ]
        },
        "message": {
          "type": "string"
        }
      }
    },
    "InternalServerErrorResponse": {
      "description": "500 Internal server error",
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "enum": [
            500
          ]
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ServiceUnavailableResponse": {
      "description": "503 Service unavailable - The server cannot handle the request (because it is overloaded or down for maintenance). Generally, this is a temporary state",
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "enum": [
            503
          ]
        },
        "message": {
          "type": "string"
        }
      }
    },
    "TrafficId": {
      "type": "string",
      "pattern": "^[\\S]{1,}$"
    },
    "UnauthorizedResponse": {
      "description": "401 Unauthorized",
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "enum": [
            401
          ]
        },
        "message": {
          "type": "string"
        }
      }
    }
  },
  "tags": [
    {
      "description": "Feature-flag service related endpoints.",
      "name": "feature"
    }
  ]
}`))
}