# 初始化
下面已经初始化过了，不需要再次初始化
```
# 将Swagger UI转换为Go源代码
go-bindata --nocompress -pkg swagger -o pkg/ui/data/swagger/datafile.go third_party/swagger-ui/...

# 生成 google 依赖
protoc -I ./proto --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:./grpc proto/google/api/*.proto

```

# 添加新的 proto
```
# 生成业务的包
protoc -I ./proto --go_out=plugins=grpc,Mgoogle/api/annotations.proto=coding-grpc-gateway/grpc/google/api:./grpc ./proto/*.proto

# 成功 gw 业务包
protoc -I ./proto --grpc-gateway_out=logtostderr=true:./grpc ./proto/*.proto

# 生成 swagger
protoc -I./proto --swagger_out=logtostderr=true:./swagger/json ./proto/*.proto
```

# swagger 使用
1、 访问 http://{domain}/swagger-ui/  
2、 输入地址：http://{domain}/swagger/{proto名称}.swagger.json  