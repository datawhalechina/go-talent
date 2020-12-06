## 7.结构体、方法、接口

### 7.1.结构体

Go 语言中没有“类”的概念，也不支持像继承这种面向对象的概念。但是Go 语言的结构体与“类”都是复合结构体，而且Go 语言中结构体的组合方式比面向对象具有更高的扩展性和灵活性。

#### 7.1.1 结构体定义

结构体一般定义如下：

```go
type identifier struct {
  field1 type1
  field2 type2
  ...
}
```

例如我们想声明一个学生的结构体类型：

```go
type Student struct {
	Name string
	Age int
}
```

结构体中字段的类型可以是任何类型，包括函数类型，接口类型，甚至结构体类型本身。例如我们声明一个链表中的节点的结构体类型。

```go
type ListNode struct {
  Val int
  Next *ListNode
}
```

在声明结构体时我们也可以不给字段指定名字，例如下面这样

```go
type Person struct {
	ID string
	int
}
```

我们可以看到其中有一个int字段没有名字，这种我们称其为**匿名字段**。

#### 7.1.2 操作结构体

声明完结构体之后我们需要创建结构体的实例，可以使用如下几种方法创建，仍然以上面的Student结构体为例。

```go
s1 := new(Student) //第一种方式
s2 := Student{"james", 35} //第二种方式
s3 := &Student { //第三种方式
	Name: "LeBron",
	Age:  36,
}
```

- 使用new函数会创建一个指向结构体类型的指针，创建过程中会自动为结构体分配内存，结构体中每个变量被赋予对应的零值。
- 也可以使用第二种方式生命结构类型，需要注意的是此时给结构体赋值的顺序需要与结构体字段声明的顺序一致。
- 第三种方式更为常用，我们创建结构体的同时显示的为结构体中每个字段进行赋值。

声明完结构体之后可以直接按如下方式操作结构体。

```go
s1.Name = "james"
s1.Age = 35
```

需要注意的是，结构体也仍然遵循可见性规则，要是定义结构体的字段时首字母为小写在其他包是不能直接访问该字段的。

如果我们将定义的结构体首字母也变为小写那么在其他包内就不能直接创建该结构体，你知道这种情况应该怎么处理么？

上面我们提到的匿名字段，可以使用如下方法对其进行操作。

```go
p := new(Person)
p.ID = "123"
p.int = 10
```

我们直接通过p.int的方式来访问结构体中的匿名字段对其赋值，通过这个例子也可以发现，对于一个结构体来说，每一种数据类型只能有一个匿名字段。

#### 7.1.3 标签

在go语言中结构体除了字段的名称和类型外还有一个可选的标签tag，标记的tag只有reflect包可以访问到，一般用于orm或者json的数据传递，下面这段代码演示了如何为结构体打标签。

```go
type Student struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```

我们可以使用go自带的json包将声明的结构体变量转变为json字符串。

```go
func ToJson(s *Student) (string, error) {
	bytes, err := json.Marshal(s)
	if err != nil {
		return "", nil
	}
	return string(bytes), nil
}
```

如果我们没有给结构体打标签输出的json字符串如下所示

```go
{"Name":"james","Age":35}
```

如果我们给结构体打过标签之后输出的json字符串如下所示

```go
{"name":"james","age":35}
```

#### 7.1.4 内嵌结构体

之前我们介绍到了匿名字段，结构体作为一种数据类型也可以将其生命为匿名字段，此时我们称其为内嵌结构体，下面这段代码中我们将结构体A嵌入到结构体B中。

```go
type A struct {
	X, Y int
}

type B struct {
	A
	Name string
}
```

通过内嵌结构体的方式我们可以在结构体B的变量下很方便的操作A中定义的字段。

```go
b := new(B)
b.X = 10
b.Y = 20
b.Name = "james"
```

可以看到在b中我们操作结构体A中定义的字段就像结构体B本身定义的字段一样自然。

但是如果存在字段的名称冲突我们该怎么办？例如我们声明如下一个结构体C。

```go
type C struct {
	A
  B
	X int
}
```

此时结构体C中也有字段X，但是内嵌的结构体A中也有字段X，如果我们使用如下这种赋值方式会将X的值赋给谁呢？你可以尝试一下

```go
c := new(C)
c.X = 10
c.Y = 11
```

如果上面结构体B也有字段X，那么程序还能成功运行么？

需要注意的是，内嵌结构体和声明一个结构体类型的字段是不同的，例如下面的结构体B的定义方式与上面是完全不同的。

```go
type B struct {
	a A
	Name string
}
```

