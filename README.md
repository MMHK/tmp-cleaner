# tmp-cleaner

[![Go Report Card](https://goreportcard.com/badge/github.com/mmhk/tmp-cleaner)](https://goreportcard.com/report/github.com/mmhk/tmp-cleaner)
[![Docker Pulls](https://img.shields.io/docker/pulls/mmhk/tmp-cleaner.svg)](https://hub.docker.com/r/mmhk/tmp-cleaner)
[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)



`tmp-cleaner` 是一个用 Go 编写的守护进程，能够定期清理指定目录中的旧文件。通过 Docker 容器运行，支持环境变量和命令行参数配置。

## 功能

- 定期删除指定目录中超过指定天数的文件。
- 支持通过环境变量和命令行参数进行配置。
- 使用 Docker 容器运行，方便部署和管理。

## 环境变量

- `TMP_DIR`：要清理的临时目录（默认：`/tmp`）。
- `DAYS`：保留文件的天数（默认：`7`）。
- `INTERVAL`：清理的时间间隔（秒）（默认：`86400`，即一天）。

## 命令行参数

- `-tmpDir`：要清理的临时目录。
- `-days`：保留文件的天数。
- `-interval`：清理的时间间隔（秒）。

命令行参数将覆盖环境变量配置。

## 使用方法

### 1. 克隆仓库

```bash
git clone https://github.com/yourusername/tmp-cleaner.git
cd tmp-cleaner
```

### 2. 构建 Docker 镜像

```bash
docker build -t tmp_cleaner .
```

### 3. 运行 Docker 容器

#### 使用环境变量

```bash
docker run -d \
  -e TMP_DIR="/mnt/tmp" \
  -e DAYS=7 \
  -e INTERVAL=86400 \
  -v /path/to/local/tmp:/mnt/tmp \
  mmhk/tmp_cleaner:latest
```

#### 使用命令行参数

```bash
docker run -d \
  -v /path/to/local/tmp:/mnt/tmp \
  mmhk/tmp_cleaner:latest -tmpDir="/mnt/tmp" -days=7 -interval=86400
```

### 4. 查看帮助信息

```bash
docker run --rm tmp_cleaner -h
```

### 5. Docker-Compose
可以使用 Docker-Compose 来管理容器，下面是一个示例：
```yaml
version: '3'
services:
  tmp_cleaner:
    image: mmhk/tmp-cleaner:latest
    container_name: tmp_cleaner
    volumes:
      - /path/to/local/tmp:/mnt/tmp
    environment:
      - TMP_DIR=/mnt/tmp
      - DAYS=7
      - INTERVAL=86400
    restart: always
```