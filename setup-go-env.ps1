# Go 环境配置脚本
# 以管理员身份运行 PowerShell，然后执行此脚本

Write-Host "开始配置 Go 环境变量..." -ForegroundColor Green

# 1. 配置 GOPATH
$gopath = "$env:USERPROFILE\go"
Write-Host "设置 GOPATH: $gopath" -ForegroundColor Yellow

# 创建 Go 工作空间目录
if (-not (Test-Path $gopath)) {
    New-Item -ItemType Directory -Force -Path $gopath | Out-Null
    Write-Host "已创建目录: $gopath" -ForegroundColor Green
} else {
    Write-Host "目录已存在: $gopath" -ForegroundColor Yellow
}

# 设置 GOPATH 环境变量
[System.Environment]::SetEnvironmentVariable("GOPATH", $gopath, "User")
Write-Host "已设置 GOPATH 环境变量" -ForegroundColor Green

# 2. 将 GOPATH\bin 添加到 PATH
$currentPath = [System.Environment]::GetEnvironmentVariable("Path", "User")
$binPath = "$gopath\bin"

if ($currentPath -notlike "*$binPath*") {
    $newPath = "$currentPath;$binPath"
    [System.Environment]::SetEnvironmentVariable("Path", $newPath, "User")
    Write-Host "已将 $binPath 添加到 PATH" -ForegroundColor Green
} else {
    Write-Host "PATH 中已包含 $binPath" -ForegroundColor Yellow
}

# 3. 配置 Go 模块代理（使用七牛云代理，加速国内下载）
$goproxy = "https://goproxy.cn,direct"
[System.Environment]::SetEnvironmentVariable("GOPROXY", $goproxy, "User")
Write-Host "已设置 GOPROXY: $goproxy" -ForegroundColor Green

# 4. 设置 GOSUMDB（可选，用于模块校验）
$gosumdb = "sum.golang.google.cn"
[System.Environment]::SetEnvironmentVariable("GOSUMDB", $gosumdb, "User")
Write-Host "已设置 GOSUMDB: $gosumdb" -ForegroundColor Green

Write-Host ""
Write-Host "配置完成！" -ForegroundColor Green
Write-Host "请关闭当前 PowerShell 窗口，重新打开后运行以下命令验证：" -ForegroundColor Yellow
Write-Host '  go version' -ForegroundColor Cyan
Write-Host '  go env GOPATH' -ForegroundColor Cyan
Write-Host '  go env GOPROXY' -ForegroundColor Cyan
