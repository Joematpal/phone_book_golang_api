# protoc -I. contact/ contact/contact.proto \
# -I$GOPATH/src \
# -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
# --go_out=plugins=grpc:contact \

# protoc -I contact/ contact/contact.proto --go_out=plugins=grpc:contact
# protoc -I routeguide/ routeguide/routeguide.proto --go_out=plugins=grpc:routeguide

protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --swagger_out=logtostderr=true:. \
  contact/contact.proto
