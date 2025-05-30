// Package mgmtapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package mgmtapi

// Defines values for LinkRelationship.
const (
	CHILD  LinkRelationship = "CHILD"
	CORE   LinkRelationship = "CORE"
	PARENT LinkRelationship = "PARENT"
	PEER   LinkRelationship = "PEER"
)

// Defines values for LinkState.
const (
	DOWN LinkState = "DOWN"
	UP   LinkState = "UP"
)

// Defines values for LogLevelLevel.
const (
	Debug LogLevelLevel = "debug"
	Error LogLevelLevel = "error"
	Info  LogLevelLevel = "info"
)

// BFD defines model for BFD.
type BFD struct {
	// DesiredMinimumTxInterval The minimum interval between transmission of BFD control packets that the operator desires. This value is advertised to the peer, however the actual interval used is specified by taking the maximum of desired-minimum-tx-interval and the value of the remote required-minimum-receive interval value.
	DesiredMinimumTxInterval string `json:"desired_minimum_tx_interval"`

	// DetectionMultiplier The number of packets that must be missed to declare this session as down. The detection interval for the BFD session is calculated by multiplying the value of the negotiated transmission interval by this value.
	DetectionMultiplier int `json:"detection_multiplier"`

	// Enabled Indication of whether BFD is enabled and configured on this interface.
	Enabled bool `json:"enabled"`

	// RequiredMinimumReceive The minimum interval between received BFD control packets that this system should support. This value is advertised to the remote peer to indicate the maximum frequency (i.e., minimum inter-packet interval) between BFD control packets that is acceptable to the local system.
	RequiredMinimumReceive string `json:"required_minimum_receive"`
}

// Interface defines model for Interface.
type Interface struct {
	Bfd BFD `json:"bfd"`

	// InterfaceId SCION interface identifier.
	InterfaceId int `json:"interface_id"`

	// InternalInterface The address of internal SCION interface of the router.
	InternalInterface string            `json:"internal_interface"`
	Neighbor          InterfaceNeighbor `json:"neighbor"`
	Relationship      LinkRelationship  `json:"relationship"`

	// ScionMtu The maximum transmission unit in bytes for SCION packets. This represents the protocol data unit (PDU) of the SCION layer and is usually calculated as maximum Ethernet payload - IP Header - UDP Header.
	ScionMtu ScionMTU  `json:"scion_mtu"`
	State    LinkState `json:"state"`
}

// InterfaceNeighbor defines model for InterfaceNeighbor.
type InterfaceNeighbor struct {
	// Address UDP/IP underlay address of the SCION Interface.
	Address string `json:"address"`
	IsdAs   IsdAs  `json:"isd_as"`
}

// InterfacesResponse defines model for InterfacesResponse.
type InterfacesResponse struct {
	Interfaces        *[]Interface        `json:"interfaces,omitempty"`
	SiblingInterfaces *[]SiblingInterface `json:"sibling_interfaces,omitempty"`
}

// IsdAs defines model for IsdAs.
type IsdAs = string

// LinkRelationship defines model for LinkRelationship.
type LinkRelationship string

// LinkState defines model for LinkState.
type LinkState string

// LogLevel defines model for LogLevel.
type LogLevel struct {
	// Level Logging level
	Level LogLevelLevel `json:"level"`
}

// LogLevelLevel Logging level
type LogLevelLevel string

// Problem defines model for Problem.
type Problem struct {
	// Detail A human readable explanation specific to this occurrence of the problem that is helpful to locate the problem and give advice on how to proceed. Written in English and readable for engineers, usually not suited for non technical stakeholders and not localized.
	Detail *string `json:"detail,omitempty"`

	// Instance A URI reference that identifies the specific occurrence of the problem, e.g. by adding a fragment identifier or sub-path to the problem type. May be used to locate the root of this problem in the source code.
	Instance *string `json:"instance,omitempty"`

	// Status The HTTP status code generated by the origin server for this occurrence of the problem.
	Status int `json:"status"`

	// Title A short summary of the problem type. Written in English and readable for engineers, usually not suited for non technical stakeholders and not localized.
	Title string `json:"title"`

	// Type A URI reference that uniquely identifies the problem type only in the context of the provided API. Opposed to the specification in RFC-7807, it is neither recommended to be dereferencable and point to a human-readable documentation nor globally unique for the problem type.
	Type *string `json:"type,omitempty"`
}

// ScionMTU The maximum transmission unit in bytes for SCION packets. This represents the protocol data unit (PDU) of the SCION layer and is usually calculated as maximum Ethernet payload - IP Header - UDP Header.
type ScionMTU = int

// SiblingInterface defines model for SiblingInterface.
type SiblingInterface struct {
	// InterfaceId SCION interface identifier.
	InterfaceId int `json:"interface_id"`

	// InternalInterface Internal address of the sibling router.
	InternalInterface string           `json:"internal_interface"`
	Neighbor          SiblingNeighbor  `json:"neighbor"`
	Relationship      LinkRelationship `json:"relationship"`

	// ScionMtu The maximum transmission unit in bytes for SCION packets. This represents the protocol data unit (PDU) of the SCION layer and is usually calculated as maximum Ethernet payload - IP Header - UDP Header.
	ScionMtu ScionMTU  `json:"scion_mtu"`
	State    LinkState `json:"state"`
}

// SiblingNeighbor defines model for SiblingNeighbor.
type SiblingNeighbor struct {
	IsdAs IsdAs `json:"isd_as"`
}

// StandardError defines model for StandardError.
type StandardError struct {
	// Error Error message
	Error string `json:"error"`
}

// BadRequest defines model for BadRequest.
type BadRequest = StandardError

// SetLogLevelJSONRequestBody defines body for SetLogLevel for application/json ContentType.
type SetLogLevelJSONRequestBody = LogLevel
