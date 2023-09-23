dev:
	air
generate:
	go generate ./ent
create-migration:
	atlas migrate diff migration_name \
		--dir "file://ent/migrate/migrations" \
		--to "ent://ent/schema" \
		--dev-url "docker://postgres/15/test?search_path=public"
docker-up:
	docker compose up -d
swag-generate:
	swag init
swag-fmt:
	swag fmt