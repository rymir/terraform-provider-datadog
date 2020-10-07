/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// TeamType Team resource type.
type TeamType string

// List of TeamType
const (
	TEAMTYPE_TEAMS TeamType = "teams"
)

func (v *TeamType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TeamType(value)
	for _, existing := range []TeamType{"teams"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TeamType", value)
}

// Ptr returns reference to TeamType value
func (v TeamType) Ptr() *TeamType {
	return &v
}

type NullableTeamType struct {
	value *TeamType
	isSet bool
}

func (v NullableTeamType) Get() *TeamType {
	return v.value
}

func (v *NullableTeamType) Set(val *TeamType) {
	v.value = val
	v.isSet = true
}

func (v NullableTeamType) IsSet() bool {
	return v.isSet
}

func (v *NullableTeamType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTeamType(val *TeamType) *NullableTeamType {
	return &NullableTeamType{value: val, isSet: true}
}

func (v NullableTeamType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTeamType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}