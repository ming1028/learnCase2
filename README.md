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

***

### RPC服务

* 远程过程调用，是分布式系统中不同节点调用的方式（进程间通信），属于C/S模式
* 核心：通信协议和序列化

### 深拷贝：

    值类型的数据，默认全部都是深复制，Array、Int、String、Struct、Float，Bool。

### 浅拷贝：

    拷贝的是数据地址，只复制指向的对象的指针，此时新对象和老对象指向的内存地址是一样的，
    新对象值修改时老对象也会变化。释放内存地址时，同时释放内存地址。引用类型的数据，默认全部都是浅复制，Slice，Map