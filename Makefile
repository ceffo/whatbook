
regen:
	oapi-codegen api/whatbook_v1.yml -generate server > api-library/whatbook.gen.go

run:
	go run main.go
	