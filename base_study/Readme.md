## 0 环境变量

GOROOT: C:\Program Files\Go
Path下bin目录: %GOROOT%\bin -->GOBIN: go intstall生成的exe存放路径, go build是直接在当前目录生成exe
GOPATH:工作目录: F:\go_space	注:记得删除用户环境变量里的默认值C:\...
以上 %GOPATH% 目录约定有三个子目录：
src 存放源代码（比如：.go .c .h .s等）
pkg 编译后生成的归档文件-->静态库文件（比如：.a）
bin 编译后生成的可执行文件(exe)
（为了方便，可以把此目录加入到 windows的PATH 变量中，在环境变量path后追加%GOPATH%\bin）
最后可以go env查看环境变量

## 1 运行程序

(package main类似于c里面的using namespace std)
在src里创建hello文件夹,在里面创建helloworld.go
go run %GOPATH%\src\hello\helloworld.go
go install %GOPATH%\src\hello,在bin里会生成exe文件,这也是我们为什么将bin放入path的原因
这样之后可以直接使用hello.exe执行
//还有个是playground运行,需要借助play.golang.org,国内被墙了

## 2 go语言依赖

cd $GOPATH/src/golang.org/x //目录不存在时可以先创建 
git clone https://github.com/golang/tools 
cd $GOPATH/src 
go install -v golang.org/x/tools/cmd/guru // 安装工具，其他同理
注:gocode-gomod和godef-gomod没法安装!!因为前面有别的插件安装了gocode和godef,
其本身也是安装gocode和godef而不是*-gomod,只要将前面的改名,这个按上改成gocode和godef就可以了

> Go语言中的引用类型有：映射（map），切片（slice），通道（channel），方法与函数。

## 3 变量.

var a int = 3
//one, 不需要分号,定义格式有区别
var a =3
//two, 可以进行类型推导
var (
a=3
b=4
)
//three, 可以进行美观的批量初始化赋值,等同于:var a, b int = 3 ,4
//等同于:
a,b := 3 ,4 
//four, 可以进行简洁的声明
//注意,:=的左边变量,不能全部都是赋新值,必须有一个是刚初始化的!!!
//会抛出no new variables on left side of := 的编译时异常
a := 3
a = "nav"
//five, go是强类型,不允许不同类型的变量赋值或者计算
//会抛出cannot use "nav" (type string) as type int in assignment的编译时异常
注:type a int	-->给变量定义类型
const b = "hhh"		-->设定常量,可以根据本身赋值给合适的类型

## 4 类型

Go 支持的基本类型:
bool:true或者false
数字类型
	int8, int16, int32, int64, int(n位有符号:-2^(n-1)到2^(n-1)-1,如-128到127)
	uint8, uint16, uint32, uint64, uint(n位有符号:0到2^n-1,如0到256)
	float32, float64(浮点)
	complex64, complex128(复数:实部和虚部都是对应/2的浮点类型,声明c := 5 + 7i或者c:=complex(5, 7) )
	byte(uint8的别名,8位,如果是int8有符号的话,第一位是符号)
	rune(int32的别名)
string:连接还是用+
注:
go是强类型,没有自动类型转换什么的,如int和float直接加减就不行,
但是可以实现手动转换后计算.sum := i + int(j)

## 5 函数

func functionname(para1,para2 int) int{
}
参考变量设定的规则,类型放在后面(这是go的一个特色)
同理的,如果参数的类型一致,只需要最后一个即可
//go的函数体特色一就是他的方法支持多个返回值
func rectProps(length, width float64)(float64, float64) {  
}
//go的函数体特色二就是他的方法支持返回值也有名字
func rectProps(length, width float64)(area, perimeter float64) {  
    area = length * width
    perimeter = (length + width) * 2
    return // 不需要明确指定返回值，默认返回 area, perimeter 的值
}
//go的函数体特色三之二的后续(空白符)
如果对于返回多个值得函数体我只需要其中一个返回值怎么说
那么就使用空白符`_`来接受你不需要的哪个返回值

## 6 包package

