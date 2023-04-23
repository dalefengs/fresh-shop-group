ssh root@fungs.cn "rm -rf /www/wwwroot/qiyun/web/*"
scp -r .\dist\* root@fungs.cn:/www/wwwroot/qiyun/web
pause
