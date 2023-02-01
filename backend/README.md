## SOJ

```shell
# first use
# intall gcc, because of the Makefile
go install github.com/zeromicro/go-zero/tools/goctl@1.4.3
go mod tidy

# second

cd serviceName/cmd
make api-gen | make model-gen | make rpc-gen
cd api && go run serviceName.go | cd rpc && go run serviceName.go
```
