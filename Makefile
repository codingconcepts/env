test:
	go test ./... -v

cover:
	go test --coverprofile=coverage.out
	go tool cover --html=coverage.out
