package transport

import (
	"context"

	grocery "github.com/IRFAN374/goRestApi/grocery"
	endpoint "github.com/go-kit/kit/endpoint"
)

func Endpoints(svc grocery.Service) EndpointsSet {
	return EndpointsSet{
		AddGroceryEndpoint:    AddGroceryEndpoint(svc),
		GetGroceryEndpoint:    GetGroceryEndpoint(svc),
		GetAllGroceryEndpoint: GetAllGroceryEndpoint(svc),
		UpdateGroceryEndpoint: UpdateGroceryEndpoint(svc),
		DeleteGroceryEndpoint: DeleteGroceryEndpoint(svc),
	}
}

func AddGroceryEndpoint(svc grocery.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(*AddGroceryRequest)
		res0 := svc.AddGrocery(arg0)
		return &AddGroceryResponse{}, res0
	}
}

func GetGroceryEndpoint(svc grocery.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(*GetGroceryRequest)
		res0 := svc.GetGrocery(arg0)

		return &GetGroceryResponse{}, res0

	}
}

func GetAllGroceryEndpoint(svc grocery.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(*GetAllGroceryRequest)
		res0 := svc.GetGrocery(arg0)

		return &GetAllGroceryResponse{}, res0

	}
}

func UpdateGroceryEndpoint(svc grocery.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(*UpdateGroceryRequest)
		res0 := svc.UpdateGrocery(arg0)

		return &UpdateGroceryResponse{}, res0

	}
}

func DeleteGroceryEndpoint(svc grocery.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(*DeleteGroceryRequest)
		res0 := svc.DeleteGrocery(arg0)

		return &DeleteGroceryResponse{}, res0

	}
}
