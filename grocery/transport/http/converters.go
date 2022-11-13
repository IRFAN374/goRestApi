package transporthttp

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	"github.com/IRFAN374/goRestApi/common/chttp"
	transport "github.com/IRFAN374/goRestApi/grocery/transport"
)

func CommonHTTPRequestEncoder(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}

	r.Body = ioutil.NopCloser(&buf)

	return nil
}

func CommonHTTPResponseEncoder(ctx context.Context, w http.ResponseWriter, response interface{}) error {

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	return chttp.EncodeResponse(ctx, w, response)
}

func _Decode_AddGrocery_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.AddGroceryRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	return &req, err
}

func _Decode_GetGrocery_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetGroceryRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	return &req, err
}

func _Decode_GetAllGrocery_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.GetAllGroceryRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	return &req, err
}

func _Decode_UpdateGrocery_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.UpdateGroceryRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	return &req, err
}

func _Decode_DeleteGrocery_Request(_ context.Context, r *http.Request) (interface{}, error) {
	var req transport.DeleteGroceryRequest
	err := json.NewDecoder(r.Body).Decode(&req)

	return &req, err
}

func _Decode_AddGrocery_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.AddGroceryResponse
	err := chttp.DecodeResponse(ctx, r, &resp)

	return &resp, err
}

func _Decode_GetGrocery_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.GetGroceryResponse
	err := chttp.DecodeResponse(ctx, r, &resp)

	return &resp, err
}

func _Decode_GetAllGrocery_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.GetAllGroceryResponse
	err := chttp.DecodeResponse(ctx, r, &resp)

	return &resp, err
}

func _Decode_UpdateGrocery_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.UpdateGroceryResponse
	err := chttp.DecodeResponse(ctx, r, &resp)

	return &resp, err
}

func _Decode_DeleteGrocery_Response(ctx context.Context, r *http.Response) (interface{}, error) {
	var resp transport.DeleteGroceryResponse
	err := chttp.DecodeResponse(ctx, r, &resp)

	return &resp, err
}



func _Encode_AddGrocery_Request(ctx context.Context, r *http.Request, request interface{}) error {
	_ = request.(*transport.AddGroceryRequest)
	r.URL.Path = path.Join(r.URL.Path, fmt.Sprintf("/addgrocery"))
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_GetGrocery_Request(ctx context.Context, r *http.Request, request interface{}) error {
	_ = request.(*transport.GetGroceryRequest)
	r.URL.Path = path.Join(r.URL.Path, fmt.Sprintf("/getgrocery"))
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_GetAllGrocery_Request(ctx context.Context, r *http.Request, request interface{}) error {
	_ = request.(*transport.GetAllGroceryRequest)
	r.URL.Path = path.Join(r.URL.Path, fmt.Sprintf("/getallgrocery"))
	return CommonHTTPRequestEncoder(ctx, r, request)
}
func _Encode_UpdateGrocery_Request(ctx context.Context, r *http.Request, request interface{}) error {
	_ = request.(*transport.UpdateGroceryRequest)
	r.URL.Path = path.Join(r.URL.Path, fmt.Sprintf("/updategrocery"))
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_DeleteGrocery_Request(ctx context.Context, r *http.Request, request interface{}) error {
	_ = request.(*transport.GetGroceryRequest)
	r.URL.Path = path.Join(r.URL.Path, fmt.Sprintf("/deletegrocery"))
	return CommonHTTPRequestEncoder(ctx, r, request)
}

func _Encode_AddGrocery_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_GetGrocery_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_GetAllGrocery_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_UpdateGrocery_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}

func _Encode_DeleteGrocery_Response(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return CommonHTTPResponseEncoder(ctx, w, response)
}
