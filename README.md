# MeowsList Server
![LICENSE](https://img.shields.io/badge/license-Apache%202.0-blue.svg)
![Go](https://img.shields.io/badge/go-1.18.3-00A8D1.svg)
![Postgres](https://img.shields.io/badge/db-PostgreSQL-32648D.svg)
![cache](https://img.shields.io/badge/cache-Redis-red.svg)
![Gin](https://img.shields.io/badge/gin-1.8.1-0090D1.svg)
![Serverless](https://img.shields.io/badge/serverless-TencentCloud-blue.svg)
![Serverless](https://img.shields.io/badge/serverless-Aliyun-orange.svg)
![Docker](https://img.shields.io/badge/docker-20.10.17-2392E6.svg)

Server of [meows-list](https://github.com/uiuing/meows-list)

## 说明

### 编译

```shell
go build
```

### 部署

1. `mv docker-compose.demo.yml docker-compose.yml`
2. `mv config.demo.yml config.yml`
3. 修改两个文件中的配置
4. `docker-compose up -d`
