# \KeysApi

All URIs are relative to *https://%7Bregion%7D.kms.cloud.ibm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActionOnKey**](KeysApi.md#ActionOnKey) | **Post** /api/v2/keys/{id} | Invoke an action on a key
[**CreateKey**](KeysApi.md#CreateKey) | **Post** /api/v2/keys | Create a new key
[**DeleteKey**](KeysApi.md#DeleteKey) | **Delete** /api/v2/keys/{id} | Delete a key by ID
[**GetKey**](KeysApi.md#GetKey) | **Get** /api/v2/keys/{id} | Retrieve a key by ID
[**GetKeyCollectionMetadata**](KeysApi.md#GetKeyCollectionMetadata) | **Head** /api/v2/keys | Retrieve the number of keys
[**GetKeys**](KeysApi.md#GetKeys) | **Get** /api/v2/keys | Retrieve a list of keys



## ActionOnKey

> map[string]interface{} ActionOnKey(ctx, id, action, bluemixInstance, body, optional)

Invoke an action on a key

Invokes an action, such as a `wrap`, `unwrap`, or `rotate` operation, on a specified root key.       **Note:** When you unwrap a wrapped data encryption key (WDEK) by using a rotated root key, the service returns a new ciphertext in the response entity-body. Each ciphertext remains available for `unwrap` actions. If you unwrap a DEK with a previous ciphertext, the service also returns the latest ciphertext in the response. Use the latest  ciphertext for future unwrap operations.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The root key that is used as the wrapping key. It must be a v4 UUID for an active key. | 
**action** | **string**| The action to perform on the specified key. | 
**bluemixInstance** | **string**| The IBM Cloud instance ID that identifies your Key Protect service instance. | 
**body** | **map[string]interface{}**| The base request for key actions. | 
 **optional** | ***ActionOnKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ActionOnKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **correlationId** | **optional.String**| The v4 UUID used to correlate and track transactions. | 
 **prefer** | **optional.String**| Alters server behavior for POST or DELETE operations. A header with &#x60;return&#x3D;minimal&#x60; causes the service to  return only the key identifier, or metadata. A header containing &#x60;return&#x3D;representation&#x60; returns both the key  material and metadata in the response entity-body. If the key has been designated as a root key, the  system cannot return the key material.      **Note:** During POST operations, Key Protect may not immediately return the key material due to key generation time. To retrieve the key material, you can perform a subsequent &#x60;GET /keys/{id}&#x60; request. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.ibm.kms.key_action+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateKey

> map[string]interface{} CreateKey(ctx, bluemixInstance, body, optional)

Create a new key

Creates a new key with specified key material.      Key Protect designates the resource as either a root key or a standard key based on the `extractable` value that you specify. A successful `POST /keys` operation adds the key to the service and  returns the details of the request in the response entity-body, if the Prefer header is set to  `return=representation`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bluemixInstance** | **string**| The IBM Cloud instance ID that identifies your Key Protect service instance. | 
**body** | **map[string]interface{}**| The base request for creating a new key. | 
 **optional** | ***CreateKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **correlationId** | **optional.String**| The v4 UUID used to correlate and track transactions. | 
 **prefer** | **optional.String**| Alters server behavior for POST or DELETE operations. A header with &#x60;return&#x3D;minimal&#x60; causes the service to  return only the key identifier, or metadata. A header containing &#x60;return&#x3D;representation&#x60; returns both the key  material and metadata in the response entity-body. If the key has been designated as a root key, the  system cannot return the key material.      **Note:** During POST operations, Key Protect may not immediately return the key material due to key generation time. To retrieve the key material, you can perform a subsequent &#x60;GET /keys/{id}&#x60; request. | 

### Return type

[**map[string]interface{}**](map[string]interface{}.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.ibm.kms.key+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteKey

> DeleteKey DeleteKey(ctx, id, bluemixInstance, optional)

Delete a key by ID

Deletes a key by specifying the ID of the key.    **Important:** When you delete a key, you permanently shred its contents and associated data. The action cannot be reversed.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The v4 UUID that uniquely identifies the key. | 
**bluemixInstance** | **string**| The IBM Cloud instance ID that identifies your Key Protect service instance. | 
 **optional** | ***DeleteKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **correlationId** | **optional.String**| The v4 UUID used to correlate and track transactions. | 
 **prefer** | **optional.String**| Alters server behavior for POST or DELETE operations. A header with &#x60;return&#x3D;minimal&#x60; causes the service to  return only the key identifier, or metadata. A header containing &#x60;return&#x3D;representation&#x60; returns both the key  material and metadata in the response entity-body. If the key has been designated as a root key, the  system cannot return the key material.      **Note:** During POST operations, Key Protect may not immediately return the key material due to key generation time. To retrieve the key material, you can perform a subsequent &#x60;GET /keys/{id}&#x60; request. | 

### Return type

[**DeleteKey**](DeleteKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKey

> GetKey GetKey(ctx, id, bluemixInstance, optional)

Retrieve a key by ID

Retrieves the details of a key by specifying the ID of the key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The v4 UUID that uniquely identifies the key. | 
**bluemixInstance** | **string**| The IBM Cloud instance ID that identifies your Key Protect service instance. | 
 **optional** | ***GetKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetKeyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **correlationId** | **optional.String**| The v4 UUID used to correlate and track transactions. | 

### Return type

[**GetKey**](GetKey.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyCollectionMetadata

> GetKeyCollectionMetadata(ctx, bluemixInstance, optional)

Retrieve the number of keys

Returns the same HTTP headers as a GET request without returning the entity-body. This operation returns the number of keys in your instance in a header called `Key-Total`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bluemixInstance** | **string**| The IBM Cloud instance ID that identifies your Key Protect service instance. | 
 **optional** | ***GetKeyCollectionMetadataOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetKeyCollectionMetadataOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **correlationId** | **optional.String**| The v4 UUID used to correlate and track transactions. | 

### Return type

 (empty response body)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeys

> ListKeys GetKeys(ctx, bluemixInstance, optional)

Retrieve a list of keys

Retrieves a list of keys that are stored in your Key Protect service instance.    **Note:** `GET /keys` will not return the key material in the response body. You can retrieve the key material for a standard key with a subsequent `GET /keys/{id}` request.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bluemixInstance** | **string**| The IBM Cloud instance ID that identifies your Key Protect service instance. | 
 **optional** | ***GetKeysOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetKeysOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **correlationId** | **optional.String**| The v4 UUID used to correlate and track transactions. | 
 **limit** | **optional.Int32**| The number of keys to retrieve. By default, &#x60;GET /keys&#x60; returns the first 200 keys. To retrieve a different set of keys, use &#x60;limit&#x60; with &#x60;offset&#x60; to page through your available resources. The maximum value for &#x60;limit&#x60; is 5000.    **Usage:** If you have 20 keys in your instance, and you want to retrieve only the first 5 keys, use  &#x60;../keys?limit&#x3D;5&#x60;. | [default to 200]
 **offset** | **optional.Int32**| The number of keys to skip. By specifying &#x60;offset&#x60;, you retrieve a subset of keys that starts with the &#x60;offset&#x60; value. Use &#x60;offset&#x60; with &#x60;limit&#x60; to page through your available resources.    **Usage:** If you have 100 keys in your instance, and you want to retrieve keys 26 through 50, use  &#x60;../keys?offset&#x3D;25&amp;limit&#x3D;25&#x60;.     | [default to 0]

### Return type

[**ListKeys**](ListKeys.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

