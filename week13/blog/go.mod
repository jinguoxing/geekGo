module geekGo/week13/blog

go 1.15

require (
	entgo.io/ent v0.8.0
	github.com/envoyproxy/protoc-gen-validate v0.6.1
	github.com/go-kratos/kratos/v2 v2.0.1
	github.com/go-redis/redis/extra/redisotel v0.3.0
	github.com/go-redis/redis/v8 v8.11.0
	github.com/go-sql-driver/mysql v1.6.0
	github.com/google/wire v0.5.0
	go.opentelemetry.io/otel v1.0.0-RC1
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0-RC1
	go.opentelemetry.io/otel/sdk v1.0.0-RC1
	go.opentelemetry.io/otel/trace v1.0.0-RC1
	google.golang.org/genproto v0.0.0-20210701191553-46259e63a0a9
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
)

replace (
	geekGo/week13 => ../../
)
