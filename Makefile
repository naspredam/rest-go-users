start:
	docker build -t rest-go-users -f docker/Dockerfile .
	docker-compose up -d

stop:
	docker-compose down

restart: stop start

logs:
	docker-compose logs -f

start-local-db:
	docker-compose up -d db
	sh ./test/start-local-database.sh

run-all-tests: start-local-db
	docker build -t app-test-image -f test/Dockerfile .
	docker run --rm --net=host --env DATABASE_URL=root:rootpasswd@\(localhost:3306\)/app app-test-image
	docker rmi app-test-image
	docker-compose stop db
	docker-compose rm -f db