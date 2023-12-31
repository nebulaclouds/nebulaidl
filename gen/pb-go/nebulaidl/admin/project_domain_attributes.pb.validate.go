// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: nebulaidl/admin/project_domain_attributes.proto

package admin

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _project_domain_attributes_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ProjectDomainAttributes with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ProjectDomainAttributes) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Project

	// no validation rules for Domain

	if v, ok := interface{}(m.GetMatchingAttributes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectDomainAttributesValidationError{
				field:  "MatchingAttributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProjectDomainAttributesValidationError is the validation error returned by
// ProjectDomainAttributes.Validate if the designated constraints aren't met.
type ProjectDomainAttributesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectDomainAttributesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectDomainAttributesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectDomainAttributesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectDomainAttributesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectDomainAttributesValidationError) ErrorName() string {
	return "ProjectDomainAttributesValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectDomainAttributesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectDomainAttributes.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectDomainAttributesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectDomainAttributesValidationError{}

// Validate checks the field values on ProjectDomainAttributesUpdateRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ProjectDomainAttributesUpdateRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetAttributes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectDomainAttributesUpdateRequestValidationError{
				field:  "Attributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProjectDomainAttributesUpdateRequestValidationError is the validation error
// returned by ProjectDomainAttributesUpdateRequest.Validate if the designated
// constraints aren't met.
type ProjectDomainAttributesUpdateRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectDomainAttributesUpdateRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectDomainAttributesUpdateRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectDomainAttributesUpdateRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectDomainAttributesUpdateRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectDomainAttributesUpdateRequestValidationError) ErrorName() string {
	return "ProjectDomainAttributesUpdateRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectDomainAttributesUpdateRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectDomainAttributesUpdateRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectDomainAttributesUpdateRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectDomainAttributesUpdateRequestValidationError{}

// Validate checks the field values on ProjectDomainAttributesUpdateResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ProjectDomainAttributesUpdateResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ProjectDomainAttributesUpdateResponseValidationError is the validation error
// returned by ProjectDomainAttributesUpdateResponse.Validate if the
// designated constraints aren't met.
type ProjectDomainAttributesUpdateResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectDomainAttributesUpdateResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectDomainAttributesUpdateResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectDomainAttributesUpdateResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectDomainAttributesUpdateResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectDomainAttributesUpdateResponseValidationError) ErrorName() string {
	return "ProjectDomainAttributesUpdateResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectDomainAttributesUpdateResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectDomainAttributesUpdateResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectDomainAttributesUpdateResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectDomainAttributesUpdateResponseValidationError{}

// Validate checks the field values on ProjectDomainAttributesGetRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ProjectDomainAttributesGetRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Project

	// no validation rules for Domain

	// no validation rules for ResourceType

	return nil
}

// ProjectDomainAttributesGetRequestValidationError is the validation error
// returned by ProjectDomainAttributesGetRequest.Validate if the designated
// constraints aren't met.
type ProjectDomainAttributesGetRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectDomainAttributesGetRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectDomainAttributesGetRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectDomainAttributesGetRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectDomainAttributesGetRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectDomainAttributesGetRequestValidationError) ErrorName() string {
	return "ProjectDomainAttributesGetRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectDomainAttributesGetRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectDomainAttributesGetRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectDomainAttributesGetRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectDomainAttributesGetRequestValidationError{}

// Validate checks the field values on ProjectDomainAttributesGetResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ProjectDomainAttributesGetResponse) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetAttributes()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ProjectDomainAttributesGetResponseValidationError{
				field:  "Attributes",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ProjectDomainAttributesGetResponseValidationError is the validation error
// returned by ProjectDomainAttributesGetResponse.Validate if the designated
// constraints aren't met.
type ProjectDomainAttributesGetResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectDomainAttributesGetResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectDomainAttributesGetResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectDomainAttributesGetResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectDomainAttributesGetResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectDomainAttributesGetResponseValidationError) ErrorName() string {
	return "ProjectDomainAttributesGetResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectDomainAttributesGetResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectDomainAttributesGetResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectDomainAttributesGetResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectDomainAttributesGetResponseValidationError{}

// Validate checks the field values on ProjectDomainAttributesDeleteRequest
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ProjectDomainAttributesDeleteRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Project

	// no validation rules for Domain

	// no validation rules for ResourceType

	return nil
}

// ProjectDomainAttributesDeleteRequestValidationError is the validation error
// returned by ProjectDomainAttributesDeleteRequest.Validate if the designated
// constraints aren't met.
type ProjectDomainAttributesDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectDomainAttributesDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectDomainAttributesDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectDomainAttributesDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectDomainAttributesDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectDomainAttributesDeleteRequestValidationError) ErrorName() string {
	return "ProjectDomainAttributesDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectDomainAttributesDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectDomainAttributesDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectDomainAttributesDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectDomainAttributesDeleteRequestValidationError{}

// Validate checks the field values on ProjectDomainAttributesDeleteResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ProjectDomainAttributesDeleteResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// ProjectDomainAttributesDeleteResponseValidationError is the validation error
// returned by ProjectDomainAttributesDeleteResponse.Validate if the
// designated constraints aren't met.
type ProjectDomainAttributesDeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ProjectDomainAttributesDeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ProjectDomainAttributesDeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ProjectDomainAttributesDeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ProjectDomainAttributesDeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ProjectDomainAttributesDeleteResponseValidationError) ErrorName() string {
	return "ProjectDomainAttributesDeleteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ProjectDomainAttributesDeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sProjectDomainAttributesDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ProjectDomainAttributesDeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ProjectDomainAttributesDeleteResponseValidationError{}
