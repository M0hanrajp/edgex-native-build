package transforms

import (
	"fmt"

	"github.com/edgexfoundry/app-functions-sdk-go/v3/pkg/interfaces"
	"github.com/edgexfoundry/go-mod-core-contracts/v3/dtos"
)

// EventProcessor contains functions to process the Event DTO
type EventProcessor struct {
	// additionalTags []dtos.Event
}

//TODO: Implementation of Event tags needs to be done i.e. include []dtos.EventTag format

// NewEventProcessor creates a new EventProcessor with additional tags to add to the Events that are processed
func NewEventProcessor(additionalTags map[string]interface{}) (*EventProcessor, error) {
	ep := &EventProcessor{}
	// for name, value := range additionalTags {
	// 	if err := dtos.ValidateMetricName(name, "Tag"); err != nil {
	// 		return nil, err
	// 	}
	// 	ep.additionalTags = append(ep.additionalTags, dtos.MetricTag{Name: name, Value: fmt.Sprintf("%v", value)})
	// }
	return ep, nil
}

// ToEventLineProtocol transforms an Event DTO to a string conforming to Line Protocol syntax
// Similar to how ToLineProtocol works for Metrics
func (ep *EventProcessor) ToEventLineProtocol(ctx interfaces.AppFunctionContext, data interface{}) (bool, interface{}) {
	lc := ctx.LoggingClient()

	lc.Debugf("ToEventLineProtocol called in pipeline '%s'", ctx.PipelineId())
	if data == nil {
		// Go here for details on Error Handle: https://docs.edgexfoundry.org/1.3/microservices/application/ErrorHandling/
		return false, fmt.Errorf("function ToEventLineProtocol in pipeline '%s': No Data Received", ctx.PipelineId())
	}

	event, ok := data.(dtos.Event)
	// lc.Debugf("The value of event: %v", event)
	if !ok {
		return false, fmt.Errorf("function ToEventLineProtocol in pipeline '%s', type received is not an Event", ctx.PipelineId())
	}
	// Convert the event to Line Protocol format
	result := fmt.Sprintln(event.ToEventLineProtocol())
	// lc.Debugf("The value of event.ToEventLineProtocol: %v", event.ToEventLineProtocol())
	lc.Debugf("Transformed Event to '%s' in pipeline '%s'", result, ctx.PipelineId())

	return true, result
}
