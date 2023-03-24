# 2.1命名

命名规则：一个名字必须以一个字母（Unicode字母）或下划线开头，后面可以跟任意数量的字母、数字或下划线。（驼峰式命名）
关键字不能做命名
25关键字:
    break      default       func     interface   select
    case       defer         go       map         struct
    chan       else          goto     package     switch
    const      fallthrough   if       range       type
    continue   for           import   return      var

# 2.3变量

变量声明的一般语法: var 变量名字 类型 = 表达式

简短变量声明: 名字 := 表达式

指针: 如果用“var x int”声明语句声明一个x变量，那么&x表达式（取x变量的内存地址）将产生一个指向该整数变量的指针，指针对应的数据类型是*int，指针被称之为“指向int类型的指针”。

new函数: 表达式new(T)将创建一个T类型的匿名变量，初始化为T类型的零值，然后返回变量地址，返回的指针类型为*T

类型: 一个类型声明语句创建了一个新的类型名称，和现有类型具有相同的底层结构
type 类型名字 底层类型
类型声明语句一般出现在包一级