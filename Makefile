run:
	go run cmd/main.go

create_task:
	curl -X POST localhost:8080/account -d '{"id": "uuid_user_1", "balance": 100}'

get_balance:
	curl -X GET 'localhost:8080/balance?id=uuid_user_1'

