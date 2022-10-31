init-schema:
	go run -mod=mod entgo.io/ent/cmd/ent init ${target}
gen-ent:
	go generate ./ent
create-db:
	docker exec -it sol-shop-pg createdb --username=root --owner=root -e sol-shop

.PHONY:
	init-schema go-ent