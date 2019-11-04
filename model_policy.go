/*
 * IBM Key Protect API
 *
 * IBM Key Protect helps you provision encrypted keys for apps across IBM Cloud. As you manage the lifecycle of your keys, you can benefit from knowing that your keys are secured by cloud-based FIPS 140-2 Level 3 hardware security modules (HSMs) that protect against theft of information. You can use the Key Protect API to store, generate, and retrieve your key material. Keys within the service can protect any type of data in your symmetric key based encryption solution.
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package keyprotect
import (
	"time"
)
// Policy Properties that are associated with policies.
type Policy struct {
	// Specifies the MIME type that represents the policy resource. Currently, only the default is supported.
	Type string `json:"type"`
	// The Cloud Resource Name (CRN) that uniquely identifies your cloud network resources.
	Crn string `json:"crn,omitempty"`
	// The date the policy was created. The date format follows RFC 3339.
	CreationDate time.Time `json:"creationDate,omitempty"`
	// The unique identifier for the resource that created the policy.
	CreatedBy string `json:"createdBy,omitempty"`
	// Updates when the policy is replaced or modified. The date format follows RFC 3339.
	LastUpdateDate time.Time `json:"lastUpdateDate,omitempty"`
	// The unique identifier for the resource that updated the policy.
	UpdatedBy string `json:"updatedBy,omitempty"`
	Rotation PolicyRotation `json:"rotation"`
}
