package v1

import "context"

//===========================================================================
// Service Interface
//===========================================================================

type FlightTrackingService interface {
	Calculate(context.Context, *CalculateRequest) (*CalculateReply, error)
}

// Reply is a generic api repsonse
type Reply struct {
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty" yaml:"error,omitempty"`
}

type CalculateRequest struct {
	Flights [][]string `json:"flights"`
}

// Calculate Reply contain the starting and ending airport
type CalculateReply struct {
	Route []string `json:"route"`
}
