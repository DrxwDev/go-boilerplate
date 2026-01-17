include .env
export $(shell sed 's/=.*//' .env)

migration:
	@echo "Creating migration: $(name)"
	@goose -dir=database/migrations create "$(name)" sql

seeder:
	@echo "Creating seeder: $(name)"
	@goose -dir=database/seeder create "$(name)" sql

mu: migrate/up
migrate/up:
	@echo "Running up migrations..."
	@goose -dir=database/migrations up

md: migrate/down
migrate/down:
	@echo "Rolling back migrations..."
	@goose -dir=database/migrations down

mf: migrate/fresh
migrate/fresh:
	@echo "Dropping..."
	@goose -dir=database/migrations reset
	@make migrate/up

mfs: migrate/fresh/seed
migrate/fresh/seed:
	@make migrate/fresh
	@echo "Seeding..."
	@goose -dir=database/seeder -no-versioning up
