DOCKER_COMPOSE = docker-compose
GO_BUILD = go build -v
BIN_DIR = $(CURDIR)/bin
RM = rm -rf

stop-air: 
	@sudo killall main
	@sudo killall air

containers: clean
	@$(DOCKER_COMPOSE) up -d

clean: stop-air
	@$(DOCKER_COMPOSE) down
	@$(RM) $(BIN_DIR)


obu-send:
	@$(GO_BUILD) -o $(BIN_DIR)/obu-send ./obu-send
	@$(BIN_DIR)/obu-send

obu-recv:
	@$(GO_BUILD) -o $(BIN_DIR)/obu-recv ./obu-recv
	@$(BIN_DIR)/obu-recv

distance-calculator:
	@$(GO_BUILD) -o $(BIN_DIR)/distance-calculator ./distance-calculator
	@$(BIN_DIR)/distance-calculator

aggregator:
	@$(GO_BUILD) -o $(BIN_DIR)/aggregator ./aggregator
	@-$(BIN_DIR)/aggregator

gateway:
	@$(GO_BUILD) -o $(BIN_DIR)/gateway ./gateway
	@$(BIN_DIR)/gateway

proto:
	@protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative shared/types/pb/*.proto


.PHONY: obu-send obu-recv distance-calculator aggregator proto gateway
