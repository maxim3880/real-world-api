generate-api:
	make generate-model
	make generate-spec
	make generate-server
	make generate-client

generate-model:
	oapi-codegen -package restapi -templates ./swagger-template/ -generate types -o ./restapi/schemas.go ./swagger.yaml

generate-spec:
	oapi-codegen -package restapi -templates ./swagger-template/ -generate spec -o ./restapi/spec.go ./swagger.yaml

generate-server:
	oapi-codegen -package restapi -import-mapping fiber:github.com/gofiber/fiber/v2 -templates ./swagger-template/ -generate server -o ./restapi/server.go ./swagger.yaml

generate-client:
	oapi-codegen -package restapi -templates ./swagger-template/ -generate client -o ./restapi/client.go ./swagger.yaml