.PHONY: generate

run: generate
	@go run cmd/api/main.go

build: generate
	@go build -o server cmd/api/main.go

generate: cmd/templates/*.templ
	@templ generate
