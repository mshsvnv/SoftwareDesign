create_db_command := psql -U postgres -d Shop -f ./sql/create.sql
delete_db_command := psql -U postgres -d Shop -f ./sql/drop.sql

.PHONY: create_db delete_db test

create_db:
	$(create_db_command)

delete_db:
	$(delete_db_command)

# test: create_db
# 	go test -v ./internal/repository/postgres
# 	$(delete_db_command)

test: 
	go test -v ./internal/repository/postgres