package main

//go:generate npx swagger-cli bundle pkg/openapi/openapi.yaml -o pkg/openapi/.openapi.bundle.yaml -t yaml
//go:generate oapi-codegen --config=pkg/openapi/oapi-codegen.yaml pkg/openapi/.openapi.bundle.yaml
