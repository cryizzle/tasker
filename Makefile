.PHONY: server_start server_log server_restart db_connect client_start client_build

server_start:
	docker compose build && docker compose up -d

server_log:
	docker compose logs -f -t

server_restart:
	docker compose build tasker_server && docker compose up -d tasker_server

server_test:
	go test -v ./...

db_connect:
	mysql -h localhost -P 3306 --protocol=tcp -u root -p

.PHONY: server_start server_log server_restart db_connect

client_start:
	cd tasker_client &&	npm start

client_build:
	cd tasker_client &&	npm run make

