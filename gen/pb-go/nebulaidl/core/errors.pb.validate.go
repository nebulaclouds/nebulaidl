// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: nebulaidl/core/errors.proto

package core

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
var _errors_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ContainerError with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ContainerError) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Code

	// no validation rules for Message

	// no validation rules for Kind

	// no validation rules for Origin

	return nil
}

// ContainerErrorValidationError is the validation error returned by
// ContainerError.Validate if the designated constraints aren't met.
type ContainerErrorValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContainerErrorValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContainerErrorValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContainerErrorValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContainerErrorValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContainerErrorValidationError) ErrorName() string { return "ContainerErrorValidationError" }

// Error satisfies the builtin error interface
func (e ContainerErrorValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContainerError.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContainerErrorValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContainerErrorValidationError{}

// Validate checks the field values on ErrorDocument with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ErrorDocument) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetError()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ErrorDocumentValidationError{
				field:  "Error",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ErrorDocumentValidationError is the validation error returned by
// ErrorDocument.Validate if the designated constraints aren't met.
type ErrorDocumentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorDocumentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorDocumentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorDocumentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorDocumentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorDocumentValidationError) ErrorName() string { return "ErrorDocumentValidationError" }

// Error satisfies the builtin error interface
func (e ErrorDocumentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sErrorDocument.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorDocumentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorDocumentValidationError{}
