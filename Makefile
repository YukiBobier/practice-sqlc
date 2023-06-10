.PHONY: mysql-run mysql-stop

mysql-run:
	@docker run --name mysql-practice-sqlc -v $(shell pwd)/docker/mysql/docker-entrypoint-initdb.d:/docker-entrypoint-initdb.d -p 127.0.0.1:3306:3306/tcp -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:8

mysql-stop:
	@docker stop mysql-practice-sqlc
