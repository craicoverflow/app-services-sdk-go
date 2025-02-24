/*
 * Kafka Admin REST API
 *
 * An API to provide REST endpoints for query Kafka for admin operations
 *
 * API version: 0.8.1-SNAPSHOT
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kafkainstanceclient

import (
	"encoding/json"
)

// AclBindingListPage A page of ACL binding entries
type AclBindingListPage struct {

	Items []AclBinding `json:"items"`

	// Total number of entries in the full result set
	Total int32 `json:"total"`

	// Number of entries per page (returned for fetch requests)
	Size *int32 `json:"size,omitempty"`

	// Current page number (returned for fetch requests)
	Page *int32 `json:"page,omitempty"`

}

// NewAclBindingListPage instantiates a new AclBindingListPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAclBindingListPage(items []AclBinding, total int32) *AclBindingListPage {
	this := AclBindingListPage{}
	this.Items = items
	this.Total = total
	return &this
}

// NewAclBindingListPageWithDefaults instantiates a new AclBindingListPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAclBindingListPageWithDefaults() *AclBindingListPage {
	this := AclBindingListPage{}





	return &this
}


// GetItems returns the Items field value
func (o *AclBindingListPage) GetItems() []AclBinding {
	if o == nil {
		var ret []AclBinding
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *AclBindingListPage) GetItemsOk() (*[]AclBinding, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *AclBindingListPage) SetItems(v []AclBinding) {
	o.Items = v
}


// GetTotal returns the Total field value
func (o *AclBindingListPage) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *AclBindingListPage) GetTotalOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *AclBindingListPage) SetTotal(v int32) {
	o.Total = v
}


// GetSize returns the Size field value if set, zero value otherwise.
func (o *AclBindingListPage) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclBindingListPage) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *AclBindingListPage) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *AclBindingListPage) SetSize(v int32) {
	o.Size = &v
}


// GetPage returns the Page field value if set, zero value otherwise.
func (o *AclBindingListPage) GetPage() int32 {
	if o == nil || o.Page == nil {
		var ret int32
		return ret
	}
	return *o.Page
}

// GetPageOk returns a tuple with the Page field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AclBindingListPage) GetPageOk() (*int32, bool) {
	if o == nil || o.Page == nil {
		return nil, false
	}
	return o.Page, true
}

// HasPage returns a boolean if a field has been set.
func (o *AclBindingListPage) HasPage() bool {
	if o != nil && o.Page != nil {
		return true
	}

	return false
}

// SetPage gets a reference to the given int32 and assigns it to the Page field.
func (o *AclBindingListPage) SetPage(v int32) {
	o.Page = &v
}


func (o AclBindingListPage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["items"] = o.Items
	}
    
	if true {
		toSerialize["total"] = o.Total
	}
    
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
    
	if o.Page != nil {
		toSerialize["page"] = o.Page
	}
    
	return json.Marshal(toSerialize)
}

type NullableAclBindingListPage struct {
	value *AclBindingListPage
	isSet bool
}

func (v NullableAclBindingListPage) Get() *AclBindingListPage {
	return v.value
}

func (v *NullableAclBindingListPage) Set(val *AclBindingListPage) {
	v.value = val
	v.isSet = true
}

func (v NullableAclBindingListPage) IsSet() bool {
	return v.isSet
}

func (v *NullableAclBindingListPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAclBindingListPage(val *AclBindingListPage) *NullableAclBindingListPage {
	return &NullableAclBindingListPage{value: val, isSet: true}
}

func (v NullableAclBindingListPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAclBindingListPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

