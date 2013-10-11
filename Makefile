all: collect_migrations test

collect_migrations:
	./gen_migration_file.sh > migrations.go

test:
	go test
