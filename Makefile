
regen:
	oapi-codegen api/whatbook_v1.yml > api-library/whatbook.gen.go

run:
	go run main.go
