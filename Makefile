build:
	docker-compose build todo-app

run:
	docker-compose up todo-app

migrate:
	migrate -path ./schema -database 'postgres://postgres:butterfly3000@0.0.0.0:5436/postgres?sslmode=disable' up