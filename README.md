# gin_learning

## 目录结构

project-root/
├── cmd/                  # 入口文件
│   └── main.go
├── internal/             # 私有实现
│   ├── api/              # 路由层
│   │   ├── handler/      # 控制器
│   │   ├── middleware/   # 中间件
│   │   └── router.go     # 路由配置
│   ├── service/          # 业务逻辑层
│   ├── repository/       # 数据访问层
│   ├── model/            # 数据模型
│   ├── config/           # 配置管理
│   └── pkg/              # 内部工具库
├── pkg/                  # 公共库
├── test/                 # 测试文件
├── scripts/              # 部署脚本
├── configs/              # 配置文件
├── docs/                 # API文档
├── web/                  # 前端资源
├── go.mod
└── Dockerfile


