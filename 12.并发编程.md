## 12.并发编程

### 12.1 并发与并行

Erlang 之父 Joe Armstrong曾经以下图解释并发与并行。

![](./img/cor.jpg)

并发在图中的解释是两队人排队接咖啡，两队切换。

并行是两个咖啡机，两队人同时接咖啡。

“Concurrency is about dealing with lots of things at once. Parallelism is about doing lots of things at once.” — Rob Pike

并发使并行变得容易,并发提供了一种构造解决方案的方法,并行一般伴随这多核。并发一般伴随这CPU切换轮训。

### 12.2 为什么需要并发？

原因有很多，其中比较重要的原因如下：
1) 不阻塞等待其他任务的执行，从而浪费时间，影响系统性能。
2) 并行可以使系统变得简单些，将复杂的大任务切换成许多小任务执行，单独测试。

在开发中，经常会遇到为什么某些进程通常会相互等待呢？为什么有些运行慢，有些快呢？

通常受限来源于**进程I/O**或**CPU**。

- 进程I/O限制

如：等待网络或磁盘访问
- CPU限制

如：大量计算

### 12.3 Go并发原语

#### 12.3.1 协程Goroutines

每个go程序至少都有一个Goroutine：主Goroutine（在运行进程时自动创建）。以及程序中其他Goroutine
例如：下面程序创建了main的Goroutine及匿名的Goroutine。

```go
func main() {
	go func() {
		fmt.Println("you forgot me !")
	}()
}
```

在go中有个package是sync，里面包含了：


WaitGroup、Mutex、Cond、Once、Pool，下面依次介绍。

**1.WaitGroup**


假设主线程要等待其余的goroutine都运行完毕，不得不在末尾添加time.Sleep()，但是这样会引发两个问题：
- 等待多长时间？
- 时间太长，影响性能？

在go的sync库中的WaitGroup可以帮助我们完成此项工作，Add(n)把计数器设置为n,Done()会将计数器每次减1，Wait()函数会阻塞代码运行，直到计数器减0。


等待多个goroutine完成，可以使用一个等待组。 例如：

```go
// 这是我们将在每个goroutine中运行的函数。
// 注意，等待组必须通过指针传递给函数。
func worker(id int, wg *sync.WaitGroup) {

	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
```

这里首先把wg 计数设置为1， 每个for循环运行完毕都把计数器减一，主函数中使用Wait() 一直阻塞，直到wg为1——也就是所有的5个for循环都运行完毕。


使用注意点：
- 计数器不能为负值
- WaitGroup对象不是引用类型

**2.Once**

sync.Once可以控制函数只能被调用一次，不能多次重复调用。

例如：

```go
var doOnce sync.Once

func main() {
	DoSomething()
	DoSomething()
}

func DoSomething() {
	doOnce.Do(func() {
		fmt.Println("Run once - first time, loading...")
	})
	fmt.Println("Run this every time")
}
```

输出:
```
Run once - first time, loading...
Run this every time
Run this every tim
```

**3.互斥锁Mutex**


互斥锁是并发程序对共享资源进行访问控制的主要手段,在go中的sync中提供了Mutex的支持。

例如：使用互斥锁解决多个Goroutine访问同一变量。

```go
// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
  c.mux.Lock()
  defer c.mux.Unlock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
  c.v[key]++
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
```

在这个例子中，我们使用了sync.Mutex的Lock与Unlock方法。


在前面例子中我们使用了sync.Mutex，读操作与写操作都会被阻塞。其实读操作的时候我们是不需要进行阻塞的，因此sync中还有另一个锁：读写锁RWMutex,这是一个单写多读模型。

sync.RWMutex分为：读、写锁。在读锁占用下，会阻止写，但不会阻止读，多个goroutine可以同时获取读锁，调用RLock()函数即可，RUnlock()函数释放。写锁会阻止任何goroutine进来，整个锁被当前goroutine，此时等价于Mutex,写锁调用Lock启用，通过UnLock()释放。

例如： 我们对上述例子进行改写，读的时候用读锁，写的时候用写锁。


```cpp
// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v     map[string]int
	rwmux sync.RWMutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	// 写操作使用写锁
	c.rwmux.Lock()
	defer c.rwmux.Unlock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
  // 读的时候加读锁
	c.rwmux.RLock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.rwmux.RUnlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)

	for i := 0; i < 10; i++ {
		fmt.Println(c.Value("somekey"))
	}
}
```

**4.条件变量Cond**


sync.Cond是条件变量，它可以让一系列的 Goroutine 都在满足特定条件时被唤醒。

