Kubernetes入门实战课
---

https://time.geekbang.org/column/intro/100114501

发布时间：2022

配套的学习项目：https://github.com/chronolaw/k8s_study

## 前言

现在Kubernetes已经没有了实际意义上的竞争对手，它的地位就如同Linux一样，成为了事实上的云原生操作系统，是构建现代应用的基石。

现代应用是什么？是微服务，是服务网格，这些统统要围绕着容器来开发、部署和运行，而使用容器就必然要用到容器编排技术，在现在只有唯一的选项，那就是Kubernetes。



### 学习Kubernetes有哪些难点

Kubernetes是一个分布式、集群化、云时代的系统，有许多新概念和新思维方式，和我们以往经验、认知的差异很大。

Kubernetes技术栈的特点，四个字概括：“**技术新、领域广、实现杂、方向深**”

- “**新**”是指Kubernetes用到的基本上都是比较前沿、陌生的技术，而且版本升级很快，经常变来变去。
- “**广**”是指Kubernetes涉及的应用领域很多、覆盖面非常广，不太好找到合适的切入点或者突破口。
- “**杂**”是指Kubernetes的各种实现比较杂乱，谁都可以上来“掺和”一下，让人看的眼晕。
- “**深**”是指Kubernetes面对的每个具体问题和方向，都需要有很深的技术背景和底蕴，想要吃透很不容易。





如果你真想做Kubernetes开发，等学会了Kubernetes的基本概念和用法，再回头去学Go语言也完全来得及。

学习Kubernetes最好的方式是**尽快建立一个全局观和大局观，等到你对这个陌生领域的全貌有了粗略但完整的认识之后，再挑选一个自己感兴趣的方向去研究，才是性价比最高的做法**。



2022年初发布的Kubernetes 1.23.3，是最后一个支持Docker的大版本，承上启下，具有很多的新特性，同时也保留了足够的历史兼容性，非常适合用来学习Kubernetes。

## 动手实践才是最好的学习方式



Ubuntu 22.04 Jammy Jellyfish 桌面版（https://ubuntu.com/download/desktop），它有足够新的特性，非常适合运行Kubernetes，而内置的浏览器、终端等工具也很方便我们的调试和测试。



# 入门

## 1 初识容器



除了 Docker，其他容器技术还有 Kata、gVisor、rkt、podman 等，但都不如 Docker 流行。

Docker Desktop之前一直是可以免费使用的，但在2021年8月31日，Docker 公司改变了策略，只对个人、教育和小型公司免费，其他形式的商业使用需要采用订阅制付费。

事实上，Docker Desktop 内部包含了 Docker Engine，也就是说，Docker Engine是Docker Desktop的核心组件之一。

我们还可以选择从 Docker 软件仓库，而不是Ubuntu 软件仓库来安装 Docker Engine,Docker 提供了官方的安装脚本，可以使用命令 `curl -fsSL
https://get.docker.com | bash -s docker`。用“docker version” 会看到有“Docker Engine - Community”的信息，表示社区版本（相应地还会有一个付费的“Docker Engine - Enterprise”）。

busybox 是一个小巧精致的“工具箱”，把诸多Linux 命令整合在一个可执行文件里，体积一般不超过2MB，非常适合测试任务或者嵌入式系统。

Moby是原来的Docker 开源项目，因为 Docker 已经成为了注册商标，所以在 2017年改了名字，作为目前 Docker 产品的试验上游而存在，类似 Fedora 与 CentOS/RHEL 的关系。



## 2 被隔离的进程：容器的本质

广义上来说，容器技术是动态的容器、静态的镜像和远端的仓库这三者的组合。

### 容器到底是什么

集装箱的作用是标准化封装各种货物，一旦打包完成之后，就可以从一个地方迁移到任意的其他地方。相比散装形式而言，集装箱隔离了箱内箱外两个世界，保持了货物的原始形态，避免了内外部相互干扰，极大地简化了商品的存储、运输、管理等工作。

