# manadb/internal
一个简单的数据库实现，支持最基础的 SQL 语法

## 使用方法
支持远程连接和内嵌式使用
#### 远程连接
go run main.go 启动服务器
go run ./cmd/manasql 启动客户端

## 特性
#### 数据类型支持
- [x] BOOLEAN
- [x] BYTE
- [x] BYTES
- [x] CHAR
- [x] FLOAT32
- [x] FLOAT64
- [x] INT8
- [x] INT16
- [x] INT32
- [x] INT64
- [x] OID
- [x] OIDARRAY
- [x] TEXT
- [x] UINT8
- [x] UINT16
- [x] UINT32
- [x] UINT64
- [x] VARCHAR

#### SQL 语法支持
- [x] CREATE TABLE
- [x] DELETE
- [x] DROP TABLE
- [x] INSERT
- [x] SELECT
- [x] UPDATE

#### 连接条件支持
- [x] 笛卡尔积
- [x] INNER JOIN
- [x] LEFT JOIN
- [x] RIGHT JOIN
- [x] FULL OUTER JOIN

