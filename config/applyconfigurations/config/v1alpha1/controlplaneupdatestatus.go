// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	v1 "k8s.io/client-go/applyconfigurations/meta/v1"
)

// ControlPlaneUpdateStatusApplyConfiguration represents an declarative configuration of the ControlPlaneUpdateStatus type for use
// with apply.
type ControlPlaneUpdateStatusApplyConfiguration struct {
	ControlPlaneUpdateStatusSummaryApplyConfiguration `json:",inline"`
	Informers                                         []UpdateInformerApplyConfiguration `json:"informers,omitempty"`
	Conditions                                        []v1.ConditionApplyConfiguration   `json:"conditions,omitempty"`
}

// ControlPlaneUpdateStatusApplyConfiguration constructs an declarative configuration of the ControlPlaneUpdateStatus type for use with
// apply.
func ControlPlaneUpdateStatus() *ControlPlaneUpdateStatusApplyConfiguration {
	return &ControlPlaneUpdateStatusApplyConfiguration{}
}

// WithAssessment sets the Assessment field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Assessment field is set to the value of the last call.
func (b *ControlPlaneUpdateStatusApplyConfiguration) WithAssessment(value string) *ControlPlaneUpdateStatusApplyConfiguration {
	b.Assessment = &value
	return b
}

// WithVersions sets the Versions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Versions field is set to the value of the last call.
func (b *ControlPlaneUpdateStatusApplyConfiguration) WithVersions(value *ControlPlaneUpdateVersionsApplyConfiguration) *ControlPlaneUpdateStatusApplyConfiguration {
	b.Versions = value
	return b
}

// WithCompletion sets the Completion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Completion field is set to the value of the last call.
func (b *ControlPlaneUpdateStatusApplyConfiguration) WithCompletion(value int32) *ControlPlaneUpdateStatusApplyConfiguration {
	b.Completion = &value
	return b
}

// WithStartedAt sets the StartedAt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the StartedAt field is set to the value of the last call.
func (b *ControlPlaneUpdateStatusApplyConfiguration) WithStartedAt(value metav1.Time) *ControlPlaneUpdateStatusApplyConfiguration {
	b.StartedAt = &value
	return b
}

// WithCompletedAt sets the CompletedAt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CompletedAt field is set to the value of the last call.
func (b *ControlPlaneUpdateStatusApplyConfiguration) WithCompletedAt(value metav1.Time) *ControlPlaneUpdateStatusApplyConfiguration {
	b.CompletedAt = &value
	return b
}

// WithEstimatedCompletedAt sets the EstimatedCompletedAt field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EstimatedCompletedAt field is set to the value of the last call.
func (b *ControlPlaneUpdateStatusApplyConfiguration) WithEstimatedCompletedAt(value metav1.Time) *ControlPlaneUpdateStatusApplyConfiguration {
	b.EstimatedCompletedAt = &value
	return b
}

// WithInformers adds the given value to the Informers field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Informers field.
func (b *ControlPlaneUpdateStatusApplyConfiguration) WithInformers(values ...*UpdateInformerApplyConfiguration) *ControlPlaneUpdateStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithInformers")
		}
		b.Informers = append(b.Informers, *values[i])
	}
	return b
}

// WithConditions adds the given value to the Conditions field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Conditions field.
func (b *ControlPlaneUpdateStatusApplyConfiguration) WithConditions(values ...*v1.ConditionApplyConfiguration) *ControlPlaneUpdateStatusApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithConditions")
		}
		b.Conditions = append(b.Conditions, *values[i])
	}
	return b
}
