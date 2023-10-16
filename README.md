# goutils
------------------
日常开发工具箱。

目前包含以下内容：
- 简单易用的AES对称加密，Base64加密和解码，UUID生成，随机字符串
- go程序后台运行(daemonize)功能，支持MacOS/Windows/Linux
- Terminal UI(tui)，目前有颜色、渐变色、简单日志打印、密码输入框、进度条、选择列表、确认等
- JSON格式配置文件管理，修复了koanf的JSON格式没有indent的问题
- 日志文件rotate封装，简单易用
- 请求封装，支持http/socks5格式代理，以及下载进度条，分段并发下载
- 基于unix套接字的http服务，有客户端和服务端，可以方便地进行进程间通信
- 支持http/socks5两种代理模式的git命令，用于加速github的clone、pull、push等操作
- 文件压缩为zip格式
- zip, rar, 7z, tar.gz, xz等格式的文件解压
- zip压缩，支持设置密码
- git增强功能，支持设置代理(socks5/http)，支持的命令有clone/pull/push/commmit+push/tag+push/tag delete+push
- ctrl+C signal捕获
- 快速排序通用版
