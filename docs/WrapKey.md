# WrapKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Aad** | **[]string** | The additional authentication data (AAD) used to further secure the key.     If you supply AAD when you make a &#x60;wrap&#x60; call, you must specify the same AAD during a subsequent &#x60;unwrap&#x60; call.  | [optional] 
**Ciphertext** | **string** | The wrapped data encryption key (DEK) that you can export to your app or service. The value is base64 encoded. | [optional] [readonly] 
**Plaintext** | **string** | The data encryption key (DEK) used in wrap actions when the query parameter is set to &#x60;wrap&#x60;. The system returns a base64 encoded plaintext in the response entity-body when you perform an &#x60;unwrap&#x60; action on a key.     To wrap an existing DEK, provide a base64 encoded plaintext during a &#x60;wrap&#x60; action. To generate a new DEK, omit the &#x60;plaintext&#x60; property. Key Protect generates a random plaintext (32 bytes) that is rooted in an HSM and then wraps that value.  | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


