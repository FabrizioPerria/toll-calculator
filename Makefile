obu-send:
	@go run obu-send/main.go

obu-recv:
	@go run obu-recv/main.go

.PHONY: obu-send obu-recv 

