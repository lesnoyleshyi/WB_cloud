run:
	go run cmd/main.go

create_user:
	curl -X POST localhost:8080/account -d '{"id": "uuid_user_1", "balance": 100000}'
	@echo ""
	curl -X POST localhost:8080/account -d '{"id": "uuid_user_2", "balance": 100000}'

get_balance:
	curl -X GET 'localhost:8080/balance?id=uuid_user_1'
	@echo ""
	curl -X GET 'localhost:8080/balance?id=uuid_user_2'

transfer:
	curl -X POST localhost:8080/transfer -d \
    				'{"from": "uuid_user_1", "to": "uuid_user_2", "amount": 0.2}'

transfer_1000:
	for run in {1..1000}; do make transfer &>/dev/null ; done
