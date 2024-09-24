//
// Copyright (C) 2020-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package dtos

import (
	"encoding/xml"
	"fmt"
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

// Implementation of ToEventLineProtocol
func (e *Event) ToEventLineProtocol() string {
	// Initialize builders for tags and fields
	var tags strings.Builder
	var fields strings.Builder
	isFirst := true

	// Add deviceName as a tag
	// tags.WriteString(",deviceName=" + e.DeviceName)

	// Optionally add profileName as a tag
	tags.WriteString(",profileName=" + e.ProfileName)

	// Iterate over readings to build fields section
	for _, reading := range e.Readings {
		if isFirst {
			isFirst = false
		} else {
			fields.WriteString(",")
		}
		// Field key is the resourceName, field value is the reading value
		fields.WriteString(reading.ResourceName + "=" + formatLineProtocolEventValue(reading.Value))
	}

	// Build the final line protocol
	result := fmt.Sprintf("%s%s %s %d", e.DeviceName, tags.String(), fields.String(), e.Origin)
	return result
}

// Helper function
func formatLineProtocolEventValue(value interface{}) string {
	switch value.(type) {
	case string:
		return fmt.Sprintf("%s", value)
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%di", value)
	case uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%du", value)
	default:
		return fmt.Sprintf("%v", value)
	}
}
