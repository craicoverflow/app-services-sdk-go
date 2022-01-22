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

// ConnectorRequestMeta struct for ConnectorRequestMeta
type ConnectorRequestMeta struct {

	Name string `json:"name"`

	ConnectorTypeId string `json:"connector_type_id"`

	Channel *Channel `json:"channel,omitempty"`

	DeploymentLocation DeploymentLocation `json:"deployment_location"`

	DesiredState ConnectorDesiredState `json:"desired_state"`

}

// NewConnectorRequestMeta instantiates a new ConnectorRequestMeta object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorRequestMeta(name string, connectorTypeId string, deploymentLocation DeploymentLocation, desiredState ConnectorDesiredState) *ConnectorRequestMeta {
	this := ConnectorRequestMeta{}
	this.Name = name
	this.ConnectorTypeId = connectorTypeId
	var channel Channel = CHANNEL_STABLE
	this.Channel = &channel
	this.DeploymentLocation = deploymentLocation
	this.DesiredState = desiredState
	return &this
}

// NewConnectorRequestMetaWithDefaults instantiates a new ConnectorRequestMeta object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorRequestMetaWithDefaults() *ConnectorRequestMeta {
	this := ConnectorRequestMeta{}



	var channel Channel = CHANNEL_STABLE
	this.Channel = &channel



	return &this
}


// GetName returns the Name field value
func (o *ConnectorRequestMeta) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConnectorRequestMeta) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ConnectorRequestMeta) SetName(v string) {
	o.Name = v
}


// GetConnectorTypeId returns the ConnectorTypeId field value
func (o *ConnectorRequestMeta) GetConnectorTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConnectorTypeId
}

// GetConnectorTypeIdOk returns a tuple with the ConnectorTypeId field value
// and a boolean to check if the value has been set.
func (o *ConnectorRequestMeta) GetConnectorTypeIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ConnectorTypeId, true
}

// SetConnectorTypeId sets field value
func (o *ConnectorRequestMeta) SetConnectorTypeId(v string) {
	o.ConnectorTypeId = v
}


// GetChannel returns the Channel field value if set, zero value otherwise.
func (o *ConnectorRequestMeta) GetChannel() Channel {
	if o == nil || o.Channel == nil {
		var ret Channel
		return ret
	}
	return *o.Channel
}

// GetChannelOk returns a tuple with the Channel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorRequestMeta) GetChannelOk() (*Channel, bool) {
	if o == nil || o.Channel == nil {
		return nil, false
	}
	return o.Channel, true
}

// HasChannel returns a boolean if a field has been set.
func (o *ConnectorRequestMeta) HasChannel() bool {
	if o != nil && o.Channel != nil {
		return true
	}

	return false
}

// SetChannel gets a reference to the given Channel and assigns it to the Channel field.
func (o *ConnectorRequestMeta) SetChannel(v Channel) {
	o.Channel = &v
}


// GetDeploymentLocation returns the DeploymentLocation field value
func (o *ConnectorRequestMeta) GetDeploymentLocation() DeploymentLocation {
	if o == nil {
		var ret DeploymentLocation
		return ret
	}

	return o.DeploymentLocation
}

// GetDeploymentLocationOk returns a tuple with the DeploymentLocation field value
// and a boolean to check if the value has been set.
func (o *ConnectorRequestMeta) GetDeploymentLocationOk() (*DeploymentLocation, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DeploymentLocation, true
}

// SetDeploymentLocation sets field value
func (o *ConnectorRequestMeta) SetDeploymentLocation(v DeploymentLocation) {
	o.DeploymentLocation = v
}


// GetDesiredState returns the DesiredState field value
func (o *ConnectorRequestMeta) GetDesiredState() ConnectorDesiredState {
	if o == nil {
		var ret ConnectorDesiredState
		return ret
	}

	return o.DesiredState
}

// GetDesiredStateOk returns a tuple with the DesiredState field value
// and a boolean to check if the value has been set.
func (o *ConnectorRequestMeta) GetDesiredStateOk() (*ConnectorDesiredState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DesiredState, true
}

// SetDesiredState sets field value
func (o *ConnectorRequestMeta) SetDesiredState(v ConnectorDesiredState) {
	o.DesiredState = v
}


func (o ConnectorRequestMeta) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	
	if true {
		toSerialize["name"] = o.Name
	}
    
	if true {
		toSerialize["connector_type_id"] = o.ConnectorTypeId
	}
    
	if o.Channel != nil {
		toSerialize["channel"] = o.Channel
	}
    
	if true {
		toSerialize["deployment_location"] = o.DeploymentLocation
	}
    
	if true {
		toSerialize["desired_state"] = o.DesiredState
	}
    
	return json.Marshal(toSerialize)
}

type NullableConnectorRequestMeta struct {
	value *ConnectorRequestMeta
	isSet bool
}

func (v NullableConnectorRequestMeta) Get() *ConnectorRequestMeta {
	return v.value
}

func (v *NullableConnectorRequestMeta) Set(val *ConnectorRequestMeta) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorRequestMeta) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorRequestMeta) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorRequestMeta(val *ConnectorRequestMeta) *NullableConnectorRequestMeta {
	return &NullableConnectorRequestMeta{value: val, isSet: true}
}

func (v NullableConnectorRequestMeta) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorRequestMeta) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

