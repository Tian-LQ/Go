protoc --proto_path=./api --proto_path=./third_party --go_out=paths=source_relative:./api --go-http_out=paths=source_relative:./api --go-grpc_out=paths=source_relative:./api --openapi_out==paths=source_relative:. ./api/realworld/v1/*.proto


protoc --proto_path=./internal --proto_path=./third_party --go_out=paths=source_relative:./internal ./internal/conf/*.proto