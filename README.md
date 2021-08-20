# hcshop_srvs

mooc学习go

## 8-11

1. 定义了简单的基础项目结构
2. 通过gorm创建了用户表，这里涉及了很多mysql基础知识
3. 在在定义user结构体的tag时候，gland对很多标签上没有支持的（比如 gorm 标签），这里可以考虑使用 kite 插件，能减少一些手写的错误
4. 在定义proto 协议时候，go_package ,以及引入的google/protobuf/empty 报红色，是因为需要配置protobuff的默认引用位置，这个在protobuff
   安装理由，mac使用brew安装，默认位置在/usr/local/include/google/protobuf， 所以需要在插件->protobuff 里添加配置路径。路径地址是 ：`/usr/local/include`。

## 8-12

1、完成了剩下创建用户，修改用户和密码校验接口，这里都比较简单，没有太多东西

## 8-13

1、本次主要是完成web端的对接处理，service端提供服务 2、本次调试对接接口为获取用户列表，有个小问题，获取用户列表的orm操作，有两个sql查询
![image](https://user-images.githubusercontent.com/4961672/129339480-0e1a7cde-bb3a-4dca-944a-453ff0d52c34.png)
这里需要看看是不是使用方式有问题，因为是第一次正式使用gorm，而且是跟着老师写，所以这个orm框架，自己还得再研究研究。

## 8-14

1、微服务用户模块主要接口已经提供，主要是端功能实现 2、web端已经用上了日志框架zap以及配置文件读取框架 viper。这边服务提供方也需要跟进了😳

## 8-15 -> 8-17
处理web层的服务，暂时跟这里没关系

## 8-18
 1. 添加全局配置文件，日志框架，这个跟web一样，这是为下一步做服务注册准备的

## 8-19
1. 添加服务注册，件web-srv 注册到consul注册中心
2. grpc本身就提供了健康检查的接口，**不需要像web层一样，需要提供指定的http接口来探测**

## 8-20
 注册中心会将ID相同的服务器覆盖掉，为了同时启动多个服务，使用uuid 作为服务的ID，服务名字可以相同，在consul中，一个名字下会存在多个实例。
 uuid生成直接用了go本身方法，没有向课程中那样引入第三方，目测问题不大，这里应该没坑。