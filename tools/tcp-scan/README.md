# 简单的 TCP 扫描器

## 安装
```shell script
go build github.com/ylh990835774/go-base/tools/tcp-scan
```

## 使用帮助

```shell script
./tcp-scan -h

Usage of ./tcp-scan:
  -end-port int
        the port on which the scanning ends (default 100)
  -hostname string
        hostname to test (default "www.baidu.com")
  -start-port int
        the port on which the scanning starts (default 80)
  -timeout duration
        timeout (default 200ms)
```

## 使用

```shell script
./tcp-scan
```

```shell script
./tcp-scan -hostname=www.baidu.com -start-port=3306 -end-port=3310 -timeout=300ms
```