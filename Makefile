
regen:
	oapi-codegen api/whatbook_v1.yml > api-library/whatbook.gen.go

build:
	go build -o bin/

run: build
	./bin/whatbook
