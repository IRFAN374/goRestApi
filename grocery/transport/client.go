package transport

import (
	"context"
	"fmt"

	model "github.com/IRFAN374/goRestApi/model"
)

func (set EndpointsSet) AddGrocery(arg0 context.Context, grocery model.Grocery) (res0 error) {

	request := AddGroceryRequest{
		Grocery: grocery,
	}

	fmt.Println("request:", request)

	_, res0 = set.AddGroceryEndpoint(arg0, &request)

	if res0 != nil {
		return
	}

	return res0
}

func (set EndpointsSet) GetGrocery(arg0 context.Context, Name string) (grocery model.Grocery, res0 error) {

	request := GetGroceryRequest{
		Name: Name,
	}

	res1, res0 := set.GetGroceryEndpoint(arg0, &request)

	if res0 != nil {
		return
	}

	return res1.(*GetGroceryResponse).Grocery, res0
}

func (set EndpointsSet) GetAllGrocery(arg0 context.Context) (groceries []model.Grocery, res0 error) {

	request := GetAllGroceryRequest{}

	response, res0 := set.GetAllGroceryEndpoint(arg0, &request)

	if res0 != nil {
		return
	}

	return response.(*GetAllGroceryResponse).Groceries, res0
}

func (set EndpointsSet) UpdateGrocery(arg0 context.Context, arg1 string) (grocery model.Grocery, res0 error) {

	request := UpdateGroceryRequest{
		Name: arg1,
	}

	response, res0 := set.UpdateGroceryEndpoint(arg0, &request)

	if res0 != nil {
		return
	}

	return response.(*UpdateGroceryResponse).Grocery, res0
}

func (set EndpointsSet) DeleteGrocery(arg0 context.Context, Name string) (grocery model.Grocery, res0 error) {

	request := DeleteGroceryRequest{
		Name: Name,
	}

	response, res0 := set.DeleteGroceryEndpoint(arg0, &request)

	if res0 != nil {
		return
	}

	return response.(*DeleteGroceryResponse).Grocery, res0
}
