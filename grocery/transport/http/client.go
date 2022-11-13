package transporthttp

import (
	"net/url"

	transport "github.com/IRFAN374/goRestApi/grocery/transport"
	httpkit "github.com/go-kit/kit/transport/http"
)

func NewHTTPClient(u *url.URL, opts ...httpkit.ClientOption) transport.EndpointsSet {

	return transport.EndpointsSet{
		AddGroceryEndpoint: httpkit.NewClient(
			"POST", u,
			_Encode_AddGrocery_Request,
			_Decode_AddGrocery_Response,
			opts...,
		).Endpoint(),
		GetGroceryEndpoint: httpkit.NewClient(
			"GET", u,
			_Encode_GetGrocery_Request,
			_Decode_GetGrocery_Response,
			opts...,
		).Endpoint(),
		GetAllGroceryEndpoint: httpkit.NewClient(
			"GET", u,
			_Encode_GetAllGrocery_Request,
			_Decode_GetAllGrocery_Response,
			opts...,
		).Endpoint(),
		UpdateGroceryEndpoint: httpkit.NewClient(
			"PUT", u,
			_Encode_UpdateGrocery_Request,
			_Decode_UpdateGrocery_Response,
			opts...,
		).Endpoint(),
		DeleteGroceryEndpoint: httpkit.NewClient(
			"DELETE", u,
			_Encode_DeleteGrocery_Request,
			_Decode_UpdateGrocery_Response,
			opts...,
		).Endpoint(),
	}

}
