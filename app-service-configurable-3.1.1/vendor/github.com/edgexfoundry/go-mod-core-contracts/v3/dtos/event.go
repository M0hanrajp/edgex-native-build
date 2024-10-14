//
// Copyright (C) 2020-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/models"

	"github.com/google/uuid"
)

type Event struct {
	common.Versionable `json:",inline"`
	Id                 string        `json:"id" validate:"required,uuid"`
	DeviceName         string        `json:"deviceName" validate:"required,edgex-dto-none-empty-string"`
	ProfileName        string        `json:"profileName" validate:"required,edgex-dto-none-empty-string"`
	SourceName         string        `json:"sourceName" validate:"required"`
	Origin             int64         `json:"origin" validate:"required"`
	Readings           []BaseReading `json:"readings" validate:"gt=0,dive,required"`
	Tags               Tags          `json:"tags,omitempty"`
}

// NewEvent creates and returns an initialized Event with no Readings
func NewEvent(profileName, deviceName, sourceName string) Event {
	return Event{
		Versionable: common.NewVersionable(),
		Id:          uuid.NewString(),
		DeviceName:  deviceName,
		ProfileName: profileName,
		SourceName:  sourceName,
		Origin:      time.Now().UnixNano(),
	}
}

// FromEventModelToDTO transforms the Event Model to the Event DTO
func FromEventModelToDTO(event models.Event) Event {
	var readings []BaseReading
	for _, reading := range event.Readings {
		readings = append(readings, FromReadingModelToDTO(reading))
	}

	tags := make(map[string]interface{})
	for tag, value := range event.Tags {
		tags[tag] = value
	}

	return Event{
		Versionable: common.NewVersionable(),
		Id:          event.Id,
		DeviceName:  event.DeviceName,
		ProfileName: event.ProfileName,
		SourceName:  event.SourceName,
		Origin:      event.Origin,
		Readings:    readings,
		Tags:        tags,
	}
}

// AddSimpleReading adds a simple reading to the Event
func (e *Event) AddSimpleReading(resourceName string, valueType string, value interface{}) error {
	reading, err := NewSimpleReading(e.ProfileName, e.DeviceName, resourceName, valueType, value)
	if err != nil {
		return err
	}
	e.Readings = append(e.Readings, reading)
	return nil
}

// AddBinaryReading adds a binary reading to the Event
func (e *Event) AddBinaryReading(resourceName string, binaryValue []byte, mediaType string) {
	e.Readings = append(e.Readings, NewBinaryReading(e.ProfileName, e.DeviceName, resourceName, binaryValue, mediaType))
}

// AddObjectReading adds a object reading to the Event
func (e *Event) AddObjectReading(resourceName string, objectValue interface{}) {
	e.Readings = append(e.Readings, NewObjectReading(e.ProfileName, e.DeviceName, resourceName, objectValue))
}

// ToXML provides a XML representation of the Event as a string
func (e *Event) ToXML() (string, error) {
	eventXml, err := xml.Marshal(e)
	if err != nil {
		return "", err
	}

	return string(eventXml), nil
}

// The following code contains Implementation of ToEventLineProtocol
// the code is written with reference of ToLineProtocol for converting metric data to Line protocol format
// the below code is written with the same apporach but to take in an Event DTO and convert to line protocol format
// Line Protocol Syntax:
// <measurement>[,<tag_key>=<tag_value>[,<tag_key>=<tag_value>]] <field_key>=<field_value>[,<field_key>=<field_value>] [<timestamp>]
// Examples:
// myMeasurement,tag1=value1,tag2=value2 fieldKey="fieldValue" 1556813561098000000
// More info related to line protocol https://docs.influxdata.com/influxdb/v2/reference/syntax/line-protocol/
func (event *Event) ToEventLineProtocol() string {
	// Initialize builders for tags and fields
	var tags strings.Builder
	var fields strings.Builder
	isFirst := true

	// Add deviceName as a tag
	// tags.WriteString(",deviceName=" + e.DeviceName)

	// Optionally add profileName as a tag
	tags.WriteString(",profileName=" + event.ProfileName)

	// Iterate over readings to build fields section
	for _, reading := range event.Readings {
		if isFirst {
			isFirst = false
		} else {
			fields.WriteString(",")
		}
		// Field key is the resourceName, field value is the reading value
		fields.WriteString(reading.ResourceName + "=" + formatLineProtocolEventValue(reading.ValueType, reading.Value))
	}
	// Build the final line protocol
	result := fmt.Sprintf("%s%s %s %d", event.DeviceName, tags.String(), fields.String(), event.Origin)
	return result
}

// Helper function
func formatLineProtocolEventValue(valueType string, value string) string {
	switch valueType {

	// Enclose strings in quotes
	case "String":
		return fmt.Sprintf("\"%s\"", value)

	// Integer data type, append i at the end
	case "Int8", "Int16", "Int32", "Int64", "Int":
		intValue, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return fmt.Sprintf("\"%s\"", value)
		}
		return fmt.Sprintf("%di", intValue)

	// unsigned-Integer data type, append U at the end
	case "Uint8", "Uint16", "Uint32", "Uint64", "Uint":
		uintValue, err := strconv.ParseUint(value, 10, 64)
		if err != nil {
			return fmt.Sprintf("\"%s\"", value)
		}
		return fmt.Sprintf("%du", uintValue)

	// Float data type
	case "Float32", "Float64":
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return fmt.Sprintf("\"%s\"", value)
		}
		return fmt.Sprintf("%f", floatValue)

	// Boolean data type
	case "Bool":
		boolValue, err := strconv.ParseBool(value)
		if err != nil {
			return fmt.Sprintf("\"%s\"", value)
		}
		return fmt.Sprintf("%t", boolValue)

	default:
		return fmt.Sprintf("\"Data type not supported!%s\"", value)
	}
}

// Old helper function where string does not work
//	func formatLineProtocolEventValue(value interface{}) string {
//		switch v := value.(type) {
//		case string:
//			return fmt.Sprintf("\"%s\"", v)
//		case int, int8, int16, int32, int64:
//			return fmt.Sprintf("%di", v) // integer values are suffixed with 'i'
//		case uint, uint8, uint16, uint32, uint64:
//			return fmt.Sprintf("%du", v) // unsigned integers
//		case float32, float64:
//			return fmt.Sprintf("%f", v) // floats do not need suffixes
//		case bool:
//			return fmt.Sprintf("%t", v) // booleans are stored as true/false
//		default:
//			return fmt.Sprintf("%v", v) // fallback case
//		}
//	}
