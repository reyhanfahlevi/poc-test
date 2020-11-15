# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [accounts.proto](#accounts.proto)
    - [AuthenticatedUser](#account.AuthenticatedUser)
    - [CheckUserIsAuthenticatedReq](#account.CheckUserIsAuthenticatedReq)
    - [CheckUserIsAuthenticatedRes](#account.CheckUserIsAuthenticatedRes)
    - [CreateUserReq](#account.CreateUserReq)
    - [CreateUserRes](#account.CreateUserRes)
    - [Date](#account.Date)
    - [GetAccessTokenReq](#account.GetAccessTokenReq)
    - [GetAccessTokenResp](#account.GetAccessTokenResp)
    - [GetUserInfoReq](#account.GetUserInfoReq)
    - [GetUserInfoResp](#account.GetUserInfoResp)
  
  
  
    - [Account](#account.Account)
  

- [accounts.proto](#accounts.proto)
    - [AuthenticatedUser](#account.AuthenticatedUser)
    - [CheckUserIsAuthenticatedReq](#account.CheckUserIsAuthenticatedReq)
    - [CheckUserIsAuthenticatedRes](#account.CheckUserIsAuthenticatedRes)
    - [CreateUserReq](#account.CreateUserReq)
    - [CreateUserRes](#account.CreateUserRes)
    - [Date](#account.Date)
    - [GetAccessTokenReq](#account.GetAccessTokenReq)
    - [GetAccessTokenResp](#account.GetAccessTokenResp)
    - [GetUserInfoReq](#account.GetUserInfoReq)
    - [GetUserInfoResp](#account.GetUserInfoResp)
  
  
  
    - [Account](#account.Account)
  

- [Scalar Value Types](#scalar-value-types)



<a name="accounts.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## accounts.proto



<a name="account.AuthenticatedUser"></a>

### AuthenticatedUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userID | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| joinDate | [Date](#account.Date) |  |  |
| status | [int32](#int32) |  |  |






<a name="account.CheckUserIsAuthenticatedReq"></a>

### CheckUserIsAuthenticatedReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authToken | [string](#string) |  |  |






<a name="account.CheckUserIsAuthenticatedRes"></a>

### CheckUserIsAuthenticatedRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| isAuthenticated | [bool](#bool) |  |  |
| user | [AuthenticatedUser](#account.AuthenticatedUser) |  |  |






<a name="account.CreateUserReq"></a>

### CreateUserReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  |  |
| name | [string](#string) |  |  |






<a name="account.CreateUserRes"></a>

### CreateUserRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="account.Date"></a>

### Date



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| unixTime | [int64](#int64) |  |  |






<a name="account.GetAccessTokenReq"></a>

### GetAccessTokenReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="account.GetAccessTokenResp"></a>

### GetAccessTokenResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| expiredAt | [int64](#int64) |  |  |






<a name="account.GetUserInfoReq"></a>

### GetUserInfoReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userID | [int64](#int64) |  |  |
| email | [string](#string) |  |  |






<a name="account.GetUserInfoResp"></a>

### GetUserInfoResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userID | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| joinDate | [Date](#account.Date) |  |  |
| status | [int32](#int32) |  |  |





 

 

 


<a name="account.Account"></a>

### Account


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetUserInfo | [GetUserInfoReq](#account.GetUserInfoReq) | [GetUserInfoResp](#account.GetUserInfoResp) |  |
| CheckUserIsAuthenticated | [CheckUserIsAuthenticatedReq](#account.CheckUserIsAuthenticatedReq) | [CheckUserIsAuthenticatedRes](#account.CheckUserIsAuthenticatedRes) |  |
| GetAccessToken | [GetAccessTokenReq](#account.GetAccessTokenReq) | [GetAccessTokenResp](#account.GetAccessTokenResp) |  |
| CreateUser | [CreateUserReq](#account.CreateUserReq) | [CreateUserRes](#account.CreateUserRes) |  |

 



<a name="accounts.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## accounts.proto



<a name="account.AuthenticatedUser"></a>

### AuthenticatedUser



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userID | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| joinDate | [Date](#account.Date) |  |  |
| status | [int32](#int32) |  |  |






<a name="account.CheckUserIsAuthenticatedReq"></a>

### CheckUserIsAuthenticatedReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| authToken | [string](#string) |  |  |






<a name="account.CheckUserIsAuthenticatedRes"></a>

### CheckUserIsAuthenticatedRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| isAuthenticated | [bool](#bool) |  |  |
| user | [AuthenticatedUser](#account.AuthenticatedUser) |  |  |






<a name="account.CreateUserReq"></a>

### CreateUserReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  |  |
| name | [string](#string) |  |  |






<a name="account.CreateUserRes"></a>

### CreateUserRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="account.Date"></a>

### Date



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| unixTime | [int64](#int64) |  |  |






<a name="account.GetAccessTokenReq"></a>

### GetAccessTokenReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| password | [string](#string) |  |  |






<a name="account.GetAccessTokenResp"></a>

### GetAccessTokenResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| expiredAt | [int64](#int64) |  |  |






<a name="account.GetUserInfoReq"></a>

### GetUserInfoReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userID | [int64](#int64) |  |  |
| email | [string](#string) |  |  |






<a name="account.GetUserInfoResp"></a>

### GetUserInfoResp



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userID | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| email | [string](#string) |  |  |
| joinDate | [Date](#account.Date) |  |  |
| status | [int32](#int32) |  |  |





 

 

 


<a name="account.Account"></a>

### Account


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetUserInfo | [GetUserInfoReq](#account.GetUserInfoReq) | [GetUserInfoResp](#account.GetUserInfoResp) |  |
| CheckUserIsAuthenticated | [CheckUserIsAuthenticatedReq](#account.CheckUserIsAuthenticatedReq) | [CheckUserIsAuthenticatedRes](#account.CheckUserIsAuthenticatedRes) |  |
| GetAccessToken | [GetAccessTokenReq](#account.GetAccessTokenReq) | [GetAccessTokenResp](#account.GetAccessTokenResp) |  |
| CreateUser | [CreateUserReq](#account.CreateUserReq) | [CreateUserRes](#account.CreateUserRes) |  |

 



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

