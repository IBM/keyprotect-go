# \PoliciesApi

All URIs are relative to *https://us-south.kms.cloud.ibm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPolicy**](PoliciesApi.md#GetPolicy) | **Get** /api/v2/keys/{id}/policies | Retrieve a list of policies
[**PutPolicy**](PoliciesApi.md#PutPolicy) | **Put** /api/v2/keys/{id}/policies | Replace an existing policy



## GetPolicy

> CreatePolicy GetPolicy(ctx, id, bluemixInstance, optional)

Retrieve a list of policies

Retrieves a list of policies that are associated with a specified key.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The v4 UUID that uniquely identifies the key. | 
**bluemixInstance** | **string**| The IBM Cloud instance ID that identifies your Key Protect service instance. | 
 **optional** | ***GetPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **correlationId** | **optional.String**| The v4 UUID used to correlate and track transactions. | 

### Return type

[**CreatePolicy**](CreatePolicy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPolicy

> CreatePolicy PutPolicy(ctx, id, bluemixInstance, createPolicy, optional)

Replace an existing policy

Replaces the policy that is associated with a specified key

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| The v4 UUID that uniquely identifies the key. | 
**bluemixInstance** | **string**| The IBM Cloud instance ID that identifies your Key Protect service instance. | 
**createPolicy** | [**CreatePolicy**](CreatePolicy.md)| The base request for creating a new policies resource. | 
 **optional** | ***PutPolicyOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutPolicyOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **correlationId** | **optional.String**| The v4 UUID used to correlate and track transactions. | 

### Return type

[**CreatePolicy**](CreatePolicy.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/vnd.ibm.kms.key+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

