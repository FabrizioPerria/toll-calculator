DOCKER_COMPOSE = docker-compose
GO_BUILD = go build -v
BIN_DIR = $(CURDIR)/bin
RM = rm -rf


containers: clean
	@$(DOCKER_COMPOSE) up -d

clean:
	@$(DOCKER_COMPOSE) down
	@$(RM) $(BIN_DIR)

stop-air: clean
	@sudo killall main
	@sudo killall air

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


.PHONY: obu-send obu-recv distance-calculator aggregator
