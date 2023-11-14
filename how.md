done: 
- 代码中的错误处理应该怎么做？是写在函数里？还是写在处理流程里？: 在函数里就应该把已知的可能报错给 raise 出来
- 多线程，包含重试，直到完成: queue 就行了
- 如无必要，不要透传!
- 每个函数只应该知道自己该做什么，不应该知道别的地方发生了什么

todo: 
- cron 没生效
- gradio
- curl
- 返回值是不是一律用字典就有最好的扩展性？
- 如何 build 一个好的 system