[TOC]



# Go基础

## 0. 写在前面

###  0.1 Go语言特性

- 自动垃圾回收
- 更丰富的内置类型
- 函数多返回值
- 错误处理
- 匿名函数和闭包
- 类型和接口
- 并发编程
- 反射
- 语言交互性
  
### 0.2 语言结构
- 包声明
- 引入包
- 函数
- 变量
- 语句 & 表达式
- 注释

```go
package main

import "fmt"
func main() {
   /* Always Hello, World! */
   fmt.Println("Hello, World!")
}
```
解释：
1. package main定义了包名。**必须**在源文件中非注释的第一行指明这个文件属于哪个包。package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
2. import "fmt"告诉编译器程序运行需要用fmt包。
3. func main() 是程序开始执行的函数，main 函数是每一个可执行程序所必须包含的，一般来说都是在启动后第一个执行的函数（如果有 init() 函数则会先执行该函数）。
4. {}中"{"不可以单独放一行。
5. /*...*/ 是注释，在程序执行时将被忽略。//单行注释， /* ... */ 多行注释也叫块注释,不可以嵌套使用，一般用于包的文档描述或注释成块的代码片段。
6. fmt.Println(...) 将字符串输出到控制台，并在最后自动增加换行字符 \n。用 fmt.Print("hello, world\n") 可以得到相同的结果。


## 1.数据类型、关键字、标识符
### 1.1 数据类型
#### 1.1.1 按类别
- 布尔型：只可以是常量 true 或者 false。
```go
eg:
var b bool = true
```
- 数字类型：整型和浮点型。
- 位的运算采用补码字符串类型：字符串就是一串固定长度的字符连接起来的字符序列，Go 的字符串是由单个字节连接起来。
- Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本
- 复数：complex128（64 位实数和虚数）和 complex64（32 位实数和虚数），其中 complex128 为复数的默认类型。
	注：
	1. 复数的值由三部分组成 RE + IMi，其中 RE 是实数部分，IM 是虚数部分，RE 和 IM 均为 float 类型，而最后的 i 是虚数单位。
    ```go
    var name complex128 = complex(x, y)
    或者
    z := complex(x, y)
    x = real(z) 
    y = imag(z)
    ```
	2. 复数也可以用==和!=进行相等比较，只有两个复数的实部和虚部都相等的时候它们才是相等的


#### 1.1.2 派生类型
- 指针类型（Pointer）
- 数组类型
- 结构化类型(struct)
-  Channel 类型
- 函数类型
- 切片类型
- 接口类型（interface）
-  Map 类型

#### 1.1.3 基于架构
1整型，同时提供了四种有符号整型，分别对应8、16、32、64bit（二进制）的有符号整数，与此对应四种无符号的整数类型
- Uint8无符号 8 位整型 (0 到 255)
- Unit16
- Unit32
- Unit64
- int8
- int16
- int32
- int64

2浮点型：
- float32
- float64
- complex64(实数虚数)
- complex128

3其他：
- byte
- rune
- uint
- int
- uintptr(无符号整型，存放一个指针)
    
注：
1. 表示 Unicode 字符的 rune 类型和 int32 类型是等价的，通常用于表示一个 Unicode 码点，是等价的。
2. byte 和 uint8 也是等价类型，byte 类型一般用于强调数值是一个原始的数据而不是一个小的整数。
3. 无符号的整数类型 uintptr，它没有指定具体的 bit 大小但是足以容纳指针。只有在底层编程时才需要，特别是Go语言和C语言函数库或操作系统接口相交互的地方。
4. 有符号整数采用 2 的补码形式表示，也就是最高 bit 位用来表示符号位，一个 n-bit 的有符号数的取值范围是从 -2(n-1) 到 2(n-1)-1。无符号整数的所有 bit 位都用于表示非负数，取值范围是 0 到 2n-1。
5. 常量 math.MaxFloat32 表示 float32 能取到的最大数值，大约是 3.4e38。
6. 常量 math.MaxFloat64 表示 float64 能取到的最大数值，大约是 1.8e308。
7. float32 和 float64 能表示的最小值分别为 1.4e-45 和 4.9e-324。
8. 浮点数在声明的时候可以只写整数部分或者小数部分。
```go
const e = .71828 // 0.71828
const f = 1.     // 1
```
9. 很小或很大的数最好用科学计数法书写，通过 e 或 E 来指定指数部分
```go 
const Avogadro = 6.02214129e23  // 阿伏伽德罗常数
const Planck   = 6.62606957e-34 // 普朗克常数
```

