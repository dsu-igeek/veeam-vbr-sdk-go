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
	"time"
)

// GeneralOptionsEmailNotificationsModel Global email notification settings.
type GeneralOptionsEmailNotificationsModel struct {
	// If *true*, global email notification settings are enabled.
	IsEnabled bool `json:"isEnabled"`
	// Full DNS name or IP address of the SMTP server.
	SmtpServerName string `json:"smtpServerName"`
	AdvancedSmtpOptions AdvancedSmtpOptionsModel `json:"advancedSmtpOptions"`
	// Email address from which email notifications must be sent.
	From string `json:"from"`
	// Recipient email addresses. Use a semicolon to separate multiple addresses.
	To string `json:"to"`
	// Notification subject. Use the following variables in the subject:<ul> <li>%Time% — completion time</li> <li>%JobName% — job name</li> <li>%JobResult% — job result</li> <li>%ObjectCount% — number of VMs in the job</li> <li>%Issues% — number of VMs in the job that have been processed with the Warning or Failed status</li></ul>
	Subject string `json:"subject"`
	// Time when Veeam Backup & Replication sends daily email reports.
	SendDailyReportsAt time.Time `json:"sendDailyReportsAt"`
	// If *true*, email notifications are sent when the job completes successfully.
	NotifyOnSuccess bool `json:"notifyOnSuccess"`
	// If *true*, email notifications are sent when the job completes with a warning.
	NotifyOnWarning bool `json:"notifyOnWarning"`
	// If *true*, email notifications are sent when the job fails.
	NotifyOnFailure bool `json:"notifyOnFailure"`
	// If *true*, email notifications are sent about the final job status only (not per every job retry).
	NotifyOnLastRetry bool `json:"notifyOnLastRetry"`
}

// NewGeneralOptionsEmailNotificationsModel instantiates a new GeneralOptionsEmailNotificationsModel object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeneralOptionsEmailNotificationsModel(isEnabled bool, smtpServerName string, advancedSmtpOptions AdvancedSmtpOptionsModel, from string, to string, subject string, sendDailyReportsAt time.Time, notifyOnSuccess bool, notifyOnWarning bool, notifyOnFailure bool, notifyOnLastRetry bool, ) *GeneralOptionsEmailNotificationsModel {
	this := GeneralOptionsEmailNotificationsModel{}
	this.IsEnabled = isEnabled
	this.SmtpServerName = smtpServerName
	this.AdvancedSmtpOptions = advancedSmtpOptions
	this.From = from
	this.To = to
	this.Subject = subject
	this.SendDailyReportsAt = sendDailyReportsAt
	this.NotifyOnSuccess = notifyOnSuccess
	this.NotifyOnWarning = notifyOnWarning
	this.NotifyOnFailure = notifyOnFailure
	this.NotifyOnLastRetry = notifyOnLastRetry
	return &this
}

// NewGeneralOptionsEmailNotificationsModelWithDefaults instantiates a new GeneralOptionsEmailNotificationsModel object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeneralOptionsEmailNotificationsModelWithDefaults() *GeneralOptionsEmailNotificationsModel {
	this := GeneralOptionsEmailNotificationsModel{}
	return &this
}

// GetIsEnabled returns the IsEnabled field value
func (o *GeneralOptionsEmailNotificationsModel) GetIsEnabled() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.IsEnabled
}

// GetIsEnabledOk returns a tuple with the IsEnabled field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsEmailNotificationsModel) GetIsEnabledOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.IsEnabled, true
}

// SetIsEnabled sets field value
func (o *GeneralOptionsEmailNotificationsModel) SetIsEnabled(v bool) {
	o.IsEnabled = v
}

// GetSmtpServerName returns the SmtpServerName field value
func (o *GeneralOptionsEmailNotificationsModel) GetSmtpServerName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SmtpServerName
}

// GetSmtpServerNameOk returns a tuple with the SmtpServerName field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsEmailNotificationsModel) GetSmtpServerNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SmtpServerName, true
}

// SetSmtpServerName sets field value
func (o *GeneralOptionsEmailNotificationsModel) SetSmtpServerName(v string) {
	o.SmtpServerName = v
}

// GetAdvancedSmtpOptions returns the AdvancedSmtpOptions field value
func (o *GeneralOptionsEmailNotificationsModel) GetAdvancedSmtpOptions() AdvancedSmtpOptionsModel {
	if o == nil  {
		var ret AdvancedSmtpOptionsModel
		return ret
	}

	return o.AdvancedSmtpOptions
}

// GetAdvancedSmtpOptionsOk returns a tuple with the AdvancedSmtpOptions field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsEmailNotificationsModel) GetAdvancedSmtpOptionsOk() (*AdvancedSmtpOptionsModel, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.AdvancedSmtpOptions, true
}

// SetAdvancedSmtpOptions sets field value
func (o *GeneralOptionsEmailNotificationsModel) SetAdvancedSmtpOptions(v AdvancedSmtpOptionsModel) {
	o.AdvancedSmtpOptions = v
}

// GetFrom returns the From field value
func (o *GeneralOptionsEmailNotificationsModel) GetFrom() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsEmailNotificationsModel) GetFromOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *GeneralOptionsEmailNotificationsModel) SetFrom(v string) {
	o.From = v
}

