---
title: Chinese_Learning_App v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.17"

---

# Chinese_Learning_App

> v1.0.0

Base URLs:

* <a href="http://127.0.0.1:8888">开发环境: http://127.0.0.1:8888</a>

# Authentication

# 健康检测接口

## GET PublicGroup的health检测

GET /health

- 用于检测PublicGroup的health

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET PrivateGroup的hello

GET /hello

# PrivateGroup的hello接口

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# TeachingVideoApi

## POST uploadApi

POST /teachingVideo/upload

- sys_teaching_video API的上传视频功能

> Body 请求参数

```yaml
video: string
videoIcon: string
videoName: string
group: string
videoType: string

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» video|body|string(binary)| 否 |none|
|» videoIcon|body|string(binary)| 否 |none|
|» videoName|body|string| 否 |none|
|» group|body|string| 否 |none|
|» videoType|body|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST getTeachingVideoListApi

POST /teachingVideo/getTeachingVideoList

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST SearchAllParentVideoList

POST /teachingVideo/getTeachingParentVideoList

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# user接口

## POST 注册接口

POST /user/register

userService的接口

> Body 请求参数

```json
{}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 登录接口

POST /user/login

# 用于用户登录

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 娱乐视频接口

## GET 获取视频的接口

GET /amusement/get

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 上传娱乐视频接口

POST /amusement/upload

# 上传娱乐视频API

> Body 请求参数

```yaml
video: file://D:\video\2023-10-31 20-46-37.mp4
videoIcon: file://D:\images\}XMF_AT$VVVU7E%AJ11YL_D.jpg
videoName: amusement-测试2
type: "1"
content: amusement-测试2

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» video|body|string(binary)| 是 |none|
|» videoIcon|body|string(binary)| 是 |none|
|» videoName|body|string| 是 |none|
|» type|body|string| 是 |none|
|» content|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

