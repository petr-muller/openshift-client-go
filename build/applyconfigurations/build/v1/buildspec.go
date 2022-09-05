// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	buildv1 "github.com/openshift/api/build/v1"
	corev1 "k8s.io/api/core/v1"
)

// BuildSpecApplyConfiguration represents an declarative configuration of the BuildSpec type for use
// with apply.
type BuildSpecApplyConfiguration struct {
	CommonSpecApplyConfiguration `json:",inline"`
	TriggeredBy                  []BuildTriggerCauseApplyConfiguration `json:"triggeredBy,omitempty"`
}

// BuildSpecApplyConfiguration constructs an declarative configuration of the BuildSpec type for use with
// apply.
func BuildSpec() *BuildSpecApplyConfiguration {
	return &BuildSpecApplyConfiguration{}
}

// WithServiceAccount sets the ServiceAccount field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ServiceAccount field is set to the value of the last call.
func (b *BuildSpecApplyConfiguration) WithServiceAccount(value string) *BuildSpecApplyConfiguration {
	b.ServiceAccount = &value
	return b
}

// WithSource sets the Source field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Source field is set to the value of the last call.
func (b *BuildSpecApplyConfiguration) WithSource(value *BuildSourceApplyConfiguration) *BuildSpecApplyConfiguration {
	b.Source = value
	return b
}

// WithRevision sets the Revision field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Revision field is set to the value of the last call.
func (b *BuildSpecApplyConfiguration) WithRevision(value *SourceRevisionApplyConfiguration) *BuildSpecApplyConfiguration {
	b.Revision = value
	return b
}

// WithStrategy sets the Strategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Strategy field is set to the value of the last call.
func (b *BuildSpecApplyConfiguration) WithStrategy(value *BuildStrategyApplyConfiguration) *BuildSpecApplyConfiguration {
	b.Strategy = value
	return b
}

// WithOutput sets the Output field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Output field is set to the value of the last call.
func (b *BuildSpecApplyConfiguration) WithOutput(value *BuildOutputApplyConfiguration) *BuildSpecApplyConfiguration {
	b.Output = value
	return b
}

// WithResources sets the Resources field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Resources field is set to the value of the last call.
func (b *BuildSpecApplyConfiguration) WithResources(value corev1.ResourceRequirements) *BuildSpecApplyConfiguration {
	b.Resources = &value
	return b
}

// WithPostCommit sets the PostCommit field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PostCommit field is set to the value of the last call.
func (b *BuildSpecApplyConfiguration) WithPostCommit(value *BuildPostCommitSpecApplyConfiguration) *BuildSpecApplyConfiguration {
	b.PostCommit = value
	return b
}

// WithCompletionDeadlineSeconds sets the CompletionDeadlineSeconds field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the CompletionDeadlineSeconds field is set to the value of the last call.
func (b *BuildSpecApplyConfiguration) WithCompletionDeadlineSeconds(value int64) *BuildSpecApplyConfiguration {
	b.CompletionDeadlineSeconds = &value
	return b
}

// WithNodeSelector sets the NodeSelector field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the NodeSelector field is set to the value of the last call.
func (b *BuildSpecApplyConfiguration) WithNodeSelector(value buildv1.OptionalNodeSelector) *BuildSpecApplyConfiguration {
	b.NodeSelector = &value
	return b
}

// WithMountTrustedCA sets the MountTrustedCA field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the MountTrustedCA field is set to the value of the last call.
func (b *BuildSpecApplyConfiguration) WithMountTrustedCA(value bool) *BuildSpecApplyConfiguration {
	b.MountTrustedCA = &value
	return b
}

// WithTriggeredBy adds the given value to the TriggeredBy field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the TriggeredBy field.
func (b *BuildSpecApplyConfiguration) WithTriggeredBy(values ...*BuildTriggerCauseApplyConfiguration) *BuildSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithTriggeredBy")
		}
		b.TriggeredBy = append(b.TriggeredBy, *values[i])
	}
	return b
}