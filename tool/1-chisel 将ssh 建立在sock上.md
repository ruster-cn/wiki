# chisel 将 ssh 建立在 sock 上

本文只简单介绍使用方法，不做原理讨论。一般互联网公司的办公网和数据中心是网络隔离的，登陆数据中心需要使用跳板机。但是一般公司都会提供一个代理，允许办公网通过这个代理统一访问数据中心的 http 服务。chisel 就可以利用这个 http 代理实现 ssh 登陆。

- 在数据中心机器上启动 chisel server：

```
/usr/bin/chisel server --port 9999 --socks5 --key JHJGHGYTUHJTUHJYHFGHJJKHJYHGHJHFFGHJHJHJKGHJGJHGJGHJ
```

- 在办公网 mac 上 启动 chisel client:

```
./chisel_1.6.0_darwin_amd64 client --fingerprint a5:7e:c8:e1:04:7a:da:8c:69:cb:12:39:fd:8a:5c:c2  https://cds-admin.corp.kuaishou.com  socks
```

- ssh 登陆指定本地 sock 代理

```
ssh -o StrictHostKeyChecking=no -o ProxyCommand="nc -X 5 -x 127.0.0.1:1080 %h %p" web_server@bjpg-rs4286.yz02
```

- 下载地址：https://github.com/jpillora/chisel
