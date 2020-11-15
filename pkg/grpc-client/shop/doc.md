# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [shop.proto](#shop.proto)
    - [GetShopInfoReq](#shop.GetShopInfoReq)
    - [GetShopInfoRes](#shop.GetShopInfoRes)
    - [RegisterShopReq](#shop.RegisterShopReq)
    - [RegisterShopResp](#shop.RegisterShopResp)
  
  
  
    - [Shop](#shop.Shop)
  

- [shop.proto](#shop.proto)
    - [GetShopInfoReq](#shop.GetShopInfoReq)
    - [GetShopInfoRes](#shop.GetShopInfoRes)
    - [RegisterShopReq](#shop.RegisterShopReq)
    - [RegisterShopResp](#shop.RegisterShopResp)
  
  
  
    - [Shop](#shop.Shop)
  

- [Scalar Value Types](#scalar-value-types)



<a name="shop.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## shop.proto



<a name="shop.GetShopInfoReq"></a>

### GetShopInfoReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shopID | [int64](#int64) |  |  |
| userID | [int64](#int64) |  |  |






<a name="shop.GetShopInfoRes"></a>

### GetShopInfoRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| userID | [int64](#int64) |  |  |
| Name | [string](#string) |  |  |
| Address | [string](#string) |  |  |
| Status | [int32](#int32) |  |  |






<a name="shop.RegisterShopReq"></a>

### RegisterShopReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userID | [int64](#int64) |  |  |
| Name | [string](#string) |  |  |
| Address | [string](#string) |  |  |






<a name="shop.RegisterShopResp"></a>

### RegisterShopResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |





 

 

 


<a name="shop.Shop"></a>

### Shop


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetShopInfo | [GetShopInfoReq](#shop.GetShopInfoReq) | [GetShopInfoRes](#shop.GetShopInfoRes) |  |
| RegisterShop | [RegisterShopReq](#shop.RegisterShopReq) | [RegisterShopResp](#shop.RegisterShopResp) |  |

 



<a name="shop.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## shop.proto



<a name="shop.GetShopInfoReq"></a>

### GetShopInfoReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shopID | [int64](#int64) |  |  |
| userID | [int64](#int64) |  |  |






<a name="shop.GetShopInfoRes"></a>

### GetShopInfoRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| userID | [int64](#int64) |  |  |
| Name | [string](#string) |  |  |
| Address | [string](#string) |  |  |
| Status | [int32](#int32) |  |  |






<a name="shop.RegisterShopReq"></a>

### RegisterShopReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userID | [int64](#int64) |  |  |
| Name | [string](#string) |  |  |
| Address | [string](#string) |  |  |






<a name="shop.RegisterShopResp"></a>

### RegisterShopResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |





 

 

 


<a name="shop.Shop"></a>

### Shop


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetShopInfo | [GetShopInfoReq](#shop.GetShopInfoReq) | [GetShopInfoRes](#shop.GetShopInfoRes) |  |
| RegisterShop | [RegisterShopReq](#shop.RegisterShopReq) | [RegisterShopResp](#shop.RegisterShopResp) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

