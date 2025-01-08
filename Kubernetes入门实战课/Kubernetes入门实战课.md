Kubernetes入门实战课
---

https://time.geekbang.org/column/intro/100114501

发布时间：2022

配套的学习项目：https://github.com/chronolaw/k8s_study

## 前言

现在Kubernetes已经没有了实际意义上的竞争对手，它的地位就如同Linux一样，成为了事实上的云原生操作系统，是构建现代应用的基石。

现代应用是什么？是微服务，是服务网格，这些统统要围绕着容器来开发、部署和运行，而使用容器就必然要用到容器编排技术，在现在只有唯一的选项，那就是Kubernetes。

```
 ;
 ;
 ;
```



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

### Docker的诞生

除了 Docker，其他容器技术还有 Kata、gVisor、rkt、podman 等，但都不如 Docker 流行。



我们还可以选择从 Docker 软件仓库，而不是Ubuntu 软件仓库来安装 Docker Engine,Docker 提供了官方的安装脚本，可以使用命令 `curl -fsSL
https://get.docker.com | bash -s docker`。用“docker version” 会看到有“Docker Engine - Community”的信息，表示社区版本（相应地还会有一个付费的“Docker Engine - Enterprise”）。

busybox 是一个小巧精致的“工具箱”，把诸多Linux 命令整合在一个可执行文件里，体积一般不超过2MB，非常适合测试任务或者嵌入式系统。

Moby是原来的Docker 开源项目，因为 Docker 已经成为了注册商标，所以在 2017年改了名字，作为目前 Docker 产品的试验上游而存在，类似 Fedora 与 CentOS/RHEL 的关系。

### Docker的形态

Docker Desktop之前一直是可以免费使用的，但在2021年8月31日，Docker 公司改变了策略，只对个人、教育和小型公司免费，其他形式的商业使用需要采用订阅制付费。

事实上，Docker Desktop 内部包含了 Docker Engine，也就是说，Docker Engine是Docker Desktop的核心组件之一。



### Docker的安装



```
{
    "dns": ["8.8.8.8", "8.8.4.4"],
    "registry-mirrors": [
        "https://docker.m.daocloud.io/",
        "https://huecker.io/",
        "https://dockerhub.timeweb.cloud",
        "https://noohub.ru/",
        "https://dockerproxy.com",
        "https://docker.mirrors.ustc.edu.cn",
        "https://docker.nju.edu.cn",
        "https://xx4bwyg2.mirror.aliyuncs.com",
        "http://f1361db2.m.daocloud.io",
        "https://registry.docker-cn.com",
        "http://hub-mirror.c.163.com"
    ]
}
```

### Docker的使用

### Docker的架构



## 2 被隔离的进程：容器的本质

广义上来说，容器技术是动态的容器、静态的镜像和远端的仓库这三者的组合。

### 容器到底是什么

集装箱的作用是标准化封装各种货物，一旦打包完成之后，就可以从一个地方迁移到任意的其他地方。相比散装形式而言，集装箱隔离了箱内箱外两个世界，保持了货物的原始形态，避免了内外部相互干扰，极大地简化了商品的存储、运输、管理等工作。

计算机世界，容器也发挥着同样的作用，不过它封装的货物是运行中的**应用程序**，也就是**进程**，同样它也会把进程与外界隔离开，让进程与外部系统互不影响。

```sh
docker pull ubuntu:18.04
docker run -it ubuntu:18.04 sh

# 下面的命令都是在容器内执行
cat /etc/os-release
apt update
apt install -y wget redis
redis-server &
```



### 为什么要隔离

计算机世界里的隔离目的也是**系统安全**。

对于Linux操作系统来说，一个不受任何限制的应用程序是十分危险的。这个进程能够看到系统里**所有的文件、所有的进程、所有的网络流量，访问内存里的任何数据**，那么恶意程序很容易就会把系统搞瘫痪，正常程序也可能会因为无意的Bug导致信息泄漏或者其他安全事故。虽然Linux提供了用户权限控制，能够限制进程只访问某些资源，但这个机制还是比较薄弱的，和真正的“隔离”需求相差得很远。

