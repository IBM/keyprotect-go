# GetImportTokenAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Nonce** | [***os.File**](*os.File.md) | The nonce value that is used to verify a key import request. Encrypt and provide the encrypted nonce value when you use &#x60;POST /keys&#x60; to securely import a key to the service. | [optional] [readonly] 
**Payload** | [***os.File**](*os.File.md) | The public encryption key that you can use to encrypt key material before you import it into the service.     This value is a PEM-encoded public key in PKIX format. Because PEM encoding is a binary format, the value is base64 encoded. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


