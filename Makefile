containers: clean
	@docker-compose up -d

clean:
	@docker-compose down

obu-send:
	@go build -o bin/obu-send ./obu-send
	@./bin/obu-send

obu-recv:
	@go build -o bin/obu-recv ./obu-recv
	@./bin/obu-recv

.PHONY: obu-send obu-recv 