### 1.2 关键字
#### 1.2.1 25个关键字或保留字
break default func interface select
case defer go map struct
chan else goto package switch
const fallthrough if range type
continue for import return var
#### 1.2.2 36 个预定义标识符
append bool byte cap close complex complex64 complex128 uint16
copy false float32 float64 imag int int8 int16 uint32
int32 int64 iota len make new nil panic uint64
print println real recover string true uint uint8 uintptr
#### 1.2.3 知识点
- 程序一般由关键字、常量、变量、运算符、类型和函数组成。
- 程序中可能会使用到这些分隔符：括号 ()，中括号 [] 和大括号 {}。
- 程序中可能会使用到这些标点符号：.、,、;、: 和 …。


### 1.3 标识符
标识符用来命名变量、类型等程序实体。一个标识符实际上就是一个或是多个字母(A~ Z和a~ z)数字(0~9)、下划线“_”组成的序列，但是第一个字符必须是字母或下划线而不能是数字。

## 2. 变量 、常量、枚举 

### 2.1 变量
变量，计算机语言能存储计算结果或表示值的抽象概念。可以通过变量名访问，变量名由字母、数字、下划线组成，其中首个字符不能为数字。

声明变量的一般形式是使用 var 关键字：
```go
var identifier type
var identifier1, identifier2 type
```
变量声明**方式**：
1. 指定变量类型，若没有初始化，数值类型（包括complex64/128）默认零值，bool默认false，字符串默认“”，“var a *int、var a []int、var a map[string] int、var a chan int、var a func(string) int、var a error // error 是接口”默认nil
2. 可根据值自行判断类型
3. “：=”声明，省略 var, 注意 := 左侧必须声明新的变量，否则产生编译错误，格式：v_name := value
4. 多变量声明：
```go
//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断
vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不能是已经被声明过的，否则会导致编译错误

// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)
```
注意:
1. "：=" 赋值操作符,高效创建新变量，初始化声明：a := 50 或 b := false，a 和 b 的类型（int 和 bool）将由编译器自动推断。
2. 这是使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。
3. 在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明，但可以赋值；
4. 声明了一个局部变量却没有在相同的代码块中使用它，同样会得到编译错误
5. 全局变量可以声明但不用。
6. _ 实际上是一个只写变量，你不能得到它的值。这样做是因为 Go 语言中必须使用所有被声明的变量，但有时你并不需要使用从一个函数得到的所有返回值。


### 2.2 常量
常量是一个简单值的标识符，在程序运行时，不会被修改的量。数据类型只可以是**布尔型、数字型（整数型、浮点型和复数）和字符串型**。
常量的定义格式：（省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。）
```go
const identifier [type] = value
const b = "abc"
```
多个相同类型的声明可以简写为：
```go
const c_name1, c_name2 = value1, value2
```

常用于枚举:
```go
const (
    Unknown = 0
    Female = 1
    Male = 2
)
0,1,2 代表未知、女、男
```
常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过。

**iota**，特殊常量，可认为是可以被编译器修改的常量。在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加 1；
```
package main

import "fmt"
const (
    i=1<<iota
    j=3<<iota
    k
    l
)

func main() {
    fmt.Println("i=",i)
    fmt.Println("j=",j)
    fmt.Println("k=",k)
    fmt.Println("l=",l)
}
结果：
i= 1
j= 6
k= 12
l= 24
```
iota 表示从 0 开始自动加 1，所以 i=1<<0, j=3<<1（<< 表示左移的意思），即：i=1, j=6，这没问题，关键在 k 和 l，从输出结果看 k=3<<2，l=3<<3。

