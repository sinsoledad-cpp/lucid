goctl api go -api app/user/api/user.api -dir app/user/api -style gozero 

goctl model mysql ddl -src="schema/sql/user/001_create_users_table.sql" -dir="data/model/user"