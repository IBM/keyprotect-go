# Policy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedBy** | **string** | The unique identifier for the resource that created the policy. | [optional] [readonly] 
**CreationDate** | [**time.Time**](time.Time.md) | The date the policy was created. The date format follows RFC 3339. | [optional] [readonly] 
**Crn** | **string** | The Cloud Resource Name (CRN) that uniquely identifies your cloud network resources. | [optional] [readonly] 
**LastUpdateDate** | [**time.Time**](time.Time.md) | Updates when the policy is replaced or modified. The date format follows RFC 3339. | [optional] [readonly] 
**Rotation** | [**PolicyRotation**](Policy_rotation.md) |  | 
**Type** | **string** | Specifies the MIME type that represents the policy resource. Currently, only the default is supported. | [default to application/vnd.ibm.kms.key+json]
**UpdatedBy** | **string** | The unique identifier for the resource that updated the policy. | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


