package main

const yaml = `
runMode: #服务运行模式
  debug
server:
  addr: #服务监听地址
    :9999
consul:
  addr: #consul监听地址（用于服务发现）
    127.0.0.1:8500
mysql:
  addr: #mysql监听地址
    127.0.0.1:3306
  username: #mysql用户名
    user
  password: #mysql密码
    pwd
  dbName: #mysql数据库
    db
redis:
  addr: #redis监听地址
    127.0.0.1:6379
  password: #redis密码
  db: #redis默认数据库
    1
  poolSize: #redis池大小
    100
mongodb:
  addr: #mongodb监听地址
    127.0.0.1:27017
  db: #mongodb数据库
    db
  cl: #mongodb表
    cl
  user: #mongodb用户名
  password: #mongodb密码
#可配置自定义配置项
`