计算机世界，容器也发挥着同样的作用，不过它封装的货物是运行中的**应用程序**，也就是**进程**，同样它也会把进程与外界隔离开，让进程与外部系统互不影响。

### 为什么要隔离

计算机世界里的隔离目的也是**系统安全**。

对于Linux操作系统来说，一个不受任何限制的应用程序是十分危险的。这个进程能够看到系统里**所有的文件、所有的进程、所有的网络流量，访问内存里的任何数据**，那么恶意程序很容易就会把系统搞瘫痪，正常程序也可能会因为无意的Bug导致信息泄漏或者其他安全事故。虽然Linux提供了用户权限控制，能够限制进程只访问某些资源，但这个机制还是比较薄弱的，和真正的“隔离”需求相差得很远。

**使用容器技术，我们就可以让应用程序运行在一个有严密防护的“沙盒”（Sandbox）环境之内**。

**容器技术的另一个本领就是为应用程序加上==资源隔离==，在系统里切分出一部分资源，让它只能使用指定的配额**。

### 与虚拟机的区别是什么

容器和虚拟机面对的都是相同的问题，使用的也都是虚拟化技术，只是所在的**层次**不同。

![](images/b734f7d91bda055236b3467bc16f6302.jpg)

（这图并不太准确，容器并不直接运行在Docker上，Docker只是辅助建立隔离环境，让容器基于Linux操作系统运行）



![](images/image-20240721194449779.png)



### 隔离是怎么实现的

容器隔离奥秘就在于Linux操作系统内核之中，为资源隔离提供了三种技术：**namespace、cgroup、chroot**，虽然这三种技术的初衷并不是为了实现容器，但它们三个结合在一起就会发生奇妙的“化学反应”。



## 3 容器化的应用：会了这些你就是Docker高手

### 什么是容器化的应用

### 常用的镜像操作有哪些

镜像的完整名字由两个部分组成，名字和标签，中间用:连接起来。

### 常用的容器操作有哪些





## 4 创建容器镜像：如何编写正确、高效的Dockerfile

### 镜像的内部机制是什么

### Dockerfile是什么

比起容器、镜像来说，Dockerfile非常普通，它就是一个纯文本，里面记录了一系列的构建指令，比如选择基础镜像、拷贝文件、运行脚本等等，每个指令都会生成一个Layer，而Docker顺序执行这个文件里的所有步骤，最后就会创建出一个新的镜像出来。



### 怎样编写正确、高效的Dockerfile



### docker build是怎么工作的

Dockerfile必须要经过 `docker build` 才能生效。





## 5 镜像仓库：该怎样用好Docker Hub这个宝藏

什么是镜像仓库（Registry）
什么是Docker Hub
如何在Docker Hub上挑选镜像
Docker Hub上镜像命名的规则是什么
该怎么上传自己的镜像
离线环境该怎么办



## 6 打破次元壁：容器该如何与外界互联互通

如何拷贝容器内的数据
如何共享主机上的文件
如何实现网络互通
如何分配服务端口号

## 7 实战演练：玩转Docker



![](images/79f8c75e018e0a82eff432786110ef16.jpg)

## 8 视频：入门篇实操总结



# 初级

## 9 走近云原生：如何在本机搭建小巧完备的Kubernetes环境

什么是容器编排
什么是Kubernetes
什么是minikube
如何搭建minikube环境
实际验证minikube环境



## 10 自动化的运维管理：探究Kubernetes工作机制的奥秘

### 云计算时代的操作系统



### Kubernetes的基本架构

![](images/image-20240721231007199.png)

### 节点内部的结构

#### Master里的组件有哪些

Master里有4个组件是**apiserver**、**etcd**、**scheduler**、**controller-manager**。

#### Node里的组件有哪些

#### 插件（Addons）有哪些





![](images/65d38ac50b4f2f1fd4b6700d5b8e7be1.jpg)



## 11 YAML：Kubernetes世界里的通用语

### 声明式与命令式是怎么回事

