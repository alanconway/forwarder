// +k8s:deepcopy-gen=package,register

// Package v1beta1 contains API Schema definitions for the logging API group
//
// The forwarder API is `scope: Namespaced`. The clusterforwarder API is
// identical but is `scope: Cluster`. The implementations for namespaced and
// cluster will be different but the user experience is the same.
//
// +groupName=logging.openshift.io
package v1beta1
