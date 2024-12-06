api:
	@oapi-codegen -generate "types,server" -o internal/api/loans.gen.go -package api internal/api/loans.openapi.yaml