代码分割联系
import (
	"fmt"
	"packages/rectangle"
)
//包的初始目录算在src下(bin算是classpath??)
在一层包下,开头就package main
二层包下,就package 二层包
//对于需要被其余程序调用的函数方法,其首字母需要大写(表示被导出的名字,导出才能被访问)
这也算很奇特了,用大写来定义访问
//go下每个go文件有一个奇特的初始化函数init(),跟java的构造器类似,无需显示声明
init 函数不应该有任何返回值类型和参数，在我们的代码中也不能显式地调用它,这点与构造器不一样,其主要是负责初始化,我们可以在里面加初始化判断。
//go里面导入了包就必须使用,编译器报错或者直接帮你删除(代码洁癖啊,写go的人)
为了避免我们导入的暂时不会用的包,影响编译,或者被编译器删除,那么我们可以使用空白符`_`暂时调用其方法
var _ = rectangle.Area // 错误屏蔽器
为了能够在项目完成时删除这些错误屏蔽器,最好直接接在import段落的后面

## 7 ifelse

if num := 10; num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
}
//go比较有趣的是可以定义作用域只在if语句里面的变量,算是吸收了for循环的特性
//此外要注意else只能放在这个位置,不然if else有分号相当于隔开了

## 8 for语句

i := 0
for i <= 10 { //semicolons are ommitted and only condition is present
fmt.Printf("%d ", i)
i += 2 
}
//go在将13两步单独挑出来的时候,两个分号可以省略
for {  
}
//死循环

## 9 switch语句

switch letter := "i"; letter {
	case "a", "e", "i", "o", "u": // 一个选项多个表达式
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
}
//go同if一样可以声明局部变量,还可以case匹配多个,default是匹配不到的情况
switch{
	..
}
//go的switch可以省略,默认匹配true的case,即在case里面写入类似于if判断的条件
//这样可以轻松地替代多重if else(case内还可以switch),但是switch一般匹配成功,执行一个case就溜了
//这时候就需要使用fallthrough,放在case语句的最后,其表示无条件执行下一条case内容

## 10 数组与切片

(我们可以把切片理解成可变长度的数组,go里面切片应用多于数组)
a := [3]int{12, 78} 
//go的数组声明也是反过来的,初始值默认是0
//但是数组使用时一样的
a := [...]string{"USA", "China", "India", "Germany", "France"}
b := a
//数组是值类型,所以b=a相当于复制了一个数组..其中长度省略,由编译器指定长度
//所以同理当数组作为方法参数的时候,也只是值传递,不会改变原有值
for i, v := range a {
}
//range机制类似于foreach/迭代器,i迭代下标,v迭代值,如果忽略索引,就使用空白符
数组切片!!!
a := [5]int{76, 77, 78, 79, 80}
var b []int = a[1:4] // creates a slice from a[1] to a[3]
//我们可以将数组赋给一个变量,那么怎么传递引用,就是切片a[1:4],即包含下标1,2,3///不包括4
//修改切片,会直接修改数组本身的值
//a[:]缺省表示匹配整个数组的切片
len(b)
cap(b)
//上述切片的长度len为3,而容量cap则从第一个到最后,即4,除去a[0]
切片有内建函数
func make([]T, len, cap)
所以上述等同
b:=make([]int,3,4)
//至于多维切片的初始化,类比于,类似切片的堆层数
f := make([][]string, 3, 3) //三行,容量3的多维切片
for i := 0; i < len(f); i++ {
	f[i] = make([]string, 3, 5) //每行长度3,容量5
}
//元素追加
切片还有元素追加函数
func append（s[]T，x ... T）
也就是在s[]后加值,x的数量可以是多个,切片超过容量会自动扩容,变2倍
//记住是容量变为2倍,长度根据你实际走
//那么切片连接呢??
veggies := []string{"potatoes", "tomatoes", "brinjal"}
fruits := []string{"oranges", "apples"}
food := append(veggies, fruits...)
//切片跟数组很大的不同就是,go中作为方法参数时,数组是值传递,而切片是引用传递
a := [3]int{5, 78, 8}
a := []int{5, 78, 8}
//数组跟切片的区别就是,数组必须是定长度的,数组可以使用...来代替具体长度
//只要数组的切片在内存中，数组就不能被垃圾回收!!!!
//一种解决方法是使用 copy 函数 func copy(dst，src[]T)int 来生成一个切片的副本。这样我们可以使用新的切片，原始数组可以被垃圾回收。

## 11 可变参数函数

