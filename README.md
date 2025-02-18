### 执行oapi-codegen

```

/scripts/api-http.sh 模块名称 包名  oapi-codegen的yml文件名 e.g. ./scripts/api-http.sh openapi internal/http http
 
```

#### 执行以上脚本前，确保有oapi-codegen命令

```
go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest
go get -u github.com/oapi-codegen/oapi-codegen/v2
```