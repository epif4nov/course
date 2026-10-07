// Package practice documents how to generate the HTTP API router from OpenAPI.
package practice

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 --config api/oapi-codegen.yaml api/openapi.yaml