**使用容器技术，我们就可以让应用程序运行在一个有严密防护的“沙盒”（Sandbox）环境之内**。

**容器技术的另一个本领就是为应用程序加上==资源隔离==，在系统里切分出一部分资源，让它只能使用指定的配额**。

### 与虚拟机的区别是什么

容器和虚拟机面对的都是相同的问题，使用的也都是虚拟化技术，只是所在的**层次**不同。

![](images/b734f7d91bda055236b3467bc16f6302.jpg)

（这图并不太准确，容器并不直接运行在Docker上，Docker只是==辅助建立隔离环境==，让容器基于Linux操作系统运行）

容器和虚拟机的目的都是隔离资源，保证系统安全，然后是尽量提高资源的利用率。

在数据中心的服务器上，虚拟机软件（即图中的`Hypervisor`）同样可以把一台物理服务器虚拟成多台逻辑服务器，这些逻辑服务器彼此独立，可以按需分隔物理服务器的资源，为不同的用户所使用。

从实现的角度来看，虚拟机虚拟化出来的是**硬件**，需要在上面再安装一个**操作系统**后才能够运行应用程序，而硬件虚拟化和操作系统都比较“重”，会消耗大量的CPU、内存、硬盘等系统资源，但这些消耗其实并没有带来什么价值，属于“重复劳动”和“无用功”，不过好处就是隔离程度非常高，每个虚拟机之间可以做到完全无干扰。

而**容器（Docker）直接利用了下层的计算机硬件和操作系统**，因为比虚拟机少了一层，所以自然就会节约CPU和内存，显得非常轻量级，能够更高效地利用硬件资源。

![](images/image-20240721194449779.png)



### 隔离是怎么实现的

容器隔离奥秘就在于Linux操作系统内核之中，为资源隔离提供了三种技术：**namespace、cgroup、chroot**，虽然这三种技术的初衷并不是为了实现容器，但它们三个结合在一起就会发生奇妙的“化学反应”。

- ==namespace==是2002年从Linux 2.4.19开始出现的，和编程语言里的namespace有点类似，它可以创建出独立的**文件系统、主机名、进程号、网络**等资源空间，相当于给进程盖了一间**小板房**，这样就实现了系统全局资源和进程局部资源的隔离。

- ==cgroup==是2008年从Linux 2.6.24开始出现的，它的全称是Linux Control Group，用来实现**对进程的CPU、内存等资源的优先级和配额限制**，相当于给进程的小板房加了一个**天花板**。

- ==chroot==的历史则要比前面的namespace、cgroup要古老得多，早在1979年的UNIX V7就已经出现了，它可以**更改进程的根目录，也就是限制访问文件系统**，相当于给进程的小板房铺上了**地砖**。





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

### 10.1 云计算时代的操作系统



### 10.2 Kubernetes的基本架构

![](images/image-20240721231007199.png)



### 10.3 节点内部的结构

#### Master里的组件有哪些

Master里有4个组件是**apiserver**、**etcd**、**scheduler**、**controller-manager**。

![](images/330e03a66f636657c0d8695397c508c6.jpg)



#### Node里的组件有哪些



![](images/87bab507ce8381325e85570f3bc1d935.jpg)





![](images/344e0c6dc2141b12f99e61252110f6b7.png)

#### 插件（Addons）有哪些



### 小结

![](images/65d38ac50b4f2f1fd4b6700d5b8e7be1.jpg)

1. Kubernetes能够在集群级别管理应用和服务器，可以认为是一种集群操作系统。它使用“控制面/数据面”的基本架构，Master节点实现管理控制功能，Worker节点运行具体业务。
2. Kubernetes由很多模块组成，可分为核心的组件和选配的插件两类。
3. Master里有4个组件，分别是apiserver、etcd、scheduler、controller-manager。
4. Node里有3个组件，分别是kubelet、kube-proxy、container-runtime。
5. 通常必备的插件有DNS和Dashboard。





## 11 YAML：Kubernetes世界里的通用语

### 声明式与命令式是怎么回事

### 什么是YAML

### 什么是API对象

### 如何描述API对象

### 如何编写YAML



