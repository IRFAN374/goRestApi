package transporthttp

import (
	http1 "net/http"

	transport "github.com/IRFAN374/goRestApi/grocery/transport"
	http "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"
)

func NewHTTPHandler(endpoints *transport.EndpointsSet, opts ...http.ServerOption) http1.Handler {
	mux := mux.NewRouter()

	mux.Methods("POST").Path("/api/addgrocery").Handler(
		http.NewServer(
			endpoints.AddGroceryEndpoint,
			_Decode_AddGrocery_Request,
			_Encode_AddGrocery_Response,
			opts...))
    mux.Methods("GET").Path("/api/getgrocery").Handler(
		http.NewServer(
			endpoints.GetGroceryEndpoint,
			_Decode_GetGrocery_Request,
			_Encode_GetGrocery_Response,
			opts...))
	mux.Methods("GET").Path("/api/getallgrocery").Handler(
		http.NewServer(
			endpoints.GetAllGroceryEndpoint,
			_Decode_GetAllGrocery_Request,
			_Encode_GetAllGrocery_Response,
			opts...))
	mux.Methods("PUT").Path("/api/updategrocery").Handler(
		http.NewServer(
			endpoints.UpdateGroceryEndpoint,
			_Decode_UpdateGrocery_Request,
			_Encode_UpdateGrocery_Response,
			opts...))
    mux.Methods("DELETE").Path("/api/deletegrocery").Handler(
		http.NewServer(
			endpoints.DeleteGroceryEndpoint,
			_Decode_DeleteGrocery_Request,
			_Encode_DeleteGrocery_Response,
			opts...))

	return mux
	
}