## Docker安装

### Docker的基本组成

![1](img/1.jpg)

**镜像（image）**

docker镜像就好比一个模板，可以通过这个模板来创建容器服务，tomcat镜像===> run ===> tomcat01容器（提供服务器），通过这个镜像可以创建多个容器（最终服务运行或项目运行就是在容器中的）。

**容器（container）**

Docker利用容器技术，独立运行一个或者一个组应用，通过镜像来创建的。

启动，停止，删除，基本命令！

目前就可以把这个容器理解为就是一个简单的linux系统

**仓库（repository）**

仓库就是存放镜像的地方！

仓库分为公有仓库和私有仓库！

### 安装Docker

> 安装

帮助文档：

```shell
# 1、卸载旧的版本
yum remove docker \
                  docker-client \
                  docker-client-latest \
                  docker-common \
                  docker-latest \
                  docker-latest-logrotate \
                  docker-logrotate \
                  docker-engine

# 2、需要的安装包
yum install -y yum-utils

# 3、设置镜像的仓库
yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo # 默认是国外的
yum-config-manager --add-repo http://mirrors.aliyun.com/docker-ce/linux/centos/docker-ce.repo # 阿里云

# 更新软件包索引
yum makecache fast

# 4、安装docker相关内容 docker-ce 社区版本  ee 企业版本
yum install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin

# 5、启动docker
systemctl start docker

# 6、查看docker是否运行成功
docker version

# 7、测试hello-word
docker run hello-world

# 8、查看一下下载的hello-world的镜像是否存在
docker images
```

了解：卸载docker

``` shell
# 1、卸载依赖
yum remove docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin docker-ce-rootless-extras

# 2、删除资源
rm -rf /var/lib/docker
rm -rf /var/lib/containerd
```

### Run的流程

![2](img/2.jpg)

### 底层原理

**Docker是怎么工作的？**

Docker是一个Client-Server结构的系统，Docker的守护进程运行在主机上。通过Socket从客户端访问！

Docker-Server接收到Docker-Client的指令，就会执行这个命令！

![3](img/3.jpg)

**Docker为什么比VM快？**

1、Docker有着比虚拟机更少的抽象层。

2、docker利用的是宿主机的内核，vm需要时Guest OS

![4](img/4.jpg)

所以说，新建一个容器的时候，docker不需要像虚拟机一样重新加载一个操作系统的内核，避免引导操作。虚拟机时加载Guest OS，分钟级别的，而docker是利用宿主机的操作系统，省略了这个过程，秒级

## Docker的常用命令

### 帮助命令

``` shell
docker version		# 显示docker的版本信息
docker info			# 显示docker的系统信息，包括镜像和容器的数量
docker 命令 --help   # 帮助命令	
```

### 镜像命令

**docker images** 查看所有本地的主机上的镜像

``` shell
[root@localhost ~]# docker images
REPOSITORY   TAG       IMAGE ID   CREATED   SIZE

# 解释
REPOSITORY 镜像的仓库源
TAG		   镜像的标签	
IMAGE ID   镜像的id
CREATED    镜像的创建时间
SIZE	   镜像的大小	

# 可选项
-a, --all		# 列出所有镜像
-q, --quiet		# 只显示镜像的id
```

**docker search** 搜索镜像

``` shell
[root@localhost ~]# docker search mysql
NAME                            DESCRIPTION                                     STARS     OFFICIAL   AUTOMATED
mysql                           MySQL is a widely used, open-source relation…   14566     [OK]       
mariadb                         MariaDB Server is a high performing open sou…   5558      [OK]   

# 可选项，通过收藏来过滤
--filter=STARS=3000   # 搜索出来的镜像就是STARS大于3000的
```

**docker pull** 下载镜像

