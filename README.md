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

## protoc-gen-customopt

```shell
cd protoc-gen-customopt
go build -o protoc-gen-customopt
protoc -I. --plugin=path/to/protoc-gen-customopt --customopt_out=. example.proto
```

# Reference

- https://github.com/yugui/protoc-gen-reqdump
- https://gist.github.com/yugui/e179eee28268e85c5036859987f8a15e
- https://qiita.com/yugui/items/87d00d77dee159e74886
- https://qiita.com/yugui/items/29adefab34f7f1a3c3c6