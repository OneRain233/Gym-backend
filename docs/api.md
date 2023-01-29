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

> Response Examples

> 200 Response

```json
"string"
```

### Responses

|HTTP Status Code |Meaning|Description|Data schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|[Gym-backend.api.v1.FacilityRes](#schemagym-backend.api.v1.facilityres)|

## POST Get Facility Detail

POST /facility_detail

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

## POST Add Facility

POST /add_facility

> Body Parameters

```json
{
  "name": "string",
  "description": "string",
  "location": "string",
  "cost": 0
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

## POST Get Facilities By searching tags

POST /facility_search

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

POST /modify_facility

> Body Parameters

```json
{
  "id": 0,
  "name": "string",
  "description": "string",
  "location": "string",
  "cost": 0
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

# Data Schema

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
  "cost": 0
}

```

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int)|true|none||none|
|name|string(string)|true|none||none|
|description|string(string)|true|none||none|
|location|string(string)|true|none||none|
|cost|number(float64)|true|none||none|

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

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|facility|[Gym-backend.internal.model.entity.Facility](#schemagym-backend.internal.model.entity.facility)|false|none||none|
|images|[[Gym-backend.internal.model.entity.FacilityImage](#schemagym-backend.internal.model.entity.facilityimage)]|false|none||none|

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

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|id|integer(int)|false|none||none|
|name|string(string)|false|none||none|
|description|string(string)|false|none||none|
|location|string(string)|false|none||none|
|cost|number(float64)|false|none||none|

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

### Attribute

|Name|Type|Required|Restrictions|Title|Description|
|---|---|---|---|---|---|
|name|string(string)|true|none||none|
|description|string(string)|true|none||none|
|location|string(string)|true|none||none|
|cost|number(float64)|true|none||none|

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

