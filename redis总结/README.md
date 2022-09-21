## redis

***

### string

key-value 结构，key 是唯一标识，value 是具体的值，value其实不仅是字符串，
也可以是数字（整数或浮点数），value 最多可以容纳的数据长度是 512M。

底层数据结构实现主要是int和SDS（简单动态字符串）

* SDS不仅可以保存文本数据，还可以保存二进制数据，通过len属性的值来判断字符串是否结束。

- 字符串内部编码：int、raw和embstr
- 如果字符串对象保存的是整数值，并且这个整数值可以用long类型来表示

```
type struct redisObject {
  type => redis-string
  encoding => int
  ptr => 数值(long)
}
```

- 如果字符串对象保存的是字符串，并且长度小于等于32字节(连续内存片)，字符串只读，修改时先将对象的编码从embstr
  转换成raw，然后执行修改命令。

```
type struct redisObject {
  type => redis-string
  encoding => embstr
  ptr => 指向SDS动态字符串
}
```

- 如果字符串对象保存的是字符串，并且长度大于32字节（两次内存分配）

```
type struct redisObject {
  type => redis-string
  encoding => raw
  ptr => 指向SDS动态字符串
}
```

***

### List

最大长度2^32 - 1,每个列表支持超过40亿个元素；底层数据结构是由双向链表或压缩列表实现的

* 列表元素个数小于512个，每个元素的值都小于64字节，会使用压缩列表作为list类型的的底层数据结构
* 否则会使用双向链表作为底层数据结构
* 3.2版本之后，底层数据结构只由quicklist实现，替代双向链表和压缩列表。
* 为了留存消息，List 类型提供了 BRPOPLPUSH 命令，这个命令的作用是让消费者程序从一个 List 中读取消息，
  同时，Redis 会把这个消息再插入到另一个 List（可以叫作备份 List）留存。

```
BRPOP key [key ...] timeout // 表尾弹出一个元素，没有就阻塞timeout秒，如果为0一直阻塞
```

### Hash

底层数据结构是由压缩列表或哈希表实现

- 如果哈希表元素个数小于512个，所有值小于64字节，会使用压缩列表作为Hash类型的底层数据结构；（redis7.0压缩列表被废弃使用listpack数据结构实现）
- 否则使用哈希表作为Hash作为底层数据结构
