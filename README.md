## 个人笔记

***

### goroutine

* 只有非休眠（non-sleepine）的Go协程才会被认为是可被调度的
* Go协程是以合作式（抢占式）调度来运作的
* 所有的Go协程都是匿名的
* goroutine是一种用户态线程，其调用栈内存被称为用户栈，也是从堆区分配
  ，分配释放都是编译器完成的。系统栈分配释放是由操作系统完成的。GMP模型中
  一个M对应一个系统栈，M上的多个goroutine会共享该系统栈。

***

### unsafe

#### unsafe.Pointer

* 任何类型的*T都可以转换为unsafe.Pointer unsafe.Pointer()
* unsafe.Pointer也可以转换为任何类型的*T (*T)(unsafe.Pointer)
* unsafe.Pointer可以转换为uintptr unintptr(unsafe.Pointer)
* uintptr也可以转换为unsafe.Pointer

unsafe.Pointer主要用于指针类型的转换，而且是各个指针类型转换的桥梁。
uintptr主要用于指针运算，尤其是通过偏移量定位不同的内存。

#### unsafe.Sizeof

返回一个类型所占用的内存大小，这个大小只与类型有关，和类型对应的变量存储的内容大小无关
指针占用8字节
***

### array

    长度、元素类型组成

* 只有数组内部元素类型和大小一致，才是同一类型
* 数组作为参数传递是值传递

***

#### slice

数组和切片都是连续的内存操作，通过索引可以快速找到元素存储的位置

```
type SliceHeader struct {
    Data uintptr // 指向存储切片元素的数组
    Len int // 切片的长度
    Cap int // 切片的容量
}
```

***

#### string和[]byte

Go语言通过先分配一个内存在复制内容的方式，实现string和[]byte之间的强制转换

```
type StringHeader struct {
    Data uintptr // 存放指向真实内容的指针
    Len int
}
```

*SliceHeader可以提供 *StringHeader所需的Data和Len字段，可以转换
*StringHeader不能转*SliceHeader 缺少Cap字段，需要自己补上

### 单元测试

* go文件必须以_test.go结尾
* 文件名_test.go
* 函数必须以Test开头，是可导出、公开的函数
* 测试函数的签名必须接受一个指向testing.T类型的指针，并且不能返回任何值
* 函数名：Test+要测试的函数名

***

### 基准测试

* 必须以Benchmark开头
* 函数的签名必须接受一个指向testing.B类型的指针，并且不返回任何值
* b.N是基准测试框架提供，表示循环次数

***

### 堆内存、栈内存

* 栈内存：由编译器自动分配和释放，一般存储函数中的局部变量、参数等，函数创建的时候，
  会被自动创建，返回的时候，会被自动释放。
* 堆内存：如果函数返回值还在其他地方被使用，那么这个值就会被编译器自动分配到堆上。只能
  通过垃圾回收器才能释放。

#### 逃逸

go build -gcflags "-m -l" -m打印逃逸分析，-l禁止内联优化

* 指针作为函数返回值的时候，一定会发生逃逸
* 被已经逃逸的指针引用的变量也会发生逃逸
* 被map、slice和chan这三种类型引用的指针一定会发生逃逸
* 逃逸分析是在编译器完成的

#### 逃逸情况

* 变量类型不确定

```
fmt.Println(a) // 参数为interface类型，编译器不能确定其参数的具体类型，所以分配到堆上
```

* 暴露给外部指针

```
func foo() *int {
  a := 44
  return &a
}
```

* 变量所占内存较大，（64KB）

```
func foo() {
  s := make([]int, 0, 10000)
  for i := 0; i < 10000; i++ {
    s = append(s, i)
  }
}
```

* 变量大小不确定

```
func foo() {
  n := 1
  s := make([]int, n) // 没有指定大小，保证内存绝对安全
  for i := 0; i < len(s); i++ {
    s[i] = i
  }
}
```

##### 优化技巧

* 栈内存效率更高，小对象的传参，array要比slice效果好
* 重用内存，使用sync.Pool
* 空间换时间
* 尽可能避免使用锁，锁粒度要小，使用stringBuilder做string和[]byte之间的转换，defer嵌套不要太多

***

### RESTful API

* POST、GET、HEAD、OPTIONS、PUT、DELETE、TRACE、PATCH、CONNECT
* GET,读取服务器上的资源
* POST,在服务器创建资源
* PUT,更新或者替换服务器上的资源
* DELETE,删除服务器上的资源
* PATCH,更新/修改资源的一部分

