# idocker

fork  from https://github.com/xianlubird/mydocker



## 环境

```bash
# cat /etc/issue
Ubuntu 20.04.1 LTS \n \l

# uname -r
5.4.0-51-generic
```



## 编译

```bash
# go build -o idocker cmd/*
# idocker/pkg/nsenter
pkg/nsenter/nsenter.go: In function ‘enter_namespace’:
pkg/nsenter/nsenter.go:36:7: warning: implicit declaration of function ‘setns’; did you mean ‘setenv’? [-Wimplicit-function-declaration]
   36 |   if (setns(fd, 0) == -1) {
      |       ^~~~~
      |       setenv
pkg/nsenter/nsenter.go:41:3: warning: implicit declaration of function ‘close’; did you mean ‘pclose’? [-Wimplicit-function-declaration]
   41 |   close(fd);
      |   ^~~~~
      |   pclose

```



## 如何运行

1. 设置docker的storage driver为aufs运行

   ```bash
   /usr/bin/dockerd -s aufs
   ```

2. 获取nginx.tar

   ```bash
   # docker run --name nginx -d nginx
   5bf9150951359707d10e1df45ba5975868858b1176f9045e88dadbe36ede178a
   
   # docker ps -a
   CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS               NAMES
   5bf915095135        nginx               "/docker-entrypoint.…"   23 seconds ago      Up 22 seconds       80/tcp              nginx
   
   # docker export -o /root/nginx.tar nginx
   ```

3. 创建容器网络

   ```bash
   # ./idocker network create --driver bridge --subnet 192.168.24.1/24 ibridge
   
   #设置filter FORWARD链为ACCEPT,不然会影响容器访问外部网络的通信
   # iptables -P FORWARD ACCEPT
   ```

4. 创建容器t1

   ```bash
   # ./idocker run -d --name t1 -net ibridge -p 9000:80 nginx top -b
   {"level":"info","msg":"createTty false","time":"2020-10-19T11:32:23+08:00"}
   {"level":"info","msg":"command all is top -b","time":"2020-10-19T11:32:23+08:00"}
   {"level":"warning","msg":"remove cgroup fail unlinkat /sys/fs/cgroup/memory/3129175603/cgroup.procs: operation not permitted","time":"2020-10-19T11:32:23+08:00"}
   {"level":"warning","msg":"remove cgroup fail unlinkat /sys/fs/cgroup/cpu,cpuacct/3129175603/cgroup.procs: operation not permitted","time":"2020-10-19T11:32:23+08:00"}
   # ./idocker exec t1 sh
   {"level":"info","msg":"container pid 132383","time":"2020-10-19T11:32:35+08:00"}
   {"level":"info",容器"msg":"command sh","time":"2020-10-19T11:32:35+08:00"}
   / # ip a
   1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1000
       link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
       inet 127.0.0.1/8 scope host lo
          valid_lft forever preferred_lft forever
       inet6 ::1/128 scope host 
          valid_lft forever preferred_lft forever
   49: cif-31291@if50: <BROADCAST,MULTICAST,UP,LOWER_UP,M-DOWN> mtu 1500 qdisc noqueue state UP qlen 1000
       link/ether b6:04:57:4a:a2:cb brd ff:ff:ff:ff:ff:ff
       inet 192.168.24.8/24 brd 192.168.24.255 scope global cif-31291
          valid_lft forever preferred_lft forever
       inet6 fe80::b404:57ff:fe4a:a2cb/64 scope link 
          valid_lft forever preferred_lft forever
   / # ping 114.114.114.114
   PING 114.114.114.114 (114.114.114.114): 56 data bytes
   64 bytes from 114.114.114.114: seq=0 ttl=61 time=29.990 ms
   64 bytes from 114.114.114.114: seq=1 ttl=70 time=29.934 ms
   ^C
   --- 114.114.114.114 ping statistics ---
   2 packets transmitted, 2 packets received, 0% packet loss
   round-trip min/avg/max = 29.934/29.962/29.990 ms
   ```

5. 创建容器t2容器

   ```bash
   # ./idocker run -d --name t2 -net ibridge -p 9001:80 nginx top -b
   {"level":"info","msg":"createTty false","time":"2020-10-19T11:32:53+08:00"}
   {"level":"info","msg":"command all is top -b","time":"2020-10-19T11:32:53+08:00"}
   {"level":"warning","msg":"remove cgroup fail unlinkat /sys/fs/cgroup/memory/5355103441/cgroup.procs: operation not permitted","time":"2020-10-19T11:32:53+08:00"}
   {"level":"warning","msg":"remove cgroup fail unlinkat /sys/fs/cgroup/cpu,cpuacct/5355103441/cgroup.procs: operation not permitted","time":"2020-10-19T11:32:53+08:00"}
   # ./idocker exec t2 sh
   {"level":"info","msg":"container pid 132462","time":"2020-10-19T11:32:56+08:00"}
   {"level":"info","msg":"command sh","time":"2020-10-19T11:32:56+08:00"}
   / # ip a
   1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN qlen 1000
       link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
       inet 127.0.0.1/8 scope host lo
          valid_lft forever preferred_lft forever
       inet6 ::1/128 scope host 
          valid_lft forever preferred_lft forever
   51: cif-53551@if52: <BROADCAST,MULTICAST,UP,LOWER_UP,M-DOWN> mtu 1500 qdisc noqueue state UP qlen 1000
       link/ether 8e:97:53:cd:8e:30 brd ff:ff:ff:ff:ff:ff
       inet 192.168.24.9/24 brd 192.168.24.255 scope global cif-53551
          valid_lft forever preferred_lft forever
       inet6 fe80::8c97:53ff:fecd:8e30/64 scope link 
          valid_lft forever preferred_lft forever
   / # ping 114.114.114.114
   PING 114.114.114.114 (114.114.114.114): 56 data bytes
   64 bytes from 114.114.114.114: seq=0 ttl=66 time=30.005 ms
   64 bytes from 114.114.114.114: seq=1 ttl=81 time=29.993 ms
   64 bytes from 114.114.114.114: seq=2 ttl=72 time=29.909 ms
   ^C
   --- 114.114.114.114 ping statistics ---
   3 packets transmitted, 3 packets received, 0% packet loss
   round-trip min/avg/max = 29.909/29.969/30.005 ms
   
   #ping t1容器ip
   / # ping 192.168.24.8
   PING 192.168.24.8 (192.168.24.8): 56 data bytes
   64 bytes from 192.168.24.8: seq=0 ttl=64 time=0.134 ms
   64 bytes from 192.168.24.8: seq=1 ttl=64 time=0.084 ms
   64 bytes from 192.168.24.8: seq=2 ttl=64 time=0.087 ms
   64 bytes from 192.168.24.8: seq=3 ttl=64 time=0.087 ms
   ```