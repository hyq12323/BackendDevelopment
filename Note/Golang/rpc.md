## RPC

### rpc概念

1、RPC（Remote Procedure Call）远程过程调用，简单理解是一个节点请求另外一个节点提供的服务。

2、对应rpc的是本地过程调用，函数是最常见的本地过程调用。

3、将本地过程调用变成远程过程调用会面临各种问题。

### 本地过程调用

``` python
def add(a,b):
    total = a + b
    return total

total = add(1,2)
print(total)
```

函数调用过程：

1、将1和2压入add函数的栈

2、进入add函数，从栈中取出1和2分别赋值给a和b

3、执行a+b将结果赋值给局部的total并压栈

4、将栈中的值取出来赋值给全局的total

### 远程过程调用带来的问题

