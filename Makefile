.PHONY: run order service
order:
	@echo "Starting order service..."
	@cd order && go run main.go

.PHONY: run notify service
notif:
	@echo "Starting notify service..."
	@cd notification && go run main.go


.PHONY: run test
test:
	@echo "Starting test..."
	@cd test && go test .

.PHONY: run
run:
	@echo "Starting run container..."
	@docker-compose up -d

.PHONY: stop
stop:
	@echo "Remove container..."
	@docker-compose down -v
	@docker-compose down