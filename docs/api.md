---
title: GymKneeThaiMay v1.0.0
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

# GymKneeThaiMay

> v1.0.0

Base URLs:

# Bank

## GET Get banks

GET /api/v1/bank/banks

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.GetBanksRes](#schemagym-backend.api.v1.getbanksres)|

## POST Update bank

POST /api/v1/bank/update

> Body Parameters

```json
{
  "id": 0,
  "name": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gym-backend.api.v1.UpdateBankReq](#schemagym-backend.api.v1.updatebankreq)| no |none|

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.UpdateBankRes](#schemagym-backend.api.v1.updatebankres)|

## POST Add bank

POST /api/v1/bank/add

> Body Parameters

```json
{
  "name": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gym-backend.api.v1.AddBankReq](#schemagym-backend.api.v1.addbankreq)| no |none|

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.AddBankRes](#schemagym-backend.api.v1.addbankres)|

# Account

## POST Login

POST /api/v1/user/login

> Body Parameters

```json
{
  "Username": "string",
  "Password": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gym-backend.api.v1.LoginReq](#schemagym-backend.api.v1.loginreq)| no |none|

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.LoginRes](#schemagym-backend.api.v1.loginres)|

## POST Change password'

POST /api/v1/change-passwd

> Body Parameters

```json
{
  "OldPassword": "string",
  "NewPassword": "string",
  "ConfirmPassword": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gym-backend.api.v1.ChangePasswdReq](#schemagym-backend.api.v1.changepasswdreq)| no |none|

> Response Examples

> 200 Response

```json
{
  "code": 0,
  "message": "string",
  "data": {},
  "redirect": "string"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» code|integer(int)|false|none||none|
|» message|string(string)|false|none||none|
|» data|object|false|none||none|
|» redirect|string(string)|false|none||none|

## GET Log out'

GET /api/v1/user/logout

> Response Examples

> 200 Response

```json
{
  "code": 0,
  "message": "string",
  "data": {},
  "redirect": "string"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» code|integer(int)|false|none||none|
|» message|string(string)|false|none||none|
|» data|object|false|none||none|
|» redirect|string(string)|false|none||none|

## POST Get user's profile'

POST /api/v1/user/profile

> Body Parameters

```json
{}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gym-backend.api.v1.ProfileReq](#schemagym-backend.api.v1.profilereq)| no |none|

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.ProfileRes](#schemagym-backend.api.v1.profileres)|

## POST Register

POST /api/v1/user/register

> Body Parameters

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

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gym-backend.api.v1.RegisterReq](#schemagym-backend.api.v1.registerreq)| no |none|

> Response Examples

> 200 Response

```json
{
  "code": 0,
  "message": "string",
  "data": {},
  "redirect": "string"
}
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### Responses Data Schema

HTTP Status Code **200**

|Name|Type|Required|Restrictions|Title|description|
|---|---|---|---|---|---|
|» code|integer(int)|false|none||none|
|» message|string(string)|false|none||none|
|» data|object|false|none||none|
|» redirect|string(string)|false|none||none|

# Facility

## POST Add Facility

POST /api/v1/facility/add

> Body Parameters

```json
{
  "name": "string",
  "description": "string",
  "location": "string",
  "image": [
    "string"
  ]
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gym-backend.api.v1.AddFacilityReq](#schemagym-backend.api.v1.addfacilityreq)| no |none|

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.AddFacilityRes](#schemagym-backend.api.v1.addfacilityres)|

## POST Get Facility Detail

POST /api/v1/facility/detail

> Body Parameters

```json
{
  "id": 0
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gym-backend.api.v1.FacilityDetailReq](#schemagym-backend.api.v1.facilitydetailreq)| no |none|

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.FacilityDetailRes](#schemagym-backend.api.v1.facilitydetailres)|

## GET Get All Facilities

GET /api/v1/facility/facility

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.FacilityRes](#schemagym-backend.api.v1.facilityres)|

## POST Add Facility Place

POST /api/v1/facility/place/add

> Body Parameters

```json
{
  "name": "string",
  "facility_id": 0,
  "cost": 0,
  "description": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gym-backend.api.v1.AddFacilityPlaceReq](#schemagym-backend.api.v1.addfacilityplacereq)| no |none|

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.AddFacilityPlaceRes](#schemagym-backend.api.v1.addfacilityplaceres)|

## POST Modify Facility Place

POST /api/v1/facility/place/update

> Body Parameters

```json
{
  "id": 0,
  "name": "string",
  "facility_id": 0,
  "cost": 0,
  "description": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gym-backend.api.v1.ModifyFacilityPlaceReq](#schemagym-backend.api.v1.modifyfacilityplacereq)| no |none|

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.ModifyFacilityPlaceRes](#schemagym-backend.api.v1.modifyfacilityplaceres)|

## POST Get Facilities By searching tags

POST /api/v1/facility/search

> Body Parameters

```json
{
  "name": "string",
  "id": 0
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gym-backend.api.v1.FacilitySearchReq](#schemagym-backend.api.v1.facilitysearchreq)| no |none|

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.FacilitySearchRes](#schemagym-backend.api.v1.facilitysearchres)|

## POST Modify Facility

POST /api/v1/facility/update

> Body Parameters

```json
{
  "id": 0,
  "name": "string",
  "description": "string",
  "location": "string",
  "image": [
    "string"
  ]
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gym-backend.api.v1.ModifyFacilityReq](#schemagym-backend.api.v1.modifyfacilityreq)| no |none|

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.ModifyFacilityRes](#schemagym-backend.api.v1.modifyfacilityres)|

# Order

## POST Create order

POST /api/v1/order/create

> Body Parameters

```json
{
  "placeId": 0,
  "startTime": "string",
  "endTime": "string"
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gym-backend.api.v1.CreateOrderReq](#schemagym-backend.api.v1.createorderreq)| no |none|

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.CreateOrderRes](#schemagym-backend.api.v1.createorderres)|

# Payment

## POST Create payment

POST /api/v1/order/payment

> Body Parameters

```json
{
  "orderCode": "string",
  "paymentType": 0,
  "cardId": 0
}
```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|[Gym-backend.api.v1.CreatePaymentReq](#schemagym-backend.api.v1.createpaymentreq)| no |none|

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.CreatePaymentRes](#schemagym-backend.api.v1.createpaymentres)|

# File

## POST Upload file

POST /api/v1/upload

> Body Parameters

```yaml
file: string
name: string
type: string

```

### Params

|Name|Location|Type|Required|Description|
|---|---|---|---|---|
|body|body|object| no |none|
|» file|body|string| no |none|
|» name|body|string| yes |none|
|» type|body|string| yes |none|

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.FileUploadRes](#schemagym-backend.api.v1.fileuploadres)|

# Data Schema

<h2 id="tocS_Gym-backend.api.v1.UpdateBankRes">Gym-backend.api.v1.UpdateBankRes</h2>

<a id="schemagym-backend.api.v1.updatebankres"></a>
<a id="schema_Gym-backend.api.v1.UpdateBankRes"></a>
<a id="tocSgym-backend.api.v1.updatebankres"></a>
<a id="tocsgym-backend.api.v1.updatebankres"></a>

```json
"string"

```

### Attribute

*None*

<h2 id="tocS_Gym-backend.api.v1.UpdateBankReq">Gym-backend.api.v1.UpdateBankReq</h2>

<a id="schemagym-backend.api.v1.updatebankreq"></a>
<a id="schema_Gym-backend.api.v1.UpdateBankReq"></a>
<a id="tocSgym-backend.api.v1.updatebankreq"></a>
<a id="tocsgym-backend.api.v1.updatebankreq"></a>

```json
{
  "id": 0,
  "name": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int)|true|none||none|
|name|string(string)|true|none||none|

<h2 id="tocS_Gym-backend.api.v1.AddBankRes">Gym-backend.api.v1.AddBankRes</h2>

<a id="schemagym-backend.api.v1.addbankres"></a>
<a id="schema_Gym-backend.api.v1.AddBankRes"></a>
<a id="tocSgym-backend.api.v1.addbankres"></a>
<a id="tocsgym-backend.api.v1.addbankres"></a>

```json
"string"

```

### Attribute

*None*

<h2 id="tocS_Gym-backend.api.v1.AddBankReq">Gym-backend.api.v1.AddBankReq</h2>

<a id="schemagym-backend.api.v1.addbankreq"></a>
<a id="schema_Gym-backend.api.v1.AddBankReq"></a>
<a id="tocSgym-backend.api.v1.addbankreq"></a>
<a id="tocsgym-backend.api.v1.addbankreq"></a>

```json
{
  "name": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|name|string(string)|true|none||none|

<h2 id="tocS_Gym-backend.internal.model.entity.Bank">Gym-backend.internal.model.entity.Bank</h2>

<a id="schemagym-backend.internal.model.entity.bank"></a>
<a id="schema_Gym-backend.internal.model.entity.Bank"></a>
<a id="tocSgym-backend.internal.model.entity.bank"></a>
<a id="tocsgym-backend.internal.model.entity.bank"></a>

```json
{
  "id": 0,
  "name": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int)|false|none||none|
|name|string(string)|false|none||none|

<h2 id="tocS_Gym-backend.api.v1.GetBanksRes">Gym-backend.api.v1.GetBanksRes</h2>

<a id="schemagym-backend.api.v1.getbanksres"></a>
<a id="schema_Gym-backend.api.v1.GetBanksRes"></a>
<a id="tocSgym-backend.api.v1.getbanksres"></a>
<a id="tocsgym-backend.api.v1.getbanksres"></a>

```json
"string"

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|Data|[[Gym-backend.internal.model.entity.Bank](#schemagym-backend.internal.model.entity.bank)]|false|none||none|

<h2 id="tocS_Gym-backend.api.v1.GetBanksReq">Gym-backend.api.v1.GetBanksReq</h2>

<a id="schemagym-backend.api.v1.getbanksreq"></a>
<a id="schema_Gym-backend.api.v1.GetBanksReq"></a>
<a id="tocSgym-backend.api.v1.getbanksreq"></a>
<a id="tocsgym-backend.api.v1.getbanksreq"></a>

```json
{}

```

### Attribute

*None*

<h2 id="tocS_Gym-backend.api.v1.CreatePaymentRes">Gym-backend.api.v1.CreatePaymentRes</h2>

<a id="schemagym-backend.api.v1.createpaymentres"></a>
<a id="schema_Gym-backend.api.v1.CreatePaymentRes"></a>
<a id="tocSgym-backend.api.v1.createpaymentres"></a>
<a id="tocsgym-backend.api.v1.createpaymentres"></a>

```json
"string"

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|orderCode|string(string)|false|none||none|
|paymentType|integer(int)|false|none||none|
|amount|number(float64)|false|none||none|
|paymentCode|string(string)|false|none||none|
|status|integer(int)|false|none||none|

<h2 id="tocS_Gym-backend.api.v1.CreatePaymentReq">Gym-backend.api.v1.CreatePaymentReq</h2>

<a id="schemagym-backend.api.v1.createpaymentreq"></a>
<a id="schema_Gym-backend.api.v1.CreatePaymentReq"></a>
<a id="tocSgym-backend.api.v1.createpaymentreq"></a>
<a id="tocsgym-backend.api.v1.createpaymentreq"></a>

```json
{
  "orderCode": "string",
  "paymentType": 0,
  "cardId": 0
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|orderCode|string(string)|true|none||none|
|paymentType|integer(int)|true|none||none|
|cardId|integer(int)|false|none||none|

<h2 id="tocS_Gym-backend.api.v1.CreateOrderRes">Gym-backend.api.v1.CreateOrderRes</h2>

<a id="schemagym-backend.api.v1.createorderres"></a>
<a id="schema_Gym-backend.api.v1.CreateOrderRes"></a>
<a id="tocSgym-backend.api.v1.createorderres"></a>
<a id="tocsgym-backend.api.v1.createorderres"></a>

```json
"string"

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|orderCode|string(string)|false|none||none|
|amount|number(float64)|false|none||none|
|placeId|integer(int)|false|none||none|
|startTime|string(string)|false|none||none|
|endTime|string(string)|false|none||none|

<h2 id="tocS_Gym-backend.api.v1.CreateOrderReq">Gym-backend.api.v1.CreateOrderReq</h2>

<a id="schemagym-backend.api.v1.createorderreq"></a>
<a id="schema_Gym-backend.api.v1.CreateOrderReq"></a>
<a id="tocSgym-backend.api.v1.createorderreq"></a>
<a id="tocsgym-backend.api.v1.createorderreq"></a>

```json
{
  "placeId": 0,
  "startTime": "string",
  "endTime": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|placeId|integer(int)|true|none||none|
|startTime|string(string)|true|none||none|
|endTime|string(string)|true|none||none|

<h2 id="tocS_Gym-backend.api.v1.ModifyFacilityPlaceRes">Gym-backend.api.v1.ModifyFacilityPlaceRes</h2>

<a id="schemagym-backend.api.v1.modifyfacilityplaceres"></a>
<a id="schema_Gym-backend.api.v1.ModifyFacilityPlaceRes"></a>
<a id="tocSgym-backend.api.v1.modifyfacilityplaceres"></a>
<a id="tocsgym-backend.api.v1.modifyfacilityplaceres"></a>

```json
"string"

```

### Attribute

*None*

<h2 id="tocS_Gym-backend.api.v1.ModifyFacilityPlaceReq">Gym-backend.api.v1.ModifyFacilityPlaceReq</h2>

<a id="schemagym-backend.api.v1.modifyfacilityplacereq"></a>
<a id="schema_Gym-backend.api.v1.ModifyFacilityPlaceReq"></a>
<a id="tocSgym-backend.api.v1.modifyfacilityplacereq"></a>
<a id="tocsgym-backend.api.v1.modifyfacilityplacereq"></a>

```json
{
  "id": 0,
  "name": "string",
  "facility_id": 0,
  "cost": 0,
  "description": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int)|true|none||none|
|name|string(string)|true|none||none|
|facility_id|integer(int)|true|none||none|
|cost|number(float64)|true|none||none|
|description|string(string)|true|none||none|

<h2 id="tocS_Gym-backend.api.v1.AddFacilityPlaceRes">Gym-backend.api.v1.AddFacilityPlaceRes</h2>

<a id="schemagym-backend.api.v1.addfacilityplaceres"></a>
<a id="schema_Gym-backend.api.v1.AddFacilityPlaceRes"></a>
<a id="tocSgym-backend.api.v1.addfacilityplaceres"></a>
<a id="tocsgym-backend.api.v1.addfacilityplaceres"></a>

```json
"string"

```

### Attribute

*None*

<h2 id="tocS_Gym-backend.api.v1.AddFacilityPlaceReq">Gym-backend.api.v1.AddFacilityPlaceReq</h2>

<a id="schemagym-backend.api.v1.addfacilityplacereq"></a>
<a id="schema_Gym-backend.api.v1.AddFacilityPlaceReq"></a>
<a id="tocSgym-backend.api.v1.addfacilityplacereq"></a>
<a id="tocsgym-backend.api.v1.addfacilityplacereq"></a>

```json
{
  "name": "string",
  "facility_id": 0,
  "cost": 0,
  "description": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|name|string(string)|true|none||none|
|facility_id|integer(int)|true|none||none|
|cost|number(float64)|true|none||none|
|description|string(string)|true|none||none|

<h2 id="tocS_Gym-backend.internal.model.entity.FacilityPlace">Gym-backend.internal.model.entity.FacilityPlace</h2>

<a id="schemagym-backend.internal.model.entity.facilityplace"></a>
<a id="schema_Gym-backend.internal.model.entity.FacilityPlace"></a>
<a id="tocSgym-backend.internal.model.entity.facilityplace"></a>
<a id="tocsgym-backend.internal.model.entity.facilityplace"></a>

```json
{
  "id": 0,
  "facilityId": 0,
  "name": "string",
  "description": "string",
  "cost": 0
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int)|false|none||none|
|facilityId|integer(int)|false|none||none|
|name|string(string)|false|none||none|
|description|string(string)|false|none||none|
|cost|number(float64)|false|none||none|

<h2 id="tocS_Gym-backend.api.v1.ChangePasswdRes">Gym-backend.api.v1.ChangePasswdRes</h2>

<a id="schemagym-backend.api.v1.changepasswdres"></a>
<a id="schema_Gym-backend.api.v1.ChangePasswdRes"></a>
<a id="tocSgym-backend.api.v1.changepasswdres"></a>
<a id="tocsgym-backend.api.v1.changepasswdres"></a>

```json
{}

```

### Attribute

*None*

<h2 id="tocS_Gym-backend.api.v1.ChangePasswdReq">Gym-backend.api.v1.ChangePasswdReq</h2>

<a id="schemagym-backend.api.v1.changepasswdreq"></a>
<a id="schema_Gym-backend.api.v1.ChangePasswdReq"></a>
<a id="tocSgym-backend.api.v1.changepasswdreq"></a>
<a id="tocsgym-backend.api.v1.changepasswdreq"></a>

```json
{
  "OldPassword": "string",
  "NewPassword": "string",
  "ConfirmPassword": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|OldPassword|string(string)|true|none||none|
|NewPassword|string(string)|true|none||none|
|ConfirmPassword|string(string)|true|none||none|

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

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int)|false|none||none|
|path|string(string)|false|none||none|
|uploadTime|string(*gtime.Time)|false|none||none|
|facilityID|integer(int)|false|none||none|

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
    "images": "string"
  },
  "places": [
    {
      "id": 0,
      "facilityId": 0,
      "name": "string",
      "description": "string",
      "cost": 0
    }
  ]
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|facility|[Gym-backend.internal.model.entity.Facility](#schemagym-backend.internal.model.entity.facility)|false|none||none|
|places|[[Gym-backend.internal.model.entity.FacilityPlace](#schemagym-backend.internal.model.entity.facilityplace)]|false|none||none|

<h2 id="tocS_Gym-backend.api.v1.FileUploadRes">Gym-backend.api.v1.FileUploadRes</h2>

<a id="schemagym-backend.api.v1.fileuploadres"></a>
<a id="schema_Gym-backend.api.v1.FileUploadRes"></a>
<a id="tocSgym-backend.api.v1.fileuploadres"></a>
<a id="tocsgym-backend.api.v1.fileuploadres"></a>

```json
"string"

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|Data|object|false|none||none|
|» url|string(string)|false|none||none|
|» name|string(string)|false|none||none|

<h2 id="tocS_github.com.gogf.gf.v2.net.ghttp.UploadFile">github.com.gogf.gf.v2.net.ghttp.UploadFile</h2>

<a id="schemagithub.com.gogf.gf.v2.net.ghttp.uploadfile"></a>
<a id="schema_github.com.gogf.gf.v2.net.ghttp.UploadFile"></a>
<a id="tocSgithub.com.gogf.gf.v2.net.ghttp.uploadfile"></a>
<a id="tocsgithub.com.gogf.gf.v2.net.ghttp.uploadfile"></a>

```json
{}

```

### Attribute

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

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|file|[github.com.gogf.gf.v2.net.ghttp.UploadFile](#schemagithub.com.gogf.gf.v2.net.ghttp.uploadfile)|false|none||none|
|name|string(string)|true|none||none|
|type|string(string)|true|none||none|

<h2 id="tocS_Gym-backend.api.v1.RegisterRes">Gym-backend.api.v1.RegisterRes</h2>

<a id="schemagym-backend.api.v1.registerres"></a>
<a id="schema_Gym-backend.api.v1.RegisterRes"></a>
<a id="tocSgym-backend.api.v1.registerres"></a>
<a id="tocsgym-backend.api.v1.registerres"></a>

```json
{}

```

### Attribute

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

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|Username|string(string)|true|none||none|
|Password|string(string)|true|none||none|
|ConfirmPassword|string(string)|true|none||none|
|Email|string(string)|true|none||none|
|Gender|string(string)|true|none||none|
|Phone|string(string)|true|none||none|

<h2 id="tocS_Gym-backend.api.v1.ProfileRes">Gym-backend.api.v1.ProfileRes</h2>

<a id="schemagym-backend.api.v1.profileres"></a>
<a id="schema_Gym-backend.api.v1.ProfileRes"></a>
<a id="tocSgym-backend.api.v1.profileres"></a>
<a id="tocsgym-backend.api.v1.profileres"></a>

```json
"string"

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|Data|object|false|none||none|
|» Username|string(string)|false|none||none|
|» Gender|integer(uint)|false|none||none|
|» Role|integer(uint)|false|none||none|
|» Email|string(string)|false|none||none|
|» Phone|string(string)|false|none||none|
|» Avatar|string(string)|false|none||none|

<h2 id="tocS_Gym-backend.api.v1.ProfileReq">Gym-backend.api.v1.ProfileReq</h2>

<a id="schemagym-backend.api.v1.profilereq"></a>
<a id="schema_Gym-backend.api.v1.ProfileReq"></a>
<a id="tocSgym-backend.api.v1.profilereq"></a>
<a id="tocsgym-backend.api.v1.profilereq"></a>

```json
{}

```

### Attribute

*None*

<h2 id="tocS_Gym-backend.api.v1.ModifyFacilityRes">Gym-backend.api.v1.ModifyFacilityRes</h2>

<a id="schemagym-backend.api.v1.modifyfacilityres"></a>
<a id="schema_Gym-backend.api.v1.ModifyFacilityRes"></a>
<a id="tocSgym-backend.api.v1.modifyfacilityres"></a>
<a id="tocsgym-backend.api.v1.modifyfacilityres"></a>

```json
"string"

```

### Attribute

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
  "image": [
    "string"
  ]
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int)|true|none||none|
|name|string(string)|true|none||none|
|description|string(string)|true|none||none|
|location|string(string)|true|none||none|
|image|[string]|true|none||none|

<h2 id="tocS_interface">interface</h2>

<a id="schemainterface"></a>
<a id="schema_interface"></a>
<a id="tocSinterface"></a>
<a id="tocsinterface"></a>

```json
{}

```

### Attribute

*None*

<h2 id="tocS_Gym-backend.api.v1.LogoutRes">Gym-backend.api.v1.LogoutRes</h2>

<a id="schemagym-backend.api.v1.logoutres"></a>
<a id="schema_Gym-backend.api.v1.LogoutRes"></a>
<a id="tocSgym-backend.api.v1.logoutres"></a>
<a id="tocsgym-backend.api.v1.logoutres"></a>

```json
{}

```

### Attribute

*None*

<h2 id="tocS_Gym-backend.api.v1.LogoutReq">Gym-backend.api.v1.LogoutReq</h2>

<a id="schemagym-backend.api.v1.logoutreq"></a>
<a id="schema_Gym-backend.api.v1.LogoutReq"></a>
<a id="tocSgym-backend.api.v1.logoutreq"></a>
<a id="tocsgym-backend.api.v1.logoutreq"></a>

```json
{}

```

### Attribute

*None*

<h2 id="tocS_Gym-backend.api.v1.LoginRes">Gym-backend.api.v1.LoginRes</h2>

<a id="schemagym-backend.api.v1.loginres"></a>
<a id="schema_Gym-backend.api.v1.LoginRes"></a>
<a id="tocSgym-backend.api.v1.loginres"></a>
<a id="tocsgym-backend.api.v1.loginres"></a>

```json
"string"

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|Data|object|false|none||none|
|» username|string(string)|false|none||none|
|» avatar|string(string)|false|none||none|
|» role|integer(int)|false|none||none|

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

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|Username|string(string)|true|none||none|
|Password|string(string)|true|none||none|

<h2 id="tocS_Gym-backend.api.v1.FacilitySearchRes">Gym-backend.api.v1.FacilitySearchRes</h2>

<a id="schemagym-backend.api.v1.facilitysearchres"></a>
<a id="schema_Gym-backend.api.v1.FacilitySearchRes"></a>
<a id="tocSgym-backend.api.v1.facilitysearchres"></a>
<a id="tocsgym-backend.api.v1.facilitysearchres"></a>

```json
"string"

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|Facility|[[Gym-backend.internal.model.FacilityEntity](#schemagym-backend.internal.model.facilityentity)]|false|none||none|

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

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|name|string(string)|false|none||none|
|id|integer(int)|false|none||none|

<h2 id="tocS_Gym-backend.api.v1.FacilityDetailRes">Gym-backend.api.v1.FacilityDetailRes</h2>

<a id="schemagym-backend.api.v1.facilitydetailres"></a>
<a id="schema_Gym-backend.api.v1.FacilityDetailRes"></a>
<a id="tocSgym-backend.api.v1.facilitydetailres"></a>
<a id="tocsgym-backend.api.v1.facilitydetailres"></a>

```json
"string"

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|Facility|[Gym-backend.internal.model.FacilityEntity](#schemagym-backend.internal.model.facilityentity)|false|none||none|

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

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int)|true|none||none|

<h2 id="tocS_Gym-backend.api.v1.FacilityRes">Gym-backend.api.v1.FacilityRes</h2>

<a id="schemagym-backend.api.v1.facilityres"></a>
<a id="schema_Gym-backend.api.v1.FacilityRes"></a>
<a id="tocSgym-backend.api.v1.facilityres"></a>
<a id="tocsgym-backend.api.v1.facilityres"></a>

```json
"string"

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|Facility|[[Gym-backend.internal.model.FacilityEntity](#schemagym-backend.internal.model.facilityentity)]|false|none||none|

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
  "images": "string"
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int)|false|none||none|
|name|string(string)|false|none||none|
|description|string(string)|false|none||none|
|location|string(string)|false|none||none|
|images|string(string)|false|none||none|

<h2 id="tocS_Gym-backend.api.v1.AddFacilityRes">Gym-backend.api.v1.AddFacilityRes</h2>

<a id="schemagym-backend.api.v1.addfacilityres"></a>
<a id="schema_Gym-backend.api.v1.AddFacilityRes"></a>
<a id="tocSgym-backend.api.v1.addfacilityres"></a>
<a id="tocsgym-backend.api.v1.addfacilityres"></a>

```json
"string"

```

### Attribute

*None*

<h2 id="tocS_Gym-backend.api.v1.FacilityReq">Gym-backend.api.v1.FacilityReq</h2>

<a id="schemagym-backend.api.v1.facilityreq"></a>
<a id="schema_Gym-backend.api.v1.FacilityReq"></a>
<a id="tocSgym-backend.api.v1.facilityreq"></a>
<a id="tocsgym-backend.api.v1.facilityreq"></a>

```json
{}

```

### Attribute

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
  "image": [
    "string"
  ]
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|name|string(string)|true|none||none|
|description|string(string)|true|none||none|
|location|string(string)|true|none||none|
|image|[string]|true|none||none|

