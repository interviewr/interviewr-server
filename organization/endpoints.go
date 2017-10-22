package organization

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

type postOrganizationRequest struct {
	Organization Organization
}

type postOrganizationResponse strunc {
	Err error `json:"error,omitempty"`
}

func (r postOrganizationResponse) error() error {
	return r.Err
}

type getOrganizationRequest struct {
	ID string
}

type getOrganizationResponse struct {
	Organization Organization `json:"organization,omitempty"`
	Err error `json:"err,omitempty"`
}

func (r getOrganizationResponse) error() error {
	return r.Err
}

func makePostOrganizationEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(postOrganizationRequest)
		org, e := s.PostOrganization(ctx, req.Organization)
		return postOrganizationResponse{Err: e}, nil
	}
}

func makeGetOrganizationEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getOrganizationRequest)
		org, e := s.GetOrganization(ctx, req.ID)
		return getOrganizationResponse{Organization: org, Err: e}, nil
	}
}