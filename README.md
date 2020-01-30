# Proposal for log forwarding API.

NOTE: XXX(alanconway) comments mark points that need review & completion.

API is written in Go and generated as  [YAML CRDs](deploy/crds).  You can browse the Go source here or as [Godoc](https://godoc.org/github.com/alanconway/forwarder/pkg/apis/logging/v1beta1)

There is one specialized config defined as an example for
[syslog](pkg/apis/logging/v1beta1/output/syslog.go). Others may be needed for other
targets.

## Design rationale

Goals:

1. No API dependency on fluentd (our implementation can use it but won't be tied to it)
2. Support for an open set of logging destination types (plugins)
3. Users don't have to learn new config for each forwarding type (in typical use cases)

`forwarder.output` defines a generic configuration that should be sufficient
for *any* output plugin in typical use cases. Most fields are optional, the
simplest case only needs `name`, `type` and `url`.

The `output.plugin` field is an optional unspecified JSON object for special
plugin-specific config. It *MUST NOT* be required for typical use-cases, but
provides an escape hatch for special cases.

Our initial 'plugins' can just be configuration facades for fluentd plugins, but
future plugins need not be restricted to fluentd.  

The Syslog plugin config is at 

##  K8s API conventions

Trying to follow the [k8s API conventions](https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#units)

In particular, changed the API name from "logforwarding" to "forwarder".  The
API name should be a noun, and it's already in the 'logging' package so the
"log" prefix is redundant.

Other relevant items from the conventions doc:

- Enum string values are CapitalMixedCase.
- Optional values have +option tag and are pointer values.
- In comments, use `` when mentioning fields, '' for string literals.
- Avoid abbreviations.
