# 函数

## 单一职责，函数如其名

曾经，debug 的时候，一个函数叫 call_bot，其中接受一个 chat_history 作为参数。
我以为这个参数是无状态的，接受 chat_history 和其它几个参数，然后输出一个 text，就完事了。
结果这个函数除了 call_bot，还修改了 chat_history。让我 debug 半天。

每个函数都是**独立**的，意即，每个函数都只做与自己的命名相符的事情。

TODO: call_bot 结果其中还维护了 history 并改变了它

## 最小依赖

每个函数所做的事情应该**只依赖其入参**，而不要依赖全局变量。

## 错误处理

每个函数应自行 try catch 尝试处理其已知的错误，无法处理的自行 raise。

这样，外部不会因为要处理各种错误而膨胀（因为有些错误函数自己可以处理并尝试恢复，如果让外部调用来做，则 1 外部不该知道函数内部的知识，2 外部代码膨胀）