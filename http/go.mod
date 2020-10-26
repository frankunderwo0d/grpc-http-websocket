module http

go 1.14

require (
	base_grpc v1.0.0
	github.com/gin-gonic/gin v1.6.3
	github.com/gorilla/websocket v1.4.2
	google.golang.org/grpc v1.33.1
)

replace base_grpc => ../grpc
