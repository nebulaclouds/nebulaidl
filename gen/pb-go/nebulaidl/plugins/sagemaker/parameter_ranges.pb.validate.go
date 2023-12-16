// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: nebulaidl/plugins/sagemaker/parameter_ranges.proto

package plugins

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
var _parameter_ranges_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on HyperparameterScalingType with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HyperparameterScalingType) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// HyperparameterScalingTypeValidationError is the validation error returned by
// HyperparameterScalingType.Validate if the designated constraints aren't met.
type HyperparameterScalingTypeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HyperparameterScalingTypeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HyperparameterScalingTypeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HyperparameterScalingTypeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HyperparameterScalingTypeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HyperparameterScalingTypeValidationError) ErrorName() string {
	return "HyperparameterScalingTypeValidationError"
}

// Error satisfies the builtin error interface
func (e HyperparameterScalingTypeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHyperparameterScalingType.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HyperparameterScalingTypeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HyperparameterScalingTypeValidationError{}

// Validate checks the field values on ContinuousParameterRange with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ContinuousParameterRange) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MaxValue

	// no validation rules for MinValue

	// no validation rules for ScalingType

	return nil
}

// ContinuousParameterRangeValidationError is the validation error returned by
// ContinuousParameterRange.Validate if the designated constraints aren't met.
type ContinuousParameterRangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContinuousParameterRangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContinuousParameterRangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContinuousParameterRangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContinuousParameterRangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContinuousParameterRangeValidationError) ErrorName() string {
	return "ContinuousParameterRangeValidationError"
}

// Error satisfies the builtin error interface
func (e ContinuousParameterRangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContinuousParameterRange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContinuousParameterRangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContinuousParameterRangeValidationError{}

// Validate checks the field values on IntegerParameterRange with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *IntegerParameterRange) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for MaxValue

	// no validation rules for MinValue

	// no validation rules for ScalingType

	return nil
}

// IntegerParameterRangeValidationError is the validation error returned by
// IntegerParameterRange.Validate if the designated constraints aren't met.
type IntegerParameterRangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IntegerParameterRangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IntegerParameterRangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IntegerParameterRangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IntegerParameterRangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IntegerParameterRangeValidationError) ErrorName() string {
	return "IntegerParameterRangeValidationError"
}

// Error satisfies the builtin error interface
func (e IntegerParameterRangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIntegerParameterRange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IntegerParameterRangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IntegerParameterRangeValidationError{}

// Validate checks the field values on CategoricalParameterRange with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CategoricalParameterRange) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CategoricalParameterRangeValidationError is the validation error returned by
// CategoricalParameterRange.Validate if the designated constraints aren't met.
type CategoricalParameterRangeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CategoricalParameterRangeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CategoricalParameterRangeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CategoricalParameterRangeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CategoricalParameterRangeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CategoricalParameterRangeValidationError) ErrorName() string {
	return "CategoricalParameterRangeValidationError"
}

// Error satisfies the builtin error interface
func (e CategoricalParameterRangeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCategoricalParameterRange.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CategoricalParameterRangeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CategoricalParameterRangeValidationError{}

// Validate checks the field values on ParameterRangeOneOf with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ParameterRangeOneOf) Validate() error {
	if m == nil {
		return nil
	}

	switch m.ParameterRangeType.(type) {

	case *ParameterRangeOneOf_ContinuousParameterRange:

		if v, ok := interface{}(m.GetContinuousParameterRange()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ParameterRangeOneOfValidationError{
					field:  "ContinuousParameterRange",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ParameterRangeOneOf_IntegerParameterRange:

		if v, ok := interface{}(m.GetIntegerParameterRange()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ParameterRangeOneOfValidationError{
					field:  "IntegerParameterRange",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ParameterRangeOneOf_CategoricalParameterRange:

		if v, ok := interface{}(m.GetCategoricalParameterRange()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ParameterRangeOneOfValidationError{
					field:  "CategoricalParameterRange",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ParameterRangeOneOfValidationError is the validation error returned by
// ParameterRangeOneOf.Validate if the designated constraints aren't met.
type ParameterRangeOneOfValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ParameterRangeOneOfValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ParameterRangeOneOfValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ParameterRangeOneOfValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ParameterRangeOneOfValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ParameterRangeOneOfValidationError) ErrorName() string {
	return "ParameterRangeOneOfValidationError"
}

// Error satisfies the builtin error interface
func (e ParameterRangeOneOfValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sParameterRangeOneOf.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ParameterRangeOneOfValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ParameterRangeOneOfValidationError{}

// Validate checks the field values on ParameterRanges with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ParameterRanges) Validate() error {
	if m == nil {
		return nil
	}

	for key, val := range m.GetParameterRangeMap() {
		_ = val

		// no validation rules for ParameterRangeMap[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ParameterRangesValidationError{
					field:  fmt.Sprintf("ParameterRangeMap[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ParameterRangesValidationError is the validation error returned by
// ParameterRanges.Validate if the designated constraints aren't met.
type ParameterRangesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ParameterRangesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ParameterRangesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ParameterRangesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ParameterRangesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ParameterRangesValidationError) ErrorName() string { return "ParameterRangesValidationError" }

// Error satisfies the builtin error interface
func (e ParameterRangesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sParameterRanges.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ParameterRangesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ParameterRangesValidationError{}
