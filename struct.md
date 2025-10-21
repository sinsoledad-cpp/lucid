/lucid (项目根目录)
├── app/
│   ├── user/                   // ⭐ 微服务: 用户服务 (一个完整的限界上下文)
│   │   │
│   │   ├── api/                // 1. 应用层 (Application): API 接入 (go-zero)
│   │   │   ├── etc/
│   │   │   │   └── user.yaml
│   │   │   ├── internal/
│   │   │   │   ├── config/
│   │   │   │   │   └── config.go
│   │   │   │   ├── handler/
│   │   │   │   │   ├── routes.go
│   │   │   │   │   └── user/
│   │   │   │   │       ├── loginhandler.go
│   │   │   │   │       └── registerhandler.go
│   │   │   │   ├── logic/
│   │   │   │   │   └── user/
│   │   │   │   │       ├── loginlogic.go
│   │   │   │   │       └── registerlogic.go
│   │   │   │   ├── svc/
│   │   │   │   │   └── servicecontext.go  // ❗ 重要: 这里会注入 domain.repository 接口
│   │   │   │   └── types/
│   │   │   │       └── types.go
│   │   │   ├── user.api
│   │   │   └── user.go               // API 服务启动入口
│   │   │
│   │   ├── rpc/                // 1. 应用层 (Application): RPC 接入 (go-zero)
│   │   │   ├── etc/
│   │   │   │   └── user.yaml
│   │   │   ├── internal/
│   │   │   │   ├── config/
│   │   │   │   │   └── config.go
│   │   │   │   ├── logic/
│   │   │   │   │   └── getuserinfologic.go
│   │   │   │   ├── server/
│   │   │   │   │   └── userserver.go
│   │   │   │   └── svc/
│   │   │   │       └── servicecontext.go
│   │   │   └── user.go               // RPC 服务启动入口
│   │   │
│   │   ├── domain/             // 2. 领域层 (Domain): 核心业务 (手写)
│   │   │   ├── entity/
│   │   │   │   └── user.go         //    - 定义领域实体 (e.g., User struct)
│   │   │   ├── repository/
│   │   │   │   └── user_repository.go //    - 定义仓储接口 (e.g., UserRepository interface)
│   │   │   └── service/
│   │   │       └── user_service.go    //    - 复杂的领域服务 (e.g., 注册时的密码策略)
│   │   │
│   │   ├── data/               // 3. 基础设施层 (Infrastructure): user 服务的具体实现
│   │   │   ├── model/          //    - goctl 生成的 DB 模型 (PO) (原 data/model/user)
│   │   │   │   ├── usersmodel.go
│   │   │   │   ├── usersmodel_gen.go
│   │   │   │   └── vars.go
│   │   │   └── repository/     //    - 仓储接口的 *实现* (手写)
│   │   │       └── user_repo_impl.go //    - (e.g., 实现 UserRepository, 内部调用 model)
│   │   │
│   │   └── schema/             // 4. 数据库定义 (Migrations)
│   │       └── sql/            //    - (原 schema/sql/user)
│   │           └── 000001_create_users_table.up.sql
│   │
│   └── domain-service/              // ⭐ 微服务: 同上
│
├── common/                     // 通用共享库 (e.g., utils, common errors, log wrappers)
│   │                           // ❗ 重要: 禁止包含任何服务的业务逻辑
│   └── (空)
│
├── gen/                        // ⭐ 生成的客户端SDK (e.g., gRPC clients)
│   └── go/
│       └── user/
│           └── v1/
│               ├── user.pb.go
│               ├── user_grpc.pb.go
│               └── user/           // zRPC 客户端
│                   └── user.go
│
├── protos/                     // ⭐ 服务契约 (公开的 .proto 文件)
│   └── user/
│       └── v1/
│           └── user.proto
│
├── scripts/                    // 自动化脚本
│   └── docker-compose.yml      // (原 deployments/docker-compose.yml)
│
├── go.mod                      // 主项目的 Go Module
├── go.sum
├── main.go                     // (根目录下的 main.go, 可能是测试或空文件)
├── Makefile                    // (需要更新 goctl 的 -dir 和 -src 路径)
└── README.md