你可以尝试一下在结构体中定义一些复杂类型例如切片，字典等是如何操作的。

### 7.2 方法

#### 7.2.1 方法定义

方法与函数类似，只不过在方法定义时会在func和方法名之间增加一个参数，如下所示：

```go
func (r Receiver)func_name(){
  // body
}
```

其中r被称为方法的接收者，例如我们下面这个例子：

```go
type Person struct {
	name string
}

func (p Person) GetName() string {
	return p.name
}
```

其中GetName方法的接收者为p是Person结构体类型，也就是说我们为结构体Person绑定了一个GetName方法，我们可以使用如下的方式进行调用。

```go
func main() {
	p := Person{
		name:"james",
	}
	fmt.Println(p.GetName())
}
```

#### 7.2.2 方法接收者

对于一个方法来说接收者分为两种类型：值接收者和指针接收者。上面的GetName的接收者就是值接收者。我们再为Person结构体定义一个指针接收者。

```go
func (p *Person)SetName(name string){
	p.name = name
}
```

使用值接收者定义的方法，在调用的时使用的其实是值接收者的一个拷贝，所以对该值的任何操作，都不会影响原来的类型变量。

但是如果使用指针接收者的话，在方法体内的修改就会影响原来的变量，因为指针传递的也是地址，但是是指针本身的地址，此时拷贝得到的指针还是指向原值的，所以对指针接收者操作的同时也会影响原来类型变量的值。

而且在go语言中还有一点比较特殊，我们使用值接收者定义的方法使用指针来调用也是可以的，反过来也是如此，如下所示：

```go
func main() {
	p := &Person{
		name: "james",
	}
	fmt.Println(p.GetName())

  p1 := Person{
		name: "james",
	}
	p1.SetName("kobe")
	fmt.Println(p1.GetName())
}
```

### 7.3 接口

#### 7.3.1 接口定义

接口相当于一种规范，它需要做的是谁想要实现我这个接口要做哪些内容，而不是怎么做。在go语言中接口的定义如下所示：

```go
type Namer interface {
    Method1(param_list) return_type
    Method2(param_list) return_type
    ...
}
```

#### 7.3.2 实现接口

在go语言中不需要显示的去实现接口，只要一个类型实现了该接口中定义的所有方法就是默认实现了该接口，而且允许多个类型都实现该接口，也允许一个类型实现多个接口。

案例如下：

```go
type Animal interface {
	Eat()
}

type Bird struct {
	Name string
}

func (b Bird) Eat() {
	fmt.Println(b.Name + "吃虫")
}

type Dog struct {
	Name string
}

func (d Dog) Eat() {
	fmt.Println(d.Name + "吃肉")
}

func EatWhat(a Animal) {
	a.Eat()
}

func main() {
	b := Bird{"Bird"}
	d := Dog{"Dog"}
	EatWhat(b)
	EatWhat(d)
}
```

在EatWaht函数中是传递一个Animal接口类型，上面的Bird和Dog结构体都实现了Animal接口，所以都可以传递到函数中去来实现多态特性。

但是还有几点需要大家去探索一下：

- 通过值接收者和指针接收者定义的方法，对于接口的实现有什么影响吗？
- 还记得我们之前说过的内嵌结构体么，如果嵌入的结构体实现了某个接口那么对于外部的结构体有什么影响吗？

#### 7.3.3 类型断言

有些时候方法传递进来的参数可能是一个接口类型，但是我们要继续判断是哪个具体的类型才能进行下一步操作，这时就用到了类型断言，下面我们通过一个例子来进行讲解：

```go
func IsDog(a Animal) bool {
	if v, ok := a.(Dog); ok {
		fmt.Println(v)
		return true
	}
	return false
}
```

上面的方法对传递进来的参数进行判断，判断其是否为Dog类型，如果是Dog类型的话就会将其进行转换为v，ok用来表示是否断言成功。

但是如果我们对于一个类型有好多种子类型要进行判断，这样写的话显然是有些复杂，可以使用如下这种方式：

```go
func WhatType(a Animal) {
	switch a.(type) {
	case Dog:
		fmt.Println("Dog")
	case Bird:
		fmt.Println("Bird")
	default:
		fmt.Println("error")
	}
}
```

#### 7.3.4 空接口

空接口是一个比较特殊的类型，因为其内部没有定义任何方法所以空接口可以表示任何一个类型，比如可以进行下面的操作：

```go
var any interface{}

any = 1
fmt.Println(any)

any = "hello"
fmt.Println(any)

any = false
fmt.Println(any)
```



