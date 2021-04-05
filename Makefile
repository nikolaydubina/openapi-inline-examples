docs:
	cat testdata/openapi.only-comment.yaml | go run main.go > testdata/openapi.inplaced.yaml

test:
	go test -covermode=atomic ./...