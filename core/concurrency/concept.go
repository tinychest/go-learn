package concurrency

// 无论是 main 函数对应的主协程，还是 test 函数对应的主协程，只要方法逻辑执行结束了，程序就结束了，不会理会在执行过程种开启的协程是否结束
