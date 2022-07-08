## 个人笔记

***

### goroutine

* 只有非休眠（non-sleepine）的Go协程才会被认为是可被调度的；
* Go协程是以合作式（抢占式）调度来运作的；
* 所有的Go协程都是匿名的；

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
