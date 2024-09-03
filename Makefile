build: verify
	go build ./...
verify:
	go mod verify
test:
	go test ./...
demo-docker-start: mysql redis
demo-docker-stop:
	docker container stop web-app-redis web-app-mysql
demo-docker-rm: demo-docker-stop
	docker container rm web-app-redis web-app-mysql
mysql:
	docker run --name web-app-mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=sql_demo -p 13306:3306 -d mysql:8.0
redis:
	docker run --name web-app-redis -p 16379:6379 -d redis:7.0.4

.PHONY: build verify test mysql redis demo-docker-start demo-docker-stop demo-docker-rm