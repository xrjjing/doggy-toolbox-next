<script setup lang="ts">
import { onMounted, reactive } from 'vue'
import { Greet, GetAppInfo } from '../wailsjs/go/main/App'

type AppInfo = {
  name: string
  stack: string
  status: string
}

const state = reactive({
  name: '',
  greetText: '点击按钮，验证 Vue 到 Go 的调用链。',
  info: {
    name: 'doggy-toolbox-next',
    stack: 'Go + Wails + Vue 3 + TypeScript',
    status: '正在读取...',
  } as AppInfo,
})

async function loadInfo() {
  state.info = await (GetAppInfo() as Promise<AppInfo>)
}

async function greet() {
  state.greetText = await Greet(state.name)
}

onMounted(() => {
  loadInfo().catch(() => {
    state.info.status = '读取失败，请检查 Wails 绑定是否正常'
  })
})
</script>

<template>
  <div class="shell">
    <aside class="sidebar">
      <div>
        <div class="brand">doggy-toolbox-next</div>
        <p class="sidebar-subtitle">重写白板 · 独立新仓库</p>
      </div>
      <nav class="nav">
        <button class="nav-item active">欢迎页</button>
        <button class="nav-item" disabled>设置（待实现）</button>
        <button class="nav-item" disabled>工具模块（待迁移）</button>
      </nav>
    </aside>

    <main class="content">
      <section class="hero card">
        <span class="eyebrow">Phase 1 / Runnable Baseline</span>
        <h1>Go + Wails + Vue 3 + TypeScript 已跑通到最小基线</h1>
        <p>
          这个页面的目标不是功能完整，而是先确认：目录、前后端绑定、Wails 构建链、Vue 页面渲染都已经具备可继续开发的基础。
        </p>
      </section>

      <section class="grid">
        <article class="card">
          <h2>应用信息</h2>
          <ul class="kv-list">
            <li><span>名称</span><strong>{{ state.info.name }}</strong></li>
            <li><span>技术栈</span><strong>{{ state.info.stack }}</strong></li>
            <li><span>状态</span><strong>{{ state.info.status }}</strong></li>
          </ul>
        </article>

        <article class="card">
          <h2>调用链验证</h2>
          <p class="muted">输入任意名字，点击按钮，验证 Vue → Wails → Go 是否正常。</p>
          <div class="action-row">
            <input v-model="state.name" class="input" type="text" placeholder="输入一个名字，例如 xrj" />
            <button class="btn" @click="greet">调用 Go 方法</button>
          </div>
          <p class="result">{{ state.greetText }}</p>
        </article>
      </section>
    </main>
  </div>
</template>
