//
// Copyright (C) 2021-2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package messaging

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/fxamacker/cbor/v2"

	"github.com/edgexfoundry/edgex-go/internal/core/data/application"
	dataContainer "github.com/edgexfoundry/edgex-go/internal/core/data/container"
	"github.com/edgexfoundry/edgex-go/internal/pkg/utils"

	"github.com/edgexfoundry/go-mod-messaging/v3/pkg/types"

	"github.com/edgexfoundry/go-mod-bootstrap/v3/bootstrap/container"
	"github.com/edgexfoundry/go-mod-bootstrap/v3/di"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/common"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos/requests"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/errors"
)

// SubscribeEvents subscribes to events from message bus
func SubscribeEvents(ctx context.Context, dic *di.Container) errors.EdgeX {
	// Retrieve message bus configuration and logging client from the DI container.
	messageBusInfo := dataContainer.ConfigurationFrom(dic.Get).MessageBus
	lc := container.LoggingClientFrom(dic.Get)
	// Retrieve the message bus client from the DI container.
	messageBus := container.MessagingClientFrom(dic.Get)
	// Added this variable to check messageBus value
	lc.Infof("MessageBus: %+v", messageBus)
	// Create channels for receiving messages and errors.
	messages := make(chan types.MessageEnvelope)
	messageErrors := make(chan error)
	// Retrieve the core data application instance from the DI container.
	app := application.CoreDataAppFrom(dic.Get)
	// Build the subscription topic using the base topic prefix and core-data event subscribe topic.
	subscribeTopic := common.BuildTopic(messageBusInfo.GetBaseTopicPrefix(), common.CoreDataEventSubscribeTopic)
	// Added this variable to check subscribeTopic value
	lc.Infof("SubscribeTopic: %s", subscribeTopic)
	// Define the topics to subscribe to, with each topic associated with a channel for receiving messages.
	topics := []types.TopicChannel{
		{
			Topic:    subscribeTopic,
			Messages: messages,
		},
	}
	// Subscribe to the message bus topics.
	err := messageBus.Subscribe(topics, messageErrors)
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	// Start a goroutine to handle incoming messages and errors.
	go func() {
		for {
			select {
			case <-ctx.Done():
				lc.Infof("Exiting waiting for MessageBus '%s' topic messages", subscribeTopic)
				return
			case e := <-messageErrors:
				// Log any errors received from the message bus.
				lc.Error(e.Error())
			case msgEnvelope := <-messages:
				// Log the received message.
				lc.Debugf("Event received from MessageBus. Topic: %s, Correlation-id: %s", msgEnvelope.ReceivedTopic, msgEnvelope.CorrelationID)
				event := &requests.AddEventRequest{}
				// decoding the large payload may cause memory issues so checking before decoding
				// Check the payload size to avoid memory issues.
				maxEventSize := dataContainer.ConfigurationFrom(dic.Get).MaxEventSize
				edgeXerr := utils.CheckPayloadSize(msgEnvelope.Payload, maxEventSize*1024)
				if edgeXerr != nil {
					lc.Errorf("event size exceed MaxEventSize(%d KB)", maxEventSize)
					break
				}
				// Unmarshal the payload into an AddEventRequest object.
				err = unmarshalPayload(msgEnvelope, event)
				if err != nil {
					lc.Errorf("fail to unmarshal event, %v", err)
					break
				}
				// Validate the event against the message topic.
				err = validateEvent(msgEnvelope.ReceivedTopic, event.Event)
				if err != nil {
					lc.Error(err.Error())
					break
				}
				// Add the event to the core-data service.
				err = app.AddEvent(requests.AddEventReqToEventModel(*event), ctx, dic)
				if err != nil {
					lc.Errorf("fail to persist the event, %v", err)
				}
			}
		}
	}()

	return nil
}

// unmarshalPayload decodes the message payload based on its content type (JSON or CBOR).
func unmarshalPayload(envelope types.MessageEnvelope, target interface{}) error {
	var err error
	switch envelope.ContentType {
	case common.ContentTypeJSON:
		err = json.Unmarshal(envelope.Payload, target)

	case common.ContentTypeCBOR:
		err = cbor.Unmarshal(envelope.Payload, target)

	default:
		err = fmt.Errorf("unsupported content-type '%s' recieved", envelope.ContentType)
	}
	return err
}

// validateEvent validates the event fields against the message topic to ensure they match.
func validateEvent(messageTopic string, e dtos.Event) errors.EdgeX {
	// Parse messageTopic by the pattern `edgex/events/device/<device-service-name>/<device-profile-name>/<device-name>/<source-name>`
	fields := strings.Split(messageTopic, "/")

	// assumes a non-empty base topic with events/device/<device-service-name>/<device-profile-name>/<device-name>/<source-name>
	if len(fields) < 6 {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid message topic %s", messageTopic), nil)
	}

	len := len(fields)
	profileName, err := url.PathUnescape(fields[len-3])
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	deviceName, err := url.PathUnescape(fields[len-2])
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}
	sourceName, err := url.PathUnescape(fields[len-1])
	if err != nil {
		return errors.NewCommonEdgeXWrapper(err)
	}

	// Check whether the event fields match the message topic
	if e.ProfileName != profileName {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("event's profileName %s mismatches with the name %s received in topic", e.ProfileName, profileName), nil)
	}
	if e.DeviceName != deviceName {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("event's deviceName %s mismatches with the name %s received in topic", e.DeviceName, deviceName), nil)
	}
	if e.SourceName != sourceName {
		return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("event's sourceName %s mismatches with the name %s received in topic", e.SourceName, sourceName), nil)
	}
	return nil
}
