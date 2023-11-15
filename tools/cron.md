# crontab


`sudo apt install cron`

`sudo systemctl start cron`

`crontab -e` 编辑定时任务

`crontab -l` 查看定时任务

https://crontab.guru/#*_3_*_*_* 用来快速查看你的 cron 是不是写对了

先写个脚本，验证正确性，然后 crontab 直接执行脚本