用户先进sendhttp，然后生成模板和发送记录  进入模板审核队列 审核消费通过，审核通过根据模板以及发送记录 组装成一条条短信队列shengchan，进入发送微服务消费 发送给短信通道


1 template must build 
t_record

call sendsvc,
call appsvc,
call 