/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type StoreApi interface {

  /*
   * DeleteOrder Delete purchase order by ID
   * For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors
   * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
   * @param orderId ID of the order that needs to be deleted
   * @return ApiDeleteOrderRequest
   */
  DeleteOrder(ctx _context.Context, orderId string) ApiDeleteOrderRequest

  /*
   * DeleteOrderExecute executes the request
   */
  DeleteOrderExecute(r ApiDeleteOrderRequest) (*_nethttp.Response, error)

  /*
   * GetInventory Returns pet inventories by status
   * Returns a map of status codes to quantities
   * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
   * @return ApiGetInventoryRequest
   */
  GetInventory(ctx _context.Context) ApiGetInventoryRequest

  /*
   * GetInventoryExecute executes the request
   * @return map[string]int32
   */
  GetInventoryExecute(r ApiGetInventoryRequest) (map[string]int32, *_nethttp.Response, error)

  /*
   * GetOrderById Find purchase order by ID
   * For valid response try integer IDs with value <= 5 or > 10. Other values will generated exceptions
   * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
   * @param orderId ID of pet that needs to be fetched
   * @return ApiGetOrderByIdRequest
   */
  GetOrderById(ctx _context.Context, orderId int64) ApiGetOrderByIdRequest

  /*
   * GetOrderByIdExecute executes the request
   * @return Order
   */
  GetOrderByIdExecute(r ApiGetOrderByIdRequest) (Order, *_nethttp.Response, error)

  /*
   * PlaceOrder Place an order for a pet
   * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
   * @return ApiPlaceOrderRequest
   */
  PlaceOrder(ctx _context.Context) ApiPlaceOrderRequest

  /*
   * PlaceOrderExecute executes the request
   * @return Order
   */
  PlaceOrderExecute(r ApiPlaceOrderRequest) (Order, *_nethttp.Response, error)
}

// StoreApiService StoreApi service
type StoreApiService service

type ApiDeleteOrderRequest struct {
	ctx _context.Context
	ApiService StoreApi
	orderId string
}


func (r ApiDeleteOrderRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteOrderExecute(r)
}

/*
 * DeleteOrder Delete purchase order by ID
 * For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orderId ID of the order that needs to be deleted
 * @return ApiDeleteOrderRequest
 */
func (a *StoreApiService) DeleteOrder(ctx _context.Context, orderId string) ApiDeleteOrderRequest {
	return ApiDeleteOrderRequest{
		ApiService: a,
		ctx: ctx,
		orderId: orderId,
	}
}

/*
 * Execute executes the request
 */
func (a *StoreApiService) DeleteOrderExecute(r ApiDeleteOrderRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreApiService.DeleteOrder")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/order/{order_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"order_id"+"}", _neturl.PathEscape(parameterToString(r.orderId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetInventoryRequest struct {
	ctx _context.Context
	ApiService StoreApi
}


func (r ApiGetInventoryRequest) Execute() (map[string]int32, *_nethttp.Response, error) {
	return r.ApiService.GetInventoryExecute(r)
}

/*
 * GetInventory Returns pet inventories by status
 * Returns a map of status codes to quantities
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiGetInventoryRequest
 */
func (a *StoreApiService) GetInventory(ctx _context.Context) ApiGetInventoryRequest {
	return ApiGetInventoryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return map[string]int32
 */
func (a *StoreApiService) GetInventoryExecute(r ApiGetInventoryRequest) (map[string]int32, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  map[string]int32
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreApiService.GetInventory")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/inventory"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["api_key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiGetOrderByIdRequest struct {
	ctx _context.Context
	ApiService StoreApi
	orderId int64
}


func (r ApiGetOrderByIdRequest) Execute() (Order, *_nethttp.Response, error) {
	return r.ApiService.GetOrderByIdExecute(r)
}

/*
 * GetOrderById Find purchase order by ID
 * For valid response try integer IDs with value <= 5 or > 10. Other values will generated exceptions
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param orderId ID of pet that needs to be fetched
 * @return ApiGetOrderByIdRequest
 */
func (a *StoreApiService) GetOrderById(ctx _context.Context, orderId int64) ApiGetOrderByIdRequest {
	return ApiGetOrderByIdRequest{
		ApiService: a,
		ctx: ctx,
		orderId: orderId,
	}
}

/*
 * Execute executes the request
 * @return Order
 */
func (a *StoreApiService) GetOrderByIdExecute(r ApiGetOrderByIdRequest) (Order, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Order
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreApiService.GetOrderById")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/order/{order_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"order_id"+"}", _neturl.PathEscape(parameterToString(r.orderId, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.orderId < 1 {
		return localVarReturnValue, nil, reportError("orderId must be greater than 1")
	}
	if r.orderId > 5 {
		return localVarReturnValue, nil, reportError("orderId must be less than 5")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiPlaceOrderRequest struct {
	ctx _context.Context
	ApiService StoreApi
	body *Order
}

func (r ApiPlaceOrderRequest) Body(body Order) ApiPlaceOrderRequest {
	r.body = &body
	return r
}

func (r ApiPlaceOrderRequest) Execute() (Order, *_nethttp.Response, error) {
	return r.ApiService.PlaceOrderExecute(r)
}

/*
 * PlaceOrder Place an order for a pet
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @return ApiPlaceOrderRequest
 */
func (a *StoreApiService) PlaceOrder(ctx _context.Context) ApiPlaceOrderRequest {
	return ApiPlaceOrderRequest{
		ApiService: a,
		ctx: ctx,
	}
}

/*
 * Execute executes the request
 * @return Order
 */
func (a *StoreApiService) PlaceOrderExecute(r ApiPlaceOrderRequest) (Order, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  Order
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StoreApiService.PlaceOrder")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/store/order"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/xml", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.body
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