简单表述:
- i=1：左移 0 位,不变仍为 1;
- j=3：左移 1 位,变为二进制 110, 即 6;
- k=3：左移 2 位,变为二进制 1100, 即 12;
- l=3：左移 3 位,变为二进制 11000,即 24。

注：<<n==*(2^n)。


### 2.3枚举
枚举，将变量的值一一列举出来，变量只限于列举出来的值的范围内取值。Go语言中没有枚举这种数据类型的，但是可以使用const配合iota模式来实现
#### 2.3.1 普通枚举
```go
const (
	 a = 0
	 b = 1
	 c = 2
	 d = 3
)
```
#### 2.3.2 自增枚举
1. iota只能用于常量表达式
2. 它默认开始值是0，const中每增加一行加1,同行值相同
```go
const (
			a = iota  //0
			c         //1
			d         //2
		  )
const (
 			e,f = iota,iota     //e=0, f=0
			g   =iota           //g=1 
			)  
```
3. 若中间中断iota，必须显式恢复。
```go
const ( 
    a = iota    //0
    b           //1
    c = 100     //100
    d           //100
    e = iota    //4
)
```

## 3. 运算符、控制语句 

### 3.1 运算符
假定 A 值为 10，B 值为 20。
#### 3.1.1 算数运算符

运算符 | 描述 |实例
:---:|:---:|:---
+| 相加| A + B 输出结果 30
- |相减 |A - B 输出结果 -10
* |相乘 |A * B 输出结果 200
/ |相除 |B / A 输出结果 2
% |求余|B % A 输出结果 0
++ |自增 |A++ 输出结果 11
-- |自减| A-- 输出结果 9

#### 3.1.2 关系运算符

运算符 | 描述 
:---:|:---
== | 检查两个值是否相等，如果相等返回 True 否则返回 False。
!= |检查两个值是否不相等，如果不相等返回 True 否则返回 False。
> |检查左边值是否大于右边值，如果是返回 True 否则返回 False。
< |检查左边值是否小于右边值，如果是返回 True 否则返回 False。
>= |检查左边值是否大于等于右边值，如果是返回 True 否则返回 False。
<= |检查左边值是否小于等于右边值，如果是返回 True 否则返回 False。

#### 3.1.3 逻辑运算符

运算符 | 描述 
:---:|:---
&& | 逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。
&#x007c; &#x007c;| 逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。
! | 逻辑 NOT 运算符。 如果条件为 True，则逻辑 NOT 条件 False，否则为 True。

#### 3.1.4 位运算符，假定 A 为60，B 为13

运算符 | 描述 
:---:|:---
& | 按位与运算符"&"是双目运算符。 其功能是参与运算的两数各对应的二进位相与。
&#x007c; | 按位或运算符"&#x007c;"是双目运算符。 其功能是参与运算的两数各对应的二进位相或
^ | 按位异或运算符"^"是双目运算符。 其功能是参与运算的两数各对应的二进位相异或，当两对应的二进位相异时，结果为1。
<<|左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。 其功能把"<<"左边的运算数的各二进位全部左移若干位，由"<<"右边的数指定移动的位数，高位丢弃，低位补0。
>>|右移运算符">>"是双目运算符。右移n位就是除以2的n次方。 其功能是把">>"左边的运算数的各二进位全部右移若干位，">>"右边的数指定移动的位数。

#### 3.1.5 赋值运算符

运算符 | 描述 |实例
:----:|:---|:---
= | 简单的赋值运算符，将一个表达式的值赋给一个左值|C = A + B 将 A + B 表达式结果赋值给 C
+= | 相加后再赋值|C += A 等于 C = C + A 
-= | 相减后再赋值 |C -= A 等于 C = C - A
*= |相乘后再赋值 |C *= A 等于 C = C * A
/= |相除后再赋值| C /= A 等于 C = C / A
%= | 求余后再赋值| C %= A 等于 C = C % A
<<= |左移后赋值 |C <<= 2 等于 C = C << 2
>>= |右移后赋值 |C >>= 2 等于 C = C >> 2
&= |按位与后赋值|C &= 2 等于 C = C & 2
^= |按位异或后赋值|C ^= 2 等于 C = C ^ 2
|= |按位或后赋值 C |= 2 等于 C = C | 2

