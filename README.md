## 项目结构
go-store/
├── api/          # 路由层
│   └── v1/       # 接口版本
├── service/      # 业务逻辑层
├── dao/          # 数据访问层
├── model/        # 数据模型
├── serializer/   # 序列化器
├── middleware/   # 中间件
├── pkg/          # 工具包
│   ├── mytools/  # 加密/JWT/日志
│   └── e/        # 错误码
└── conf/         # 配置管理