函数最后一个参数被记作xxx ...T, 这时函数可以接受任意个 T 类型参数作为最后一个参数
// 但是不能直接接受[]T的输入, 因为本身T传入, 编译器就是将其转换为[]T , 所以[]T传入就是[]{[]T}, 解决方法就是加...如: find(89, nums...)
fmt.Printf("type of nums is %T\n", nums) // %T获取参数类型,这里是[]int
// printf类似c, println类型java..
切片是引用类型:
参数传进来使用x[0]是直接改变的, 但是使用append数组的方法, 就是值传递了(数组相关的方法是值传递).

## 12 maps

通过make函数传入键和值得类型, 创建maps
personSalary := make(map[string]int)
personSalary["zhj"]=23
value,ok=personSalary["zhj"]
// 获取值会返回2个值, 第一个是value, 如果不存在就返回0, 同时ok接收到false
delete(personSalary, "steve")
// 直接删除, 没有返回值
map也是引用类型:
传递给新map, 新map的值改变, 原有map也会改变
// map 之间不能使用 == 操作符判断，== 只能用来检查 map 是否为 nil
判断两个 map 是否相等的方法是遍历比较两个 map 中的每个元素。我建议你写一段这样的程序实现这个功能

## 13 string

// 对于特殊字符遍历输出, 如果占用了两个字节, 就会输出异常
rune 是 Go 语言的内建类型，它也是 int32 的别称。在 Go 语言中，rune 表示一个代码点。代码点无论占用多少个字节，都可以用一个 rune 来表示。
//使用rune来转换字符串, 遍历rune即可
runes := []rune(s)
//可以使用字节切片, 来构造
byteSlice := []byte{0xe4, 0xbd, 0xa0, 0xe5, 0xa5, 0xbd}
//字符串rune的长度
utf8.RuneCountInString(str)
// 字符串无法直接被改变, 如果要改变需要转为rune切片入参再转回去
[]rune(str)

## 14 指针

// 跟c一样的创建方式 a *int=&b
*a可以取到b的值
// 数组传参到函数, 可以使用数组指针, 但最好用切片
不支持指针运算如 a++从b[0]指向b1

## 15 结构体, 嵌套 , 匿名

结构体可以匿名结构体//参照java匿名内部类
可以缺省字段名, 就是类型本身
可以使用指针, 但是没啥用
可以结构体嵌套结构体
嵌套结构体+匿名字段=提升字段, 也就是不需要p.Address.city, p.city即可
// 这就达到了一个继承的错觉
!!!
甚至p可以直接调用匿名字段Address作为接收器的方法-->继承
!!!
如果结构体名称以大写字母开头，则它是其他包可以访问的导出类型（Exported Type）。
同样，如果结构体里的字段首字母大写，它也能被其他包访问到。
~~
结构体如果字段都是值类型, 就可以==比较, 如果含有map类无法比较的字段, 就无法==

## 16 方法

func (t Type) methodName() {
}
而函数的声明
func functionname(t Type) returntype {  
}
使用起来:
t.methodName
functionname(t)
主要作用: 
1.将方法与接收器关联, 实现类似面向对象的作用(连接点类似于类,方法类似非静态方法)
2.函数无法重载, 而相同名字的方法, 可以指定不同的接收器类型
3.函数参数是值传递就必须是值传递, 是指针就必须是指针, 而方法都可以, 易用性高
!!!
值类型接收器调用方法是无法修改的,是值传递(拷贝),对外部不可见
指针类型调用方法是可见的, 是引用传递, 对外可见
!!!
方法的接收器类型定义和方法的定义应该在同一个包中

## 17 接口

接口指定了一个类型应该具有的方法，并由该类型决定如何实现这些方法。
// 在go中, 指定的方法, 叫做方法签名
定义一个接口, 然后定义一个无方法体的方法
在外面进行方法的完整定义, 其连接点就算是实现了接口的类
// 多态使用
比如方法传入一个接口数组, 然后可以传入所有实现了接口方法签名的接收器数组
// 空接口(没有任何方法签名)
可以接受所有类型, 因为其没有什么方法需要实现
// 类型断言i.(T), 同时检查能获取到类型, 减少异常产生
s, ok := i.(int) //获取i变量底层int的值,ok获取是否类型匹配, 能否获取到值
// 1.使用值接受者声明的方法，既可以用值来调用，也能用指针调用
// 2.对于使用指针接受者的方法，用一个指针或者一个可取得地址的值来调用都是合法的。
!!!
接口是可以嵌套的
// EmployeeOperations :接口嵌套
type EmployeeOperations interface {
	SalaryCalculator
	LeaveCalculator
}
EmployeeOperations可以接受实现了其嵌套的两个接口的接收器类型

