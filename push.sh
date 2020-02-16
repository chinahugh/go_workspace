#! /bin/bash 

#echo "******************************************************************************************************"
#echo "备份mysql数据库"
# 备份mysql数据库
# 备份文件在同级目录下的mysql.sql文件中
# 向数据库导入备份文件时，先删除文件中的mysq数据库
#echo "开始备份数据库..."
#set  MYSQL_HOME=C:/Program Files/MySQL/MySQL Server 8.0
#"$MYSQL_HOME/bin/mysqldump.exe" -uroot -proot -A >mysql.sql
#echo "备份数据库完成..."
#read -t 36000 -p "按任意键继续"

echo "******************************************************************************************************">>push.log
echo start push...
#read -t 36000  -p "请输入本次提交信息 " msg
date >>push.log
#echo "message >>>>> $msg">>push.log
git add . >>push.log 2>&1
git commit -m "自动提交" >>push.log 2>&1
git push origin master >>push.log 2>&1
read -t 36000 -p "按任意键结束"
