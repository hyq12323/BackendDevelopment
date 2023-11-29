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

### 常用其他命令

**后台启动容器**

``` shell
# 命令 docker run -d 镜像名
[root@localhost ~]# docker run -d centos

# 问题docker ps，发现centos 停止了

# 常见的坑：docker 容器使用后台运行，就必须要有一个前台进程，docker发现没有应用，就会自动停止
# nginx，容器后，发现自己没有提供服务，就会立刻停止，就是没有程序了
```

**查看日志**

``` shell
docker logs -f -t --tail 容器ID

[root@localhost ~]# docker run -d centos /bin/bash -c "while true;do echo hyq;sleep 1;done"

[root@localhost ~]# docker ps
CONTAINER ID   IMAGE     
74769c92d123   centos   

# 显示日志
-tf 			# 显示全部日志
--tail number   # 要显示日志条数
[root@localhost ~]# docker logs -tf --tail 10 74769c92d123
```

**查看容器中进程信息**

``` shell
# 命令 docker top 容器ID
[root@localhost ~]# docker top 74769c92d123
UID                 PID                 PPID                C                   STIME               TTY     
root                2810                2791                0                   01:19               ?       
root                3205                2810                0                   01:24               ?       
```

**查看镜像原数据**

``` shell
# 命令
docker inspect 容器ID

# 测试
[root@localhost ~]# docker inspect 74769c92d123
```

**进入当前正在运行的容器**

``` shell
# 我们通常容器都是使用后台方式运行的，需要进入容器，修改一些配置

# 命令1
docker exec -it 容器ID bashShell

# 测试
[root@localhost ~]# docker exec -it 74769c92d123 /bin/bash
[root@74769c92d123 /]# ls
bin  dev  etc  home  lib  lib64  lost+found  media  mnt  opt  proc  root  run  sbin  srv  sys  tmp  usr  var

# 命令2
docker attach 容器ID

# 测试
[root@localhost ~]# docker attach 74769c92d123
正在执行当前的代码...


# docker exec		# 进入容器后开启一个新的终端，可以在里面操作（常用）
# docker attach		# 进入容器正在执行的终端，不会启动新的进程！
```

**从容器内拷贝文件到主机上**

``` shell
docker cp 容器id:容器内路径 目的的主机路径

[root@localhost ~]# docker cp 74769c92d123:/srv/test.java /srv
```

### 小结

![5](/img/5.jpg)

``` shell
attach		 	# 当前shell下attach连接指定运行镜像
build	 		# 通过Dockerfile定制镜像
commit			# 提交当前容器为新的镜像
cp				# 从容器中拷贝指定文件或者目录到宿主机中
create			# 创建一个新的容器，同run，但不启动容器
diff			# 查看docker容器变化
events			# 从docker服务获取容器实时事件
exec			# 在已存在的容器上运行命令
export			# 导出容器的内容流作为一个tar归档文件【对应import】
history			# 展示一个镜像形成历史
images			# 列出系统当前镜像
import			# 从tar包的内容创建一个新的文件系统映像【对应export】
info			# 显示系统相关信息
inspect			# 查看容器详细信息
kill			# kill指定docker容器
load			# 从一个tar包中加载一个镜像【对应save】
login			# 注册或登录一个docker源服务器
logout			# 从当前Docker registry退出
logs			# 输出当前容器日志消息
port			# 查看映射端口对应的容器内部源端口
pause			# 暂停容器
ps				# 列出容器列表
pull			# 从docker镜像源服务器拉取指定镜像或者库镜像
push			# 推送指定镜像或者库镜像至docker源服务器
restart			# 重启运行的容器
rm				# 移除一个或者多个容器
rmi				# 移除一个或多个镜像【无容器使用该镜像才可删除，否则需删除相关容器才可继续，或-f强制删除】
run				# 创建一个新的容器并运行一个命令
save			# 保存一个镜像为一个tar包【对应load】
search			# 在docker hub中搜索镜像
start			# 启动容器
stop			# 停止容器
tag				# 给源中镜像打标签
top				# 查看容器中运行的进程信息
unpause			# 取消暂停容器
version			# 查看docker版本号
wait			# 截止容器停止时的退出状态值
```

