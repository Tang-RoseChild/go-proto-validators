// Code generated by protoc-gen-gogo.
// source: examples/nested.proto
// DO NOT EDIT!

/*
Package validator_examples is a generated protocol buffer package.

It is generated from these files:
	examples/nested.proto

It has these top-level messages:
	InnerMessage
	OuterMessage
*/
package validator_examples

import regexp "regexp"
import fmt "fmt"
import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import math "math"
import _ "github.com/mwitkow/go-proto-validators"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *InnerMessage) Validate() error {
	if !(this.SomeInteger > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("SomeInteger", fmt.Errorf(`value '%v' must be greater than '0'`, this.SomeInteger))
	}
	if !(this.SomeInteger < 100) {
		return github_com_mwitkow_go_proto_validators.FieldError("SomeInteger", fmt.Errorf(`value '%v' must be smaller than '100'`, this.SomeInteger))
	}
	if !(this.Type >= 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must be  equal or greater than  '0'`, this.Type))
	}
	if !(this.Type <= 5) {
		return github_com_mwitkow_go_proto_validators.FieldError("Type", fmt.Errorf(`value '%v' must be equal or smaller than '5'`, this.Type))
	}
	if !(len(this.Content) >= 3) {
		return github_com_mwitkow_go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must string length equal or greater than 3`, this.Content))
	}
	if !(len(this.Content) <= 15) {
		return github_com_mwitkow_go_proto_validators.FieldError("Content", fmt.Errorf(`value '%v' must string length equal or smaller than 15`, this.Content))
	}
	return nil
}

var _regex_OuterMessage_ImportantString = regexp.MustCompile("^[a-z]{2,5}$")

func (this *OuterMessage) Validate() error {
	if !_regex_OuterMessage_ImportantString.MatchString(this.ImportantString) {
		return github_com_mwitkow_go_proto_validators.FieldError("ImportantString", fmt.Errorf(`value '%v' must be a string conforming to regex "^[a-z]{2,5}$"`, this.ImportantString))
	}
	if nil == this.Inner {
		return github_com_mwitkow_go_proto_validators.FieldError("Inner", fmt.Errorf("message must exist"))
	}
	if this.Inner != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Inner); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Inner", err)
		}
	}
	return nil
}
