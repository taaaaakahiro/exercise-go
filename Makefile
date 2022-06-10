
up:
	docker-compose up --build -d

down:
	docker-compose down -v

up-test:
	docker-compose -f docker-compose.test.yml up --build -d

down-test:
	docker-compose -f docker-compose.test.yml down -v

test: 
	go test ./... -count=1