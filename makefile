export MYSQL_URL = 'mysql://root:example@tcp(database:3306)/fastcampus'

migrate-create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up: 
	@ migrate -database ${MYSQL_URL} -path scripts/migrations up

migrate-down:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations down ${version}

migrate-force:
	@ migrate -path scripts/migrations -database ${MYSQL_URL} force ${version}