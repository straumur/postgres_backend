all: collect_migrations test

collect_migrations:
	./gen_migration_file.sh > migrations.go
	go fmt migration.go

test:
	go test
