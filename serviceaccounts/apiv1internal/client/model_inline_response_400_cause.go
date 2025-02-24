/*
 * Service Accounts API Documentation
 *
 * This is the API documentation for Service Accounts
 *
 * API version: 1.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceaccountsclient

import (
	"encoding/json"
)

// InlineResponse400Cause struct for InlineResponse400Cause
type InlineResponse400Cause struct {

	StackTrace *[]InlineResponse400CauseStackTrace `json:"stackTrace,omitempty"`

	Message *string `json:"message,omitempty"`

	LocalizedMessage *string `json:"localizedMessage,omitempty"`

	Suppressed *[]InlineResponse400CauseSuppressed `json:"suppressed,omitempty"`

}

// NewInlineResponse400Cause instantiates a new InlineResponse400Cause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse400Cause() *InlineResponse400Cause {
	this := InlineResponse400Cause{}
	return &this
}

// NewInlineResponse400CauseWithDefaults instantiates a new InlineResponse400Cause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse400CauseWithDefaults() *InlineResponse400Cause {
	this := InlineResponse400Cause{}





	return &this
}


// GetStackTrace returns the StackTrace field value if set, zero value otherwise.
func (o *InlineResponse400Cause) GetStackTrace() []InlineResponse400CauseStackTrace {
	if o == nil || o.StackTrace == nil {
		var ret []InlineResponse400CauseStackTrace
		return ret
	}
	return *o.StackTrace
}

// GetStackTraceOk returns a tuple with the StackTrace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400Cause) GetStackTraceOk() (*[]InlineResponse400CauseStackTrace, bool) {
	if o == nil || o.StackTrace == nil {
		return nil, false
	}
	return o.StackTrace, true
}

// HasStackTrace returns a boolean if a field has been set.
func (o *InlineResponse400Cause) HasStackTrace() bool {
	if o != nil && o.StackTrace != nil {
		return true
	}

	return false
}

// SetStackTrace gets a reference to the given []InlineResponse400CauseStackTrace and assigns it to the StackTrace field.
func (o *InlineResponse400Cause) SetStackTrace(v []InlineResponse400CauseStackTrace) {
	o.StackTrace = &v
}


// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineResponse400Cause) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400Cause) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineResponse400Cause) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InlineResponse400Cause) SetMessage(v string) {
	o.Message = &v
}


// GetLocalizedMessage returns the LocalizedMessage field value if set, zero value otherwise.
func (o *InlineResponse400Cause) GetLocalizedMessage() string {
	if o == nil || o.LocalizedMessage == nil {
		var ret string
		return ret
	}
	return *o.LocalizedMessage
}

// GetLocalizedMessageOk returns a tuple with the LocalizedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400Cause) GetLocalizedMessageOk() (*string, bool) {
	if o == nil || o.LocalizedMessage == nil {
		return nil, false
	}
	return o.LocalizedMessage, true
}

// HasLocalizedMessage returns a boolean if a field has been set.
func (o *InlineResponse400Cause) HasLocalizedMessage() bool {
	if o != nil && o.LocalizedMessage != nil {
		return true
	}

	return false
}

// SetLocalizedMessage gets a reference to the given string and assigns it to the LocalizedMessage field.
func (o *InlineResponse400Cause) SetLocalizedMessage(v string) {
	o.LocalizedMessage = &v
}


// GetSuppressed returns the Suppressed field value if set, zero value otherwise.
func (o *InlineResponse400Cause) GetSuppressed() []InlineResponse400CauseSuppressed {
	if o == nil || o.Suppressed == nil {
		var ret []InlineResponse400CauseSuppressed
		return ret
	}
	return *o.Suppressed
}

// GetSuppressedOk returns a tuple with the Suppressed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponse400Cause) GetSuppressedOk() (*[]InlineResponse400CauseSuppressed, bool) {
	if o == nil || o.Suppressed == nil {
		return nil, false
	}
	return o.Suppressed, true
}

// HasSuppressed returns a boolean if a field has been set.
func (o *InlineResponse400Cause) HasSuppressed() bool {
	if o != nil && o.Suppressed != nil {
		return true
	}

	return false
}

// SetSuppressed gets a reference to the given []InlineResponse400CauseSuppressed and assigns it to the Suppressed field.
func (o *InlineResponse400Cause) SetSuppressed(v []InlineResponse400CauseSuppressed) {
	o.Suppressed = &v
}


func (o InlineResponse400Cause) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if o.StackTrace != nil {
		toSerialize["stackTrace"] = o.StackTrace
	}
    
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
    
	if o.LocalizedMessage != nil {
		toSerialize["localizedMessage"] = o.LocalizedMessage
	}
    
	if o.Suppressed != nil {
		toSerialize["suppressed"] = o.Suppressed
	}
    
	return json.Marshal(toSerialize)
}

type NullableInlineResponse400Cause struct {
	value *InlineResponse400Cause
	isSet bool
}

func (v NullableInlineResponse400Cause) Get() *InlineResponse400Cause {
	return v.value
}

func (v *NullableInlineResponse400Cause) Set(val *InlineResponse400Cause) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse400Cause) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse400Cause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse400Cause(val *InlineResponse400Cause) *NullableInlineResponse400Cause {
	return &NullableInlineResponse400Cause{value: val, isSet: true}
}

func (v NullableInlineResponse400Cause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse400Cause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

