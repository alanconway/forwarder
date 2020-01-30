package output

// XXX(alanconway)  this is copied from existing syslog plugin config,
// maybe needs review for GA?

// Syslog is the optional `plugin` field for `output.type` of 'syslog'.
//
// URL schemes are: syslog+udp:// syslog+tcp:// syslog+tls://
//
// If the URL scheme is 'syslog+tls' the `outpout.tls` field provides
// TLS configuration.
//
type Syslog struct {
	// Severity for outgoing syslog records. Special value 'COPY' means copy from incoming
	// log record.
	//
	// +optional
	Severity Severity `json:"severity,omitempty"`

	// Facility for outgoing syslog records. Special value 'COPY' means copy from incoming
	// log record.
	//
	// +optional
	Facility Facility `json:"severity,omitempty"`

	// TrimPrefix is a prefix to trim from the tag.
	//
	// +optional
	TrimPrefix string `json:"severity,omitempty"`

	// TagKey specifies a record field  to  use as as the tag on the syslog message.
	//
	// +optional
	TagKey string `json:"severity,omitempty"`

	// PayloadKey specifies a record field  to use as as the payload on the syslog message.
	PayloadKey string `json:"severity,omitempty"`
}

// XXX(alanconway) define enum for syslog severities
type Severity string

// XXX(alanconway) define enum for syslog facilities
type Facility string
