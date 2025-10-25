# urB00ks API 文档
#### 由apifox自动生成

## GET 查询用户信息

GET /user/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |用户的ID|

> 返回示例

> 200 Response

```json
{
  "status": 200,
  "msg": "Successfully get the user!",
  "data": {
    "name": "缪玲",
    "birth": "2003-04-05",
    "favs": [
      "in commodo minim Lorem occaecat",
      "et elit",
      "enim"
    ]
  }
}
```

> 400 Response

```json
{
  "status": 400,
  "msg": "failed to get a book",
  "data": 1
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» name|string|true|none||none|
|»» birth|string|true|none||none|
|»» favs|[string]|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

## GET 查询书籍信息

GET /book/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |书本的id|

> 返回示例

> 200 Response

```json
{
  "status": 31,
  "msg": "officia",
  "data": {
    "id": 102,
    "title": "大地巡旅",
    "author": "鹰角网络",
    "rating": 10,
    "url": "https://j.jpg"
  }
}
```

> 400 Response

```json
{
  "status": 400,
  "msg": "failed to get a book",
  "data": 1
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» id|integer|true|none||none|
|»» title|string|true|none||none|
|»» author|string|true|none||none|
|»» rating|integer|true|none||none|
|»» url|string|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

## GET 获得一定数量的书籍

GET /book/list

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|num|query|string| 否 |获得所有书本|
|offset|query|integer| 否 |在多少书之后进行后续的获取|

> 返回示例

> 200 Response

```json
{
  "status": 200,
  "msg": "Get bookinfo successfully",
  "data": [
    {
      "id": 900,
      "title": "The Selfish Gene",
      "author": "Richard Dawkins",
      "rating": 4,
      "views": 8210,
      "url": "https://images.gr-assets.com/books/1366758096m/61535.jpg\r"
    },
    {
      "id": 901,
      "title": "The Amy",
      "author": "Andy Max",
      "rating": 4,
      "views": 8210,
      "url": "https://images.gr-assets.com/books/1366758096m/61525.jpg\r"
    }
  ]
}
```

> 400 Response

```json
{
  "status": 400,
  "msg": "failed to get a book",
  "data": 1
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|[object]|true|none||none|
|»» id|integer|true|none||none|
|»» title|string|true|none||none|
|»» author|string|true|none||none|
|»» rating|integer|true|none||none|
|»» views|integer|true|none||none|
|»» url|string|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

## GET 获得书库内书的数量

GET /book/count

> 返回示例

> 200 Response

```json
{
  "status": 200,
  "msg": "Successfully get the count",
  "data": 10
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

## POST 创建用户

POST /user/register

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|name|query|string| 否 |none|
|birth|query|string| 否 |none|
|pswd|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status": 200,
  "msg": "Successfully create a user!",
  "data": 0
}
```

> 400 Response

```json
{
  "status": 400,
  "msg": "Failed to create a user!",
  "data": 1
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

## POST 新增书本

POST /book/add

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|title|query|string| 否 |none|
|author|query|string| 否 |none|
|url|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status": 200,
  "msg": "Successfully create a book!",
  "data": 0
}
```

> 400 Response

```json
{
  "status": 400,
  "msg": "Failed to create a book!",
  "data": 1
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

## DELETE 删除书本

DELETE /book/remove/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status": 200,
  "msg": "successfully delete!",
  "data": 0
}
```

> 400 Response

```json
{
  "status": 400,
  "msg": "delete failed",
  "data": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

## DELETE 删除用户

DELETE /user/remove/{id}

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|id|path|integer| 是 |none|

> 返回示例

> 200 Response

```json
{
  "status": 200,
  "msg": "successfully delete!",
  "data": 0
}
```

> 400 Response

```json
{
  "status": 400,
  "msg": "delete failed",
  "data": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

## PATCH 更新用户相关信息

PATCH /user/update

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|target|query|integer| 否 |要修改的用户ID|
|type|query|string| 否 |更新类型，是更新用户名称还是用户密码 etc.|
|content|query|string| 否 |更新的字段内容|

> 返回示例

> 200 Response

```json
{
  "status": 200,
  "msg": "update name succfully",
  "data": 0
}
```

> 400 Response

```json
{
  "status": 400,
  "msg": "update failed",
  "data": 0
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

## PATCH 更新书本相关信息

PATCH /book/update

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|target|query|integer| 否 |要修改的书本ID|
|type|query|string| 否 |更新类型，是更新书本标题还是书本作者 etc.|
|content|query|string| 否 |更新的字段内容|
|rating|query|integer| 否 |none|
|views|query|integer| 否 |none|
|url|query|string| 否 |none|

> 返回示例

> 200 Response

```json
{
  "status": 200,
  "msg": "update name succfully",
  "data": 0
}
```

> 400 Response

```json
{
  "status": 400,
  "msg": "update failed",
  "data": 1
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» status|integer|true|none||none|
|» msg|string|true|none||none|
|» data|integer|true|none||none|

# 数据模型

