// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: product.proto

package product

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *CreateProductReq) Validate() error {
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Description == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Description", fmt.Errorf(`value '%v' must not be an empty string`, this.Description))
	}
	if !(this.ShopID > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ShopID", fmt.Errorf(`value '%v' must be greater than '0'`, this.ShopID))
	}
	if len(this.Images) < 1 {
		return github_com_mwitkow_go_proto_validators.FieldError("Images", fmt.Errorf(`value '%v' must contain at least 1 elements`, this.Images))
	}
	for _, item := range this.Images {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Images", err)
			}
		}
	}
	return nil
}
func (this *ProductImage) Validate() error {
	return nil
}
func (this *CreateProductRes) Validate() error {
	return nil
}
func (this *UpdateProductReq) Validate() error {
	if !(this.ProductID > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProductID", fmt.Errorf(`value '%v' must be greater than '0'`, this.ProductID))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Description == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Description", fmt.Errorf(`value '%v' must not be an empty string`, this.Description))
	}
	if !(this.ShopID > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ShopID", fmt.Errorf(`value '%v' must be greater than '0'`, this.ShopID))
	}
	if len(this.Images) < 1 {
		return github_com_mwitkow_go_proto_validators.FieldError("Images", fmt.Errorf(`value '%v' must contain at least 1 elements`, this.Images))
	}
	for _, item := range this.Images {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Images", err)
			}
		}
	}
	return nil
}
func (this *UpdateProductRes) Validate() error {
	return nil
}
func (this *UpdateImageReq) Validate() error {
	if !(this.ProductID > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProductID", fmt.Errorf(`value '%v' must be greater than '0'`, this.ProductID))
	}
	if len(this.Images) < 1 {
		return github_com_mwitkow_go_proto_validators.FieldError("Images", fmt.Errorf(`value '%v' must contain at least 1 elements`, this.Images))
	}
	for _, item := range this.Images {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Images", err)
			}
		}
	}
	return nil
}
func (this *UpdateImageRes) Validate() error {
	return nil
}
func (this *GetProductInfoReq) Validate() error {
	if !(this.ProductID > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProductID", fmt.Errorf(`value '%v' must be greater than '0'`, this.ProductID))
	}
	return nil
}
func (this *GetProductInfoRes) Validate() error {
	if this.Product != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Product); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Product", err)
		}
	}
	return nil
}
func (this *GetProductListReq) Validate() error {
	if !(this.Limit > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Limit", fmt.Errorf(`value '%v' must be greater than '0'`, this.Limit))
	}
	if !(this.ShopID > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ShopID", fmt.Errorf(`value '%v' must be greater than '0'`, this.ShopID))
	}
	return nil
}
func (this *GetProductListRes) Validate() error {
	for _, item := range this.Products {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Products", err)
			}
		}
	}
	return nil
}
func (this *ProductInfo) Validate() error {
	if !(this.ProductID > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ProductID", fmt.Errorf(`value '%v' must be greater than '0'`, this.ProductID))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	if this.Description == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Description", fmt.Errorf(`value '%v' must not be an empty string`, this.Description))
	}
	if !(this.ShopID > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("ShopID", fmt.Errorf(`value '%v' must be greater than '0'`, this.ShopID))
	}
	for _, item := range this.Images {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Images", err)
			}
		}
	}
	return nil
}
