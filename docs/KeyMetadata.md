# KeyMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlgorithmMetadata** | [**KeyMetadataAlgorithmMetadata**](KeyMetadata_algorithmMetadata.md) |  | [optional] 
**AlgorithmType** | **string** | The algorithm type used to generate the key. Currently, AES is supported. | [optional] [readonly] [default to ALGORITHM_TYPE_AES]
**CreatedBy** | **string** | The unique identifier for the resource that created the key. | [optional] [readonly] 
**CreationDate** | [**time.Time**](time.Time.md) | The date the key material was created. The date format follows RFC 3339. | [optional] [readonly] 
**LastRotateDate** | [**time.Time**](time.Time.md) | Updates to show when the key was last rotated. The date format follows RFC 3339. | [optional] [readonly] 
**LastUpdateDate** | [**time.Time**](time.Time.md) | Updates when any part of the key metadata is modified. The date format follows RFC 3339. | [optional] [readonly] 
**NonactiveStateReason** | **int32** | A code indicating the reason the key is not in the activation state. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


