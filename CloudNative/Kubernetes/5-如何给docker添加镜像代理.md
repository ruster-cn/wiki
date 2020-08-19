# 5-如何给docker添加镜像代理
- 新建一个docker.service.d目录，配置http proxy
```
mkdir /lib/systemd/system/docker.service.d
在该目录下新建文件http.conf
```
```
[Service]
Environment="HTTP_PROXY=http://proxy:38080/" "NO_PROXY=localhost"
```
重新夹在systemd配置，重启docker
```
systemctl daemon-reload 
systemctl restart docker
```