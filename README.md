# hcshop_srvs
mooc学习go

## 8-11
 1. 定义了简单的基础项目结构
 2. 通过gorm创建了用户表，这里涉及了很多mysql基础知识
 3. 在在定义user结构体的tag时候，gland对很多标签上没有支持的（比如 gorm 标签），这里可以考虑使用 kite 插件，能减少一些手写的错误
 4. 在定义proto 协议时候，go_package ,以及引入的google/protobuf/empty 报红色，是因为需要配置protobuff的默认引用位置，这个在protobuff 安装理由，mac使用brew安装，默认位置在/usr/local/include/google/protobuf，
 所以需要在插件->protobuff 里添加配置路径。路径地址是 ：`/usr/local/include`。

## 8-12
1、完成了剩下创建用户，修改用户和密码校验接口，这里都比较简单，没有太多东西

## 8-13
1、本次主要是完成web端的对接处理，service端提供服务
2、本次调试对接接口为获取用户列表，有个小问题，获取用户列表的orm操作，有两个sql查询
![image](https://user-images.githubusercontent.com/4961672/129339480-0e1a7cde-bb3a-4dca-944a-453ff0d52c34.png)
这里需要看看是不是使用方式有问题，因为是第一次正式使用gorm，而且是跟着老师写，所以这个orm框架，自己还得再研究研究。
