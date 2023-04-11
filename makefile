run:
	CompileDaemon -command="./gocrud"

migrate:
	go run migrate/migrate.go