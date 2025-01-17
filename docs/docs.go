// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Massdriver",
            "url": "https://github.com/massdriver-cloud/mass"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "https://github.com/massdriver-cloud/mass/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/bundle/connections": {
            "get": {
                "description": "Get bundle connections",
                "produces": [
                    "application/json"
                ],
                "summary": "Get bundle connections",
                "operationId": "get-bundle-connections",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/bundle.Connections"
                        }
                    }
                }
            },
            "post": {
                "description": "Post bundle connections",
                "consumes": [
                    "application/json"
                ],
                "summary": "Post bundle connections",
                "operationId": "post-bundle-connections",
                "parameters": [
                    {
                        "description": "Connections",
                        "name": "connectons",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/bundle.Connections"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/bundle/secrets": {
            "get": {
                "description": "Get bundle secrets",
                "produces": [
                    "application/json"
                ],
                "summary": "Get bundle secrets",
                "operationId": "get-bundle-secrets",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/bundle.AppSpec"
                        }
                    }
                }
            }
        },
        "/config": {
            "get": {
                "description": "Get the users config",
                "produces": [
                    "application/json"
                ],
                "summary": "Get the users config",
                "operationId": "get-config",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_massdriver-cloud_mass_pkg_config.Config"
                        }
                    }
                }
            }
        },
        "/containers/list": {
            "get": {
                "description": "List containers searches using the name param, defaults to 'mass' if none provided.",
                "produces": [
                    "application/json"
                ],
                "summary": "List containers",
                "operationId": "list-containers",
                "parameters": [
                    {
                        "type": "boolean",
                        "default": false,
                        "description": "all containers, even stopped",
                        "name": "all",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "default": 0,
                        "description": "number of containers to return, 0 is all",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "default": "mass",
                        "description": "name of container to search with",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/types.Container"
                            }
                        }
                    }
                }
            }
        },
        "/containers/logs": {
            "get": {
                "description": "Stream the logs from a container using a websocket",
                "produces": [
                    "text/plain"
                ],
                "summary": "Stream logs",
                "operationId": "stream-logs",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of the container",
                        "name": "id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "101": {
                        "description": "Switching Protocols"
                    }
                }
            }
        }
    },
    "definitions": {
        "bundle.AppSpec": {
            "type": "object",
            "properties": {
                "envs": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "policies": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "secrets": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/bundle.Secret"
                    }
                }
            }
        },
        "bundle.Connections": {
            "type": "object",
            "additionalProperties": {}
        },
        "bundle.Secret": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "json": {
                    "type": "boolean"
                },
                "required": {
                    "type": "boolean"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "github_com_massdriver-cloud_mass_pkg_config.Config": {
            "type": "object",
            "properties": {
                "apiKey": {
                    "type": "string"
                },
                "orgID": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "mount.Propagation": {
            "type": "string",
            "enum": [
                "rprivate",
                "private",
                "rshared",
                "shared",
                "rslave",
                "slave"
            ],
            "x-enum-varnames": [
                "PropagationRPrivate",
                "PropagationPrivate",
                "PropagationRShared",
                "PropagationShared",
                "PropagationRSlave",
                "PropagationSlave"
            ]
        },
        "mount.Type": {
            "type": "string",
            "enum": [
                "bind",
                "volume",
                "tmpfs",
                "npipe",
                "cluster"
            ],
            "x-enum-varnames": [
                "TypeBind",
                "TypeVolume",
                "TypeTmpfs",
                "TypeNamedPipe",
                "TypeCluster"
            ]
        },
        "network.EndpointIPAMConfig": {
            "type": "object",
            "properties": {
                "ipv4Address": {
                    "type": "string"
                },
                "ipv6Address": {
                    "type": "string"
                },
                "linkLocalIPs": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "network.EndpointSettings": {
            "type": "object",
            "properties": {
                "aliases": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "driverOpts": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "endpointID": {
                    "type": "string"
                },
                "gateway": {
                    "type": "string"
                },
                "globalIPv6Address": {
                    "type": "string"
                },
                "globalIPv6PrefixLen": {
                    "type": "integer"
                },
                "ipaddress": {
                    "type": "string"
                },
                "ipamconfig": {
                    "description": "Configurations",
                    "allOf": [
                        {
                            "$ref": "#/definitions/network.EndpointIPAMConfig"
                        }
                    ]
                },
                "ipprefixLen": {
                    "type": "integer"
                },
                "ipv6Gateway": {
                    "type": "string"
                },
                "links": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "macAddress": {
                    "type": "string"
                },
                "networkID": {
                    "description": "Operational data",
                    "type": "string"
                }
            }
        },
        "types.Container": {
            "type": "object",
            "properties": {
                "Id": {
                    "type": "string"
                },
                "command": {
                    "type": "string"
                },
                "created": {
                    "type": "integer"
                },
                "hostConfig": {
                    "type": "object",
                    "properties": {
                        "networkMode": {
                            "type": "string"
                        }
                    }
                },
                "image": {
                    "type": "string"
                },
                "imageID": {
                    "type": "string"
                },
                "labels": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "mounts": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.MountPoint"
                    }
                },
                "names": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "networkSettings": {
                    "$ref": "#/definitions/types.SummaryNetworkSettings"
                },
                "ports": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/types.Port"
                    }
                },
                "sizeRootFs": {
                    "type": "integer"
                },
                "sizeRw": {
                    "type": "integer"
                },
                "state": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "types.MountPoint": {
            "type": "object",
            "properties": {
                "destination": {
                    "description": "Destination is the path relative to the container root (` + "`" + `/` + "`" + `) where the\nSource is mounted inside the container.",
                    "type": "string"
                },
                "driver": {
                    "description": "Driver is the volume driver used to create the volume (if it is a volume).",
                    "type": "string"
                },
                "mode": {
                    "description": "Mode is a comma separated list of options supplied by the user when\ncreating the bind/volume mount.\n\nThe default is platform-specific (` + "`" + `\"z\"` + "`" + ` on Linux, empty on Windows).",
                    "type": "string"
                },
                "name": {
                    "description": "Name is the name reference to the underlying data defined by ` + "`" + `Source` + "`" + `\ne.g., the volume name.",
                    "type": "string"
                },
                "propagation": {
                    "description": "Propagation describes how mounts are propagated from the host into the\nmount point, and vice-versa. Refer to the Linux kernel documentation\nfor details:\nhttps://www.kernel.org/doc/Documentation/filesystems/sharedsubtree.txt\n\nThis field is not used on Windows.",
                    "allOf": [
                        {
                            "$ref": "#/definitions/mount.Propagation"
                        }
                    ]
                },
                "rw": {
                    "description": "RW indicates whether the mount is mounted writable (read-write).",
                    "type": "boolean"
                },
                "source": {
                    "description": "Source is the source location of the mount.\n\nFor volumes, this contains the storage location of the volume (within\n` + "`" + `/var/lib/docker/volumes/` + "`" + `). For bind-mounts, and ` + "`" + `npipe` + "`" + `, this contains\nthe source (host) part of the bind-mount. For ` + "`" + `tmpfs` + "`" + ` mount points, this\nfield is empty.",
                    "type": "string"
                },
                "type": {
                    "description": "Type is the type of mount, see ` + "`" + `Type\u003cfoo\u003e` + "`" + ` definitions in\ngithub.com/docker/docker/api/types/mount.Type",
                    "allOf": [
                        {
                            "$ref": "#/definitions/mount.Type"
                        }
                    ]
                }
            }
        },
        "types.Port": {
            "type": "object",
            "properties": {
                "IP": {
                    "description": "Host IP address that the container's port is mapped to",
                    "type": "string"
                },
                "PrivatePort": {
                    "description": "Port on the container\nRequired: true",
                    "type": "integer"
                },
                "PublicPort": {
                    "description": "Port exposed on the host",
                    "type": "integer"
                },
                "Type": {
                    "description": "type\nRequired: true",
                    "type": "string"
                }
            }
        },
        "types.SummaryNetworkSettings": {
            "type": "object",
            "properties": {
                "networks": {
                    "type": "object",
                    "additionalProperties": {
                        "$ref": "#/definitions/network.EndpointSettings"
                    }
                }
            }
        }
    },
    "externalDocs": {
        "description": "OpenAPI",
        "url": "https://swagger.io/resources/open-api/"
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "127.0.0.1:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Massdriver API",
	Description:      "Massdriver Bundle Development Server API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
