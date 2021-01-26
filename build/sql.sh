
#!/bin/sh

#####################################
# ./sql.sh 生成mysql环境的sql
# ./sql.sh oracle  生成oracle环境的sql
#####################################

PATH=$PATH:$GOPATH/bin
rootdir=$(dirname $(pwd)) 

echo ""
echo "---------generate.db--------------------" 
echo ""

dbtype="mysql"
if [ $# -eq 1 ] && [ $1 = "oracle" ];then   
    dbtype=$1
fi
echo "当前生成sql的数据库类型为:$dbtype" 

cd $rootdir
which phorcys > /dev/null
if [ $? -ne 0 ]; then
	echo "phorcys未安装"
	echo "请到https://gitlab.100bm.cn/devtools/phorcys/phorcys.git下载安装"
	exit 1	
fi

if [ -d $rootdir/modules/const/$dbtype ] ; then 
	echo  $rootdir/modules/const/$dbtype"文件夹已存在，如需生成，请删除该文件夹"
	echo ""
	echo ""
	exit 1 
fi 


echo "1. 生成数据库sql 脚本"
phorcys markdown create sql -db $dbtype -file $rootdir/wtserver/db.oracle.md -o $rootdir/modules/const/$dbtype -c
if [ $? -ne 0 ]; then
	echo "phorcys 生成数据库脚本失败."
	exit 1
fi

echo ""
echo "---------generate.db-success-----------------" 
echo ""