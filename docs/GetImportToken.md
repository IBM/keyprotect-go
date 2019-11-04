# GetImportToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expiration** | **float32** | The time in seconds from the creation of an import token that determines how long its associated public key  remains valid.     The minimum value is &#x60;300&#x60; seconds (5 minutes), and the maximum value is &#x60;86400&#x60; (24 hours). The default value is &#x60;600&#x60; (10 minutes). | [optional] [default to 600]
**MaxAllowedRetrievals** | **float32** | The number of times that an import token can be retrieved within its expiration time before it is no longer accessible.  | [optional] [default to 1]
**CreationDate** | [**time.Time**](time.Time.md) | The date the import token was created. The date format follows RFC 3339. | [optional] [readonly] 
**ExpirationDate** | [**time.Time**](time.Time.md) | The date the import token expires. The date format follows RFC 3339. | [optional] [readonly] 
**RemainingRetrievals** | **float32** | The number of retrievals that are available for the import token before it is no longer accessible.   | [optional] [readonly] [default to 1]
**Payload** | [***os.File**](*os.File.md) | The public encryption key that you can use to encrypt key material before you import it into the service.     This value is a PEM-encoded public key in PKIX format. Because PEM encoding is a binary format, the value is base64 encoded. | [optional] [readonly] 
**Nonce** | [***os.File**](*os.File.md) | The nonce value that is used to verify a key import request. Encrypt and provide the encrypted nonce value when you use &#x60;POST /keys&#x60; to securely import a key to the service. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


