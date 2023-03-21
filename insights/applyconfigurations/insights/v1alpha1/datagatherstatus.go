// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/openshift/api/insights/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DataGatherStatusApplyConfiguration represents an declarative configuration of the DataGatherStatus type for use
// with apply.
type DataGatherStatusApplyConfiguration struct {
	Conditions        []v1.Condition                      `json:"conditions,omitempty"`
	State             *v1alpha1.DataGatherState           `json:"dataGatherState,omitempty"`
	Gatherers         []GathererStatusApplyConfiguration  `json:"gatherers,omitempty"`
	StartTime         *v1.Time                            `json:"startTime,omitempty"`
	FinishTime        *v1.Time                            `json:"finishTime,omitempty"`
	RelatedObjects    []ObjectReferenceApplyConfiguration `json:"relatedObjects,omitempty"`
	InsightsRequestID *string                             `json:"insightsRequestID,omitempty"`
	InsightsReport    *InsightsReportApplyConfiguration   `json:"insightsReport,omitempty"`
}

// DataGatherStatusApplyConfiguration constructs an declarative configuration of the DataGatherStatus type for use with
// apply.
func DataGatherStatus() *DataGatherStatusApplyConfiguration {
	return &DataGatherStatusApplyConfiguration{}
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *DataGatherStatusApplyConfiguration) WithConditions(values ...v1.Condition) *DataGatherStatusApplyConfiguration {
	for i := range values {
		b.Conditions = append(b.Conditions, values[i])
	}
	return b
}

// WithState sets the State field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the State field is set to the value of the last call.
func (b *DataGatherStatusApplyConfiguration) WithState(value v1alpha1.DataGatherState) *DataGatherStatusApplyConfiguration {
	b.State = &value
	return b
}

// WithGatherers adds the given value to the Gatherers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Gatherers field.
func (b *DataGatherStatusApplyConfiguration) WithGatherers(values ...*GathererStatusApplyConfiguration) *DataGatherStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithGatherers")
		}
		b.Gatherers = append(b.Gatherers, *values[i])
	}
	return b
}

// WithStartTime sets the StartTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartTime field is set to the value of the last call.
func (b *DataGatherStatusApplyConfiguration) WithStartTime(value v1.Time) *DataGatherStatusApplyConfiguration {
	b.StartTime = &value
	return b
}

// WithFinishTime sets the FinishTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the FinishTime field is set to the value of the last call.
func (b *DataGatherStatusApplyConfiguration) WithFinishTime(value v1.Time) *DataGatherStatusApplyConfiguration {
	b.FinishTime = &value
	return b
}

// WithRelatedObjects adds the given value to the RelatedObjects field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the RelatedObjects field.
func (b *DataGatherStatusApplyConfiguration) WithRelatedObjects(values ...*ObjectReferenceApplyConfiguration) *DataGatherStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithRelatedObjects")
		}
		b.RelatedObjects = append(b.RelatedObjects, *values[i])
	}
	return b
}

// WithInsightsRequestID sets the InsightsRequestID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InsightsRequestID field is set to the value of the last call.
func (b *DataGatherStatusApplyConfiguration) WithInsightsRequestID(value string) *DataGatherStatusApplyConfiguration {
	b.InsightsRequestID = &value
	return b
}

// WithInsightsReport sets the InsightsReport field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the InsightsReport field is set to the value of the last call.
func (b *DataGatherStatusApplyConfiguration) WithInsightsReport(value *InsightsReportApplyConfiguration) *DataGatherStatusApplyConfiguration {
	b.InsightsReport = value
	return b
}