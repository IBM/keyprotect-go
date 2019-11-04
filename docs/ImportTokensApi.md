# \ImportTokensApi

All URIs are relative to *https://us-south.kms.cloud.ibm.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetImportToken**](ImportTokensApi.md#GetImportToken) | **Get** /api/v2/import_token | Retrieve an import token
[**PostImportToken**](ImportTokensApi.md#PostImportToken) | **Post** /api/v2/import_token | Create an import token



## GetImportToken

> GetImportToken GetImportToken(ctx, bluemixInstance, optional)

Retrieve an import token

Retrieves the import token that is associated with your service instance.    When you call `GET /import_token`, Key Protect returns the public key that you can use to encrypt and import key material to the service, along with details about the key.     **Note:** After you reach the `maxAllowedRetrievals` or `expirationDate` for the import token, the import token and its associated public key can no longer be used for key operations. To create a new import token, use  `POST /import_token`.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bluemixInstance** | **string**| The IBM Cloud instance ID that identifies your Key Protect service instance. | 
 **optional** | ***GetImportTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetImportTokenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **correlationId** | **optional.String**| The v4 UUID used to correlate and track transactions. | 

### Return type

[**GetImportToken**](GetImportToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.ibm.kms.import_token+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostImportToken

> ImportToken PostImportToken(ctx, bluemixInstance, importToken, optional)

Create an import token

Creates an import token that you can use to encrypt and import root keys into the service.  [Learn more](/docs/services/key-protect?topic=key-protect-importing-keys#using-import-tokens)     When you call `POST /import_token`, Key Protect creates an RSA key-pair from its HSMs. The service encrypts and  stores the private key in the HSM, and returns the corresponding public key when you call `GET /import_token`.  You can create only one import token per service instance. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**bluemixInstance** | **string**| The IBM Cloud instance ID that identifies your Key Protect service instance. | 
**importToken** | [**ImportToken**](ImportToken.md)| The base request to create an import token. | 
 **optional** | ***PostImportTokenOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostImportTokenOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **correlationId** | **optional.String**| The v4 UUID used to correlate and track transactions. | 

### Return type

[**ImportToken**](ImportToken.md)

### Authorization

[BearerAuth](../README.md#BearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/vnd.ibm.kms.import_token+json, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

