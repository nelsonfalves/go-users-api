run-db:
	docker compose -f deployments/docker-compose.yaml up -d ;

kill-db:
	docker compose -f deployments/docker-compose.yaml down ;
	