条件变量通常与互斥锁一起使用，条件变量可以在共享资源的状态变化时通知相关协程。
经常使用的函数如下：
- NewCond

创建一个Cond的条件变量。
```go
func NewCond(l Locker) *Cond
```

- Broadcast

广播通知，调用时可以加锁，也可以不加。

```go
func (c *Cond) Broadcast()
```
- Signal

单播通知，只唤醒一个等待c的goroutine。

```go
func (c *Cond) Signal()
```
- Wait
等待通知, Wait()会自动释放c.L，并挂起调用者的goroutine。之后恢复执行，Wait()会在返回时对c.L加锁。

除非被Signal或者Broadcast唤醒，否则Wait()不会返回。

```go
func (c *Cond) Wait()
```

例如：使用WaitGroup等待两个Goroutine完成，
Goroutine1与Goroutine2进入Wait状态，main函数在2s后改变共享数据状态，调用Broadcast函数，此时c.Wait从中恢复并判断条件变量是否已经满足，满足后消费条件，解锁，wg.Done()。

**5.原子操作**

原子操作即是进行过程中不能被中断的操作。针对某个值的原子操作在被进行的过程中，CPU绝不会再去进行其他的针对该值的操作。
为了实现这样的严谨性，原子操作仅会由一个独立的CPU指令代表和完成。


在sync/atomic 中，提供了一些原子操作，包括加法（Add）、比较并交换（Compare And Swap，简称 CAS）、加载（Load）、存储（Store）和交换（Swap）。

1.加法操作
提供了32/64位有符号与无符号加减操作

```go
var i int64
atomic.AddInt64(&i, 1)
fmt.Println("i = i + 1 =", i)
atomic.AddInt64(&i, -1)
fmt.Println("i = i - 1 =", i
```

2.比较并交换

CAS: Compare And Swap

如果addr和old相同,就用new代替addr。

```go
func CompareAndSwapInt32(addr *int32, old, new int32) (swapped bool)
```
例如：
```go
var a int32 = 1
var b int32 = 2
var c int32 = 3
ok := atomic.CompareAndSwapInt32(&a, a, b)
fmt.Printf("ok = %v, a = %v, b = %v\n", ok, a, b)
ok = atomic.CompareAndSwapInt32(&a, c, b)
fmt.Printf("ok = %v, a = %v, b = %v, c=%v\n", ok, a, b, c)
```
输出：
```
ok = true, a = 2, b = 2
ok = false, a = 2, b = 2, c = 3
```
3.交换

不管旧值与新值是否相等，都会通过新值替换旧值，返回的值是旧值。

```go
func SwapInt32(addr *int32, new int32) (old int32)
```

例如：
```go
var x int32 = 1
var y int32 = 2
old := atomic.SwapInt32(&x, y)
fmt.Println(x, old)
```
输出：2 1


3.加载

当读取该指针指向的值时，CPU 不会执行任何其它针对此值的读写操作
```go
func LoadInt32(addr *int32) (val int32)
```

例如：
```go
var x1 int32 = 1
y1 := atomic.LoadInt32(&x)
fmt.Println("x1, y1:", x1, y1)
```

4.存储

加载逆向操作。

例如：
```go
var xx int32 = 1
var yy int32 = 2
atomic.StoreInt32(&yy, atomic.LoadInt32(&xx))
fmt.Println(xx, yy)
```

5.原子类型

sync/atomic中添加了一个新的类型Value。
例如：
```go
v := atomic.Value{}
v.Store(1)
fmt.Println(v.Load())
```

**6.临时对象池Pool**


sync.Pool 可以作为临时对象的保存和复用的集合


P是Goroutine中的重要组成之一，例如：P实际上在操作时会为它的每一个goroutine相关的P生成一个本地P。 本地池没有，则会从其它的 P 本地池获取，或者全局P取。


sync.Pool对于需要重复分配、回收内存的地方，sync.Pool 是一个很好的选择。减少GC负担,如果Pool中有对象，下次直接取，不断服用对象内存，减轻 GC 的压力，提升系统的性能。

例如：

```go
var pool *sync.Pool

type Foo struct {
	Name string
}

func Init() {
	pool = &sync.Pool{
		New: func() interface{} {
			return new(Foo)
		},
	}
}

func main() {
	fmt.Println("Init p")
	Init()

	p := pool.Get().(*Foo)
	fmt.Println("第一次取：", p)
	p.Name = "bob"
	pool.Put(p)

	fmt.Println("池子有对象了，调用获取", pool.Get().(*Foo))
	fmt.Println("池子空了", pool.Get().(*Foo))
}
```

