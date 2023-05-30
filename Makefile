up_run:
	docker-compose up -d && go run main.go
up:
	docker-compose up -d
down:
	docker-compose down
go:
	air