## 18 协程和信道

//Go 使用 Go 协程（Goroutine） 和信道（Channel）来处理并发
//go function, 就可以启动一个子协程
//sqrch := make(chan int), 就可以创建一个双向的信道
//所有信道都关联了一个类型。信道只能运输这种类型的数据，而运输其他类型的数据都是非法的。
发送与接收默认是阻塞的。这是什么意思？当把数据发送到信道时，程序控制会在发送数据的语句处发生阻塞，直到有其它 Go 协程从信道读取到数据，才会解除阻塞。
与此类似，当读取信道的数据时，如果没有其它的协程把数据写入到这个信道，那么读取过程就会一直阻塞着。
//这是java里很典型的一个生产者消费者案例的公共资源, 阻塞帮助 Go 协程之间进行高效的通信，不需要用到其他编程语言常见的显式锁或条件变量。
!!!
死锁, 既然是阻塞的, 如果只有生产, 或者只有消费, 那么就会一直阻塞住! 
// 这里的死锁与其余语言的死锁还不是一种情况.
所以, go语言引入了单向信道的概念,这样即使"死锁"也不会报错
cha1 := make(chan <-int) 这个表示只能往内送值的信道(唯送send信道)
cha1 := make(<-chan int) 这个表示只能从里面收值的信道(唯收receive信道)
//把一个双向信道转换成唯送信道或者唯收（Receive Only）信道都是行得通的，但是反过来就不行
所以通常做法是针对方法, 参数设置为单向, 但是信道声明实际还是双向
// 关闭信道,告诉接收方再有信息
close(sendch)
v, ok := <-cha1

## 19 缓冲信道

// 前面无缓冲信道的发送和接收过程是阻塞的。这样会造成死锁的问题, 通过单向信道解决.
我们还可以创建一个有缓冲（Buffer）的信道。
只在缓冲已满的情况，才会阻塞向缓冲信道（Buffered Channel）发送数据。
同样，只有在缓冲为空的时候，才会阻塞从缓冲信道接收数据。
ch := make(chan type, capacity)
// 顾名思义, 在缓存范围内不会阻塞, 但是超过阻塞范围就还是会阻塞,产生死锁.
缓存信道类比队列, cap是容量, len是长度.
cap(ch) len(ch)
// 前面我们有让主线程进行time.Sleep(2 * time.Second)来阻止main协程关闭
我们可以使用sync.WaitGroup来让主协程等待, wg.Add(1)伴随每一个go协程
然后在main的最后wg.Wait()等待, 将指针地址wg *sync.WaitGroup传入协程方法,
每个协程执行完就wg.Done()
// WaitGroup的设计思想类似于Java匿名的CountDownlatch, 不过那个是等待, 然后并发执行, 这是等待协程结束再关闭主协程
// java有线程池, go里面也有, 叫工作池, 可以使用缓冲信道来实现
我们工作池的核心功能如下：
	创建一个 Go 协程池，监听一个等待作业分配的输入型缓冲信道。
	将作业添加到该输入型缓冲信道中。
	作业完成后，再将结果写入一个输出型缓冲信道。
	从输出型缓冲信道读取并打印结果。
	
// 还有sync.Once来实现代码只执行一次, 比如实现单例模式
```go
// 使用sync.Once来实现, 与java使用内部类编译实现单例模式异曲同工
var singleTon4 *singleTonDemo4
var once sync.Once

func NewSingleTon4() *singleTonDemo4 {
	once.Do(func() {
		singleTon4 = &singleTonDemo4{1}
		fmt.Println("singleTonDemo4 created")
	})
	if singleTon4.count != 1 {
		singleTon4.count++
	}
	return singleTon4
}
```

## 20 select

select 会阻塞, 只要case内有一个协程完成, 程序终止
//具体使用案例:
select 会选择首先响应的服务器，而忽略其它的响应。使用这种方法，我们可以向多个服务器发送请求，并给用户返回最快的响应了。:）
//default用法: 不仅可以用在switch, 在信道阻塞死锁时, 也可以实现默认操作,而不报错中止
default:
		fmt.Println("默认操作, 防止没有信道有返回值而产生死锁")
