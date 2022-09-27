gen-swagger:
	cd cmd/internal-blockchain-api; swag init; cd ../..;

postgres:
	docker-compose up -d postgres-db-blockchain-internal

migrateup:
	migrate -path db/blockchain-internal/postgres/migration -database "postgresql://root:secret@localhost:5432/golang?sslmode=disable" -verbose up

migratedown:
	migrate -path db/blockchain-internal/postgres/migration -database "postgresql://root:secret@localhost:5432/golang?sslmode=disable" -verbose down

sqlcinternal:
	cd db/blockchain-internal; sqlc generate; cd ../..;

env:
	go mod tidy
	cd db/blockchain-internal; sqlc generate cd ../..;
	cd cmd/blockchain-internal; swag init; cd ../..;

swagger:
	cd cmd/internal-blockchain-api; swag init; cd ../..

genpb:
	protoc --proto_path=proto --go-grpc_opt=require_unimplemented_servers=false --go-grpc_out=proto/pb --go_out=proto/pb proto/*.proto

genpb-old:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:proto/pb

.PHONY: postgres migrateup migratedown sqlcuser sqlcwishlist genpb cleanpb test api user cert swagger ngrok env