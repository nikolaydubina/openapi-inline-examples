docs:
	cat testdata/openapi.only-comment.yaml | go run main.go > testdata/openapi.inplaced.yaml

test:
	go test -timeout 300ms -covermode=atomic -coverprofile=coverage.out ./...

.PHONY: docs test
