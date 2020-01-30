package v1beta1

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Forwarder is the schema for log forwarder API.
type Forwarder struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ForwarderSpec   `json:"spec,omitempty"`
	Status ForwarderStatus `json:"status,omitempty"`
}

// ForwarderSpec defines the desired state of Forwarder
type ForwarderSpec struct {
	// Outputs are named output destinations for log messages.
	Outputs []Output `json:"outputs,omitempty"`

	// Pipelines select log messages to send to outputs.
	Pipelines []Pipeline `json:"pipelines,omitempty"`
}

// Inputs defines the observed state of Forwarder
type ForwarderStatus struct {
	// FIXME(alanconway) TODO
}

// Output defines a destination for log messages.
// Most fields are optional, the simplest working configurations need
// only name, type and url.
type Output struct {
	// Name used to refer to the output from a pipeline.
	Name string `json:"name,required"`

	// Type of output plugin needed to communicate with the desired target.
	//
	// XXX(alanconway) how to document the list of available types.
	Type string `json:"type,required"`

	// URL to connect to the destination, valid schemes depend on the type.  `url`
	// may include 'username:password', but the password is visible to anyone who
	// can view the spec.  For more detail see `authSecret`.
	URL string `json:"url,required"`

	// AuthSecret is the name of a secret with keys `username` and/or `password`
	// for outputs that have password or shared-key authentication.
	// The secret is in the same namespace as this forwarder.
	//
	// If present, the secret username and password are used in preference to
	// the 'username:password' from the `url`. Outputs with shared-key authentication
	// use just the `password` key and ignore `username`.
	//
	// +optional
	AuthSecretRef string `json:"authSecretRef,omitempty"`

	// TLSSecret is the name of a secret with keys `tls.crt`, `tls.key` and `ca.crt`.
	// If present it enables TLS client authentication.
	// It is an error to specify this field if the output does not support TLS.
	// The secret is in the same namespace as this forwarder.
	//
	// +optional
	TLSSecret string `json:"tlsSecretRef,omitempty"`

	// Insecure must be explicitly set to 'true' if no encryption is configured.
	//
	// Encryption can be configured via the `url` for server-only authentication
	// (e.g. 'https://...', 'syslog+tls://...') and/or the `tlsSecret` field for
	// client authentication.
	//
	// +optional
	Insecure bool `json:"insecure,omitempty"`

	// Reconnect defines how to reconnect if a connection closes.
	//
	// +optional
	Reconnect *Reconnect `json:"reconnect,omitempty"`

	// Resend if true un-acknowledged data will be re-sent after reconnect, so
	// data may be duplicated. If false (the default) there is no re-send so data
	// may be lost.
	//
	// +optional
	Resend bool `json:"resend,omitempty"`

	// TimeoutMilliseconds is the max time to wait for a connection to be
	// established, or for any other "waiting" that is done by the plugin; for
	// example waiting for an acknowledgement. 0 or unspecified means use the
	// default timeout determined by the plugin.
	//
	// +optional
	TimeoutMilliseconds int64 `json:"timeoutSeconds,omitempty"`

	// Plugin provides configuration specific to the output plugin `type`.
	// You should not need to use this in most cases.
	//
	// XXX(alanconway) document and validate type-specifics, // +docLink
	// to documentation. Should this be an explicit union?
	//
	// +optional
	Plugin *json.RawMessage `json:"plugin,omitempty"`

	// MasterNamespace should not normally be set, use the default.
	// The 'forwarder' object in this namespace is the master forwarder.
	//
	// +kubebuilder:Validation:Default="openshift-logging"
	// +optional
	MasterNamespace string `json:"masterNamespace,omitempty"`
}

type Reconnect struct {
	// XXX(alanconway) review reconnect in other k8s APIs and follow conventions.
	// Need something like: minDelayMilliseconds, maxDelaySeconds, maxRetryTime
	// for exponential backoff.
}

// Pipeline selects log messages and directs them to outputs.
type Pipeline struct {

	// OutputRefs lists the names of outputs for selected log messages.
	//
	// XXX(alanconway) is it correct to use Output_Refs_ convention here, since
	// these are not k8s object references? Should we use OutputNames or Outputs?
	OutputRefs []string `json:"outputrefs,required"`

	// Name of the pipeline, for patch updates.
	//
	// +optional
	Name string `json:"name, omitempty"`

	// Selector restricts the containers used as log sources.
	//
	// +optional
	Selector *Selector `json:"selector,omitempty"`

	// Balance determines how to select among multiple outputs: choices are
	// 'Random', 'RoundRobin' or 'FanOut' to send to all outputs concurrently.
	//
	// +kubebuilder:Validation:Default="Random"
	// +optional
	Balance Balance `json:"balance,omitempty"`
}

// +kubebuilder:validation:Enum=Random;RoundRobin;All

type Balance string

const (
	BalanceRandom     = "Random"
	BalanceRoundRobin = "RoundRobin"
	BalanceFanOut     = "FanOut"
)

// Selector restricts containers to use as log sources.
// A container must match all criteria to be included.
//
type Selector struct {
	// Labels is an equality-based label selector.
	//
	// +optional
	Labels map[string]string `json:"labels,omitempty"`

	// Expressions is a list of label selector expressions.
	//
	// +optional
	Expressions []string `json:"expression,omitempty"`

	// Source restricts forwarding to message by type of log source.
	// If omitted or empty, forward all source types.
	//
	// +optional
	Source []SourceType `json:"source,omitempty"`

	// Namespaces restricts forwarding to the listed namespaces.
	//
	// Only allowed for the master forwarder, namespace forwarders
	// forward only from the namespace they are deployed in.
	//
	// +optional
	Namespaces []string `json:"namespaces,omitempty"`
}

// +kubebuilder:validation:Enum=Application;Infrastructure;Audit

// SourceType of log message.
type SourceType string

const (
	Application    SourceType = "Application"
	Infrastructure            = "Infrastructure"
	Audit                     = "Audit"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ForwarderList is a list of Forwarders
type ForwarderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Forwarder `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Forwarder{}, &ForwarderList{})
}
