# goutils
------------------
日常开发工具箱。

目前包含以下内容：
- 简单易用的AES对称加密，Base64解码
- go程序后台运行(daemonize)功能，支持MacOS/Windows/Linux
- 终端UI(tui)，目前有颜色、渐变色、简单日志打印、密码输入框、进度条等
- JSON格式配置文件管理，修复了koanf的JSON格式没有indent的问题
- 日志文件rotate封装，简单易用
- 请求封装，支持socks5格式代理，以及下载进度条
- 基于unix套接字的http服务，有客户端和服务端，可以方便地进行进程间通信
