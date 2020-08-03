leetcode practice

# 链表与指针

### 解题思路

+ 正向思考：从题目解答出发，想通思路，抽象方法
+ 逆向思考：根据“守则”，倒退边界场景出现的场景，以及应该进行的操作，保证边界

### 守则
+ 守则1：(空指针引用)凡是涉及指针的使用，前面必须做非空判断

+ 守则2：(边界问题)根据其他指针规则，倒推边界条件出现的场景，以及应该进行的操作，保证边界正确处理用例

+ 守则3：(有效遍历)循环中遍历链表且指针位移在括号中，保证每次循环都让指针向下移动一次，尤其是分支中出现continue时

+ 守则4：(保留链表头)凡是还要用到的链表，在遍历时必须保留其头部节点，一般开游标遍历; 如果要一份新的链表，使用深拷贝

  
### 快慢指针/多指针
#### 使用场景
+ 单向链表中涉及链表中两个节点需要相互比较的场景
+ 单项链表指针修改（翻转，两侧插入等涉及多点指针修改场景）

#### 常见题目

+ 找到链表的中间节点-快慢指针
+ 翻转链表-pre,op, post多指针



### dummy node/哨兵节点

#### 使用场景

+ 处理链表的边界条件(尤其是多指针/快慢指针场景)
+ 哑巴节点可以很好的处理在链表头前部插入的场景

#### 常见题目

+ 合并链表
  + 链表合并时，关注被合入目标链表与拿来操作的链表的长短关系
+ 翻转链表
  + dummy节点的next指向需要置空，防止尾结点回环
  + 翻转完成后删除dummy节点
  + 翻转结束后，关注尾结点是否完成翻转--边界条件



### golang函数传参

#### 传参为指针

+ 只能对传入的链表进行修改，新建链表的内存块，会在出函数时被释放
+ 但是可以通过将计算出的值更新传入的链表，就可以把修改值带出去

> 核心：指针参数在出入子函数中其值（指针地址）是保持不变的 

# Stack && Queue

### 应用场景

#### 计算器场景

+ 调度场算法，中序转后序表达式--factors queue+operator stack(运算优先级低的进栈)
+ 后序表达式进栈出栈完成结果计算-calculate stack

#### [字符串解码类操作](https://leetcode-cn.com/problems/decode-string/)

+ 解法一：简单入栈出栈处理字符串编码
+ 解法二：划归为万能公式“计算器场景”（把"\[""\]"当做运算符）
  - 先做一次翻译：把原有“中序表达式”，"\[""\]"看做“-”，数字与字母或"\[”间加上乘法，这里减号的意义是字符串连接，乘法的意义是字符串翻倍。
  - 调度场算法转成逆波兰表达式
  - 逆波兰表达式入栈出栈，根据step1中的翻译规则做连接与倍乘（利用减法运算需要分算子的前后顺序来保障字符串的顺序正确）

> 总结
>
> + 调度场算法+逆波兰表达式入栈出站可以解决大多数的复杂操作符（分优先级）+算子的计算，但是要注意先后顺序
> + 简化算法暴力对中序进行计算能够降低调度场算法和逆波兰进出站计算的代码复杂性，这个取舍要看题目的操作符分级的复杂程度

# 字符串操作

### Tips：

+ 翻转字符串中点为i < len(s)/2
+ 判断字符串中字符是否落在某个区间[charA, charZ]，可先split，再字符串做比（不是char类型比较）

### 字符串操作
``` golang
func Join(a []string, sep string) string

// Split slices s into all substrings separated by sep and returns a slice of
// the substrings between those separators.
//
// If s does not contain sep and sep is not empty, Split returns a
// slice of length 1 whose only element is s.
//
// If sep is empty, Split splits after each UTF-8 sequence. If both s
// and sep are empty, Split returns an empty slice.
//
// It is equivalent to SplitN with a count of -1.
func Split(s, sep string) []string { return genSplit(s, sep, 0, -1) }

// Generic split: splits after each instance of sep,
// including sepSave bytes of sep in the subarrays.
func genSplit(s, sep string, sepSave, n int) []string
```