### 小结

声明式的语言能够更准确更清晰地描述系统状态，避免引入繁琐的操作步骤扰乱系统，与Kubernetes高度自动化的内部结构相得益彰，而且纯文本形式的YAML也很容易版本化，适合CI/CD。

内容要点：

1. YAML是JSON的超集，支持数组和对象，能够描述复杂的状态，可读性也很好。
2. Kubernetes把集群里的一切资源都定义为API对象，通过RESTful接口来管理。描述API对象需要使用YAML语言，必须的字段是**apiVersion、kind、metadata**。
3. 命令 `kubectl api-resources` 可以查看对象的apiVersion和kind，命令 `kubectl explain` 可以查看对象字段的说明文档。
4. 命令 `kubectl apply`、`kubectl delete` 发送HTTP请求，管理API对象。
5. 使用参数 `--dry-run=client -o yaml` 可以生成对象的YAML模板，简化编写工作。





## 12 Pod：如何理解这个Kubernetes里最核心的概念？

### 为什么要有Pod



### 为什么Pod是Kubernetes的核心对象

![](images/9ebab7d513a211a926dd69f7535ac175.png)

![](images/image-20240721234815253.png)



### 如何使用YAML描述Pod





### 如何使用kubectl操作Pod





### 小结

Pod屏蔽了容器的一些底层细节，同时又具有足够的控制管理能力，比起容器的“细粒度”、虚拟机的“粗粒度”，Pod可以说是“中粒度”，灵活又轻便，非常适合在云计算领域作为应用调度的基本单元，因而成为了Kubernetes世界里构建一切业务的“原子”。

知识要点：

1. 现实中经常会有多个进程密切协作才能完成任务的应用，而仅使用容器很难描述这种关系，所以就出现了Pod，它“打包”一个或多个容器，保证里面的进程能够被整体调度。
2. Pod是Kubernetes管理应用的最小单位，其他的所有概念都是从Pod衍生出来的。
3. Pod也应该使用YAML“声明式”描述，关键字段是“spec.containers”，列出名字、镜像、端口等要素，定义内部的容器运行状态。
4. 操作Pod的命令很多与Docker类似，如 `kubectl run`、`kubectl cp`、`kubectl exec` 等，但有的命令有些小差异，使用的时候需要注意。



在Kubernetes里通常并不会直接创建Pod，因为它只是对容器做了简单的包装，比较脆弱，离复杂的业务需求还有些距离，需要Job、CronJob、Deployment等其他对象增添更多的功能才能投入生产使用。



## 13 Job_CronJob：为什么不直接用Pod来处理业务？

### 为什么不直接使用Pod

### 为什么要有Job/CronJob

#### 如何使用YAML描述Job

![](images/image-20250107150724162.png)

#### 如何在Kubernetes里操作Job





### 如何使用YAML描述CronJob

![](images/image-20250107150906084.png)



### 小结



1. Pod是Kubernetes的最小调度单元，但为了保持它的独立性，不应该向它添加多余的功能。
2. Kubernetes为离线业务提供了Job和CronJob两种API对象，分别处理“临时任务”和“定时任务”。
3. Job的关键字段是 `spec.template`，里面定义了用来运行业务的Pod模板，其他的重要字段有 `completions`、`parallelism` 等
4. CronJob的关键字段是 `spec.jobTemplate` 和 `spec.schedule`，分别定义了Job模板和定时运行的规则。



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



### 小结

Deployment，它表示的是在线业务，和Job/CronJob的结构类似，也包装了Pod对象，通过添加额外的控制功能实现了应用永不宕机。



1. Pod只能管理容器，不能管理自身，所以就出现了Deployment，由它来管理Pod。
2. Deployment里有三个关键字段，其中的template和Job一样，定义了要运行的Pod模板。
3. replicas字段定义了Pod的“期望数量”，Kubernetes会自动维护Pod数量到正常水平。
4. selector字段定义了基于labels筛选Pod的规则，它必须与template里Pod的labels一致。
5. 创建Deployment使用命令 `kubectl apply`，应用的扩容、缩容使用命令 `kubectl scale`。







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