### 什么是YAML

### 什么是API对象

### 如何描述API对象

### 如何编写YAML





## 12 Pod：如何理解这个Kubernetes里最核心的概念？

### 为什么要有Pod

### 为什么Pod是Kubernetes的核心对象

![](images/9ebab7d513a211a926dd69f7535ac175.png)

![](images/image-20240721234815253.png)



### 如何使用YAML描述Pod





### 如何使用kubectl操作Pod





## 13 Job_CronJob：为什么不直接用Pod来处理业务？

### 为什么不直接使用Pod

### 为什么要有Job/CronJob

如何使用YAML描述Job
如何在Kubernetes里操作Job

### 如何使用YAML描述CronJob





## 14 ConfigMap_Secret：怎样配置、定制我的应用

### ConfigMap/Secret

什么是ConfigMap
什么是Secret



### 如何使用

如何以环境变量的方式使用ConfigMap/Secret
如何以Volume的方式使用ConfigMap/Secret



## 15 实战演练：玩转Kubernetes（1）





## 16 视频：初级篇实操总结



# 中级

## 17 更真实的云原生：实际搭建多节点的Kubernetes集群





## 18 Deployment：让应用永不宕机

### 为什么要有Deployment

### 如何使用YAML描述Deployment

### Deployment的关键字段

### 如何使用kubectl操作Deployment



## 19 Daemonset：忠实可靠的看门狗

为什么要有DaemonSet
如何使用YAML描述DaemonSet
如何在Kubernetes里使用DaemonSet
什么是污点（taint）和容忍度（toleration）
什么是静态Pod



## 20 Service：微服务架构的应对之道

为什么要有Service
如何使用YAML描述Service
如何在Kubernetes里使用Service
如何以域名的方式使用Service
如何让Service对外暴露服务



## 21 Ingress：集群进出流量的总管

为什么要有Ingress
为什么要有Ingress Controller
为什么要有IngressClass
如何使用YAML描述Ingress/Ingress Class
如何在Kubernetes里使用Ingress/Ingress Class
如何在Kubernetes里使用Ingress Controller





## 22 实战演练：玩转Kubernetes（2）





## 23 视频：中级篇实操总结



# 高级

## 24 PersistentVolume：怎么解决数据持久化的难题？

什么是PersistentVolume
什么是PersistentVolumeClaim/StorageClass
如何使用YAML描述PersistentVolume
如何使用YAML描述PersistentVolumeClaim
如何在Kubernetes里使用PersistentVolume
如何为Pod挂载PersistentVolume



## 25 PersistentVolume + NFS：怎么使用网络共享存储？

如何安装NFS服务器
如何安装NFS客户端
如何使用NFS存储卷
如何部署NFS Provisoner
如何使用NFS动态存储卷

## 26 StatefulSet：怎么管理有状态的应用？

什么是有状态的应用
如何使用YAML描述StatefulSet
如何在Kubernetes里使用StatefulSet
如何实现StatefulSet的数据持久化



## 27 滚动更新：如何做到平滑的应用升级降级？

Kubernetes如何定义应用版本
Kubernetes如何实现应用更新
Kubernetes如何管理应用更新
Kubernetes如何添加更新描述

## 28 应用保障：如何让Pod运行得更健康？

容器资源配额
什么是容器状态探针
如何使用容器状态探针

## 29 集群管理：如何用名字空间分隔系统资源？

为什么要有名字空间
如何使用名字空间
什么是资源配额
如何使用资源配额
默认资源配额

## 30 系统监控：如何使用Metrics Server和Prometheus？





## 31 网络通信：CNI是怎么回事？又是怎么工作的？

Kubernetes的网络模型
什么是CNI
CNI插件是怎么工作的
使用Calico网络插件



## 32 实战演练：玩转Kubernetes（3）



## 33 视频：高级篇实操总结



## docker-compose：单机环境下的容器编排工具



## 谈谈Kong Ingress Controller

结束语 是终点，更是起点
