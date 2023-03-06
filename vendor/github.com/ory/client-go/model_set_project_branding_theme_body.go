/*
Ory APIs

Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

API version: v1.1.21
Contact: support@ory.sh
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// SetProjectBrandingThemeBody Set Project Branding Theme Request Parameters
type SetProjectBrandingThemeBody struct {
	// Logo type
	LogoType *string `json:"logo_type,omitempty"`
	// Logo URL
	LogoUrl *string `json:"logo_url,omitempty"`
	// Branding name
	Name *string `json:"name,omitempty"`
	Theme *ProjectBrandingColors `json:"theme,omitempty"`
}

// NewSetProjectBrandingThemeBody instantiates a new SetProjectBrandingThemeBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetProjectBrandingThemeBody() *SetProjectBrandingThemeBody {
	this := SetProjectBrandingThemeBody{}
	return &this
}

// NewSetProjectBrandingThemeBodyWithDefaults instantiates a new SetProjectBrandingThemeBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetProjectBrandingThemeBodyWithDefaults() *SetProjectBrandingThemeBody {
	this := SetProjectBrandingThemeBody{}
	return &this
}

// GetLogoType returns the LogoType field value if set, zero value otherwise.
func (o *SetProjectBrandingThemeBody) GetLogoType() string {
	if o == nil || o.LogoType == nil {
		var ret string
		return ret
	}
	return *o.LogoType
}

// GetLogoTypeOk returns a tuple with the LogoType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetProjectBrandingThemeBody) GetLogoTypeOk() (*string, bool) {
	if o == nil || o.LogoType == nil {
		return nil, false
	}
	return o.LogoType, true
}

// HasLogoType returns a boolean if a field has been set.
func (o *SetProjectBrandingThemeBody) HasLogoType() bool {
	if o != nil && o.LogoType != nil {
		return true
	}

	return false
}

// SetLogoType gets a reference to the given string and assigns it to the LogoType field.
func (o *SetProjectBrandingThemeBody) SetLogoType(v string) {
	o.LogoType = &v
}

// GetLogoUrl returns the LogoUrl field value if set, zero value otherwise.
func (o *SetProjectBrandingThemeBody) GetLogoUrl() string {
	if o == nil || o.LogoUrl == nil {
		var ret string
		return ret
	}
	return *o.LogoUrl
}

// GetLogoUrlOk returns a tuple with the LogoUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetProjectBrandingThemeBody) GetLogoUrlOk() (*string, bool) {
	if o == nil || o.LogoUrl == nil {
		return nil, false
	}
	return o.LogoUrl, true
}

// HasLogoUrl returns a boolean if a field has been set.
func (o *SetProjectBrandingThemeBody) HasLogoUrl() bool {
	if o != nil && o.LogoUrl != nil {
		return true
	}

	return false
}

// SetLogoUrl gets a reference to the given string and assigns it to the LogoUrl field.
func (o *SetProjectBrandingThemeBody) SetLogoUrl(v string) {
	o.LogoUrl = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SetProjectBrandingThemeBody) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetProjectBrandingThemeBody) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SetProjectBrandingThemeBody) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SetProjectBrandingThemeBody) SetName(v string) {
	o.Name = &v
}

// GetTheme returns the Theme field value if set, zero value otherwise.
func (o *SetProjectBrandingThemeBody) GetTheme() ProjectBrandingColors {
	if o == nil || o.Theme == nil {
		var ret ProjectBrandingColors
		return ret
	}
	return *o.Theme
}

// GetThemeOk returns a tuple with the Theme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SetProjectBrandingThemeBody) GetThemeOk() (*ProjectBrandingColors, bool) {
	if o == nil || o.Theme == nil {
		return nil, false
	}
	return o.Theme, true
}

// HasTheme returns a boolean if a field has been set.
func (o *SetProjectBrandingThemeBody) HasTheme() bool {
	if o != nil && o.Theme != nil {
		return true
	}

	return false
}

// SetTheme gets a reference to the given ProjectBrandingColors and assigns it to the Theme field.
func (o *SetProjectBrandingThemeBody) SetTheme(v ProjectBrandingColors) {
	o.Theme = &v
}

func (o SetProjectBrandingThemeBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LogoType != nil {
		toSerialize["logo_type"] = o.LogoType
	}
	if o.LogoUrl != nil {
		toSerialize["logo_url"] = o.LogoUrl
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Theme != nil {
		toSerialize["theme"] = o.Theme
	}
	return json.Marshal(toSerialize)
}

type NullableSetProjectBrandingThemeBody struct {
	value *SetProjectBrandingThemeBody
	isSet bool
}

func (v NullableSetProjectBrandingThemeBody) Get() *SetProjectBrandingThemeBody {
	return v.value
}

func (v *NullableSetProjectBrandingThemeBody) Set(val *SetProjectBrandingThemeBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSetProjectBrandingThemeBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSetProjectBrandingThemeBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetProjectBrandingThemeBody(val *SetProjectBrandingThemeBody) *NullableSetProjectBrandingThemeBody {
	return &NullableSetProjectBrandingThemeBody{value: val, isSet: true}
}

func (v NullableSetProjectBrandingThemeBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetProjectBrandingThemeBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


