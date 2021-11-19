.PHONY: up
up:
	# docker network create --subnet=172.88.0.0/16 lynet
	docker-compose up -d
	
.PHONY: down
down:
	docker-compose down
