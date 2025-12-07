# Cursor IDE 快捷键指南

## 通用快捷键

### 文件操作
- `Ctrl + N` - 新建文件
- `Ctrl + O` - 打开文件
- `Ctrl + S` - 保存文件
- `Ctrl + Shift + S` - 另存为
- `Ctrl + W` - 关闭当前标签页
- `Ctrl + K W` - 关闭所有标签页
- `Ctrl + Shift + T` - 重新打开已关闭的文件

### 编辑操作
- `Ctrl + Z` - 撤销
- `Ctrl + Y` / `Ctrl + Shift + Z` - 重做
- `Ctrl + X` - 剪切当前行（未选中文本时）
- `Ctrl + C` - 复制当前行（未选中文本时）
- `Ctrl + V` - 粘贴
- `Alt + ↑/↓` - 向上/下移动当前行
- `Shift + Alt + ↑/↓` - 向上/下复制当前行
- `Ctrl + Shift + K` - 删除当前行
- `Ctrl + Enter` - 在下方插入新行
- `Ctrl + Shift + Enter` - 在上方插入新行

### 多光标编辑
- `Alt + Click` - 添加光标
- `Ctrl + Alt + ↑/↓` - 在上/下添加光标
- `Ctrl + D` - 选中下一个相同单词（多选）
- `Ctrl + Shift + L` - 选中所有相同单词
- `Ctrl + U` - 撤销上一次光标操作

### 查找和替换
- `Ctrl + F` - 在当前文件中查找
- `Ctrl + H` - 在当前文件中替换
- `Ctrl + Shift + F` - 在整个工作区中查找
- `Ctrl + Shift + H` - 在整个工作区中替换
- `F3` / `Shift + F3` - 查找下一个/上一个
- `Ctrl + Shift + .` - 替换焦点（在查找框和编辑器间切换）

### 代码导航
- `Ctrl + P` - 快速打开文件（按名称搜索）
- `Ctrl + Shift + P` / `F1` - 命令面板
- `Ctrl + G` - 跳转到指定行
- `Ctrl + T` - 跳转到符号（函数、变量等）
- `F12` - 跳转到定义
- `Alt + F12` - 查看定义（悬浮窗）
- `Ctrl + Click` - 跳转到定义
- `Ctrl + -` - 返回上一位置
- `Ctrl + Shift + -` - 前进到下一位置
- `Ctrl + Shift + O` - 在当前文件中查找符号

### 代码格式化
- `Shift + Alt + F` - 格式化整个文档
- `Ctrl + K Ctrl + F` - 格式化选中内容

### 注释
- `Ctrl + /` - 切换行注释
- `Shift + Alt + A` - 切换块注释

### 代码折叠
- `Ctrl + Shift + [` - 折叠区域
- `Ctrl + Shift + ]` - 展开区域
- `Ctrl + K Ctrl + 0` - 折叠所有区域
- `Ctrl + K Ctrl + J` - 展开所有区域

## Go 开发专用快捷键

### 代码补全和提示
- `Ctrl + Space` - 触发代码补全
- `Ctrl + Shift + Space` - 触发参数提示
- `Tab` - 接受代码补全建议
- `Ctrl + .` - 快速修复 / 显示建议操作

### Go 语言服务器 (gopls)
- `F2` - 重命名符号
- `Alt + Shift + F12` - 查找所有引用
- `Ctrl + K F2` - 更改所有匹配项

### 运行和调试
- `F5` - 开始调试
- `Ctrl + F5` - 运行（不调试）
- `F9` - 切换断点
- `F10` - 单步跳过
- `F11` - 单步进入
- `Shift + F11` - 单步跳出
- `Ctrl + Shift + F5` - 重启调试

### 终端
- `` Ctrl + ` `` - 切换终端显示/隐藏
- `Ctrl + Shift + ` ` `` - 创建新终端
- `Ctrl + Shift + ↑/↓` - 滚动终端输出

## 界面操作

### 侧边栏
- `Ctrl + B` - 切换侧边栏显示/隐藏
- `Ctrl + Shift + E` - 切换到资源管理器
- `Ctrl + Shift + F` - 切换到搜索
- `Ctrl + Shift + G` - 切换到源代码管理
- `Ctrl + Shift + D` - 切换到调试

### 编辑器布局
- `Ctrl + \` - 分割编辑器
- `Ctrl + 1/2/3` - 切换到第 1/2/3 编辑器组
- `Ctrl + K Ctrl + ←/→` - 在编辑器组间移动

### 标签页
- `Ctrl + Tab` - 切换到下一个标签页
- `Ctrl + Shift + Tab` - 切换到上一个标签页
- `Ctrl + K Ctrl + ←/→` - 在编辑器间移动
- `Ctrl + K P` - 复制当前文件路径

## Cursor AI 特有功能

### AI 聊天和代码生成
- `Ctrl + L` - 打开/聚焦 Cursor Tab（AI 聊天）
- `Ctrl + K` - 内联编辑（选中代码后使用）
- `Ctrl + Shift + L` - 在新标签页打开 Cursor Tab
- `Tab` - 接受 AI 建议（代码生成时）

### 代码操作
- `Ctrl + I` - 打开 Composer（组合式 AI 编辑）
- `Alt + C` - 在光标处插入代码（Cursor Tab）
- `Alt + Z` - 打开上下文菜单（AI 建议）

## 实用技巧

### 针对你当前的代码（map 操作）

在编辑 Go 代码时，特别是处理 map 时：

1. **快速重命名变量**
   - 选中 `map` → 按 `F2` → 输入新名称（如 `fruitMap`）

2. **查找所有使用**
   - 选中 `map` → 按 `Alt + Shift + F12` → 查看所有引用

3. **多光标编辑**
   - 选中 `map` → 按 `Ctrl + D` → 选中下一个相同单词
   - 然后同时编辑多个位置

4. **格式化代码**
   - 按 `Shift + Alt + F` 自动格式化 Go 代码

5. **跳转到定义**
   - 在 `map` 上按 `F12` → 跳转到变量定义
   - 或 `Ctrl + Click` 点击变量名

### 常见工作流

**编写 Go 函数：**
1. 输入函数名 → `Ctrl + Space` 触发补全
2. 输入参数 → `Ctrl + Shift + Space` 查看参数提示
3. 写完函数 → `Shift + Alt + F` 格式化
4. `Ctrl + S` 保存 → `Ctrl + ` ` 打开终端 → `go run main.go` 运行

**调试代码：**
1. 在行号左侧点击设置断点（或按 `F9`）
2. 按 `F5` 开始调试
3. 使用 `F10`、`F11` 单步执行

**重构代码：**
1. 选中变量/函数名
2. 按 `F2` 重命名（会自动更新所有引用）
3. 或按 `Ctrl + .` 查看更多重构选项

## 自定义快捷键

可以通过以下方式自定义快捷键：

1. 按 `Ctrl + Shift + P` 打开命令面板
2. 输入 "Preferences: Open Keyboard Shortcuts"
3. 搜索要修改的命令
4. 双击设置新的快捷键

## 提示

- 大多数快捷键可以在命令面板（`Ctrl + Shift + P`）中找到
- 如果快捷键冲突，可以在快捷键设置中查看和修改
- Cursor 基于 VS Code，所以大部分 VS Code 快捷键都适用
- 按住 `Ctrl + K` 然后按 `Ctrl + S` 可以打开快捷键参考卡片

