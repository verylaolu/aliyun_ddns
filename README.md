# aliyun_ddns
#最傻逼
#最简单
#最方便的 
###阿里云ddns修改代码
找了一大堆， 不是废话多，就是功能复杂， 还他妈的有人搞docker
就这么个破玩意， 搞毛线docker
#资源不是钱么

复制 config.default.yml  更名为 config.yml
自行修改config.yml配置，然后用就完了

##编译就用，最好自己加个cronjob
#听说，下雨天，得得嗖嗖和cronjob更般配哦！！！！
*/10 * * * * /home/aliyun_ddns/aliyun_ddns>>/home/aliyun_ddns/cron.log
10分钟判断IP是否更改而选择同步域名解析一次

本来编译了一大堆版本， 后来嫌弃麻烦，自己去编译吧
#不会编译， 或者嫌麻烦不想安装GO环境的，本人有偿提供编译
#微信号：veryLaolu
#价格5元，童叟无欺