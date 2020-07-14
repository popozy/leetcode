leetcode practice

# 链表与指针

### 解题思路

+ 正向思考：从题目解答出发，想通思路，抽象方法
+ 逆向思考：根据“守则”，倒退边界场景出现的场景，以及应该进行的操作，保证边界

### 守则
+ 守则1：(空指针引用)凡是涉及指针的使用，前面必须做非空判断

+ 守则2：(边界问题)根据其他指针规则，倒推边界条件出现的场景，以及应该进行的操作，保证边界正确处理用例

  
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

