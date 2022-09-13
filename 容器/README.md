## 容器

容器是镜像的运行实体，镜像是静态的只读文件，而容器带有运行时需要的可写文件层，
并且容器中的进程属于运行状态。 容器运行着真正的应用进程，有初建、运行、停止、暂停和删除五种状态。
容器本质上是主机上运行的一个进程，但是拥有自己独立的命名空间隔离和资源限制。

镜像包含了容器运行所需要的文件系统结构和内容，是静态的只读文件，而容器则是在镜像的只读层上创建可写层，
并且容器中的进程属于运行状态，容器是真正的应用载体。

* Namespace：隔离进程ID、主机名、用户ID、文件名、网络访问和进程间通信等相关资源
    - pid namespace:用于隔离进程ID。
    - net namespace:隔离网络接口，在虚拟的net namespace内用户可以拥有自己独立的IP、路由、端口等。
    - mnt namespace:文件系统挂载点隔离。
    - ipc namespace:信号量，消息队列和共享内存的隔离。
    - uts namespace:主机名和域名的隔离。
* Cgroups：对进程或者进程组做资源（例如：CPU、内存等）的限制
* 联合文件系统（unionFS）：用于镜像构建和容器运行环境

### docker基础操作命令

* docker pull 镜像名 （先从本地搜索，搜索不到则从Docker Hub下载镜像）
* docker images （docker image ls）查看本地所有的镜像
* docker tag 原镜像：tag 镜像名：tag （重命名，image ID一样）
* docker rmi 镜像名
* 构建镜像
    - docker commit 镜像 镜像名称 从运行镜像提交镜像
    - docker build从Dockerfile构建镜像

Docker镜像是静态的分层管理的文件组合，镜像底层的实现依赖于联合文件系统。