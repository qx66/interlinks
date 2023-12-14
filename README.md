# interlinks

此项目还在构思中...

interlinks 是一个关联工具

接口 -> SQL
SQL -> 接口

接口 -> API 接口
API 接口 -> 页面

openTelemetry ？

redis key 的抽象
SQL 的抽象

## 流程

1. 输入 SQL 语句
2. SQL 语句转换成指纹

## 场景

慢查SQL:

    1. 输入慢查SQL
    2. SQL转换指纹
    3. 通过指纹查询对应的服务

SQL 95, 99 执行时间:

    1. 输入慢查SQL
    2. SQL转换指纹
    3. 查询对应的95，99值

## 验证SQL

sqlparser.Parse 用于验证SQL是否有效，无效SQL返回error

## 引用库

https://github.com/percona/go-mysql

https://github.com/vitessio/vitess

https://github.com/XiaoMi/soar