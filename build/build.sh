#!/bin/sh

rm -rf ./out
#获取当前目录
publishsql="mysql"

for i in $@;  
do  
if [ $i = "oracle" ]; then 
        publishsql=$i 
fi;
done 

publishtags="$publishsql"
echo "-----------编译的tags:$publishtags--------"
echo "---------------打包开始--------------"
echo "-----------编译wtserver项目--------"
cd ../wtserver
echo "go build -mod=mod -tags "$publishtags" -o ../build/out/wtserver/bin/wtserver"
go build -mod=mod -tags "$publishtags" -o "../build/out/wtserver/bin/wtserver"

if [ $? -ne 0 ]; then
	echo "wtserver 项目编译出错,请检查"
	exit 1
fi
echo "-----------打包完成-------------"
echo "-----------都放在out目录中-------------"