### 练习

> Docker安装Nginx

``` shell
# 1、搜索镜像
docker search nginx
# 2、下载镜像
docker pull nginx
# 3、 运行测试
[root@localhost ~]# docker images
REPOSITORY   TAG       IMAGE ID       CREATED        SIZE
nginx        latest    a6bd71f48f68   45 hours ago   187MB

# -d 后台运行
# --name 给容器命名
# -p 宿主机端口:容器内部端口
[root@localhost ~]# docker run -d --name nginx01 -p 3344:80 nginx
7b7a66e33cefb575b24d7f69387e11e40a5b57189f9887f61a52450f70d8f213
[root@localhost ~]# docker ps
CONTAINER ID   IMAGE     STATUS                     PORTS                                   NAMES
7b7a66e33cef   nginx     "/docker-entrypoint.…"     0.0.0.0:3344->80/tcp, :::3344->80/tcp   nginx01
[root@localhost ~]# curl localhost:3344

[root@localhost ~]# docker exec -it nginx01 /bin/bash
root@7b7a66e33cef:/# whereis nginx
nginx: /usr/sbin/nginx /usr/lib/nginx /etc/nginx /usr/share/nginx
root@7b7a66e33cef:/# cd /etc/nginx/
root@7b7a66e33cef:/etc/nginx# ls
conf.d  fastcgi_params  mime.types  modules  nginx.conf  scgi_params  uwsgi_params
```

**端口暴露概念**

![6](/img/6.jpg)

思考问题：我们每次改动nginx配置文件，都需要进入容器内部？十分的麻烦，要是可以在容器外部提供一个映射路径，达到在容器修改文件，容器内部就可以自动修改？-v 数据卷！

> docker安装tomcat

``` shell
# 官方的使用
docker run -it --rm tomcat:9.0

# 我们之前的启动都是后台，停止容器之后，容器还是可以查到 docker run -it --rm tomcat:9.0，一般用来测试，用完即删除

# 下载在启动
docker pull tomcat

# 启动运行
docker run -d -p 3355:8080 --name tomcat01 tomcat
```

> 部署es+kibana

``` shell
# es 暴露的端口很多！
# es 十分耗费内存
# es 的数据一般需要放置到安全目录！挂载
# --net somenetwork	网络配置

# 启动
docker run -d --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node" elasticsearch:7.6.2

# 启动了 linux就卡住了

docker stats 		# 查看cpu状态  

# 增加内存限制，修改配置文件 -e 环境配置修改
docker run -d --name elasticsearch -p 9200:9200 -p 9300:9300 -e "discovery.type=single-node"  -e ES_JAVA_OPTS="-Xms64m -Xmx512m" elasticsearch:7.6.2
```

使用kibana连接es？思考网络如何才能连接过去！

![7](/img/7.jpg)

## Docker镜像讲解

### 镜像是什么

镜像是一种轻量级、可执行的独立软件包，用来打包软件运行环境和基于运行环境开发的软件，它包含运行某个软件所需的所有内容，包括代码、运行时、库、环境变量和配置文件。

所有的应用，直接打包docker镜像，就可以直接跑起来！

如何得到镜像：

+ 从远程仓库下载
+ 朋友拷贝给你
+ 自己制作一个镜像DockerFile

### Docker镜像加载原理

> UnionFS（联合文件系统）

我们下载的时候看到的一层层就是这个！

UnionFS（联合文件系统）：Union文件系统（UnionFS）是一种分层、轻量级并且高性能的文件系统，它支持对文件系统的修改作为一次提交来一层层的叠加，同时可以将不同目录挂载到同一个虚拟文件系统下（unite serveral directories into a single virtual filesystem）。Union文件系统是Docker镜像的基础。镜像可以通过分层来进行继承，基于基础镜像（没有父镜像），可以制作各种具体的应用镜像。

