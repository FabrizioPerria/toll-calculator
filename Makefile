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

distance-calculator:
	@go build -o bin/distance-calculator ./distance_calculator
	@./bin/distance-calculator

.PHONY: obu-send obu-recv distance-calculator 

