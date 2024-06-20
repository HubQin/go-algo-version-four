## 1.3.6
- 下面这段代码对队列 q 进行了什么操作？
```java
Stack<String> stack = new Stack<String>();
while (!q.isEmpty())
    stack.push(q.dequeue());
while (!stack.isEmpty())
    q.enqueue(stack.pop());
```
> 利用stack的FIFO特性， 将队列顺序反转。