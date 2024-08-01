# channel

```golang
chan T     // 可以接收和发送类型为 T 的数据
chan<- T   // 只可以用来发送 T 类型的数据
<-chan T   // 只可以用来接收 T 类型的数据
```