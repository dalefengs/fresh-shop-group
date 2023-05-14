# 部署前记得修改正式配置文件

ssh root@fungs.cn "rm -rf /www/wwwroot/qiyun/server/*"
scp -r .\* root@fungs.cn:/www/wwwroot/qiyun/server
pause