***

### 互斥锁

* 互斥锁两种模式：正常模式、饥饿模式，饥饿模式是为了优化正常模式下刚被唤起的goroutine
  于新创建的goroutine竞争时长时间获取不到锁，如果一个goroutine获取锁失败超过1ms，则会
  将Mutex切换为饥饿模式，如果一个goroutine获得了锁，并且他在等待队列队尾或者小于1ms，则
  会将Mutex的模式切换为正常模式
* 防止刚被唤醒的goroutine在和新到达的goroutine竞争不过锁，当刚被唤醒的goroutine超过1ms没有获取到锁，
  会将当前互斥锁切换到饥饿模式，防止被唤醒的goroutine被饿死；
    - 饥饿模式下，锁的所有权直接从解锁的goroutine转移到等待队列的中的队头，新来的goroutine不会尝试去获取锁，
      也不会自旋，就在等待队列的队尾。
    - 如果某个goroutine是等待队列的最后一个goroutine，或者等待获取锁的时间小于1ms，将从饥饿模式切换回正常模式。
* 加锁过程：
    - 锁处于完全空闲状态，通过CAS直接加锁
    - 当锁处于正常模式、加锁状态下，并且符合自旋条件，则会尝试最多4次的自旋
    - 若当前goroutine不满足自旋条件时，计算当前goroutine的锁期望状态
    - 尝试使用CAS更新锁状态，若更新锁状态成功判断当前goroutine是否可以获取到锁，
      获取到锁直接出去，若获取不到锁则陷入睡眠等待被唤醒
    - goroutine被唤醒后，如果锁处于饥饿模式，则直接拿到锁，否则重置自旋次数、标志
      唤醒位，重新走for循环自旋、获取锁逻辑；
* 解锁的过程
    - 原子操作mutexLocked，如果锁为完全空闲状态，直接解锁成功
    - 如果锁不是完全空闲状态，那么进入unlockedslow逻辑
    - 如果解锁一个未上锁的锁直接panic，因为没加锁mutexLocked的值为0，
      解锁时进行mutexLocked-1操作，这个操作会让整个互斥锁混乱，所以需要这个判断
    - 如果锁处于饥饿模式直接唤醒等待队列对头的waiter
    - 如果锁处于正常模式下，没有等待的goroutine可以直接退出，如果锁已经处于
      锁定状态、唤醒状态、饥饿模式则可以直接退出，因为所有被唤醒的goroutine获得了锁
* 使用互斥锁时切忌拷贝Mutex，因为拷贝Mutex时会连带状态一起拷贝，因为Lock时只有锁
  在完全空闲时才会获取锁成功，拷贝时连带状态一起拷贝后，会造成死锁
* TryLock的实现逻辑很简单，主要判断当前锁处于加锁状态、饥饿模式就会直接获取锁失败，
  尝试获取锁失败直接返回；
* 自旋是自旋锁的行为，它通过忙等待，让线程在某段时间内一直保持运行，从而避免线程上下文的调度开销，自旋锁对于线程
  只会阻塞很短时间的场景是非常合适。单核cpu不适合使用自旋锁，线程A获取不到锁，如果处于自旋状态，不挂起，其他持有锁
  的线程没有办法进入运行状态，只能等操作系统分给A的时间片用完，才有机会被调度。
  - 自旋条件：当前互斥锁处于正常模式、当前运行机器是多核CPU、至少存在一个正在运行的处理器P，并且本地运行队列为空、
    当前goroutine进行自旋次数小于4；

***

### RPC服务

* 远程过程调用，是分布式系统中不同节点调用的方式（进程间通信），属于C/S模式
* 核心：通信协议和序列化

### 深拷贝：

    值类型的数据，默认全部都是深复制，Array、Int、String、Struct、Float，Bool。

### 浅拷贝：

    拷贝的是数据地址，只复制指向的对象的指针，此时新对象和老对象指向的内存地址是一样的，
    新对象值修改时老对象也会变化。释放内存地址时，同时释放内存地址。引用类型的数据，默认全部都是浅复制，Slice，Map

### 内存对齐

* 当空结构体类型作为结构体的最后一个字段时，如果有指向该字段的指针，那么就会返回该结构体之外的地址。为了避免内存泄露会额外进行一次内存对齐。

***

### G M P

G：goroutine协程 M：thread线程 P：processor处理器

