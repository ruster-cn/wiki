# 5 - Haproxy Limit Method

配置：
```
global
    # to have these messages end up in /var/log/haproxy.log you will
    # need to:
    #
    # 1) configure syslog to accept network log events.  This is done
    #    by adding the '-r' option to the SYSLOGD_OPTIONS in
    #    /etc/sysconfig/syslog
    #
    # 2) configure local2 events to go to the /var/log/haproxy.log
    #   file. A line like the following can be added to
    #   /etc/sysconfig/syslog
    #
    #    local2.*                       /var/log/haproxy.log
    #
    log         127.0.0.1 local2

    chroot      /var/lib/haproxy
    pidfile     /var/run/haproxy.pid
    maxconn     4000 #可接受的最大链接数
    user        haproxy
    group       haproxy
    daemon
    nbproc 56 #开启56线程
    # turn on stats unix socket
    stats socket /var/lib/haproxy/stats
    # 实现绑核用法
    #stats socket /var/run/haproxy.stat1 mode 600 level admin  process 1
    #stats socket /var/run/haproxy.stat2 mode 600 level admin  process 2
    
 
 
 listen admin_stats
        stats   enable
        bind    *:8080    #监听的ip端口号
        mode    http      #开关
        option  httplog
        log     global
        maxconn 10
        stats   refresh 30s   #统计页面自动刷新时间
        stats   uri /admin    #访问的uri   ip:8080/admin
        stats   realm haproxy
        stats   auth admin:admin  #认证用户名和密码
        stats   hide-version      #隐藏HAProxy的版本号
        stats   admin if TRUE     #管理界面，如果认证成功了，可通过webui管理节点

 
defaults
    log     global
    mode    tcp
    retries 3
    timeout connect      5000
    timeout client      50000
    timeout server      50000
 
 
 
frontend tcp
    bind *:2550
    default_backend test1
 
 
 
 
 
frontend tcp
    bind *:6443
    default_backend test1

frontend tcp
    bind *:6443
    default_backend test1

backend test1
    stick-table type ip   size 200k   expire 30s   store conn_rate(100s),bytes_out_rate(60s)
    acl whitelist src 192.168.1.154



    # values below are specific to the backend
    tcp-request content  track-sc2 src
    acl conn_rate_abuse  sc2_conn_rate gt 3
    acl data_rate_abuse  sc2_bytes_out_rate  gt 20000000

    # abuse is marked in the frontend so that it's shared between all sites
    acl mark_as_abuser   sc1_inc_gpc0 gt 0
    tcp-request content  reject if conn_rate_abuse !whitelist mark_as_abuser
    tcp-request content  reject if data_rate_abuse mark_as_abuser

    server s1 bjzyx-rs4842.zqy:6443
```




```
global
    log 127.0.0.1   local0
    log 127.0.0.1   local1 notice
    stats socket /var/run/haproxy.stat mode 600 level operator
    maxconn 4096
    user haproxy
    group haproxy
    daemon



defaults
    log     global
    mode    http
    option  httplog
    option  dontlognull
    retries 3
    option redispatch
    maxconn 2000
    timeout connect 5000ms
    timeout client 50000ms
    timeout server 50000ms


frontend 6443
    bind *:6443
    mode tcp
    option tcplog


    stick-table type ip size 200k expire 10m store gpc0

    # check the source before tracking counters, that will allow it to
    # expire the entry even if there is still activity.
    acl source_is_abuser src_get_gpc0(6443) gt 0
    use_backend ease-up-y0 if source_is_abuser
    tcp-request connection track-sc1 src if ! source_is_abuser


    use_backend test1


backend test1
    mode tcp
    stick-table type ip   size 200k   expire 30s   store conn_rate(100s),bytes_out_rate(60s)


    # values below are specific to the backend
    tcp-request content  track-sc2 src
    acl conn_rate_abuse  sc2_conn_rate gt 3
    acl data_rate_abuse  sc2_bytes_out_rate  gt 2

    # abuse is marked in the frontend so that it's shared between all sites
    acl mark_as_abuser   sc1_inc_gpc0 gt 0
    tcp-request content  reject if conn_rate_abuse mark_as_abuser
    tcp-request content  reject if data_rate_abuse mark_as_abuser

    server local_apache bjzyx-rs4842.zqy:6443


backend ease-up-y0
    errorfile 503 /etc/haproxy/errors/503rate.http

```

如何开始haproxy access log：[centos7.x之haproxy开启日志](https://blog.51cto.com/yanconggod/2062213)


参考文章：

1- [Better Rate Limiting For All with HAProxy](https://blog.serverfault.com/2010/08/26/1016491873/)

2- [haproxy document](http://cbonte.github.io/haproxy-dconv/1.7/configuration.html#7)