#### 3.1.6 其他运算符

运算符 | 描述 |实例
:----:|:----:|:----:
& |返回变量存储地址 |&a; 将给出变量的实际地址。
* |指针变量| *a; 是一个指针变量

#### 3.1.7 优先级

优先级 | 运算符
:----:|:----:
5|* / % << >> & &^
4| + - &#x007c; ^
3| == != < <= > >=
2| &&
1| &#x007c;&#x007c;|

### 3.2 控制语句
#### 3.2.1 条件语句
指定一个或多个条件，并通过测试条件是否为 true 来决定是否执行指定语句，并在条件为 false 的情况在执行另外的语句。
##### 3.2.1.1 if语句
- if 语句 由一个布尔表达式后紧跟一个或多个语句组成。
- if 语句 后可以使用可选的 else 语句, else 语句中的表达式在布尔表达式为 false 时执行。
- if 或 else if 语句中可嵌入一个或多个 if 或 else if 语句。
- 同各类主流语言，不多赘述。但注意，Go 没有三目运算符，所以不支持 ?: 形式的条件判断

##### 3.2.1.2 switch语句
- 用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上至下逐一测试，直到匹配为止。
- 匹配项后面不需要再加 break。
- switch 默认情况下 case 最后自带 break 语句，匹配成功后就不会执行其他 case，如果我们需要执行后面的 case，可以使用 fallthrough 。
- fallthrough:强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true
```go
switch x.(type){
    case type:
       statement(s);      
    case type:
       statement(s); 
    default: // 可选 
       statement(s);
}
```
解释：从第一个判断表达式为 true 的 case 开始执行，如果 case 带有 fallthrough，程序会继续执行下一条 case，且它不会去判断下一个 case 的表达式是否为 true。
1. 支持多条件匹配
2. 不同的 case 之间不使用 break 分隔，默认只会执行一个 case
3. 如果想要执行多个 case，需要使用 fallthrough 关键字，也可用 break 终止

##### 3.2.1.3 select语句

```go
select {
    case communication clause  :
       statement(s);      
    case communication clause  :
       statement(s);
    default : // 可选 
       statement(s);
}
```
- 每个 case 都必须是一个通信
- 所有 channel 表达式都会被求值
- 所有被发送的表达式都会被求值
- 如果任意某个通信可以进行，它就执行，其他被忽略。
- 如果有多个 case 都可以运行，Select 会随机公平地选出一个执行。其他不会执行。
否则：
    1. 如果有 default 子句，则执行该语句。
    2. 如果没有 default 子句，select 将阻塞，直到某个通信可以运行；Go 不会重新对 channel 或值进行求值。

#### 3.2.3 循环语句

##### 3.2.3.1 for循环
```go
for init; condition; post { } //for
for condition { } //while
for {}
init： 一般为赋值表达式，给控制变量赋初值；
condition： 关系表达式或逻辑表达式，循环控制条件；
post： 一般为赋值表达式，给控制变量增量或减量。
```
for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环：
```go
for key, value := range oldMap {
    newMap[key] = value
```
##### 3.2.3.2 循环嵌套
循环套循环，格式：
```go
for [condition |  ( init; condition; increment ) | Range]
{
   for [condition |  ( init; condition; increment ) | Range]
   {
      statement(s);
   }
   statement(s);
}
```
##### 3.2.3.2 循环控制语句
1. break语句：
    - 用于循环语句中跳出循环，并开始执行循环之后的语句。
    - break 在 switch（开关语句）中在执行一条 case 后跳出语句的作用。
    - 在多重循环中，可以用标号 label 标出想 break 的循环。
2. continue语句：跳过当前循环的剩余语句，然后继续进行下一轮循环。
3. goto：无条件转移到过程中指定行，与条件语句配合，实现条件转移、构成循环、跳出循环体等（不建议用，造成混乱）