输出：

```
Init p
第一次取： &{}
池子有对象了，调用获取 &{bob}
池子空了 &{}
```

#### 12.3.2 通道Channel

**1.Channel**

这里引入一下CSP模型，CSP 是 Communicating Sequential Process 的简称，中文可以叫做通信顺序进程，是一种并发编程模型，由 [Tony Hoare](https://en.wikipedia.org/wiki/Tony_Hoare) 于 1977 年提出。

简单来说是实体之间通过发送消息进行通信，这里发送消息时使用的就是通道，或者叫 Channel。Goroutine对应并发实体。

**1) 使用**

Channel的使用需要通过make创建，例如：

```go
unBufferChan := make(chan int) 
bufferChan := make(chan int, x) 
```

上述创建了无缓冲的Channel与有缓冲的Channel，创建完成之后，需要进行读写操作，如下：

```go
ch := make(chan int, 1)

// 读操作
x <- ch

// 写操作
ch <- x
```

最终要正确关闭，只需要调用close即可。

```go
// 关闭
close(ch)
```

当channel关闭后会引发下面相关问题：

- 重复关闭Channel 会 panic
- 向关闭的Channel发数据 会 Panic，读关闭的Channel不会Panic，但读取的是默认值

对于最后一点读操作默认值怎么区分呢？例如：Channel本身的值是默认值又或者是读到的是关闭后的默认值，可以通过下面进行区分：

```go
val, ok := <-ch
if ok == false {
    // channel closed
}
```

**2) Channel分类**

- 无缓冲的Channel

发送与接受同时进行。如果没有Goroutine读取Channel(<-Channel)，发送者(Channel<-x)会一直阻塞。

![](./img/unbufferedchannel.png)

- 有缓冲的Channel

发送与接受并非同时进行。当队列为空，接受者阻塞;队列满，发送者阻塞。

![](./img/bufferedchannel.png)

**2.Select**

- 每个case 都必须是一个通信
- 所有channel表达式都会被求值
- 如果没有default语句，select将阻塞，直到某个通信可以运行
- 如果多个case都可以运行，select会随机选择一个执行

**1) 随机选择**

select特性之一：随机选择，下面会随机打印不同的case结果。
例如：
```go
ch := make(chan int, 1)
ch <- 1
select {
case <-ch:
	fmt.Println("ch 1")
case <-ch:
	fmt.Println("ch 2")
default:
	fmt.Println("ch default")
}
```
假设chan中没有值，有可能引发死锁。

例如： 下面执行后会引发死锁。

```go
ch := make(chan int, 1)
select {
case <-ch:
	fmt.Println("ch 1")
case <-ch:
	fmt.Println("ch 2")
}
```
此时可以加上default即可解决。
```go
default:
	fmt.Println("ch default")
```

另外，还可以添加超时。

```go
timeout := make(chan bool, 1)
go func() {
	time.Sleep(2 * time.Second)
	timeout <- true
}()
ch := make(chan int, 1)

select {
case <-ch:
	fmt.Println("ch 1")
case <-timeout:
	fmt.Println("timeout 1")
case <-time.After(time.Second * 1):
	fmt.Println("timeout 2")
}
```

**2) 检查chan**

select+defaul方式来确保channel是否满

```go
ch := make(chan int, 1)
ch <- 1
select {
case ch <- 1:
	fmt.Println("channel value is ", <-ch)
	fmt.Println("channel value is ", <-ch)
default:
	fmt.Println("channel blocking")
}
```
如果要调整channel大小，可以在make的时候改变size，这样就可以在case中往channel继续写数据。

**3) 选择循环**


当多个channel需要读取数据的时候，就必须使用 for+select

例如：下面例子需要从两个channel中读取数据，当从channel1中数据读取完毕后，会像signal channel中输入stop，此时终止for+select。
```go
func f1(c chan int, s chan string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		c <- i
	}
	s <- "stop"
}

func f2(c chan int, s chan string) {
	for i := 20; i >= 0; i-- {
		time.Sleep(time.Second)
		c <- i
	}
	s <- "stop"
}

func main() {
	c1 := make(chan int)
	c2 := make(chan int)
	signal := make(chan string, 10)

	go f1(c1, signal)
	go f2(c2, signal)
LOOP:
	for {
		select {
		case data := <-c1:
			fmt.Println("c1 data is ", data)
		case data := <-c2:
			fmt.Println("c2 data is ", data)
		case data := <-signal:
			fmt.Println("signal is ", data)
			break LOOP
		}
	}
}

```