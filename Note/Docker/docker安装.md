## 安装docker

### 安装

``` shell
curl -fsSL https://get.docker.com | bash -s docker -- mirror Aliyun
```

### 设置开机启动docker

``` shell
systemctl enable docker		# 设置开机启动docker

systemctl start docker		# 启动docker

systemctl restart docker	# 重启docker

docker ps -a				# 查看docker信息
```

#### 配置阿里云镜像

1、登录阿里云

2、进入控制台

![1](/img/1.jpg)

3、然后在产品与服务中选择容器镜像服务

![2](/img/2.jpg)

4、选择镜像加速器

![3](/img/3.jpg)