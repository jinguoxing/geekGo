module geekGo/grpc_demo/productinfo/client

go 1.16

replace geekGo/grpc_demo/productinfo/client/ecommerce => ./ecommerce

require (
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
)
