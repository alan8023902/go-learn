# 修复 gopls 安装问题 - 配置代理并安装 gopls
# 以管理员身份运行 PowerShell

Write-Host "========================================" -ForegroundColor Cyan
Write-Host "修复 gopls 安装问题" -ForegroundColor Cyan
Write-Host "========================================" -ForegroundColor Cyan
Write-Host ""

# 检查 Go 是否可用
$goPath = "C:\Program Files\Go\bin\go.exe"
if (-not (Test-Path $goPath)) {
    Write-Host "错误：未找到 Go 可执行文件！" -ForegroundColor Red
    Write-Host "请确认 Go 已正确安装到: C:\Program Files\Go" -ForegroundColor Yellow
    exit 1
}

Write-Host "[1/3] 配置 Go 模块代理..." -ForegroundColor Yellow

# 设置 GOPROXY（当前会话）
$env:GOPROXY = "https://goproxy.cn,direct"
$env:GOSUMDB = "sum.golang.google.cn"

Write-Host "  已设置 GOPROXY = $env:GOPROXY" -ForegroundColor Green

# 永久设置 GOPROXY
[System.Environment]::SetEnvironmentVariable("GOPROXY", "https://goproxy.cn,direct", "User")
[System.Environment]::SetEnvironmentVariable("GOSUMDB", "sum.golang.google.cn", "User")
Write-Host "  已永久设置环境变量" -ForegroundColor Green

Write-Host ""
Write-Host "[2/3] 验证 Go 环境..." -ForegroundColor Yellow

# 验证 Go 命令
& $goPath env GOPROXY
$currentProxy = & $goPath env GOPROXY
if ($currentProxy -notlike "*goproxy.cn*") {
    Write-Host "  警告：GOPROXY 可能未正确设置" -ForegroundColor Yellow
    Write-Host "  当前值: $currentProxy" -ForegroundColor Yellow
}

Write-Host ""
Write-Host "[3/3] 安装 gopls..." -ForegroundColor Yellow

# 安装 gopls
Write-Host "  正在下载并安装 gopls，请稍候..." -ForegroundColor Cyan
& $goPath install golang.org/x/tools/gopls@latest

if ($LASTEXITCODE -eq 0) {
    Write-Host ""
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host "gopls 安装成功！" -ForegroundColor Green
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "如果 IDE 仍然无法使用 gopls，请：" -ForegroundColor Yellow
    Write-Host "1. 重启 IDE" -ForegroundColor White
    Write-Host "2. 检查 gopls 是否在 PATH 中：$env:USERPROFILE\go\bin\gopls.exe" -ForegroundColor White
} else {
    Write-Host ""
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host "gopls 安装失败" -ForegroundColor Red
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host ""
    Write-Host "可能的解决方案：" -ForegroundColor Yellow
    Write-Host "1. 检查网络连接" -ForegroundColor White
    Write-Host "2. 尝试其他代理：" -ForegroundColor White
    Write-Host "   - 阿里云: https://mirrors.aliyun.com/goproxy/,direct" -ForegroundColor Cyan
    Write-Host "   - 中科大: https://goproxy.io,direct" -ForegroundColor Cyan
    Write-Host "3. 手动设置代理后重试：" -ForegroundColor White
    Write-Host '   $env:GOPROXY = "https://goproxy.cn,direct"' -ForegroundColor Cyan
    Write-Host '   go install golang.org/x/tools/gopls@latest' -ForegroundColor Cyan
}

