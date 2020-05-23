start:
	docker build -t rest-go-users -f docker/Dockerfile .
	docker-compose up -d

stop:
	docker-compose down

restart:
	make stop && make start

logs:
	docker-compose logs -f