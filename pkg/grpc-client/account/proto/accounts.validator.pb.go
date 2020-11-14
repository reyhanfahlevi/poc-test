// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: accounts.proto

package account

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *GetUserInfoReq) Validate() error {
	if !(this.UserID > -1) {
		return github_com_mwitkow_go_proto_validators.FieldError("UserID", fmt.Errorf(`value '%v' must be greater than '-1'`, this.UserID))
	}
	return nil
}
func (this *GetUserInfoResp) Validate() error {
	if this.JoinDate != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.JoinDate); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("JoinDate", err)
		}
	}
	return nil
}
func (this *CheckUserIsAuthenticatedReq) Validate() error {
	return nil
}
func (this *CheckUserIsAuthenticatedRes) Validate() error {
	return nil
}
func (this *Date) Validate() error {
	return nil
}
func (this *GetAccessTokenReq) Validate() error {
	return nil
}
func (this *GetAccessTokenResp) Validate() error {
	return nil
}
func (this *CreateUserReq) Validate() error {
	if this.Email == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Email", fmt.Errorf(`value '%v' must not be an empty string`, this.Email))
	}
	if this.Password == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Password", fmt.Errorf(`value '%v' must not be an empty string`, this.Password))
	}
	if this.Name == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Name", fmt.Errorf(`value '%v' must not be an empty string`, this.Name))
	}
	return nil
}
func (this *CreateUserRes) Validate() error {
	return nil
}