// GetTo returns the To field value
func (o *GeneralOptionsEmailNotificationsModel) GetTo() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsEmailNotificationsModel) GetToOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *GeneralOptionsEmailNotificationsModel) SetTo(v string) {
	o.To = v
}

// GetSubject returns the Subject field value
func (o *GeneralOptionsEmailNotificationsModel) GetSubject() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsEmailNotificationsModel) GetSubjectOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *GeneralOptionsEmailNotificationsModel) SetSubject(v string) {
	o.Subject = v
}

// GetSendDailyReportsAt returns the SendDailyReportsAt field value
func (o *GeneralOptionsEmailNotificationsModel) GetSendDailyReportsAt() time.Time {
	if o == nil  {
		var ret time.Time
		return ret
	}

	return o.SendDailyReportsAt
}

// GetSendDailyReportsAtOk returns a tuple with the SendDailyReportsAt field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsEmailNotificationsModel) GetSendDailyReportsAtOk() (*time.Time, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SendDailyReportsAt, true
}

// SetSendDailyReportsAt sets field value
func (o *GeneralOptionsEmailNotificationsModel) SetSendDailyReportsAt(v time.Time) {
	o.SendDailyReportsAt = v
}

// GetNotifyOnSuccess returns the NotifyOnSuccess field value
func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnSuccess() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.NotifyOnSuccess
}

// GetNotifyOnSuccessOk returns a tuple with the NotifyOnSuccess field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnSuccessOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NotifyOnSuccess, true
}

// SetNotifyOnSuccess sets field value
func (o *GeneralOptionsEmailNotificationsModel) SetNotifyOnSuccess(v bool) {
	o.NotifyOnSuccess = v
}

// GetNotifyOnWarning returns the NotifyOnWarning field value
func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnWarning() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.NotifyOnWarning
}

// GetNotifyOnWarningOk returns a tuple with the NotifyOnWarning field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnWarningOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NotifyOnWarning, true
}

// SetNotifyOnWarning sets field value
func (o *GeneralOptionsEmailNotificationsModel) SetNotifyOnWarning(v bool) {
	o.NotifyOnWarning = v
}

// GetNotifyOnFailure returns the NotifyOnFailure field value
func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnFailure() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.NotifyOnFailure
}

// GetNotifyOnFailureOk returns a tuple with the NotifyOnFailure field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnFailureOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NotifyOnFailure, true
}

// SetNotifyOnFailure sets field value
func (o *GeneralOptionsEmailNotificationsModel) SetNotifyOnFailure(v bool) {
	o.NotifyOnFailure = v
}

// GetNotifyOnLastRetry returns the NotifyOnLastRetry field value
func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnLastRetry() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.NotifyOnLastRetry
}

// GetNotifyOnLastRetryOk returns a tuple with the NotifyOnLastRetry field value
// and a boolean to check if the value has been set.
func (o *GeneralOptionsEmailNotificationsModel) GetNotifyOnLastRetryOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NotifyOnLastRetry, true
}

// SetNotifyOnLastRetry sets field value
func (o *GeneralOptionsEmailNotificationsModel) SetNotifyOnLastRetry(v bool) {
	o.NotifyOnLastRetry = v
}

func (o GeneralOptionsEmailNotificationsModel) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isEnabled"] = o.IsEnabled
	}
	if true {
		toSerialize["smtpServerName"] = o.SmtpServerName
	}
	if true {
		toSerialize["advancedSmtpOptions"] = o.AdvancedSmtpOptions
	}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if true {
		toSerialize["sendDailyReportsAt"] = o.SendDailyReportsAt
	}
	if true {
		toSerialize["notifyOnSuccess"] = o.NotifyOnSuccess
	}
	if true {
		toSerialize["notifyOnWarning"] = o.NotifyOnWarning
	}
	if true {
		toSerialize["notifyOnFailure"] = o.NotifyOnFailure
	}
	if true {
		toSerialize["notifyOnLastRetry"] = o.NotifyOnLastRetry
	}
	return json.Marshal(toSerialize)
}

type NullableGeneralOptionsEmailNotificationsModel struct {
	value *GeneralOptionsEmailNotificationsModel
	isSet bool
}

func (v NullableGeneralOptionsEmailNotificationsModel) Get() *GeneralOptionsEmailNotificationsModel {
	return v.value
}

func (v *NullableGeneralOptionsEmailNotificationsModel) Set(val *GeneralOptionsEmailNotificationsModel) {
	v.value = val
	v.isSet = true
}

func (v NullableGeneralOptionsEmailNotificationsModel) IsSet() bool {
	return v.isSet
}

func (v *NullableGeneralOptionsEmailNotificationsModel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeneralOptionsEmailNotificationsModel(val *GeneralOptionsEmailNotificationsModel) *NullableGeneralOptionsEmailNotificationsModel {
	return &NullableGeneralOptionsEmailNotificationsModel{value: val, isSet: true}
}

func (v NullableGeneralOptionsEmailNotificationsModel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeneralOptionsEmailNotificationsModel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