特性：一次同时加载多个文件系统，但从外面看起来，只能看到一个文件系统，联合加载会把各层文件系统叠加起来，这样最终的文件系统会包含所有底层的文件和目录。

> Docker镜像加载原理

Docker的镜像实际上由一层一层的文件系统组成，这种层级的文件系统UnionFS。

bootfs（boot file system）主要包含bootloader和kernel，bootloader主要是引导加载kernel，Linux刚启动时会加载bootfs文件系统，在Docker镜像的最底层是bootfs。这一层与我们典型的Linux/Unix系统是一样的，包含boot加载器和内核。当boot加载完成之后整个内核就都在内核中了，此时内存的使用权已由bootfs转交给内核，此时系统也会卸载bootfs。

rootfs（root file system），在bootfs之上。包含的就是典型Linux系统中的/dev,/proc,/bin,/etc等标准目录和文件。rootfs就是各种不同的操作系统发行版，比如Ubuntu，Centos等等。

![8](/img/8.jpg)

平时我们安装进虚拟机的CentOS都是好几个G，为什么Docker这里才200M？

![9](/img/9.jpg)

对于一个精简的OS，roofs可以很小，只需要包含最基本的命令，工具和程序库就可以了，因为底层直接用Host的kernel，自己只需要提供rootfs就可以了。由此可见对于不同的linux发行版，bootfs基本是一致的，rootfs会有差别，因此不同的发行版可以公用bootfs。

虚拟机是分钟级别，容器是秒级！

### 分层理解

> 分层的镜像

我们可以去下载一个镜像，注意观察下载的日志输出，可以看到是一层一层的在下载！

![10](/img/10.jpg)

思考：为什么Docker镜像要采用这种分层的结构呢？

最大的好处，我觉得莫过于是资源分享了！比如有多个镜像都从相同的Base镜像构建而来，那么宿主机只需在磁盘上保留一份base镜像，同时内存中也只需要加载一份base镜像，这样就可以为所有容器服务了，而且镜像的每一层都可以被共享。

查看镜像分层的方式可以通过docker image inspect 命令！

``` shell
[root@localhost ~]# docker image inspect redis:latest
[
	// ....
		"RootFS": {
            "Type": "layers",
            "Layers": [
                "sha256:92770f546e065c4942829b1f0d7d1f02c2eb1e6acf0d1bc08ef0bf6be4972839",
                "sha256:7eac7642cb301bfe7a028b97a60bb00b78a420dbc6237d27f383ea2042898b55",
                "sha256:732c5a0378d6d8dff665f0266ca0bec87f6f3473ebd27d8173205e10ea38075b",
                "sha256:13543e30dd49992a6a5913725430402a23d5f4694c699acb89ce73f43bb99271",
                "sha256:53dce266e70f8c322188d4d39cedf84ce4219e11622fc8b5f4efd653aab3a13c",
                "sha256:5f70bf18a086007016e948b04aed3b82103a36bea41755b6cddfaf10ace3c6ef",
                "sha256:0c118c4e6b60107378ae161b29f45146ecf5df21b80063b6286f7724780b0e0c"
            ]
        },
]
```

**理解：**

所有的Docker镜像都起始于一个基础镜像层，当进行修改或增加新的内容时，就会在当前镜像层之上，创建新的镜像层。

举一个简单的例子，假如基于Ubuntu Linux 16.04创建一个新的镜像，这就是新镜像的第一层；如果在该镜像添加python包，就会在基础镜像层之上创建第二个镜像层；如果继续添加一个安全补丁，就会创建第三个镜像层。

该镜像当前已经包含3个镜像层，如下图所示（这只是一个用于演示的简单的例子）。

![11](/img/11.jpg)

在添加额外的镜像层的同时，镜像始终保持是当前所有镜像的组合，理解这一点很重要。下面举一个简单的例子，每个镜像包含3个文件，而镜像包含了来自两个镜像层的6个文件。

