default:
	echo "no complete yet!"

openapi:
	go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen \
	-generate types,chi-server \
	-package openapi \
	-o internal/openapi/api.gen.go \
	api/openapi.yaml

debug: openapi
	go run cmd/main/main.go