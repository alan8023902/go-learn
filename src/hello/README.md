# Go 学习项目

## 安装 Go（解决错误 2755）

如果遇到安装错误 2755，请尝试以下解决方案：

### 方法 1：以管理员身份运行安装程序
1. 右键点击 Go 安装程序
2. 选择"以管理员身份运行"

### 方法 2：清理临时文件
1. 按 `Win + R`，输入 `%temp%` 并回车
2. 删除所有临时文件
3. 重新运行安装程序

### 方法 3：使用命令行安装
```powershell
# 下载 Go 安装包后，以管理员身份运行 PowerShell
msiexec /i go1.25.5.windows-amd64.msi /quiet
```

### 方法 4：从 .zip 文件手动安装（推荐解决 2755 错误）

#### 步骤 1：下载 .zip 文件
1. 访问 [Go 官方下载页面](https://golang.org/dl/)
2. 下载 Windows 版本的 `.zip` 文件（例如：`go1.25.5.windows-amd64.zip`）

#### 步骤 2：解压文件
1. 右键点击下载的 `.zip` 文件
2. 选择"解压到..." 或 "全部提取"
3. 解压后会得到一个 `go` 文件夹

#### 步骤 3：放置到安装目录
**推荐位置：`C:\Program Files\Go`**

**方法 A：使用 PowerShell（需要管理员权限）**
```powershell
# 以管理员身份运行 PowerShell
# 假设解压后的 go 文件夹在 D:\Downloads\go
# 移动到 Program Files
Move-Item -Path "D:\Downloads\go" -Destination "C:\Program Files\Go" -Force
```

**方法 B：手动操作**
1. 将解压后的 `go` 文件夹复制或移动到 `C:\Program Files\Go`
2. 确保最终路径是 `C:\Program Files\Go\bin\go.exe`

#### 步骤 4：配置环境变量

**方法 A：使用 PowerShell 脚本（推荐）**
```powershell
# 以管理员身份运行 PowerShell，执行：
.\setup-go-from-zip.ps1
```

**方法 B：手动配置**

1. **设置 GOROOT 环境变量：**
   - 右键"此电脑" → "属性" → "高级系统设置"
   - 点击"环境变量"
   - 在"系统变量"中点击"新建"
   - 变量名：`GOROOT`
   - 变量值：`C:\Program Files\Go`
   - 点击"确定"

2. **将 Go 添加到 PATH：**
   - 在"系统变量"中找到 `Path`，点击"编辑"
   - 点击"新建"，添加：`C:\Program Files\Go\bin`
   - 点击"确定"保存所有更改

**方法 C：使用 PowerShell 命令**
```powershell
# 以管理员身份运行 PowerShell

# 设置 GOROOT
[System.Environment]::SetEnvironmentVariable("GOROOT", "C:\Program Files\Go", "Machine")

# 将 Go\bin 添加到 PATH
$currentPath = [System.Environment]::GetEnvironmentVariable("Path", "Machine")
$goBinPath = "C:\Program Files\Go\bin"

if ($currentPath -notlike "*$goBinPath*") {
    $newPath = "$currentPath;$goBinPath"
    [System.Environment]::SetEnvironmentVariable("Path", $newPath, "Machine")
    Write-Host "已添加 Go 到 PATH" -ForegroundColor Green
} else {
    Write-Host "PATH 中已包含 Go" -ForegroundColor Yellow
}
```

#### 步骤 5：验证安装
**重要：关闭所有 PowerShell 窗口，重新打开一个新的 PowerShell 窗口**

```powershell
# 检查 Go 版本
go version

# 应该看到类似输出：go version go1.25.5 windows/amd64

# 检查 GOROOT
go env GOROOT
# 应该显示：C:\Program Files\Go
```

如果 `go version` 命令不工作，请检查：
- 环境变量是否正确设置
- 是否重新打开了 PowerShell 窗口
- `C:\Program Files\Go\bin\go.exe` 文件是否存在

## 配置环境变量

### 1. 配置 GOPATH（工作空间）

**Windows PowerShell:**
```powershell
# 创建 Go 工作空间目录
New-Item -ItemType Directory -Force -Path "$env:USERPROFILE\go"

# 设置 GOPATH 环境变量（当前会话）
$env:GOPATH = "$env:USERPROFILE\go"

# 永久设置 GOPATH（需要管理员权限）
[System.Environment]::SetEnvironmentVariable("GOPATH", "$env:USERPROFILE\go", "User")

# 将 GOPATH\bin 添加到 PATH
$currentPath = [System.Environment]::GetEnvironmentVariable("Path", "User")
$newPath = "$currentPath;$env:USERPROFILE\go\bin"
[System.Environment]::SetEnvironmentVariable("Path", $newPath, "User")
```

**或者使用图形界面：**
1. 右键"此电脑" → "属性" → "高级系统设置"
2. 点击"环境变量"
3. 在"用户变量"中：
   - 新建 `GOPATH`，值为 `C:\Users\你的用户名\go`
   - 编辑 `Path`，添加 `%GOPATH%\bin`

### 2. 配置 Go 模块代理（加速下载）

**Windows PowerShell:**
```powershell
# 设置 Go 模块代理（当前会话）
$env:GOPROXY = "https://goproxy.cn,direct"

# 永久设置（需要管理员权限）
[System.Environment]::SetEnvironmentVariable("GOPROXY", "https://goproxy.cn,direct", "User")

# 验证设置
go env GOPROXY
```

**常用代理地址：**
- 七牛云：`https://goproxy.cn,direct`
- 阿里云：`https://mirrors.aliyun.com/goproxy/,direct`
- 官方：`https://proxy.golang.org,direct`

### 3. 验证安装

打开新的 PowerShell 窗口，运行：

```powershell
# 检查 Go 版本
go version

# 检查环境变量
go env GOROOT
go env GOPATH
go env GOPROXY

# 查看所有 Go 环境变量
go env
```

## 使用 go run 和 go build

### go run - 直接运行程序

```powershell
# 运行 main.go 文件
go run main.go

# 运行多个 Go 文件
go run main.go utils.go

# 运行当前目录所有 .go 文件
go run .
```

**输出：**
```
Hello World
```

### go build - 编译程序

```powershell
# 编译当前目录的 Go 程序
go build

# 编译并指定输出文件名
go build -o hello.exe

# 编译并显示详细信息
go build -v

# 编译时禁用优化（用于调试）
go build -gcflags="-N -l"
```

**编译后运行：**
```powershell
# Windows
.\hello.exe

# 或直接
hello.exe
```

### 其他常用命令

```powershell
# 格式化代码
go fmt main.go

# 运行测试
go test

# 安装依赖包
go get package-name

# 初始化 Go 模块（用于新项目）
go mod init project-name
```

## 项目结构

```
golearn/
├── main.go          # 主程序文件
└── README.md        # 说明文档
```

## 下一步学习

1. **Go 基础语法**：变量、函数、结构体
2. **包管理**：使用 `go mod` 管理依赖
3. **并发编程**：goroutine 和 channel
4. **标准库**：fmt、net/http、database/sql 等

## 参考资源

- [Go 官方文档](https://golang.org/doc/)
- [Go 中文网](https://studygolang.com/)
- [Go by Example](https://gobyexample.com/)