* 全局队列：存放等待运行的G
* P的本地队列：存放等待运行的G，数量不超过256个，新建G时，优先加入到P的本地队列，如果队列满了，会将本地
  队列中一半的G移到全局队列
* P：所有的P在程序启动时创建，并保存在数组中，最多有GOMAXPROCS个，包含了运行goroutine的资源
* M：线程想运行任务就需要获取P，从P的本地队列获取G，P队列为空时，M会尝试从全局队列拿一批G放到P的本地队列
  或者从其他P的本地队列拿一半放到自己P的本地队列。go语言程序启动时，会设置M的最大数量，默认10000,。
* 一个M阻塞，P就会创建或者切换另外一个M

goroutine调度器和OS调度器是通过M结合起来，每个M都代表了一个内核线程，OS调度器负责把内核线程分配到CPU的核上执行。

#### 调度器设计策略

* work stealing机制：当本线程没有可运行的G时，会尝试从其他线程绑定的P偷取G，而不是把线程销毁。
* hand off机制：当本线程因为G进行系统调用阻塞时，线程释放绑定的P，把P转移给其他空闲的线程执行
* 并行：最多有GOMAXPROCS个线程分布在多个CPU上同时运行
* 抢占：为防止其他goroutine被饿死，一个goroutine最多占用CPU10ms。
* 全局G队列，当M从其他P偷不到G时，就可以从全局G队列获取G

#### go func调度流程

* go func()创建一个goroutine
* 新创建的goroutine会先保存在调度器P的本地队列，如果P的本地队列已满就会保存在全局的队列中
* goroutine只能运行在M中，一个M必须持有一个P，是1：1的关系。M会从P的本地队列弹出
  一个可执行状态的G来执行，如果拿不到就从全局队列获取
* 一个M调度G执行的过程是一个循环机制
* 当M执行某一个G的时候如果发生了syscall（系统调用）等操作，M就会阻塞，如果当前正好有一些G
  在执行，runtime会把这个线程M从P中摘除，然后在创建一个新的操作系统线程（有空闲复用空闲线程）
  来服务这个P
* 当M系统调用结束时，这个G会尝试获取一个空闲的P执行，并放入到这个P的本地队列。如果获取不到P，那么
  这个线程M变成休眠状态，加入到空闲线程中，然后这个G回被放入到全局队列中。

#### 协程

* 线程分为"内核态"线程和"用户态"线程,一个"用户态线程"必须绑定一个"内核态线程"，CPU只知道运行的是
  一个"内核态线程"（linux的PCB进程控制块），内核线程依然叫"线程"（thread），用户线程叫"协程"（co-routine）
* 线程由CPU调度是抢占式的，协程由用户态调度是协作式的，一个协程让出CPU后，才执行下一个协程。

***

### 三色标记

* 新创建的对象，默认颜色都是标记为白色
* 从根节点开始遍历所有对象，把遍历到的对象从白色集合放入灰色集合，非递归形势，只遍历一次。
* 遍历灰色集合，将灰色对象引用的对象从白色集合放入灰色集合，之后将此灰色集合（放入前）放入黑色集合
* 重复遍历灰色集合，直到灰色中没有任何对象
* 回收所有的白色标记的对象，也就是垃圾回收

#### 标记过程不使用STW影响

* 一个白色对象被黑色对象使用（白色被挂在黑色下）
* 灰色对象与它之间的可达关系的白色对象遭到破坏（灰色丢了该白色）

#### 屏障机制

* "强-弱 三色不变式"
    + 强三色不变式
        - 不存在黑色对象引用到白色对象的指针。
    + 弱三色不变式
        - 所有被黑色对象引用的白色对象都处于灰色保护状态，也就是白色对象存在其他灰色对象对它的引用，或者可达它的
          链路上游存在灰色对象。
* 插入屏障

  黑色对象的内存槽有两种位置，栈和堆，栈空间容量小，调用弹出频繁，所以插入屏障在栈空间的对象操作中不使用，
  仅仅使用在堆空间对象的操作中。栈上有可能依然存在白色对象被引用的情况，所以要对栈重新进行三色标记扫描，
  为了对象不丢失，要对本次标记扫描启动STW暂停，直到栈空间的三色标记结束。

    + A对象引用B对象，B对象被标记为灰色，白色会强制变成灰色
