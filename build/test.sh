#!/bin/sh
  
set -o errexit

dt=$(date "+%Y%m%d%H%M%S")

rootdir=$(dirname $(dirname $(pwd)))

echo "0. rootdir:${rootdir}" 

#------------------------------------" 

echo "1. 编译项目" 
cd  ${rootdir}/wetoken/wtserver

go build --tags="test,mysql" -mod=mod

echo "2. 复制文件"

scp ./wtserver root@192.168.0.82:/tmp

ssh -t  root@192.168.0.82 "cd /root/wetoken/wtserver/bin;./wtserver stop;mv ./wtserver ./wtserver_${dt};cp /tmp/wtserver ./;sleep 3;./wtserver start;rm -rf /tmp/wtserver"


rm -rf ./wtserver

echo "3. 发布成功"
