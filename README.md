## IP Search - Golang 版本 ##
[![Develop on Okteto](https://okteto.com/develop-okteto.svg)](https://cloud.okteto.com/deploy?repository=https://github.com/nekomi-cn/ipsearch-go)

## 功能 ##
- 使用 IPIP 、 Maxmind 、 IP2Location 、 CZ88 IP数据库(数据库请自行更新，本项目使用数据均为网络公开数据库)
- 支持批量查询和单个API接口
- 支持 一键部署至Okteto

## 待完善 ##
- IP 数据库自动更新

## 使用 ## 
    git clone https://github.com/nekomi-cn/ipsearch-go.git
    go run cmd/app.go

    批量查询 http(s)://your.domain/
    单个API  http(s)://your.domain/api/1.1.1.1