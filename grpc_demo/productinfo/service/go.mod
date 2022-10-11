module geekGo/grpc_demo/productinfo/service

go 1.16

require (
	github.com/gofrs/uuid v4.2.0+incompatible
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
)

replace (
	geekGo/grpc_demo/productinfo/service/ecommerce => ./ecommerce
)