//空select也会一直死锁
select {}

## 21 mutex

mutex: java的synchronized的monitor对象其实底层还是用的mutex锁
x=x+1简单可以分为三步:
	y=x
	y+1
	x=y
当两个协程对于x进行+1的时候,可能最后结果是1或者2两种情况.
//类似于数据库里面的串行化异常
Mutex可以在sync包中找到, mutex.Lock(), 和mutex.Unlock(), 避免协程竞争冲突
//如果有一个 Go 协程已经持有了锁（Lock），当其他协程试图获得该锁时，这些协程会被阻塞，直到 Mutex 解除锁定为止。
除了使用Mutex锁, 还可以使用缓冲信道来解决协程竞态
//设立容量为1的缓存信道, 同一时间只有一个协程能够输入值到信道, 也是很巧妙
总结: 至于Mutex和信道的选择, 如果只是加锁, 我觉得是Mutex好用, 但如果需要协程交互, 就可以使用缓冲信道

## 22 struct-->面向对象

结构体取代类(封装->结构体首字母小写, 继承->嵌套组合结构体, 多态->接口+结构体)
go既可以说是面向对象, 也可以说不是.

> 面向对象三大特性: 封装继承多态.
封装与多态, 使用接口与结构体(方法通过接收器来绑定结构体方法), 我们都可以实现, 但是继承呢?
实际继承也有, 在结构体的<嵌套结构体+匿名字段=提升字段>中, 我们对于嵌套进来的结构体封装元素可以看作是集成了.
//当然, Go用起来很是变扭, 因为个人习惯了java, 但是go的信道, 协程, 在我看来, 感觉甩了java几个世纪.
有一个问题就是Go匿名很多值得默认值要么是nil, 要么是0, 当然java也一样
但是java可以通过构造器来指定初始值, Go的话也可以通过类似的方式.
通过结构体名首字母小写, 关闭接口, 然后创建NewT(parameters)T{}, 来返回我们的结构体T对象
!!!!
所以, 第一步, 人造类, 面向对象就完成了.
同时将结构体的首字母改成小写, 也就是实现了封装
!!!!
记下来可以使用<嵌套结构体+匿名字段=提升字段>来实现"继承"
在go里面叫做组合, 手工完成继承效果
!!!!
最后是多态, 主要通过<接口+结构体>来实现, 编译时类型->接口, 运行时类型->结构体
!!!!
Go中什么对应java的Object类: interface{}
a interface{}作为参数表示可以接受所有的类型参数

## 23 defer

defer: 翻译中文是延缓,延迟
-->含有 defer 语句的函数，会在该函数将要返回之前，调用另一个函数。
比如func1{defer func2}, 那么在func1结束前的一刻会调用func2方法
// 当然方法,也是函数, 只是绑定了类型, 也可以使用defer o.method
!!!!
注意点1就是, 如果defer的方法有参数, 在defer后背修改了, 那么生效的是改之后还是之前的参数?
--> 是之前, defer只是延缓了方法执行的时间, 但是参数其实早就传进去了
!!!!
注意点2, 如果调用了多次defer, 那么到底是队列形式还是栈呢?
--> 是栈, 所以可以实现一些逆向的操作, 比如倒序打印字符串字符 defer fmt.Printf("%c", v)
!!!!
最常见使用, 比如每一个return前都需要执行的操作, 如waitGroup.Done()
可以使用defer waitGroup.Done()放到函数第一行, 简化代码

## 24 error

