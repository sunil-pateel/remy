.PHONY: generate tailwind_generate templ_generate

run: build
	@./server

air_build: generate
	@go build -o ./tmp/main cmd/api/main.go

build: generate
	@go build -o server cmd/api/main.go

generate: templ_generate tailwind_generate

templ_generate: web/templates/*.templ web/templates/components/*.templ
	@templ generate

tailwind_generate: web/css/*.css
	@tailwindcss -i web/css/input.css -o web/css/output.css
