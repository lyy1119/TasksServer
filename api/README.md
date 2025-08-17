# API 约定与变更记录

本文件用于记录 **TasksServer** 的 API 设计约定与变更历史。  
在开发过程中，请按照本文的规范撰写和修改 `openapi.yaml`，并在每次修改后更新“变更记录”部分。

---

## 基本规范

### 1. OpenAPI 版本
- 使用 **OpenAPI 3.0.3** 规范 (`openapi: 3.0.3`)。

### 2. 路径与方法约定
- 路径命名统一使用 **小写 + 复数**，如：  
  - `/users` → 用户集合  
  - `/tasks` → 任务集合
- 路径参数统一使用大括号：  
  - `/users/{id}`  
  - `/tasks/{id}`
- 不使用 `/xxx/get`、`/xxx/create` 这类冗余路径，通过 **HTTP 方法** 表达动作：
  - `GET /users` → 获取用户列表  
  - `POST /users` → 创建用户  
  - `GET /users/{id}` → 获取单个用户  
  - `PUT /users/{id}` → 更新用户  
  - `DELETE /users/{id}` → 删除用户

### 3. 请求与响应格式
- 统一使用 `application/json`。
- 所有请求体（POST/PUT）必须在 `components/schemas` 定义数据结构，并通过 `$ref` 引用。
- 所有响应必须定义状态码及数据结构，至少包含：
  - `200 OK`（成功返回数据）
  - `201 Created`（创建成功）
  - `204 No Content`（删除成功）
  - `400 Bad Request`（请求错误）
  - `404 Not Found`（资源不存在）

### 4. 数据模型（Schemas）
- 在 `components/schemas` 下定义可复用的对象。
- 示例：
  ```yaml
  components:
    schemas:
      User:
        type: object
        properties:
          id:
            type: integer
          username:
            type: string
          email:
            type: string
            format: email

## v1.0.0 - 2025-08-17
### Added
- 定义 `/users` 相关接口：
  - `GET /users` 获取所有用户
  - `POST /users` 创建用户
  - `GET /users/{id}` 获取单个用户
  - `PUT /users/{id}` 更新用户
  - `DELETE /users/{id}` 删除用户
- 定义 `/tasks` 相关接口：
  - `GET /tasks` 获取所有任务
  - `POST /tasks` 创建任务
  - `GET /tasks/{id}` 获取单个任务
  - `PUT /tasks/{id}` 更新任务
  - `DELETE /tasks/{id}` 删除任务
- 数据模型（Schemas）：
  - `User`, `UserCreate`, `UserUpdate`
  - `Task`, `TaskCreate`, `TaskUpdate`
- 服务器配置：
  - `http://localhost:8080`（开发环境）
  - `https://api.example.com`（生产环境）
