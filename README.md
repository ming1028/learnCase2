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


