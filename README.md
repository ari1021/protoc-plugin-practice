# protoc-plugin-practice

```shell
brew install protobuf
```

## protoc-gen-message

```shell
cd protoc-gen-message
go build -o protoc-gen-message
protoc -I. --plugin=path/to/protoc-gen-message --message_out=. example.proto
```