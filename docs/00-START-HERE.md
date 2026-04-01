# 00 START HERE

## 1. 当前目标

你现在不是要立刻把旧项目全重写完，而是要先完成这三件事：

1. **跑通新技术栈**：Go + Wails + Vue 3 + TypeScript
2. **建立清晰工程骨架**：目录、配置、启动方式、开发流程
3. **把重写任务拆成可执行阶段**：先简单，后复杂

---

## 2. 当前机器上的已知现状

基于初始化时的本地检查：

- Go：`go1.24.3`
- Node：`v23.11.0`
- npm：`11.6.2`
- Wails CLI：**未安装**

这意味着：

- Go / Node / npm 基本够用
- 你真正还缺的是 **Wails CLI**

---
## 2.5 当前仓库边界

这个项目现在已经是独立 GitHub 仓库：

- `https://github.com/xrjjing/doggy-toolbox-next`

请把它理解为：

- 旧项目 `doggy-toolbox`：功能参考与迁移对照
- 新项目 `doggy-toolbox-next`：正式重写主战场

---

## 3. 为什么这次选 Wails + Vue 3 + TypeScript + Go

### Wails

- 让 Go 负责桌面应用后端 / 系统能力 / 业务逻辑
- 让前端继续用现代 Web 技术写界面
- 和你当前 `PyWebView + HTML/JS` 的思路接近，迁移心智成本较低

### Vue 3

- 模板风格更接近 HTML
- 对非前端出身但要补前端的人更友好
- 社区资料多，AI 生成和审查都比较方便

### TypeScript

- 类型信息更清晰
- 对 Java 背景更友好
- 更容易审查 AI 生成代码

### Go

- 是你真正想补的语言
- 可以借这个项目练工程结构、接口设计、并发、文件处理、配置管理

---

## 4. 第一天就该做什么

按下面顺序做，不要跳：

### Step 1：确认 PATH

```bash
printf '%s\n' "$PATH" | tr ':' '\n' | grep 'go/bin'
```

如果没有输出，先把 `~/go/bin` 加到 shell 配置里。

### Step 2：安装 Wails CLI

官方安装命令：

```bash
go install github.com/wailsapp/wails/v2/cmd/wails@latest
```

### Step 3：验证 Wails CLI

```bash
wails version
wails doctor
```

> `wails doctor` 是官方建议的系统检查命令。

### Step 4：先生成一份官方脚手架作参考

因为当前仓库已经先建了白板目录，不建议直接在本目录里盲目跑脚手架覆盖。

推荐在临时目录先生成一份官方模板，对照学习：

```bash
mkdir -p /tmp/wails-sandbox
cd /tmp/wails-sandbox
wails init -n doggy-toolbox-next-scaffold -t vue-ts
```

你要看的不是它能不能直接拿来用，而是：

- 官方默认生成了哪些文件
- `main.go` 怎么写
- `wails.json` 怎么写
- `frontend/` 默认长什么样
- Vue + TS 默认入口文件在哪

### Step 5：对照本文档，把正式仓库按推荐结构补齐

不要一开始就迁业务。先把：

- 目录结构
- 配置文件位置
- 启动方式
- 最小示例页面
- 一个 Go 绑定方法

全部跑通。

---

## 5. Vue 3 的脚手架到底怎么装

Vue 官方推荐用 `create-vue`，命令是：

```bash
npm create vue@latest
```

如果你是单独做一个纯前端项目，就走这个命令。

但在这次重写里，更推荐你：

- **桌面壳用 Wails 初始化**
- **前端模板直接选 `vue-ts`**

也就是优先：

```bash
wails init -n doggy-toolbox-next -t vue-ts
```

而不是先 `npm create vue@latest` 再自己手动接 Wails。

原因：

- 少走一层集成成本
- 目录和配置更贴近 Wails 默认习惯
- 后续文档和社区经验更容易复用

---

## 6. 现阶段不要做的事

先不要：

- 直接迁移 AI 聊天
- 直接迁移节点转换
- 直接重写全部 CRUD 页面
- 先追求 UI 漂亮
- 先做多平台打包优化

先只要做到：

- 新项目能启动
- 有一个 Vue 页面
- 有一个 Go 方法可被前端调用
- 构建能成功

这才算真正开始。