``` shell
# 下载镜像 docker pull 镜像名[:tag]
[root@localhost ~]# docker pull mysql
Using default tag: latest				# 如果不写tag，默认就是最新版本
latest: Pulling from library/mysql
8e0176adc18c: Pull complete 			# 分层下载，docker image的核心 联合文件系统
2d2c52718f65: Pull complete 
d88d03ce139b: Pull complete 
4a7d7f11aa1e: Pull complete 
ce5949193e4c: Pull complete 
f7f024dfb329: Pull complete 
5fc3c840facc: Pull complete 
509068e49488: Pull complete 
cbc847bab598: Pull complete 
942bef62a146: Pull complete 
Digest: sha256:1773f3c7aa9522f0014d0ad2bbdaf597ea3b1643c64c8ccc2123c64afd8b82b1		# 签名
Status: Downloaded newer image for mysql:latest
docker.io/library/mysql:latest			# 真实地址

# 等价
docker pull mysql
docker pull docker.io/library/mysql:latest

# 指定版本下载
[root@localhost ~]# docker pull mysql:5.7
5.7: Pulling from library/mysql
9ad776bc3934: Pull complete 
a280ac4a8665: Pull complete 
4047a3b08336: Pull complete 
435611dd4999: Pull complete 
f84f2572cb0b: Pull complete 
ef893e58839b: Pull complete 
42897f531783: Pull complete 
8a8aad27e96b: Pull complete 
6b2751f26202: Pull complete 
b0e9b86ed64c: Pull complete 
bfef93045c96: Pull complete 
Digest: sha256:880063e8acda81825f0b946eff47c45235840480da03e71a22113ebafe166a3d
Status: Downloaded newer image for mysql:5.7
docker.io/library/mysql:5.7
```

**docker rmi** 删除镜像

``` shell
[root@localhost ~]# docker rmi -f 镜像id                   # 删除指定的镜像
[root@localhost ~]# docker rmi -f 镜像id 镜像id 镜像id 	  # 删除多个的镜像
[root@localhost ~]# docker rmi -f $(docker images -aq) 	  # 删除全部的镜像
```

### 容器命令

说明：我们有了镜像才可以创建容器

``` shell
docker pull centos
```

**新建容器并启动**

``` shell
docker run [可选参数] image

# 参数说明
--name="Name"		# 容器名字，用来区分容器
-d					# 后台方式运行
-it					# 使用交互方式运行，进入容器查看内容
-p					# 指定容器端口 -p 8080：8080
	-p ip:主机端口：容器端口
	-p 主机端口：容器端口（常用）
	-p 容器端口
	容器端口
-P					# 随机指定端口

# 测试，启动并进入容器
[root@localhost ~]# docker run -it centos /bin/bash
[root@9dde9aecd373 /]# ls	# 查看容器内的centos，基础版本，很多命令都是不完善的！ 			
bin  dev  etc  home  lib  lib64  lost+found  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var

# 从容器中退回主机
[root@9dde9aecd373 /]# exit
exit
[root@localhost ~]# ls
anaconda-ks.cfg  original-ks.cfg
```

**列出所有运行的容器**

``` shell
docker run [可选参数]

# 参数说明
		# 列出当前正在运行的容器
-a		# 列出当前正在运行的容器和带出历史运行过的容器
-n=?	# 显示最近创建的容器（？是个数）
-q		# 只显示容器的编号

[root@localhost ~]# docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
[root@localhost ~]# docker ps -a
CONTAINER ID   IMAGE     COMMAND       CREATED         STATUS                          PORTS     NAMES
9dde9aecd373   centos    "/bin/bash"   4 minutes ago   Exited (0) About a minute ago             determined_golick
```

**退出容器**

``` shell
exit		# 直接容器停止并退出
Ctrl +P +Q 	# 容器不停止退出
```

**删除容器**

``` shell
docker rm 容器id						# 删除指定的容器，不能删除正在运行的容器，如果强制删除，使用rm -f
docker rm -f $(docker ps -aq) 	 	 # 删除所有的容器
docker ps -a -q|xargs docker rm 	 # 删除所有的容器
```

**启动和停止容器的操作**

``` shell
docker start 容器id		# 启动容器
docker restart 容器id		# 重启容器
docker stop 容器id		# 停止当前正在运行的容器
docker kill 容器id		# 强制停止当前容器
```

