# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [product.proto](#product.proto)
    - [CreateProductReq](#product.CreateProductReq)
    - [CreateProductRes](#product.CreateProductRes)
    - [GetProductInfoReq](#product.GetProductInfoReq)
    - [GetProductInfoRes](#product.GetProductInfoRes)
    - [GetProductListReq](#product.GetProductListReq)
    - [GetProductListRes](#product.GetProductListRes)
    - [ProductImage](#product.ProductImage)
    - [ProductInfo](#product.ProductInfo)
    - [UpdateImageReq](#product.UpdateImageReq)
    - [UpdateImageRes](#product.UpdateImageRes)
    - [UpdateProductReq](#product.UpdateProductReq)
    - [UpdateProductRes](#product.UpdateProductRes)
  
  
  
    - [Product](#product.Product)
  

- [product.proto](#product.proto)
    - [CreateProductReq](#product.CreateProductReq)
    - [CreateProductRes](#product.CreateProductRes)
    - [GetProductInfoReq](#product.GetProductInfoReq)
    - [GetProductInfoRes](#product.GetProductInfoRes)
    - [GetProductListReq](#product.GetProductListReq)
    - [GetProductListRes](#product.GetProductListRes)
    - [ProductImage](#product.ProductImage)
    - [ProductInfo](#product.ProductInfo)
    - [UpdateImageReq](#product.UpdateImageReq)
    - [UpdateImageRes](#product.UpdateImageRes)
    - [UpdateProductReq](#product.UpdateProductReq)
    - [UpdateProductRes](#product.UpdateProductRes)
  
  
  
    - [Product](#product.Product)
  

- [Scalar Value Types](#scalar-value-types)



<a name="product.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## product.proto



<a name="product.CreateProductReq"></a>

### CreateProductReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| price | [int64](#int64) |  |  |
| shopID | [int64](#int64) |  |  |
| images | [ProductImage](#product.ProductImage) | repeated |  |






<a name="product.CreateProductRes"></a>

### CreateProductRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="product.GetProductInfoReq"></a>

### GetProductInfoReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| productID | [int64](#int64) |  |  |






<a name="product.GetProductInfoRes"></a>

### GetProductInfoRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product | [ProductInfo](#product.ProductInfo) |  |  |






<a name="product.GetProductListReq"></a>

### GetProductListReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |
| shopID | [int64](#int64) |  |  |






<a name="product.GetProductListRes"></a>

### GetProductListRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| products | [ProductInfo](#product.ProductInfo) | repeated |  |






<a name="product.ProductImage"></a>

### ProductImage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| productID | [int64](#int64) |  |  |
| ImageURL | [string](#string) |  |  |






<a name="product.ProductInfo"></a>

### ProductInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| productID | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| price | [int64](#int64) |  |  |
| shopID | [int64](#int64) |  |  |
| images | [ProductImage](#product.ProductImage) | repeated |  |
| status | [int32](#int32) |  |  |






<a name="product.UpdateImageReq"></a>

### UpdateImageReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| productID | [int64](#int64) |  |  |
| images | [ProductImage](#product.ProductImage) | repeated |  |






<a name="product.UpdateImageRes"></a>

### UpdateImageRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="product.UpdateProductReq"></a>

### UpdateProductReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| productID | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| price | [int64](#int64) |  |  |
| shopID | [int64](#int64) |  |  |
| images | [ProductImage](#product.ProductImage) | repeated |  |
| status | [int32](#int32) |  |  |






<a name="product.UpdateProductRes"></a>

### UpdateProductRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |





 

 

 


<a name="product.Product"></a>

### Product


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateProduct | [CreateProductReq](#product.CreateProductReq) | [CreateProductRes](#product.CreateProductRes) |  |
| UpdateProduct | [UpdateProductReq](#product.UpdateProductReq) | [UpdateProductRes](#product.UpdateProductRes) |  |
| UpdateImage | [UpdateImageReq](#product.UpdateImageReq) | [UpdateImageRes](#product.UpdateImageRes) |  |
| GetProductInfo | [GetProductInfoReq](#product.GetProductInfoReq) | [GetProductInfoRes](#product.GetProductInfoRes) |  |
| GetProductList | [GetProductListReq](#product.GetProductListReq) | [GetProductListRes](#product.GetProductListRes) |  |

 



<a name="product.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## product.proto



<a name="product.CreateProductReq"></a>

### CreateProductReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| price | [int64](#int64) |  |  |
| shopID | [int64](#int64) |  |  |
| images | [ProductImage](#product.ProductImage) | repeated |  |






<a name="product.CreateProductRes"></a>

### CreateProductRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="product.GetProductInfoReq"></a>

### GetProductInfoReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| productID | [int64](#int64) |  |  |






<a name="product.GetProductInfoRes"></a>

### GetProductInfoRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| product | [ProductInfo](#product.ProductInfo) |  |  |






<a name="product.GetProductListReq"></a>

### GetProductListReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int32](#int32) |  |  |
| offset | [int32](#int32) |  |  |
| shopID | [int64](#int64) |  |  |






<a name="product.GetProductListRes"></a>

### GetProductListRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| products | [ProductInfo](#product.ProductInfo) | repeated |  |






<a name="product.ProductImage"></a>

### ProductImage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| productID | [int64](#int64) |  |  |
| ImageURL | [string](#string) |  |  |






<a name="product.ProductInfo"></a>

### ProductInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| productID | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| price | [int64](#int64) |  |  |
| shopID | [int64](#int64) |  |  |
| images | [ProductImage](#product.ProductImage) | repeated |  |
| status | [int32](#int32) |  |  |






<a name="product.UpdateImageReq"></a>

### UpdateImageReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| productID | [int64](#int64) |  |  |
| images | [ProductImage](#product.ProductImage) | repeated |  |






<a name="product.UpdateImageRes"></a>

### UpdateImageRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |






<a name="product.UpdateProductReq"></a>

### UpdateProductReq



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| productID | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| description | [string](#string) |  |  |
| price | [int64](#int64) |  |  |
| shopID | [int64](#int64) |  |  |
| images | [ProductImage](#product.ProductImage) | repeated |  |
| status | [int32](#int32) |  |  |






<a name="product.UpdateProductRes"></a>

### UpdateProductRes



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| success | [bool](#bool) |  |  |





 

 

 


<a name="product.Product"></a>

### Product


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateProduct | [CreateProductReq](#product.CreateProductReq) | [CreateProductRes](#product.CreateProductRes) |  |
| UpdateProduct | [UpdateProductReq](#product.UpdateProductReq) | [UpdateProductRes](#product.UpdateProductRes) |  |
| UpdateImage | [UpdateImageReq](#product.UpdateImageReq) | [UpdateImageRes](#product.UpdateImageRes) |  |
| GetProductInfo | [GetProductInfoReq](#product.GetProductInfoReq) | [GetProductInfoRes](#product.GetProductInfoRes) |  |
| GetProductList | [GetProductListReq](#product.GetProductListReq) | [GetProductListRes](#product.GetProductListRes) |  |

 



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

