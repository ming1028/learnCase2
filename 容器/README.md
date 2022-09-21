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

### 命令

#### CMD和ENTRYPOINT：都是容器运行的命令入口

相同之处：

- CMD/ENTRYPOINT["command", "param"]。使用Linux的exec实现的，称为exec模式。
- CMD/ENTRYPOINT command param，基于shell实现的，称为shell模式。会以/bin/sh -c command方式执行命令。

区别：

- Dockerfile 中如果使用了ENTRYPOINT指令，启动 Docker 容器时需要使用 --entrypoint参数才能覆盖,
  Dockerfile 中的ENTRYPOINT指令 ，而使用CMD设置的命令则可以被docker run后面的参数直接覆盖。
- ENTRYPOINT指令可以结合CMD指令使用，也可以单独使用，而CMD指令只能单独使用。

#### ADD 和 COPY

COPY指令只支持基本的文件和文件夹拷贝功能，ADD则支持更多文件来源类型，比如自动提取 tar 包，
并且可以支持源文件为 URL 格式。

### 虚拟机、Docker区别

- 虚拟机通过管理系统（hypervisor）模拟出CPU、内存、网络等硬件，虚拟机有自己的内核和操作系统，隔离性和安全性更好。
- Docker容器则是通过Linux内核的Namespace技术实现了文件系统、进程、设备以及网络的隔离，然后再通过Cgroups对CPU、
  内存资源限制，实现容器之间相互不影响。容器隔离性仅仅依靠内核提供，隔离性方面弱于虚拟机。

### Namespace

实现在同一主机系统中对进程ID(PID)、主机名(UTS)、用户ID(User)、文件名(MOUNT)、网络(Net)和进程间通信(IPC)等资源的隔离。

### cgroups

- 资源限制：限制资源的使用量
- 优先级控制：不同组可以有不同的资源使用优先级。
- 审计：计算控制组的资源使用情况。
- 控制：控制进程的挂起或恢复。