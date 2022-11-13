package transport

import "context"

func (set EndpointsSet) AddGrocery(arg0 context.Context) (res0 error) {

	request := AddGroceryRequest{}

	_, res0 = set.AddGroceryEndpoint(arg0, &request)

	if res0 != nil {
		return
	}

	return res0
}

func (set EndpointsSet) GetGrocery(arg0 context.Context) (res0 error) {

	request := GetGroceryRequest{}

	_, res0 = set.GetGroceryEndpoint(arg0, &request)

	if res0 != nil {
		return
	}

	return res0
}

func (set EndpointsSet) GetAllGrocery(arg0 context.Context) (res0 error) {

	request := GetAllGroceryRequest{}

	_, res0 = set.GetAllGroceryEndpoint(arg0, &request)

	if res0 != nil {
		return
	}

	return res0
}

func (set EndpointsSet) UpdateGrocery(arg0 context.Context) (res0 error) {

	request := UpdateGroceryRequest{}

	_, res0 = set.UpdateGroceryEndpoint(arg0, &request)

	if res0 != nil {
		return
	}

	return res0
}

func (set EndpointsSet) DeleteGrocery(arg0 context.Context) (res0 error) {

	request := DeleteGroceryRequest{}

	_, res0 = set.DeleteGroceryEndpoint(arg0, &request)

	if res0 != nil {
		return
	}

	return res0
}
