# kutil

一个用于Kubernetes资源管理的命令行工具

## 功能

- 创建、修改、删除和查询Kubernetes部署的资源限制
- 支持CPU和内存资源的配置
- 支持命名空间级别的操作

## 安装

```bash
go install github.com/barrettzjh/kutil@latest
```

## 使用示例

### 修改资源限制
```bash
kutil resource modify my-deploy 100Mi -n default -l limit -t memory
```

### 创建资源限制
```bash
kutil resource create my-deploy '{"limits":{"cpu":"1","memory":"1Gi"},"requests":{"cpu":"0.5","memory":"512Mi"}}' -n default
```

### 删除资源限制
```bash
kutil resource delete my-deploy -n default
```

### 查询资源限制
```bash
kutil resource list -n default
```