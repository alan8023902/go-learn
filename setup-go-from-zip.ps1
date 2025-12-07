# Go .zip 安装后环境变量配置脚本
# 以管理员身份运行 PowerShell，然后执行此脚本

param(
    [string]$GoRootPath = "C:\Program Files\Go"
)

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "Go .zip 安装后环境变量配置" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# 检查 Go 是否已安装
$goExePath = Join-Path $GoRootPath "bin\go.exe"
if (-not (Test-Path $goExePath)) {
    Write-Host "错误：未找到 Go 可执行文件！" -ForegroundColor Red
    Write-Host "请确认 Go 已解压到: $GoRootPath" -ForegroundColor Yellow
    Write-Host "或者使用 -GoRootPath 参数指定正确的路径" -ForegroundColor Yellow
    Write-Host "例如: .\setup-go-from-zip.ps1 -GoRootPath 'D:\Go'" -ForegroundColor Yellow
    exit 1
}

Write-Host "检测到 Go 安装路径: $GoRootPath" -ForegroundColor Green
Write-Host ""

# 1. 配置 GOROOT
Write-Host "[1/4] 配置 GOROOT 环境变量..." -ForegroundColor Yellow
$currentGOROOT = [System.Environment]::GetEnvironmentVariable("GOROOT", "Machine")
if ($currentGOROOT -ne $GoRootPath) {
    [System.Environment]::SetEnvironmentVariable("GOROOT", $GoRootPath, "Machine")
    Write-Host "  已设置 GOROOT = $GoRootPath" -ForegroundColor Green
} else {
    Write-Host "  GOROOT 已正确设置" -ForegroundColor Green
}

# 2. 将 Go\bin 添加到 PATH
Write-Host "[2/4] 配置 PATH 环境变量..." -ForegroundColor Yellow
$goBinPath = Join-Path $GoRootPath "bin"
$currentPath = [System.Environment]::GetEnvironmentVariable("Path", "Machine")

if ($currentPath -notlike "*$goBinPath*") {
    $newPath = "$currentPath;$goBinPath"
    [System.Environment]::SetEnvironmentVariable("Path", $newPath, "Machine")
    Write-Host "  已将 $goBinPath 添加到 PATH" -ForegroundColor Green
} else {
    Write-Host "  PATH 中已包含 $goBinPath" -ForegroundColor Green
}

# 3. 配置 GOPATH（工作空间）
Write-Host "[3/4] 配置 GOPATH 环境变量..." -ForegroundColor Yellow
$gopath = "$env:USERPROFILE\go"

# 创建 Go 工作空间目录
if (-not (Test-Path $gopath)) {
    New-Item -ItemType Directory -Force -Path $gopath | Out-Null
    Write-Host "  已创建目录: $gopath" -ForegroundColor Green
}

# 设置 GOPATH（用户变量）
[System.Environment]::SetEnvironmentVariable("GOPATH", $gopath, "User")
Write-Host "  已设置 GOPATH = $gopath" -ForegroundColor Green

# 将 GOPATH\bin 添加到用户 PATH
$userPath = [System.Environment]::GetEnvironmentVariable("Path", "User")
$gopathBin = "$gopath\bin"

if ($userPath -notlike "*$gopathBin*") {
    $newUserPath = if ($userPath) { "$userPath;$gopathBin" } else { $gopathBin }
    [System.Environment]::SetEnvironmentVariable("Path", $newUserPath, "User")
    Write-Host "  已将 $gopathBin 添加到用户 PATH" -ForegroundColor Green
} else {
    Write-Host "  用户 PATH 中已包含 $gopathBin" -ForegroundColor Green
}

# 4. 配置 Go 模块代理
Write-Host "[4/4] 配置 Go 模块代理..." -ForegroundColor Yellow
$goproxy = "https://goproxy.cn,direct"
[System.Environment]::SetEnvironmentVariable("GOPROXY", $goproxy, "User")
Write-Host "  已设置 GOPROXY = $goproxy" -ForegroundColor Green

# 设置 GOSUMDB（可选）
$gosumdb = "sum.golang.google.cn"
[System.Environment]::SetEnvironmentVariable("GOSUMDB", $gosumdb, "User")
Write-Host "  已设置 GOSUMDB = $gosumdb" -ForegroundColor Green

Write-Host ""
Write-Host "========================================" -ForegroundColor Cyan
Write-Host "配置完成！" -ForegroundColor Green
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""
Write-Host "重要提示：" -ForegroundColor Yellow
Write-Host "1. 请关闭所有 PowerShell 窗口" -ForegroundColor White
Write-Host "2. 重新打开一个新的 PowerShell 窗口" -ForegroundColor White
Write-Host "3. 运行以下命令验证安装：" -ForegroundColor White
Write-Host ""
Write-Host '   go version' -ForegroundColor Cyan
Write-Host '   go env GOROOT' -ForegroundColor Cyan
Write-Host '   go env GOPATH' -ForegroundColor Cyan
Write-Host '   go env GOPROXY' -ForegroundColor Cyan
Write-Host ""

