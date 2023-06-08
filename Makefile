migrate_up:
	migrate -path db/migration -database "postgresql://mac:password@localhost:5432/restaurant_management?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://mac:password@localhost:5432/restaurant_management?sslmode=disable" -verbose down

migrate_force:
	migrate -path db/migration -database "postgresql://mac:password@localhost:5432/restaurant_management?sslmode=disable" force $(version)

run:	
	go build restaurant-management-api.go && ./restaurant-management-api

run_migrate:
	go build restaurant-management-api.go && ./restaurant-management-api -m

run_migrate_casbin:
	go build restaurant-management-api.go && ./restaurant-management-api -m -c

run_casbin:
	go build restaurant-management-api.go && ./restaurant-management-api -c

gotidy:
	go mod tidy

goinit:
	go mod init

swag:
	swag init -g restaurant-management-api.go -ot go,yaml 

migrate_init:
	migrate create -ext sql -dir db/migration -seq init_schema

launch:
	flyctl launch

docker_init:
	docker build -t everybody8/restaurant-management:v$(version) .

docker_push:
	docker push everybody8/restaurant-management:v$(version)

deploy:
	flyctl deploy

test:
	go test ./tests/repo_test && go test ./tests/service_test && go test ./tests/handler_test

repo_test:
	go test ./tests/repo_test

service_test:
	go test ./tests/service_test

handler_test:
	go test ./tests/handler_test 