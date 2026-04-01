# frontend

这里将来放 Vue 3 + TypeScript + Vite 前端工程。

建议重点目录：

- `src/views/`：页面级视图
- `src/components/`：通用组件
- `src/stores/`：Pinia 状态
- `src/router/`：路由定义
- `src/services/`：前端调用 Wails / API 的适配层
- `src/types/`：前端共享类型

注意：

- 组件负责展示与交互
- 复杂业务规则尽量下沉到 Go 服务层，不要全部堆在前端
