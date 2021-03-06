// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: aggregation/v1/aggregation.proto

package aggregationv1

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
var _aggregation_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on KeyerConfiguration with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *KeyerConfiguration) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetFragments()) < 1 {
		return KeyerConfigurationValidationError{
			field:  "Fragments",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetFragments() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return KeyerConfigurationValidationError{
					field:  fmt.Sprintf("Fragments[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// KeyerConfigurationValidationError is the validation error returned by
// KeyerConfiguration.Validate if the designated constraints aren't met.
type KeyerConfigurationValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeyerConfigurationValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeyerConfigurationValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeyerConfigurationValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeyerConfigurationValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeyerConfigurationValidationError) ErrorName() string {
	return "KeyerConfigurationValidationError"
}

// Error satisfies the builtin error interface
func (e KeyerConfigurationValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeyerConfiguration.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeyerConfigurationValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeyerConfigurationValidationError{}

// Validate checks the field values on MatchPredicate with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *MatchPredicate) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Type.(type) {

	case *MatchPredicate_AndMatch:

		if v, ok := interface{}(m.GetAndMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					field:  "AndMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchPredicate_OrMatch:

		if v, ok := interface{}(m.GetOrMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					field:  "OrMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchPredicate_NotMatch:

		if v, ok := interface{}(m.GetNotMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					field:  "NotMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchPredicate_AnyMatch:

		if m.GetAnyMatch() != true {
			return MatchPredicateValidationError{
				field:  "AnyMatch",
				reason: "value must equal true",
			}
		}

	case *MatchPredicate_RequestTypeMatch_:

		if v, ok := interface{}(m.GetRequestTypeMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					field:  "RequestTypeMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *MatchPredicate_RequestNodeMatch_:

		if v, ok := interface{}(m.GetRequestNodeMatch()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicateValidationError{
					field:  "RequestNodeMatch",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return MatchPredicateValidationError{
			field:  "Type",
			reason: "value is required",
		}

	}

	return nil
}

// MatchPredicateValidationError is the validation error returned by
// MatchPredicate.Validate if the designated constraints aren't met.
type MatchPredicateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MatchPredicateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MatchPredicateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MatchPredicateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MatchPredicateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MatchPredicateValidationError) ErrorName() string { return "MatchPredicateValidationError" }

// Error satisfies the builtin error interface
func (e MatchPredicateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatchPredicate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MatchPredicateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MatchPredicateValidationError{}

// Validate checks the field values on ResultPredicate with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ResultPredicate) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Type.(type) {

	case *ResultPredicate_AndResult_:

		if v, ok := interface{}(m.GetAndResult()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResultPredicateValidationError{
					field:  "AndResult",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ResultPredicate_RequestNodeFragment_:

		if v, ok := interface{}(m.GetRequestNodeFragment()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResultPredicateValidationError{
					field:  "RequestNodeFragment",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ResultPredicate_ResourceNamesFragment_:

		if v, ok := interface{}(m.GetResourceNamesFragment()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResultPredicateValidationError{
					field:  "ResourceNamesFragment",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *ResultPredicate_StringFragment:
		// no validation rules for StringFragment

	default:
		return ResultPredicateValidationError{
			field:  "Type",
			reason: "value is required",
		}

	}

	return nil
}

// ResultPredicateValidationError is the validation error returned by
// ResultPredicate.Validate if the designated constraints aren't met.
type ResultPredicateValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultPredicateValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultPredicateValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultPredicateValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultPredicateValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultPredicateValidationError) ErrorName() string { return "ResultPredicateValidationError" }

// Error satisfies the builtin error interface
func (e ResultPredicateValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResultPredicate.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultPredicateValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultPredicateValidationError{}

// Validate checks the field values on KeyerConfiguration_Fragment with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *KeyerConfiguration_Fragment) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRules()) < 1 {
		return KeyerConfiguration_FragmentValidationError{
			field:  "Rules",
			reason: "value must contain at least 1 item(s)",
		}
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return KeyerConfiguration_FragmentValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// KeyerConfiguration_FragmentValidationError is the validation error returned
// by KeyerConfiguration_Fragment.Validate if the designated constraints
// aren't met.
type KeyerConfiguration_FragmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeyerConfiguration_FragmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeyerConfiguration_FragmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeyerConfiguration_FragmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeyerConfiguration_FragmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeyerConfiguration_FragmentValidationError) ErrorName() string {
	return "KeyerConfiguration_FragmentValidationError"
}

// Error satisfies the builtin error interface
func (e KeyerConfiguration_FragmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeyerConfiguration_Fragment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeyerConfiguration_FragmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeyerConfiguration_FragmentValidationError{}

// Validate checks the field values on KeyerConfiguration_Fragment_Rule with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *KeyerConfiguration_Fragment_Rule) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetMatch() == nil {
		return KeyerConfiguration_Fragment_RuleValidationError{
			field:  "Match",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetMatch()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return KeyerConfiguration_Fragment_RuleValidationError{
				field:  "Match",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetResult() == nil {
		return KeyerConfiguration_Fragment_RuleValidationError{
			field:  "Result",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetResult()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return KeyerConfiguration_Fragment_RuleValidationError{
				field:  "Result",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// KeyerConfiguration_Fragment_RuleValidationError is the validation error
// returned by KeyerConfiguration_Fragment_Rule.Validate if the designated
// constraints aren't met.
type KeyerConfiguration_Fragment_RuleValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e KeyerConfiguration_Fragment_RuleValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e KeyerConfiguration_Fragment_RuleValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e KeyerConfiguration_Fragment_RuleValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e KeyerConfiguration_Fragment_RuleValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e KeyerConfiguration_Fragment_RuleValidationError) ErrorName() string {
	return "KeyerConfiguration_Fragment_RuleValidationError"
}

// Error satisfies the builtin error interface
func (e KeyerConfiguration_Fragment_RuleValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sKeyerConfiguration_Fragment_Rule.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = KeyerConfiguration_Fragment_RuleValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = KeyerConfiguration_Fragment_RuleValidationError{}

// Validate checks the field values on MatchPredicate_RequestTypeMatch with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MatchPredicate_RequestTypeMatch) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetTypes()) < 1 {
		return MatchPredicate_RequestTypeMatchValidationError{
			field:  "Types",
			reason: "value must contain at least 1 item(s)",
		}
	}

	return nil
}

// MatchPredicate_RequestTypeMatchValidationError is the validation error
// returned by MatchPredicate_RequestTypeMatch.Validate if the designated
// constraints aren't met.
type MatchPredicate_RequestTypeMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MatchPredicate_RequestTypeMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MatchPredicate_RequestTypeMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MatchPredicate_RequestTypeMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MatchPredicate_RequestTypeMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MatchPredicate_RequestTypeMatchValidationError) ErrorName() string {
	return "MatchPredicate_RequestTypeMatchValidationError"
}

// Error satisfies the builtin error interface
func (e MatchPredicate_RequestTypeMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatchPredicate_RequestTypeMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MatchPredicate_RequestTypeMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MatchPredicate_RequestTypeMatchValidationError{}

// Validate checks the field values on MatchPredicate_RequestNodeMatch with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MatchPredicate_RequestNodeMatch) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := NodeFieldType_name[int32(m.GetField())]; !ok {
		return MatchPredicate_RequestNodeMatchValidationError{
			field:  "Field",
			reason: "value must be one of the defined enum values",
		}
	}

	switch m.Type.(type) {

	case *MatchPredicate_RequestNodeMatch_ExactMatch:
		// no validation rules for ExactMatch

	case *MatchPredicate_RequestNodeMatch_RegexMatch:
		// no validation rules for RegexMatch

	default:
		return MatchPredicate_RequestNodeMatchValidationError{
			field:  "Type",
			reason: "value is required",
		}

	}

	return nil
}

// MatchPredicate_RequestNodeMatchValidationError is the validation error
// returned by MatchPredicate_RequestNodeMatch.Validate if the designated
// constraints aren't met.
type MatchPredicate_RequestNodeMatchValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MatchPredicate_RequestNodeMatchValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MatchPredicate_RequestNodeMatchValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MatchPredicate_RequestNodeMatchValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MatchPredicate_RequestNodeMatchValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MatchPredicate_RequestNodeMatchValidationError) ErrorName() string {
	return "MatchPredicate_RequestNodeMatchValidationError"
}

// Error satisfies the builtin error interface
func (e MatchPredicate_RequestNodeMatchValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatchPredicate_RequestNodeMatch.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MatchPredicate_RequestNodeMatchValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MatchPredicate_RequestNodeMatchValidationError{}

// Validate checks the field values on MatchPredicate_MatchSet with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MatchPredicate_MatchSet) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetRules()) < 2 {
		return MatchPredicate_MatchSetValidationError{
			field:  "Rules",
			reason: "value must contain at least 2 item(s)",
		}
	}

	for idx, item := range m.GetRules() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MatchPredicate_MatchSetValidationError{
					field:  fmt.Sprintf("Rules[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MatchPredicate_MatchSetValidationError is the validation error returned by
// MatchPredicate_MatchSet.Validate if the designated constraints aren't met.
type MatchPredicate_MatchSetValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MatchPredicate_MatchSetValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MatchPredicate_MatchSetValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MatchPredicate_MatchSetValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MatchPredicate_MatchSetValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MatchPredicate_MatchSetValidationError) ErrorName() string {
	return "MatchPredicate_MatchSetValidationError"
}

// Error satisfies the builtin error interface
func (e MatchPredicate_MatchSetValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMatchPredicate_MatchSet.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MatchPredicate_MatchSetValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MatchPredicate_MatchSetValidationError{}

// Validate checks the field values on ResultPredicate_ResultAction with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ResultPredicate_ResultAction) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Action.(type) {

	case *ResultPredicate_ResultAction_Exact:

		if m.GetExact() != true {
			return ResultPredicate_ResultActionValidationError{
				field:  "Exact",
				reason: "value must equal true",
			}
		}

	case *ResultPredicate_ResultAction_RegexAction_:

		if v, ok := interface{}(m.GetRegexAction()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResultPredicate_ResultActionValidationError{
					field:  "RegexAction",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	default:
		return ResultPredicate_ResultActionValidationError{
			field:  "Action",
			reason: "value is required",
		}

	}

	return nil
}

// ResultPredicate_ResultActionValidationError is the validation error returned
// by ResultPredicate_ResultAction.Validate if the designated constraints
// aren't met.
type ResultPredicate_ResultActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultPredicate_ResultActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultPredicate_ResultActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultPredicate_ResultActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultPredicate_ResultActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultPredicate_ResultActionValidationError) ErrorName() string {
	return "ResultPredicate_ResultActionValidationError"
}

// Error satisfies the builtin error interface
func (e ResultPredicate_ResultActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResultPredicate_ResultAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultPredicate_ResultActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultPredicate_ResultActionValidationError{}

// Validate checks the field values on ResultPredicate_AndResult with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ResultPredicate_AndResult) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetResultPredicates()) < 2 {
		return ResultPredicate_AndResultValidationError{
			field:  "ResultPredicates",
			reason: "value must contain at least 2 item(s)",
		}
	}

	for idx, item := range m.GetResultPredicates() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ResultPredicate_AndResultValidationError{
					field:  fmt.Sprintf("ResultPredicates[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ResultPredicate_AndResultValidationError is the validation error returned by
// ResultPredicate_AndResult.Validate if the designated constraints aren't met.
type ResultPredicate_AndResultValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultPredicate_AndResultValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultPredicate_AndResultValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultPredicate_AndResultValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultPredicate_AndResultValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultPredicate_AndResultValidationError) ErrorName() string {
	return "ResultPredicate_AndResultValidationError"
}

// Error satisfies the builtin error interface
func (e ResultPredicate_AndResultValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResultPredicate_AndResult.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultPredicate_AndResultValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultPredicate_AndResultValidationError{}

// Validate checks the field values on ResultPredicate_RequestNodeFragment with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *ResultPredicate_RequestNodeFragment) Validate() error {
	if m == nil {
		return nil
	}

	if _, ok := NodeFieldType_name[int32(m.GetField())]; !ok {
		return ResultPredicate_RequestNodeFragmentValidationError{
			field:  "Field",
			reason: "value must be one of the defined enum values",
		}
	}

	if m.GetAction() == nil {
		return ResultPredicate_RequestNodeFragmentValidationError{
			field:  "Action",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAction()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResultPredicate_RequestNodeFragmentValidationError{
				field:  "Action",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ResultPredicate_RequestNodeFragmentValidationError is the validation error
// returned by ResultPredicate_RequestNodeFragment.Validate if the designated
// constraints aren't met.
type ResultPredicate_RequestNodeFragmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultPredicate_RequestNodeFragmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultPredicate_RequestNodeFragmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultPredicate_RequestNodeFragmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultPredicate_RequestNodeFragmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultPredicate_RequestNodeFragmentValidationError) ErrorName() string {
	return "ResultPredicate_RequestNodeFragmentValidationError"
}

// Error satisfies the builtin error interface
func (e ResultPredicate_RequestNodeFragmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResultPredicate_RequestNodeFragment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultPredicate_RequestNodeFragmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultPredicate_RequestNodeFragmentValidationError{}

// Validate checks the field values on ResultPredicate_ResourceNamesFragment
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ResultPredicate_ResourceNamesFragment) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetElement() < 0 {
		return ResultPredicate_ResourceNamesFragmentValidationError{
			field:  "Element",
			reason: "value must be greater than or equal to 0",
		}
	}

	if m.GetAction() == nil {
		return ResultPredicate_ResourceNamesFragmentValidationError{
			field:  "Action",
			reason: "value is required",
		}
	}

	if v, ok := interface{}(m.GetAction()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ResultPredicate_ResourceNamesFragmentValidationError{
				field:  "Action",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// ResultPredicate_ResourceNamesFragmentValidationError is the validation error
// returned by ResultPredicate_ResourceNamesFragment.Validate if the
// designated constraints aren't met.
type ResultPredicate_ResourceNamesFragmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultPredicate_ResourceNamesFragmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultPredicate_ResourceNamesFragmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultPredicate_ResourceNamesFragmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultPredicate_ResourceNamesFragmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultPredicate_ResourceNamesFragmentValidationError) ErrorName() string {
	return "ResultPredicate_ResourceNamesFragmentValidationError"
}

// Error satisfies the builtin error interface
func (e ResultPredicate_ResourceNamesFragmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResultPredicate_ResourceNamesFragment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultPredicate_ResourceNamesFragmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultPredicate_ResourceNamesFragmentValidationError{}

// Validate checks the field values on ResultPredicate_ResultAction_RegexAction
// with the rules defined in the proto definition for this message. If any
// rules are violated, an error is returned.
func (m *ResultPredicate_ResultAction_RegexAction) Validate() error {
	if m == nil {
		return nil
	}

	if utf8.RuneCountInString(m.GetPattern()) < 1 {
		return ResultPredicate_ResultAction_RegexActionValidationError{
			field:  "Pattern",
			reason: "value length must be at least 1 runes",
		}
	}

	if utf8.RuneCountInString(m.GetReplace()) < 0 {
		return ResultPredicate_ResultAction_RegexActionValidationError{
			field:  "Replace",
			reason: "value length must be at least 0 runes",
		}
	}

	return nil
}

// ResultPredicate_ResultAction_RegexActionValidationError is the validation
// error returned by ResultPredicate_ResultAction_RegexAction.Validate if the
// designated constraints aren't met.
type ResultPredicate_ResultAction_RegexActionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResultPredicate_ResultAction_RegexActionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResultPredicate_ResultAction_RegexActionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResultPredicate_ResultAction_RegexActionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResultPredicate_ResultAction_RegexActionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResultPredicate_ResultAction_RegexActionValidationError) ErrorName() string {
	return "ResultPredicate_ResultAction_RegexActionValidationError"
}

// Error satisfies the builtin error interface
func (e ResultPredicate_ResultAction_RegexActionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResultPredicate_ResultAction_RegexAction.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResultPredicate_ResultAction_RegexActionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResultPredicate_ResultAction_RegexActionValidationError{}
