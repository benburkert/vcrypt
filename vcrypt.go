package vcrypt

//go:generate -command protoc protoc --proto_path=$GOPATH/src:$GOPATH/src/github.com/gogo/protobuf/protobuf:. --gogo_out=.
//go:generate protoc cryptex/cryptex.proto cryptex/sss.proto cryptex/xor.proto cryptex/secretbox.proto
