include $(PWD)/.env

start:
	$(info Make: starting server)
	go build -o out && ./out

gen:
	$(info Make: generating sqlc code)
	sqlc generate

migrate_up:
	cd sql/schema && goose postgres $(PG_CONN) up

migrate_down:
	cd sql/schema && goose postgres $(PG_CONN) down