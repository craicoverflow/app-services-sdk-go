/*
 * Connector Service Fleet Manager
 *
 * Connector Service Fleet Manager is a Rest API to manage connectors.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package connectormgmtclient

import (
	"encoding/json"
)

// ConnectorConfiguration struct for ConnectorConfiguration
type ConnectorConfiguration struct {

	Kafka KafkaConnectionSettings `json:"kafka"`

	ServiceAccount ServiceAccount `json:"service_account"`

	SchemaRegistry *SchemaRegistryConnectionSettings `json:"schema_registry,omitempty"`

	Connector map[string]interface{} `json:"connector"`

}

// NewConnectorConfiguration instantiates a new ConnectorConfiguration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorConfiguration(kafka KafkaConnectionSettings, serviceAccount ServiceAccount, connector map[string]interface{}) *ConnectorConfiguration {
	this := ConnectorConfiguration{}
	this.Kafka = kafka
	this.ServiceAccount = serviceAccount
	this.Connector = connector
	return &this
}

// NewConnectorConfigurationWithDefaults instantiates a new ConnectorConfiguration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorConfigurationWithDefaults() *ConnectorConfiguration {
	this := ConnectorConfiguration{}





	return &this
}


// GetKafka returns the Kafka field value
func (o *ConnectorConfiguration) GetKafka() KafkaConnectionSettings {
	if o == nil {
		var ret KafkaConnectionSettings
		return ret
	}

	return o.Kafka
}

// GetKafkaOk returns a tuple with the Kafka field value
// and a boolean to check if the value has been set.
func (o *ConnectorConfiguration) GetKafkaOk() (*KafkaConnectionSettings, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Kafka, true
}

// SetKafka sets field value
func (o *ConnectorConfiguration) SetKafka(v KafkaConnectionSettings) {
	o.Kafka = v
}


// GetServiceAccount returns the ServiceAccount field value
func (o *ConnectorConfiguration) GetServiceAccount() ServiceAccount {
	if o == nil {
		var ret ServiceAccount
		return ret
	}

	return o.ServiceAccount
}

// GetServiceAccountOk returns a tuple with the ServiceAccount field value
// and a boolean to check if the value has been set.
func (o *ConnectorConfiguration) GetServiceAccountOk() (*ServiceAccount, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ServiceAccount, true
}

// SetServiceAccount sets field value
func (o *ConnectorConfiguration) SetServiceAccount(v ServiceAccount) {
	o.ServiceAccount = v
}


// GetSchemaRegistry returns the SchemaRegistry field value if set, zero value otherwise.
func (o *ConnectorConfiguration) GetSchemaRegistry() SchemaRegistryConnectionSettings {
	if o == nil || o.SchemaRegistry == nil {
		var ret SchemaRegistryConnectionSettings
		return ret
	}
	return *o.SchemaRegistry
}

// GetSchemaRegistryOk returns a tuple with the SchemaRegistry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorConfiguration) GetSchemaRegistryOk() (*SchemaRegistryConnectionSettings, bool) {
	if o == nil || o.SchemaRegistry == nil {
		return nil, false
	}
	return o.SchemaRegistry, true
}

// HasSchemaRegistry returns a boolean if a field has been set.
func (o *ConnectorConfiguration) HasSchemaRegistry() bool {
	if o != nil && o.SchemaRegistry != nil {
		return true
	}

	return false
}

// SetSchemaRegistry gets a reference to the given SchemaRegistryConnectionSettings and assigns it to the SchemaRegistry field.
func (o *ConnectorConfiguration) SetSchemaRegistry(v SchemaRegistryConnectionSettings) {
	o.SchemaRegistry = &v
}


// GetConnector returns the Connector field value
func (o *ConnectorConfiguration) GetConnector() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value
// and a boolean to check if the value has been set.
func (o *ConnectorConfiguration) GetConnectorOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Connector, true
}

// SetConnector sets field value
func (o *ConnectorConfiguration) SetConnector(v map[string]interface{}) {
	o.Connector = v
}


func (o ConnectorConfiguration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["kafka"] = o.Kafka
	}
    
	if true {
		toSerialize["service_account"] = o.ServiceAccount
	}
    
	if o.SchemaRegistry != nil {
		toSerialize["schema_registry"] = o.SchemaRegistry
	}
    
	if true {
		toSerialize["connector"] = o.Connector
	}
    
	return json.Marshal(toSerialize)
}

type NullableConnectorConfiguration struct {
	value *ConnectorConfiguration
	isSet bool
}

func (v NullableConnectorConfiguration) Get() *ConnectorConfiguration {
	return v.value
}

func (v *NullableConnectorConfiguration) Set(val *ConnectorConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorConfiguration(val *ConnectorConfiguration) *NullableConnectorConfiguration {
	return &NullableConnectorConfiguration{value: val, isSet: true}
}

func (v NullableConnectorConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

