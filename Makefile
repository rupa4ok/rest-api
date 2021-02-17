init: docker-network docker-build docker-up wait-db

docker-build:
	docker-compose build

docker-up:
	docker-compose up -d

docker-network:
	docker network create --driver=bridge rest-go-network || true

wait-db:
	until docker-compose exec -T rest-go-mysql /usr/bin/mysql -uroot -proot -e "SHOW DATABASES;" ; do sleep 1 ; done
	docker-compose exec -T rest-go-mysql /usr/bin/mysql -uroot -proot -e "GRANT ALL ON *.* TO user@'%';FLUSH PRIVILEGES;"