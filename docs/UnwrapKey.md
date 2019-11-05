# UnwrapKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aad** | **[]string** | The additional authentication data (AAD) used to further secure the key.     If you supply AAD when you make a &#x60;wrap&#x60; call, you must specify the same AAD during a subsequent &#x60;unwrap&#x60; call.  | [optional] 
**Ciphertext** | **string** | The wrapped data encryption key (DEK) used in wrap actions when the query parameter is set to &#x60;unwrap&#x60;. The system requires a base64 encoded ciphertext and returns a base64 encoded plaintext in the response entity-body when you perform an &#x60;unwrap&#x60; action on a key.       **Note:** When you unwrap a wrapped data encryption key (WDEK) by using a rotated root key, the service returns a new ciphertext in the response entity-body. Each ciphertext remains available for &#x60;unwrap&#x60; actions.  If you unwrap a DEK with a previous ciphertext, the service also returns the latest ciphertext in the response.  Use the latest ciphertext for future unwrap operations. | 
**Plaintext** | **string** | The original data encryption key (DEK) that was used in wrap action. The value is base64 encoded. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


