web:
	cd frontend && turbo run dev
proto:
	cd backend/api/enterprise/v1 && protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative enterprise.proto
