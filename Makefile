start:
	docker build -t rest-go-users -f docker/Dockerfile .
	docker-compose up -d

stop:
	docker-compose down

restart: stop start

logs:
	docker-compose logs -f

run-tests:
	docker build -t app-test-image -f test/Dockerfile .
	docker run --rm app-test-image
	docker rmi app-test-image