* 删除屏障

  被删除的对象，如果自身为灰色或者白色，那么被标记为灰色。满足若三色不变式，保护灰色对象到白色对象的路径不会断。
  一个对象即使被删除了最后一个指向它的指针也依旧可以活过这一轮，在下一轮GC中被清理掉。

#### 1.8混合写屏障机制

* 插入写屏障和删除写屏障短板
    + 插入写屏障：结束时需要STW来重新扫描栈，标记栈上引用的白色对象的存活
    + 删除写屏障：回收精度低，GC开始时***STW扫描堆栈来记录初始快照***，这个过程会保护开始时刻的所有存活对象

* 混合写屏障规则
    + GC开始将栈上的对象全部扫描并标记为黑色（之后不再进行第二次重复扫描，无需STW）
    + GC期间，任何在栈上创建的新对象，均为黑色
    + 被删除的对象标记为灰色
    + 被添加的对象标记为灰色

***

### 进程、线程

* 进程：程序在操作系统中的一次执行过程，系统进行资源分配和调度的一个独立单位
* 线程是进程的一个执行实体，是CPU调度和分派的基本单位，能独立运行的基本单位
* 一个进程可以创建和撤销多个线程，同一个进程中的多个线程之间可以并发进行。

***

### channel

* 对一个关闭的通道在发送值就会导致panic
* 对一个关闭的通道进行接收会一直获取值直到通道为空
* 对一个关闭的并且没有值的通道执行接收会得到对应类型的零值
* 关闭一个已经关闭的通道会导致panic
* 判断channel是否关闭

```
for range // 通道关闭自动退出for range
_, ok := <-ch // ok为false
```

### 锁

* 互斥锁：能够保证同一时间有且只有一个goroutine进入临界区，其他的goroutine则在等待锁，多个goroutine
  等待同一个锁时，唤醒的策略是随机的

```
sync.Mutex // 互斥锁
```

* 读写互斥锁：读多写少场景，一个goroutine获取读锁之后，其他goroutine如果获取读锁会获得锁，获取写锁就会等待；
  一个goroutine获取写锁，其他goroutine无论读锁写锁都会等待

```
rwLock sync.RWMutex
rwLock.Lock // 写锁 rwLock.Unlock
rwLock.RLock // 读锁
rwLock.RUnlock //解读锁
```

### Kafka

* Topic: 用于区分不同类别信息的类别名称，由producer指定
* producer：将消息发布到Kafka特定的topic的对象
* Consumers: 订阅并处理特定的Topic中的消息的对象
* Broker(服务集群)：已发布的消息保存在一组服务器中，集群中的每一个服务器都是一个代理（broker）
  消费者可以订阅一个或者多个Topic，并从Broker拉数据，从而消费这些已发布的消息。
* Partition: Topic物理上的分组，一个topic可以分为多个partition，每个partition是一个有序的队列，
  partition中的每条消息都会被分配一个有序的id

#### 工作流程

* ⽣产者从Kafka集群获取分区leader信息
* ⽣产者将消息发送给leader
* leader将消息写入本地磁盘
* follower从leader拉取消息数据
* follower将消息写入本地磁盘后向leader发送ACK
* leader收到所有的follower的ACK之后向生产者发送ACK

#### 选择partition的原则

* partition在写入的时候可以指定需要写入的partition，如果有指定，则写入对应的partition。
* 如果没有指定partition，但是设置了数据的key，则会根据key的值hash出一个partition。
* 如果既没指定partition，又没有设置key，则会采用轮询⽅式，即每次取一小段时间的数据写入某
  个partition，下一小段的时间写入下一个partition

#### ACK应答机制

* 代表producer往集群发送数据不需要等到集群的返回，不确保消息发送成功。安全性最低但是效 率最高。
* 代表producer往集群发送数据只要leader应答就可以发送下一条，只确保leader发送成功。
* 代表producer往集群发送数据需要所有的follower都完成从leader的同步才会发送下一条，
  确保 leader发送成功和所有的副本都完成备份。安全性最⾼高，但是效率最低。

***

### 依赖管理

* go get -u升级到最新的次要版本或者修订版本
* go get -u=patch将会升级到最新的修订版本
* go get package@version升级到指定版本

***

### 加密

* 在密码学中，一个密钥只能加密长度等于密钥长度的数据，所以需要对数据进行合理的分组，数据长度
  不足密钥长度时，则需要使用合适的填充模式进行填充
* 16,24,32位字符串的话，分别对应AES-128，AES-192，AES-256 加密方法