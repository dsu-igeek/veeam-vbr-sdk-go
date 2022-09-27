/*
 * Veeam Backup & Replication REST API
 *
 * This document lists paths (endpoints) of the Veeam Backup & Replication REST API and operations that you can perform by sending HTTP requests to the paths.<br>Requests can contain parameters in their path, query and header. POST and PUT requests can include a request body with resource payload. In response, you receive a conventional HTTP response code, HTTP response header and an optional response body schema that contains a result model.<br>Parameters, request bodies, and response bodies are defined inline or refer to schemas defined globally. Some schemas are polymorphic.
 *
 * API version: 1.1-rev0
 * Contact: support@veeam.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// EmailNotificationSettingsModel Email notification settings for the backup job.
type EmailNotificationSettingsModel struct {
	// If *true*, email notifications are enabled for this job.
	IsEnabled bool `json:"isEnabled"`
	// Array of recipient’s email addresses.
	Recipients *[]string `json:"recipients,omitempty"`
	NotificationType *EEmailNotificationType `json:"notificationType,omitempty"`
	CustomNotificationSettings *EmailCustomNotificationType `json:"customNotificationSettings,omitempty"`
}

// NewEmailNotificationSettingsModel instantiates a new EmailNotificationSettingsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailNotificationSettingsModel(isEnabled bool, ) *EmailNotificationSettingsModel {
	this := EmailNotificationSettingsModel{}
	this.IsEnabled = isEnabled
	return &this
}

// NewEmailNotificationSettingsModelWithDefaults instantiates a new EmailNotificationSettingsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailNotificationSettingsModelWithDefaults() *EmailNotificationSettingsModel {
	this := EmailNotificationSettingsModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *EmailNotificationSettingsModel) GetIsEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *EmailNotificationSettingsModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *EmailNotificationSettingsModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetRecipients returns the Recipients field value if set, zero value otherwise.
func (o *EmailNotificationSettingsModel) GetRecipients() []string {
	if o == nil || o.Recipients == nil {
		var ret []string
		return ret
	}
	return *o.Recipients
}

// GetRecipientsOk returns a tuple with the Recipients field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationSettingsModel) GetRecipientsOk() (*[]string, bool) {
	if o == nil || o.Recipients == nil {
		return nil, false
	}
	return o.Recipients, true
}

// HasRecipients returns a boolean if a field has been set.
func (o *EmailNotificationSettingsModel) HasRecipients() bool {
	if o != nil && o.Recipients != nil {
		return true
	}

	return false
}

// SetRecipients gets a reference to the given []string and assigns it to the Recipients field.
func (o *EmailNotificationSettingsModel) SetRecipients(v []string) {
	o.Recipients = &v
}

// GetNotificationType returns the NotificationType field value if set, zero value otherwise.
func (o *EmailNotificationSettingsModel) GetNotificationType() EEmailNotificationType {
	if o == nil || o.NotificationType == nil {
		var ret EEmailNotificationType
		return ret
	}
	return *o.NotificationType
}

// GetNotificationTypeOk returns a tuple with the NotificationType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationSettingsModel) GetNotificationTypeOk() (*EEmailNotificationType, bool) {
	if o == nil || o.NotificationType == nil {
		return nil, false
	}
	return o.NotificationType, true
}

// HasNotificationType returns a boolean if a field has been set.
func (o *EmailNotificationSettingsModel) HasNotificationType() bool {
	if o != nil && o.NotificationType != nil {
		return true
	}

	return false
}

// SetNotificationType gets a reference to the given EEmailNotificationType and assigns it to the NotificationType field.
func (o *EmailNotificationSettingsModel) SetNotificationType(v EEmailNotificationType) {
	o.NotificationType = &v
}

// GetCustomNotificationSettings returns the CustomNotificationSettings field value if set, zero value otherwise.
func (o *EmailNotificationSettingsModel) GetCustomNotificationSettings() EmailCustomNotificationType {
	if o == nil || o.CustomNotificationSettings == nil {
		var ret EmailCustomNotificationType
		return ret
	}
	return *o.CustomNotificationSettings
}

// GetCustomNotificationSettingsOk returns a tuple with the CustomNotificationSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailNotificationSettingsModel) GetCustomNotificationSettingsOk() (*EmailCustomNotificationType, bool) {
	if o == nil || o.CustomNotificationSettings == nil {
		return nil, false
	}
	return o.CustomNotificationSettings, true
}

// HasCustomNotificationSettings returns a boolean if a field has been set.
func (o *EmailNotificationSettingsModel) HasCustomNotificationSettings() bool {
	if o != nil && o.CustomNotificationSettings != nil {
		return true
	}

	return false
}

// SetCustomNotificationSettings gets a reference to the given EmailCustomNotificationType and assigns it to the CustomNotificationSettings field.
func (o *EmailNotificationSettingsModel) SetCustomNotificationSettings(v EmailCustomNotificationType) {
	o.CustomNotificationSettings = &v
}

func (o EmailNotificationSettingsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if o.Recipients != nil {
		toSerialize["recipients"] = o.Recipients
	}
	if o.NotificationType != nil {
		toSerialize["notificationType"] = o.NotificationType
	}
	if o.CustomNotificationSettings != nil {
		toSerialize["customNotificationSettings"] = o.CustomNotificationSettings
	}
	return json.Marshal(toSerialize)
}

type NullableEmailNotificationSettingsModel struct {
	value *EmailNotificationSettingsModel
	isSet bool
}

func (v NullableEmailNotificationSettingsModel) Get() *EmailNotificationSettingsModel {
	return v.value
}

func (v *NullableEmailNotificationSettingsModel) Set(val *EmailNotificationSettingsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailNotificationSettingsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailNotificationSettingsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailNotificationSettingsModel(val *EmailNotificationSettingsModel) *NullableEmailNotificationSettingsModel {
	return &NullableEmailNotificationSettingsModel{value: val, isSet: true}
}

func (v NullableEmailNotificationSettingsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailNotificationSettingsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


