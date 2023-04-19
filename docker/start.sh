go mod init lykafe/news

go get gopkg.in/yaml.v3
go get github.com/kelseyhightower/envconfig

## grpc protoc
export PATH="$PATH:$HOME/.local/bin"
export PATH="$PATH:$(go env GOPATH)/bin"
protoc --go_out=. --go_opt=paths=import --go-grpc_out=. --go-grpc_opt=paths=import grpc/news/proto/*.proto --go-grpc_opt=paths=import grpc/news/proto/msg/*.proto
## grpc package
go get google.golang.org/grpc
go get google.golang.org/grpc/codes
go get google.golang.org/grpc/status
go get google.golang.org/protobuf/reflect/protoreflect
go get google.golang.org/protobuf/runtime/protoimpl
go mod download github.com/golang/protobuf
go get github.com/google/uuid
go get github.com/go-sql-driver/mysql

go get github.com/segmentio/kafka-go

go version

go run main.go