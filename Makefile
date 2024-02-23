run-dev:
	docker-compose -f docker-compose.local.yml up --build

docker-build:
	docker build --tag ivmello/rinha-backend-2024-q1:latest .

docker-push:
	docker push ivmello/rinha-backend-2024-q1:latest