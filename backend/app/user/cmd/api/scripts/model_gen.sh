#!/bin/bash

#生成的表名
tables=$2
#表生成的genmodel目录
modeldir=../../model

# 数据库配置
host=127.0.0.1
port=3306
dbname=$1
username=root
passwd=123456

echo "开始创建库：$dbname 的表：$2"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" --home ../../../../common/goctl -cache=true --style=go_zero
