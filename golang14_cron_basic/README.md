## 参考资料
https://github.com/robfig/cron <br>
http://chuquanl.com/golang-cron%E7%AE%80%E4%BB%8B%E5%8F%8A%E4%BD%BFcron%E6%94%AF%E6%8C%81%E5%B8%A6%E5%8F%82%E6%95%B0%E4%BB%BB%E5%8A%A1%E8%B0%83%E7%94%A8/ <br>
## 依赖包概述
依赖包主要有两个文件:cron.go、parser.go。<br>
主要结构体:Cron、entry、schedule。<br>
主要技术点:channel进行传递。<br>