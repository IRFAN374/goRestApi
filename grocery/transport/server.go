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
		req := request.(*AddGroceryRequest)
		res0 := svc.AddGrocery(arg0, req.grocery)
		return &AddGroceryResponse{}, res0
	}
}

func GetGroceryEndpoint(svc grocery.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*GetGroceryRequest)
		res0, res1 := svc.GetGrocery(arg0, req.Name)

		return &GetGroceryResponse{
			grocery: res0,
		}, res1

	}
}

func GetAllGroceryEndpoint(svc grocery.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		_ = request.(*GetAllGroceryRequest)
		res0, res1 := svc.GetAllGrocery(arg0)

		return &GetAllGroceryResponse{
			groceries: res0,
		}, res1

	}
}

func UpdateGroceryEndpoint(svc grocery.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*UpdateGroceryRequest)
		res0, res1 := svc.UpdateGrocery(arg0, req.Name)

		return &UpdateGroceryResponse{
			grocery: res0,
		}, res1

	}
}

func DeleteGroceryEndpoint(svc grocery.Service) endpoint.Endpoint {
	return func(arg0 context.Context, request interface{}) (response interface{}, err error) {
		req := request.(*DeleteGroceryRequest)
		res0, res1 := svc.DeleteGrocery(arg0, req.Name)

		return &DeleteGroceryResponse{
			grocery: res0,
		}, res1

	}
}
