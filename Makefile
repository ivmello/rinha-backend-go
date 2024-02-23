down:
	docker-compose down

up:
	docker-compose up --build

stats:
	docker stats

build:
	docker build --tag ivmello/rinha-backend-2024-q1:latest .

push:
	docker push ivmello/rinha-backend-2024-q1:latest