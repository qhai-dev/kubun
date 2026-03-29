proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative api/foundation/v1/user.proto

sql-gen:
	cd ./internal/foundation && sqlc generate