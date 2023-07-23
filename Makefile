init:
	@go install github.com/Songmu/gocredits/cmd/gocredits@latest
	@go mod tidy

credits:
	@gocredits . > CREDITS

.PHONY: init credits