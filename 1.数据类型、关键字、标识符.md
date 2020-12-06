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
- Channel 类型
- 函数类型
- 切片类型
- 接口类型（interface）
- Map 类型

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

