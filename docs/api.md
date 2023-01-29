---
title: Gym v1.0.0
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

# Gym

> v1.0.0

Base URLs:

# Facility

## GET Get All Facilities

GET /facility

> 返回示例

> 200 Response

```json
"string"
```

### 返回结果

| 状态码 | 状态码含义                                                   | 说明  | 数据模型                                                                    |
| --- | ------------------------------------------------------- | --- | ----------------------------------------------------------------------- |
| 200 | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功  | [Gym-backend.api.v1.FacilityRes](#schemagym-backend.api.v1.facilityres) |

## POST Get Facility Detail

POST /facility_detail

> Body 请求参数

```json
{
  "id": 0
}
```

### 请求参数

| 名称   | 位置   | 类型                                                                                  | 必选  | 说明   |
| ---- | ---- | ----------------------------------------------------------------------------------- | --- | ---- |
| body | body | [Gym-backend.api.v1.FacilityDetailReq](#schemagym-backend.api.v1.facilitydetailreq) | 否   | none |

> 返回示例

> 200 Response

```json
"string"
```

### 返回结果

| 状态码 | 状态码含义                                                   | 说明  | 数据模型                                                                                |
| --- | ------------------------------------------------------- | --- | ----------------------------------------------------------------------------------- |
| 200 | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功  | [Gym-backend.api.v1.FacilityDetailRes](#schemagym-backend.api.v1.facilitydetailres) |

## POST Add Facility

POST /add_facility

> Body 请求参数

```json
{
  "name": "string",
  "description": "string",
  "location": "string",
  "cost": 0
}
```

### 请求参数

| 名称   | 位置   | 类型                                                                            | 必选  | 说明   |
| ---- | ---- | ----------------------------------------------------------------------------- | --- | ---- |
| body | body | [Gym-backend.api.v1.AddFacilityReq](#schemagym-backend.api.v1.addfacilityreq) | 否   | none |

> 返回示例

> 200 Response

```json
"string"
```

### 返回结果

| 状态码 | 状态码含义                                                   | 说明  | 数据模型                                                                          |
| --- | ------------------------------------------------------- | --- | ----------------------------------------------------------------------------- |
| 200 | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功  | [Gym-backend.api.v1.AddFacilityRes](#schemagym-backend.api.v1.addfacilityres) |

## POST Get Facilities By searching tags

POST /facility_search

> Body 请求参数

```json
{
  "name": "string",
  "id": 0
}
```

### 请求参数

| 名称   | 位置   | 类型                                                                                  | 必选  | 说明   |
| ---- | ---- | ----------------------------------------------------------------------------------- | --- | ---- |
| body | body | [Gym-backend.api.v1.FacilitySearchReq](#schemagym-backend.api.v1.facilitysearchreq) | 否   | none |

> 返回示例

> 200 Response

```json
"string"
```

### 返回结果

| 状态码 | 状态码含义                                                   | 说明  | 数据模型                                                                                |
| --- | ------------------------------------------------------- | --- | ----------------------------------------------------------------------------------- |
| 200 | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功  | [Gym-backend.api.v1.FacilitySearchRes](#schemagym-backend.api.v1.facilitysearchres) |

## POST Modify Facility

POST /modify_facility

> Body 请求参数

```json
{
  "id": 0,
  "name": "string",
  "description": "string",
  "location": "string",
  "cost": 0
}
```

### 请求参数

| 名称   | 位置   | 类型                                                                                  | 必选  | 说明   |
| ---- | ---- | ----------------------------------------------------------------------------------- | --- | ---- |
| body | body | [Gym-backend.api.v1.ModifyFacilityReq](#schemagym-backend.api.v1.modifyfacilityreq) | 否   | none |

> 返回示例

> 200 Response

```json
"string"
```

### 返回结果

| 状态码 | 状态码含义                                                   | 说明  | 数据模型                                                                                |
| --- | ------------------------------------------------------- | --- | ----------------------------------------------------------------------------------- |
| 200 | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功  | [Gym-backend.api.v1.ModifyFacilityRes](#schemagym-backend.api.v1.modifyfacilityres) |

# Account

## POST Login

POST /login

> Body 请求参数

```json
{
  "Username": "string",
  "Password": "string"
}
```

### 请求参数

| 名称   | 位置   | 类型                                                                | 必选  | 说明   |
| ---- | ---- | ----------------------------------------------------------------- | --- | ---- |
| body | body | [Gym-backend.api.v1.LoginReq](#schemagym-backend.api.v1.loginreq) | 否   | none |

> 返回示例

> 200 Response

```json
"string"
```

### 返回结果

| 状态码 | 状态码含义                                                   | 说明  | 数据模型                                                              |
| --- | ------------------------------------------------------- | --- | ----------------------------------------------------------------- |
| 200 | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功  | [Gym-backend.api.v1.LoginRes](#schemagym-backend.api.v1.loginres) |

## GET Log out'

GET /logout

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "message": "string",
  "data": {},
  "redirect": "string"
}
```

### 返回结果

| 状态码 | 状态码含义                                                   | 说明  | 数据模型   |
| --- | ------------------------------------------------------- | --- | ------ |
| 200 | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功  | Inline |

### 返回数据结构

状态码 **200**

| 名称         | 类型             | 必选    | 约束   | 中文名 | 说明   |
| ---------- | -------------- | ----- | ---- | --- | ---- |
| » code     | integer(int)   | false | none |     | none |
| » message  | string(string) | false | none |     | none |
| » data     | object         | false | none |     | none |
| » redirect | string(string) | false | none |     | none |

## POST Register

POST /register

> Body 请求参数

```json
{
  "Username": "string",
  "Password": "string",
  "ConfirmPassword": "string",
  "Email": "string",
  "Gender": "string",
  "Phone": "string"
}
```

### 请求参数

| 名称   | 位置   | 类型                                                                      | 必选  | 说明   |
| ---- | ---- | ----------------------------------------------------------------------- | --- | ---- |
| body | body | [Gym-backend.api.v1.RegisterReq](#schemagym-backend.api.v1.registerreq) | 否   | none |

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "message": "string",
  "data": {},
  "redirect": "string"
}
```

### 返回结果

| 状态码 | 状态码含义                                                   | 说明  | 数据模型   |
| --- | ------------------------------------------------------- | --- | ------ |
| 200 | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功  | Inline |

### 返回数据结构

状态码 **200**

| 名称         | 类型             | 必选    | 约束   | 中文名 | 说明   |
| ---------- | -------------- | ----- | ---- | --- | ---- |
| » code     | integer(int)   | false | none |     | none |
| » message  | string(string) | false | none |     | none |
| » data     | object         | false | none |     | none |
| » redirect | string(string) | false | none |     | none |

## POST Get user's profile'

POST /profile

> Body 请求参数

```json
{}
```

### 请求参数

| 名称   | 位置   | 类型                                                                    | 必选  | 说明   |
| ---- | ---- | --------------------------------------------------------------------- | --- | ---- |
| body | body | [Gym-backend.api.v1.ProfileReq](#schemagym-backend.api.v1.profilereq) | 否   | none |

> 返回示例

> 200 Response

```json
"string"
```

### 返回结果

| 状态码 | 状态码含义                                                   | 说明  | 数据模型                                                                  |
| --- | ------------------------------------------------------- | --- | --------------------------------------------------------------------- |
| 200 | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功  | [Gym-backend.api.v1.ProfileRes](#schemagym-backend.api.v1.profileres) |

# File

## POST Upload file

POST /upload

> Body 请求参数

```yaml
file: string
name: string
type: string
```

### 请求参数

| 名称     | 位置   | 类型     | 必选  | 说明   |
| ------ | ---- | ------ | --- | ---- |
| body   | body | object | 否   | none |
| » file | body | string | 否   | none |
| » name | body | string | 是   | none |
| » type | body | string | 是   | none |

> 返回示例

> 200 Response

```json
"string"
```

### 返回结果

| 状态码 | 状态码含义                                                   | 说明  | 数据模型                                                                        |
| --- | ------------------------------------------------------- | --- | --------------------------------------------------------------------------- |
| 200 | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | 成功  | [Gym-backend.api.v1.FileUploadRes](#schemagym-backend.api.v1.fileuploadres) |

# 数据模型

<h2 id="tocS_Gym-backend.internal.model.entity.FacilityImage">Gym-backend.internal.model.entity.FacilityImage</h2>

<a id="schemagym-backend.internal.model.entity.facilityimage"></a>
<a id="schema_Gym-backend.internal.model.entity.FacilityImage"></a>
<a id="tocSgym-backend.internal.model.entity.facilityimage"></a>
<a id="tocsgym-backend.internal.model.entity.facilityimage"></a>

```json
{
  "id": 0,
  "path": "string",
  "uploadTime": "string",
  "facilityID": 0
}
```

### 属性

| 名称         | 类型                  | 必选    | 约束   | 中文名 | 说明   |
| ---------- | ------------------- | ----- | ---- | --- | ---- |
| id         | integer(int)        | false | none |     | none |
| path       | string(string)      | false | none |     | none |
| uploadTime | string(*gtime.Time) | false | none |     | none |
| facilityID | integer(int)        | false | none |     | none |

<h2 id="tocS_Gym-backend.internal.model.FacilityEntity">Gym-backend.internal.model.FacilityEntity</h2>

<a id="schemagym-backend.internal.model.facilityentity"></a>
<a id="schema_Gym-backend.internal.model.FacilityEntity"></a>
<a id="tocSgym-backend.internal.model.facilityentity"></a>
<a id="tocsgym-backend.internal.model.facilityentity"></a>

```json
{
  "facility": {
    "id": 0,
    "name": "string",
    "description": "string",
    "location": "string",
    "cost": 0
  },
  "images": [
    {
      "id": 0,
      "path": "string",
      "uploadTime": "string",
      "facilityID": 0
    }
  ]
}
```

### 属性

| 名称       | 类型                                                                                                          | 必选    | 约束   | 中文名 | 说明   |
| -------- | ----------------------------------------------------------------------------------------------------------- | ----- | ---- | --- | ---- |
| facility | [Gym-backend.internal.model.entity.Facility](#schemagym-backend.internal.model.entity.facility)             | false | none |     | none |
| images   | [[Gym-backend.internal.model.entity.FacilityImage](#schemagym-backend.internal.model.entity.facilityimage)] | false | none |     | none |

<h2 id="tocS_Gym-backend.api.v1.FileUploadRes">Gym-backend.api.v1.FileUploadRes</h2>

<a id="schemagym-backend.api.v1.fileuploadres"></a>
<a id="schema_Gym-backend.api.v1.FileUploadRes"></a>
<a id="tocSgym-backend.api.v1.fileuploadres"></a>
<a id="tocsgym-backend.api.v1.fileuploadres"></a>

```json
"string"
```

### 属性

| 名称     | 类型             | 必选    | 约束   | 中文名 | 说明   |
| ------ | -------------- | ----- | ---- | --- | ---- |
| Data   | object         | false | none |     | none |
| » url  | string(string) | false | none |     | none |
| » name | string(string) | false | none |     | none |

<h2 id="tocS_github.com.gogf.gf.v2.net.ghttp.UploadFile">github.com.gogf.gf.v2.net.ghttp.UploadFile</h2>

<a id="schemagithub.com.gogf.gf.v2.net.ghttp.uploadfile"></a>
<a id="schema_github.com.gogf.gf.v2.net.ghttp.UploadFile"></a>
<a id="tocSgithub.com.gogf.gf.v2.net.ghttp.uploadfile"></a>
<a id="tocsgithub.com.gogf.gf.v2.net.ghttp.uploadfile"></a>

```json
{}
```

### 属性

*None*

<h2 id="tocS_Gym-backend.api.v1.FileUploadReq">Gym-backend.api.v1.FileUploadReq</h2>

<a id="schemagym-backend.api.v1.fileuploadreq"></a>
<a id="schema_Gym-backend.api.v1.FileUploadReq"></a>
<a id="tocSgym-backend.api.v1.fileuploadreq"></a>
<a id="tocsgym-backend.api.v1.fileuploadreq"></a>

```json
{
  "file": {},
  "name": "string",
  "type": "string"
}
```

### 属性

| 名称   | 类型                                                                                              | 必选    | 约束   | 中文名 | 说明   |
| ---- | ----------------------------------------------------------------------------------------------- | ----- | ---- | --- | ---- |
| file | [github.com.gogf.gf.v2.net.ghttp.UploadFile](#schemagithub.com.gogf.gf.v2.net.ghttp.uploadfile) | false | none |     | none |
| name | string(string)                                                                                  | true  | none |     | none |
| type | string(string)                                                                                  | true  | none |     | none |

<h2 id="tocS_Gym-backend.api.v1.RegisterRes">Gym-backend.api.v1.RegisterRes</h2>

<a id="schemagym-backend.api.v1.registerres"></a>
<a id="schema_Gym-backend.api.v1.RegisterRes"></a>
<a id="tocSgym-backend.api.v1.registerres"></a>
<a id="tocsgym-backend.api.v1.registerres"></a>

```json
{}
```

### 属性

*None*

<h2 id="tocS_Gym-backend.api.v1.RegisterReq">Gym-backend.api.v1.RegisterReq</h2>

<a id="schemagym-backend.api.v1.registerreq"></a>
<a id="schema_Gym-backend.api.v1.RegisterReq"></a>
<a id="tocSgym-backend.api.v1.registerreq"></a>
<a id="tocsgym-backend.api.v1.registerreq"></a>

```json
{
  "Username": "string",
  "Password": "string",
  "ConfirmPassword": "string",
  "Email": "string",
  "Gender": "string",
  "Phone": "string"
}
```

### 属性

| 名称              | 类型             | 必选   | 约束   | 中文名 | 说明   |
| --------------- | -------------- | ---- | ---- | --- | ---- |
| Username        | string(string) | true | none |     | none |
| Password        | string(string) | true | none |     | none |
| ConfirmPassword | string(string) | true | none |     | none |
| Email           | string(string) | true | none |     | none |
| Gender          | string(string) | true | none |     | none |
| Phone           | string(string) | true | none |     | none |

<h2 id="tocS_Gym-backend.api.v1.ProfileRes">Gym-backend.api.v1.ProfileRes</h2>

<a id="schemagym-backend.api.v1.profileres"></a>
<a id="schema_Gym-backend.api.v1.ProfileRes"></a>
<a id="tocSgym-backend.api.v1.profileres"></a>
<a id="tocsgym-backend.api.v1.profileres"></a>

```json
"string"
```

### 属性

| 名称         | 类型             | 必选    | 约束   | 中文名 | 说明   |
| ---------- | -------------- | ----- | ---- | --- | ---- |
| Data       | object         | false | none |     | none |
| » Username | string(string) | false | none |     | none |
| » Gender   | string(string) | false | none |     | none |
| » Role     | integer(uint)  | false | none |     | none |
| » Email    | string(string) | false | none |     | none |
| » Phone    | string(string) | false | none |     | none |
| » Avatar   | string(string) | false | none |     | none |

<h2 id="tocS_Gym-backend.api.v1.ProfileReq">Gym-backend.api.v1.ProfileReq</h2>

<a id="schemagym-backend.api.v1.profilereq"></a>
<a id="schema_Gym-backend.api.v1.ProfileReq"></a>
<a id="tocSgym-backend.api.v1.profilereq"></a>
<a id="tocsgym-backend.api.v1.profilereq"></a>

```json
{}
```

### 属性

*None*

<h2 id="tocS_Gym-backend.api.v1.ModifyFacilityRes">Gym-backend.api.v1.ModifyFacilityRes</h2>

<a id="schemagym-backend.api.v1.modifyfacilityres"></a>
<a id="schema_Gym-backend.api.v1.ModifyFacilityRes"></a>
<a id="tocSgym-backend.api.v1.modifyfacilityres"></a>
<a id="tocsgym-backend.api.v1.modifyfacilityres"></a>

```json
"string"
```

### 属性

*None*

<h2 id="tocS_Gym-backend.api.v1.ModifyFacilityReq">Gym-backend.api.v1.ModifyFacilityReq</h2>

<a id="schemagym-backend.api.v1.modifyfacilityreq"></a>
<a id="schema_Gym-backend.api.v1.ModifyFacilityReq"></a>
<a id="tocSgym-backend.api.v1.modifyfacilityreq"></a>
<a id="tocsgym-backend.api.v1.modifyfacilityreq"></a>

```json
{
  "id": 0,
  "name": "string",
  "description": "string",
  "location": "string",
  "cost": 0
}
```

### 属性

| 名称          | 类型              | 必选   | 约束   | 中文名 | 说明   |
| ----------- | --------------- | ---- | ---- | --- | ---- |
| id          | integer(int)    | true | none |     | none |
| name        | string(string)  | true | none |     | none |
| description | string(string)  | true | none |     | none |
| location    | string(string)  | true | none |     | none |
| cost        | number(float64) | true | none |     | none |

<h2 id="tocS_interface">interface</h2>

<a id="schemainterface"></a>
<a id="schema_interface"></a>
<a id="tocSinterface"></a>
<a id="tocsinterface"></a>

```json
{}
```

### 属性

*None*

<h2 id="tocS_Gym-backend.api.v1.LogoutRes">Gym-backend.api.v1.LogoutRes</h2>

<a id="schemagym-backend.api.v1.logoutres"></a>
<a id="schema_Gym-backend.api.v1.LogoutRes"></a>
<a id="tocSgym-backend.api.v1.logoutres"></a>
<a id="tocsgym-backend.api.v1.logoutres"></a>

```json
{}
```

### 属性

*None*

<h2 id="tocS_Gym-backend.api.v1.LogoutReq">Gym-backend.api.v1.LogoutReq</h2>

<a id="schemagym-backend.api.v1.logoutreq"></a>
<a id="schema_Gym-backend.api.v1.LogoutReq"></a>
<a id="tocSgym-backend.api.v1.logoutreq"></a>
<a id="tocsgym-backend.api.v1.logoutreq"></a>

```json
{}
```

### 属性

*None*

<h2 id="tocS_Gym-backend.api.v1.LoginRes">Gym-backend.api.v1.LoginRes</h2>

<a id="schemagym-backend.api.v1.loginres"></a>
<a id="schema_Gym-backend.api.v1.LoginRes"></a>
<a id="tocSgym-backend.api.v1.loginres"></a>
<a id="tocsgym-backend.api.v1.loginres"></a>

```json
"string"
```

### 属性

| 名称         | 类型             | 必选    | 约束   | 中文名 | 说明   |
| ---------- | -------------- | ----- | ---- | --- | ---- |
| Data       | object         | false | none |     | none |
| » username | string(string) | false | none |     | none |
| » avatar   | string(string) | false | none |     | none |
| » role     | integer(int)   | false | none |     | none |

<h2 id="tocS_Gym-backend.api.v1.LoginReq">Gym-backend.api.v1.LoginReq</h2>

<a id="schemagym-backend.api.v1.loginreq"></a>
<a id="schema_Gym-backend.api.v1.LoginReq"></a>
<a id="tocSgym-backend.api.v1.loginreq"></a>
<a id="tocsgym-backend.api.v1.loginreq"></a>

```json
{
  "Username": "string",
  "Password": "string"
}
```

### 属性

| 名称       | 类型             | 必选   | 约束   | 中文名 | 说明   |
| -------- | -------------- | ---- | ---- | --- | ---- |
| Username | string(string) | true | none |     | none |
| Password | string(string) | true | none |     | none |

<h2 id="tocS_Gym-backend.api.v1.FacilitySearchRes">Gym-backend.api.v1.FacilitySearchRes</h2>

<a id="schemagym-backend.api.v1.facilitysearchres"></a>
<a id="schema_Gym-backend.api.v1.FacilitySearchRes"></a>
<a id="tocSgym-backend.api.v1.facilitysearchres"></a>
<a id="tocsgym-backend.api.v1.facilitysearchres"></a>

```json
"string"
```

### 属性

| 名称       | 类型                                                                                              | 必选    | 约束   | 中文名 | 说明   |
| -------- | ----------------------------------------------------------------------------------------------- | ----- | ---- | --- | ---- |
| Facility | [[Gym-backend.internal.model.FacilityEntity](#schemagym-backend.internal.model.facilityentity)] | false | none |     | none |

<h2 id="tocS_Gym-backend.api.v1.FacilitySearchReq">Gym-backend.api.v1.FacilitySearchReq</h2>

<a id="schemagym-backend.api.v1.facilitysearchreq"></a>
<a id="schema_Gym-backend.api.v1.FacilitySearchReq"></a>
<a id="tocSgym-backend.api.v1.facilitysearchreq"></a>
<a id="tocsgym-backend.api.v1.facilitysearchreq"></a>

```json
{
  "name": "string",
  "id": 0
}
```

### 属性

| 名称   | 类型             | 必选    | 约束   | 中文名 | 说明   |
| ---- | -------------- | ----- | ---- | --- | ---- |
| name | string(string) | false | none |     | none |
| id   | integer(int)   | false | none |     | none |

<h2 id="tocS_Gym-backend.api.v1.FacilityDetailRes">Gym-backend.api.v1.FacilityDetailRes</h2>

<a id="schemagym-backend.api.v1.facilitydetailres"></a>
<a id="schema_Gym-backend.api.v1.FacilityDetailRes"></a>
<a id="tocSgym-backend.api.v1.facilitydetailres"></a>
<a id="tocsgym-backend.api.v1.facilitydetailres"></a>

```json
"string"
```

### 属性

| 名称       | 类型                                                                                            | 必选    | 约束   | 中文名 | 说明   |
| -------- | --------------------------------------------------------------------------------------------- | ----- | ---- | --- | ---- |
| Facility | [Gym-backend.internal.model.FacilityEntity](#schemagym-backend.internal.model.facilityentity) | false | none |     | none |

<h2 id="tocS_Gym-backend.api.v1.FacilityDetailReq">Gym-backend.api.v1.FacilityDetailReq</h2>

<a id="schemagym-backend.api.v1.facilitydetailreq"></a>
<a id="schema_Gym-backend.api.v1.FacilityDetailReq"></a>
<a id="tocSgym-backend.api.v1.facilitydetailreq"></a>
<a id="tocsgym-backend.api.v1.facilitydetailreq"></a>

```json
{
  "id": 0
}
```

### 属性

| 名称  | 类型           | 必选   | 约束   | 中文名 | 说明   |
| --- | ------------ | ---- | ---- | --- | ---- |
| id  | integer(int) | true | none |     | none |

<h2 id="tocS_Gym-backend.api.v1.FacilityRes">Gym-backend.api.v1.FacilityRes</h2>

<a id="schemagym-backend.api.v1.facilityres"></a>
<a id="schema_Gym-backend.api.v1.FacilityRes"></a>
<a id="tocSgym-backend.api.v1.facilityres"></a>
<a id="tocsgym-backend.api.v1.facilityres"></a>

```json
"string"
```

### 属性

| 名称       | 类型                                                                                              | 必选    | 约束   | 中文名 | 说明   |
| -------- | ----------------------------------------------------------------------------------------------- | ----- | ---- | --- | ---- |
| Facility | [[Gym-backend.internal.model.FacilityEntity](#schemagym-backend.internal.model.facilityentity)] | false | none |     | none |

<h2 id="tocS_Gym-backend.internal.model.entity.Facility">Gym-backend.internal.model.entity.Facility</h2>

<a id="schemagym-backend.internal.model.entity.facility"></a>
<a id="schema_Gym-backend.internal.model.entity.Facility"></a>
<a id="tocSgym-backend.internal.model.entity.facility"></a>
<a id="tocsgym-backend.internal.model.entity.facility"></a>

```json
{
  "id": 0,
  "name": "string",
  "description": "string",
  "location": "string",
  "cost": 0
}
```

### 属性

| 名称          | 类型              | 必选    | 约束   | 中文名 | 说明   |
| ----------- | --------------- | ----- | ---- | --- | ---- |
| id          | integer(int)    | false | none |     | none |
| name        | string(string)  | false | none |     | none |
| description | string(string)  | false | none |     | none |
| location    | string(string)  | false | none |     | none |
| cost        | number(float64) | false | none |     | none |

<h2 id="tocS_Gym-backend.api.v1.AddFacilityRes">Gym-backend.api.v1.AddFacilityRes</h2>

<a id="schemagym-backend.api.v1.addfacilityres"></a>
<a id="schema_Gym-backend.api.v1.AddFacilityRes"></a>
<a id="tocSgym-backend.api.v1.addfacilityres"></a>
<a id="tocsgym-backend.api.v1.addfacilityres"></a>

```json
"string"
```

### 属性

*None*

<h2 id="tocS_Gym-backend.api.v1.FacilityReq">Gym-backend.api.v1.FacilityReq</h2>

<a id="schemagym-backend.api.v1.facilityreq"></a>
<a id="schema_Gym-backend.api.v1.FacilityReq"></a>
<a id="tocSgym-backend.api.v1.facilityreq"></a>
<a id="tocsgym-backend.api.v1.facilityreq"></a>

```json
{}
```

### 属性

*None*

<h2 id="tocS_Gym-backend.api.v1.AddFacilityReq">Gym-backend.api.v1.AddFacilityReq</h2>

<a id="schemagym-backend.api.v1.addfacilityreq"></a>
<a id="schema_Gym-backend.api.v1.AddFacilityReq"></a>
<a id="tocSgym-backend.api.v1.addfacilityreq"></a>
<a id="tocsgym-backend.api.v1.addfacilityreq"></a>

```json
{
  "name": "string",
  "description": "string",
  "location": "string",
  "cost": 0
}
```

### 属性

| 名称          | 类型              | 必选   | 约束   | 中文名 | 说明   |
| ----------- | --------------- | ---- | ---- | --- | ---- |
| name        | string(string)  | true | none |     | none |
| description | string(string)  | true | none |     | none |
| location    | string(string)  | true | none |     | none |
| cost        | number(float64) | true | none |     | none |