![12](/img/12.jpg)

上图中的镜像层跟之前图中的略有区别，主要是便于展示文件。

下图展示了一个稍微复杂的三次镜像，在外部看来整个镜像只有6个文件，这是因为最上层中的文件7是文件5的一个更新版本。

![13](/img/13.jpg)

这种情况下，上层镜像层中的文件覆盖了底层镜像层中的文件。这样就使得文件的更新版本作为一个新镜像层添加到镜像当中。

Docker通过存储引擎（新版本采用快照机制）的方式来实现镜像层堆栈，并保证镜像层对外展示位统一的文件系统。

Linux可用的存储引擎有AUFS、Overlay2、Device Mapper、Btrfs以及ZFS。顾名思义，每种存储引擎都基于Linux中对应的文件系统或者块设备技术，并且每种存储引擎都有其独有的性能特点。

Docker在Windows上仅支持windowsfilter一种存储引擎，该引擎基于NTFS文件系统之上实现了分层和CoW[1]。

下图展示了与系统显示相同的三层镜像。所有镜像层堆叠合并，对外提供统一的视图。

![14](/img/14.jpg)

> 特点

Docker镜像都是只读的，当容器启动时，一个新的可写层被加载到镜像的顶部！

这一层就是我们通常说的容器层，容器之下都叫镜像层！

![15](/img/15.jpg)

### Commit镜像

``` shell
docker commit 提交容器成为一个新的副本

docker commit -m="提交的描述信息" -a="作者"  容器id 目标镜像名:[TAG]
```

**实战测试**

``` shell
# 1、启动一个默认的tomcat

# 2、发现这个默认的tomcat是没有webapps应用，镜像的原因，官方的镜像默认webapps下面是没有文件的！

# 3、自己拷贝进去文件

# 4、将操作过的容器通过commit提交为一个镜像！我们以后就使用我们修改过的镜像即可，这就是我们自己的一个修改过的镜像
```

![16](/img/16.jpg)

## 容器数据卷

### 什么是容器数据卷

**docker的理念回顾**

将应用和环境打包为一个镜像！

数据？如果数据都在容器中，那么我们容器删除，数据就会丢失！<font color = red>需求：数据可以持久化</font>

MySQL，容器删了，删库跑路！<font color = red>需求：MySQL数据可以存储在本地！</font>

容器之间可以有一个数据共享的技术！Docker容器中产生的数据，同步到本地！

这就是卷技术！目录的挂载，将我们容器内的目录，挂载到Linux上面！

![17](/img/17.jpg)

**总结一句话：容器的持久化和同步操作！容器间也是可以数据共享的！**

### 使用数据卷

> 方式一：直接使用命令来挂载 -v

``` shell
docker run -it -v 主机目录:容器内目录

# 测试
[root@localhost /]# docker run -it -v /home/ceshi:/home centos /bin/bash

docker inspect 容器ID		# 查看挂载是否成功 
```

![18](/img/18.jpg)

好处：我们以后修改只需要在本地修改即可，容器内会自动同步！

### 实战：安装MySQL

思考：MySQL的数据持久化

``` shell
# 获取镜像
[root@localhost ~]# docker pull mysql:5.7

# 运行容器，需要数据挂载
# 官方使用：docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=my-secret-pw -d mysql:tag

# 启动容器
-d 后台运行
-p 端口映射
-v 卷挂载
-e 环境配置
--name 容器名字
[root@localhost ~]# docker run -d -p 3310:3306 -v /home/mysql/conf:/etc/mysql/conf.d -v /home/mysql/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=990725 --name mysql01 mysql:5.7

# 启动成功之后，我们在本地使用navicat来测试一下
# navicat--连接到服务器3310---3310和3306映射，这时候就可以连接上

# 在本地测试创建一个数据库，查看一下我们的映射路径是否成功
```

假设我们容器删除，发现我们挂载到本地的数据卷没有丢失，这就实现了容器持久化功能！
