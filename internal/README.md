# internal

这里建议放 **真正的 Go 业务逻辑**。

这是学习 Go 的主战场，建议不要把逻辑都写在 Wails 绑定层。

推荐子目录方向：

- `internal/bootstrap/`：启动初始化
- `internal/settings/`：配置管理
- `internal/storage/`：本地存储与持久化
- `internal/tools/`：各工具业务逻辑
- `internal/ai/`：AI 相关服务
- `internal/common/`：通用能力

建议原则：

- UI 放前端
- 业务放 Go
- 绑定层只做适配与桥接
