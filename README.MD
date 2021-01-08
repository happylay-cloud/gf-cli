# GF-CLI

`gf-cli` 是一个功能强大的命令行工具，可以方便地构建[GF](https://goframe.org)应用。

## 1. 安装

### 1). 二进制
建议使用预编译的二进制文件：
- `Linux   (amd64)`: https://goframe.org/cli/linux_amd64/gf
- `Mac     (amd64)`: https://goframe.org/cli/darwin_amd64/gf
- `Windows (amd64)`: https://goframe.org/cli/windows_amd64/gf.exe
- [更多...](https://goframe.org/cli)

下载后，请使用`gf install`命令将gf二进制文件安装到系统路径。

1. `Mac`
    ```shell
    wget https://goframe.org/cli/darwin_amd64/gf && chmod +x gf && ./gf install
    ```
   > 如果使用的是`zsh`，你可以通过`alias gf=gf`命令重命名别名，来解决`gf`和`git fetch`直接的冲突。
                                                                                                                          
2. `Linux` 
    ```shell
    wget https://goframe.org/cli/linux_amd64/gf && chmod +x gf && ./gf install
    ```
                                                                                                                  
3. `Windows`

    手动下载并执行，然后按照说明进行操作。

### 2). 手动安装

> 您可能需要设置GOPROXY才能完成构建。

1. 最新版本
    ```
    git clone https://github.com/gogf/gf-cli 
    cd gf-cli && go build main.go && ./main install
    ```
   
2. 指定版本
    ```
    git clone --branch v0.7.3 https://github.com/gogf/gf-cli 
    cd gf-cli && go build main.go && ./main install
    ```
   
3. 默认情况下，`gf gen model`命令不支持数据库`sqlite`和`oracle`，因为它需要`cgo`和`gcc`，你可以手动更改源代码并进行构建。


4. 当前mine分支发行版 [gfctl](https://github.com/happylay-cloud/gf-cli/releases) 已预编译sqlite包，支持中文显示，与主库分支同步。

## 2. 命令
```html
$ gf

用法
    gf 命令 [主题] [选项]

命令
    env        显示当前Golang环境变量
    get        默认情况下安装或更新GF到系统...
    gen        自动为ORM模型生成go文件...
    mod        Go模块的附加功能...
    run        运行具有类似热编译功能的go代码...
    init       在当前工作目录中初始化空GF项目...
    help       显示有关指定命令的详细信息
    pack       正在将任何文件/目录打包到资源文件或go文件...
    build      交叉编译跨平台go项目...
    docker     为当前的GF项目创建docker镜像...
    swagger    当前项目的swagger功能...
    update     更新当前的gf二进制文件到最新的(可能需要root/admin权限)
    install    安装gf二进制到系统(可能需要root/admin权限)
    version    显示当前二进制版本信息

选项
    -y         所有命令都是"yes"，没有提示询问
    -?,-h      显示指定命令的帮助或详细信息
    -v,-i      显示版本信息

附加
    使用"gf help 命令"或"gf 命令 -h"获取（注释末尾带有"..."）命令的详细信息
```

## 3. 常见问题

### 1). 命令`gf run`返回`pipe: too many open files`

请使用`ulimit -n 65535`来扩大当前终端shell会话的最大打开文件数的系统配置，然后使用`gf run`。







