/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package rest

import (
	"github.com/hyperledger/aries-framework-go/cmd/aries-agent-mobile/pkg/wrappers/models"
	cmdintroduce "github.com/hyperledger/aries-framework-go/pkg/controller/command/introduce"
)

// IntroduceREST contains an http client and endpoints for each of its operations
type IntroduceREST struct {
	httpClient httpClient
	endpoints  map[string]*Endpoint

	URL   string
	Token string
}

// Actions returns unfinished actions for the async usage.
// This creates an http request based on the provided method arguments.
func (ir *IntroduceREST) Actions(request *models.RequestEnvelope) *models.ResponseEnvelope {
	respEnvelope := execREST(&restOperation{
		url:        ir.URL,
		token:      ir.Token,
		httpClient: ir.httpClient,
		endpoint:   *ir.endpoints[cmdintroduce.Actions],
		request:    request,
	})

	return respEnvelope
}

// SendProposal sends a proposal to the introducees (the client has not published an out-of-band message) via HTTP.
func (ir *IntroduceREST) SendProposal(request *models.RequestEnvelope) *models.ResponseEnvelope {
	respEnvelope := execREST(&restOperation{
		url:        ir.URL,
		token:      ir.Token,
		httpClient: ir.httpClient,
		endpoint:   *ir.endpoints[cmdintroduce.SendProposal],
		request:    request,
	})

	return respEnvelope
}

// SendProposalWithOOBRequest sends a proposal to the introducee
// (the client has published an out-of-band request) via HTTP.
func (ir *IntroduceREST) SendProposalWithOOBRequest(request *models.RequestEnvelope) *models.ResponseEnvelope {
	respEnvelope := execREST(&restOperation{
		url:        ir.URL,
		token:      ir.Token,
		httpClient: ir.httpClient,
		endpoint:   *ir.endpoints[cmdintroduce.SendProposalWithOOBRequest],
		request:    request,
	})

	return respEnvelope
}

// SendRequest sends a request showing that the introducee is willing to share their own out-of-band message (via HTTP).
func (ir *IntroduceREST) SendRequest(request *models.RequestEnvelope) *models.ResponseEnvelope {
	respEnvelope := execREST(&restOperation{
		url:        ir.URL,
		token:      ir.Token,
		httpClient: ir.httpClient,
		endpoint:   *ir.endpoints[cmdintroduce.SendRequest],
		request:    request,
	})

	return respEnvelope
}

// AcceptProposalWithOOBRequest is used when introducee wants to provide an out-of-band request (via HTTP).
func (ir *IntroduceREST) AcceptProposalWithOOBRequest(request *models.RequestEnvelope) *models.ResponseEnvelope {
	respEnvelope := execREST(&restOperation{
		url:        ir.URL,
		token:      ir.Token,
		httpClient: ir.httpClient,
		endpoint:   *ir.endpoints[cmdintroduce.AcceptProposalWithOOBRequest],
		request:    request,
	})

	return respEnvelope
}

// AcceptProposal is used when introducee wants to accept a proposal without providing a OOBRequest (via HTTP).
func (ir *IntroduceREST) AcceptProposal(request *models.RequestEnvelope) *models.ResponseEnvelope {
	respEnvelope := execREST(&restOperation{
		url:        ir.URL,
		token:      ir.Token,
		httpClient: ir.httpClient,
		endpoint:   *ir.endpoints[cmdintroduce.AcceptProposal],
		request:    request,
	})

	return respEnvelope
}

// AcceptRequestWithPublicOOBRequest is used when an introducer
// wants to provide a published out-of-band request (via HTTP).
func (ir *IntroduceREST) AcceptRequestWithPublicOOBRequest(request *models.RequestEnvelope) *models.ResponseEnvelope {
	respEnvelope := execREST(&restOperation{
		url:        ir.URL,
		token:      ir.Token,
		httpClient: ir.httpClient,
		endpoint:   *ir.endpoints[cmdintroduce.AcceptRequestWithPublicOOBRequest],
		request:    request,
	})

	return respEnvelope
}

// AcceptRequestWithRecipients is used when the introducer does not have a published out-of-band message on hand
// but they are willing to introduce agents to each other. This is done via HTTP.
func (ir *IntroduceREST) AcceptRequestWithRecipients(request *models.RequestEnvelope) *models.ResponseEnvelope {
	respEnvelope := execREST(&restOperation{
		url:        ir.URL,
		token:      ir.Token,
		httpClient: ir.httpClient,
		endpoint:   *ir.endpoints[cmdintroduce.AcceptRequestWithRecipients],
		request:    request,
	})

	return respEnvelope
}

// DeclineProposal is used to reject the proposal (via HTTP).
func (ir *IntroduceREST) DeclineProposal(request *models.RequestEnvelope) *models.ResponseEnvelope {
	respEnvelope := execREST(&restOperation{
		url:        ir.URL,
		token:      ir.Token,
		httpClient: ir.httpClient,
		endpoint:   *ir.endpoints[cmdintroduce.DeclineProposal],
		request:    request,
	})

	return respEnvelope
}

// DeclineRequest is used to reject the request (via HTTP).
func (ir *IntroduceREST) DeclineRequest(request *models.RequestEnvelope) *models.ResponseEnvelope {
	respEnvelope := execREST(&restOperation{
		url:        ir.URL,
		token:      ir.Token,
		httpClient: ir.httpClient,
		endpoint:   *ir.endpoints[cmdintroduce.DeclineRequest],
		request:    request,
	})

	return respEnvelope
}
