# find

典型用法：

`find . -name 1699268919122-u6544fbd0e7ab82321.jpg -type f`


# 无限循环执行某命令

`while true; do echo hello; sleep 1; done`

# jobs

ctrl z 将前台进程移到后台 

jobs 查看所有被 ctrl z 的后台

fg %number 将指定后台任务再次前台运行

# 标准输出重定向

`/data/cron_tasks/nodebb_db_backup.sh >> /data/cron_tasks/log.txt 2>&1`

# set

set -eux
set -o pipefail

用来保证脚本正确性，并方便调试