error: go的错误消息:不同于其余语言的异常机制, 并不会导致程序终止
// 1.比如打开一个文件
f, err := os.Open("/test.txt")
// 然后将err与nil比较, 采取操作(手动try catch)
fmt.Println(err)实际调用的是err的Error() string , 方法打印错误消息,类似数据库po文件
//open /test.txt: The system cannot find the file specified.
//err的结构
type error interface {  
    Error() string
}
// 2.怎么定制异常返回?也就是获取异常详情
go断言底层有各种error的实现结构体
如文件相关的PathError:
type PathError struct {  
    Op   string
    Path string
    Err  error
}
func (e *PathError) Error() string { return e.Op + " " + e.Path + ": " + e.Err.Error() }
// 其实现了error接口
// 其中e.Op就是open, e.Path就是/test.txt, e.Err就是The system cannot find the file specified.
// 又如DnsError可以判断是请求超时还是域名错误, 还是暂时性地错误
// 3.还有直接比较错误的做法
比如使用path/filepath包的Glob来返回一个符合glob类型的文件名, 如果不对, 就会返回一个错误ErrBadPattern
然后使用e == filepath.ErrBadPattern来直接比较获得的异常, 进行相关操作
!!!!
千万不要使用_来忽略错误, 不对错误进行处理, 按照默认正确的方式处理, 就会返回错误的结果并没有提示.
//感觉这一点go不如java, 应该在有返回异常的地方强制进行接受并处理, 指定一个处理模块如try catch-->后面会了解到有panic可recover,但是基本不用...
// 4.自定义简单错误
类比前面直接比较的异常, 那个异常底层就是var ErrBadPattern = errors.New("syntax error in pattern")
所以很简单, 但是个人觉得好简陋啊...
// 5.给错误添加更多信息
errors.New(str)固然简单可用, 但是只能写死, 要写活怎么搞
fmt.Errorf(str,param)这样, 类型printf的语法, 可以实现动态的错误返回
// 底层实现: 实际上还是前者,不过用了一些列方法
// 6.自定义结构体类型的错误!!
func (e *areaError) Error() string {}
优先续约实现error接口的方法, 然后可以在此基础上拓展方法

## 25 panic,recover

类似于其他语言的try-catch-finally, panic出错情况下会打印堆栈信息
//但是大部分情况下, 我们只需要使用error进行错误的返回即可.除非程序无法运行!
这时候, defer延迟的函数怎么算, 是算程序终止也算结束, 会运行延迟函数, 还是不会呢?
!!!!
panic在发生时, 这个产生错误的命令会暂停, 等待执行执行完所有的延迟函数后, 再继续, 并打印堆栈
//前面是捕获到了堆栈异常, 但是程序还是终止了, 那么有什么办法可以catch到并处理呢?
那就是使用defer+recover(因为panic会阻塞住不往后面走, 我们只能在延迟函数中获取)
进行处理后, 发生异常的那一行会跳过, 继续后续的作业
!!!!
记住, recover和panic必须在同一个协程中才有效
比如你main你们defer recover, 然后go xxx在xxx你们panic这样是不行的.

>
有一些运行时panic(不是自己定义的panic), 即runtime.Error, 比如数组越界.
修复方法一样.
!!!
但是, 我们只是recover, 并没有堆栈信息, 如何获取堆栈信息呢?
>
使用 Debug 包中的 PrintStack 函数。
debug.PrintStack()
这一句放在recover里面, 同理java的e.printstack

## 26 头等函数/高阶函数

Go可以将函数当作变量.
可以赋值给变量, 也可以作为其他函数的参数, 或者返回值.

> !!!这个特性叫做头等函数.
还可以吧函数当作类型, 自定义一个类型:
type add func(a int, b int) int
!!!
使用函数作为参数, 或者将函数返回的函数, 我们都将其称为高阶函数
> 用一个变量接受返回的函数, 然后可以进行函数调用
//闭包, 匿名函数(lambda表达式)的特例: 匿名函数访问的变量定义在函数体的外部, 这个匿名函数就叫闭包....
每一个闭包都会绑定一个它自己的外围变量,也就是外部变量传进去的就是本身, 你如果函数a+=param, 那么就会进行累加

## 27 反射


Go语言也可以实现反射, 在运行时检查变量和值!

> 仿佛又看到了Spring的身影, やめで
reflect.TypeOf(face): 返回结构体声明的名称
v := reflect.ValueOf(face) : 返回结构体对象内的值
// 我们应该使用反射吗???
清晰优于聪明。而反射并不是一目了然的。
所以, 能不用就不用!
但是!! 正如c/c++使用goto一样, 如果能够掌控, 请使用

## 28 文件读写

ioutil.ReadFile("test.txt")
// 可以使用相对路径和绝对路径
至于写文件: https://studygolang.com/static/pkgdoc/pkg/os.htm#OpenFile
func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
//其中FileMode是决定文件类型和权限的: type FileMode uint32

## 29 Web框架

https://github.com/speedwheel/awesome-go-web-frameworks
gin中文文档: https://github.com/skyhee/gin-doc-cn

beego: <https://beego.me/>

## 30 AOP

go get -v github.com/gogap/aop