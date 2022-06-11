package messages

import "github.com/CascadiaShaman/go/protocol/interfaces"

// Specification:  https://identity.foundation/decentralized-web-node/spec/#request-objects
// Request Objects are JSON object envelopes used to pass messages to Decentralized Web Nodes.
type RequestObject struct {

	// The Request Object MUST include a requestId property, and its value MUST be a [RFC4122] UUID Version 4 string to identify the request.
	ID string `json:"requestId"`

	// The Request Object MUST include a target property, and its value MUST be the Decentralized Identifier base URI of the DID-relative URL.
	Target string `json:"target"`

	// The Request Object MUST include a messages property, and its value MUST be an array composed of Message objects.
	Messages []Message `json:"messages"`
}

// Specification:  https://identity.foundation/decentralized-web-node/spec/#messages
// All Decentralized Web Node messaging is transacted via Messages JSON objects. These
// objects contain message execution parameters, authorization material, authorization
// signatures, and signing/encryption information. For various purposes Messages rely on
// IPFS CIDs and DAG APIs.
type Message struct {

	// Message objects MUST contain a descriptor property, and its value MUST be an object, as defined by the Message Descriptors section of this specification.
	Descriptor MessageDescriptor `json:"descriptor"`

	// Message objects MAY contain a data property, and if present its value MUST be a base64Url encoded string of the Message’s data.
	Data string `json:"data"`

	// Message objects MAY contain an attestation property, and if present its value MUST be an object, as defined by the Signed Data section of this specification.
	Attestation SignedData `json:"attestation"`

	// If a Message object requires signatory and/or permission-evaluated authorization, it MUST include an authorization property, and its value MUST be a valid
	// Permission Grant invocation, as described in the Permissions interface section.
	Authorization PermissionsGrant `json:"authorization"`
}

// Specification:  https://identity.foundation/decentralized-web-node/spec/#message-descriptors
// The Decentralized Web Node data structure that resides in the descriptor property of the
// Message is comprised of a common JSON structure that contains the following properties regardless
// of whether the message data is signed/encrypted
type MessageDescriptor struct {

	// The object MUST contain a nonce property, and its value MUST be a cryptographically random string that ensures each object is unique.
	Nonce string `json:"nonce"`

	// The object MUST contain a method property, and its value MUST be a string that matches a Decentralized Web Node Interface method.
	Method interfaces.InterfaceMethod `json:"method"`

	// If the Message has data associated with it, passed directly via the data property of the Message or an external channel (e.g. IPFS fetch),
	// the descriptor object MUST contain a dataCid property, and its value MUST be the stringified Version 1 CID of the DAG PB encoded data.
	DataCID string `json:"dataCid"`

	// If the Message has data associated with it, passed directly via the data property of the Message object or through a channel external to the
	// message object, the descriptor object MUST contain a dataFormat property, and its value MUST be a string that corresponds with a registered
	// IANA Media Type data format (the most common being plain JSON, which is indicated by setting the value of the dataFormat property to application/json),
	// or one of the following format strings pending registration:
	// application/vc+jwt - the data is a JSON Web Token (JWT) [RFC7519] formatted variant of a W3C Verifiable Credential.
	// application/vc+ldp - the data is a JSON-LD formatted W3C Verifiable Credential.
	DataFormat interfaces.DataFormat `json:"dataFormat"`
}

// Specification: https://identity.foundation/decentralized-web-node/spec/#signed-data
// If the object is to be attested by a signer (e.g the Node owner via signature with their DID
// key), the object MUST contain the following additional properties to produce a [RFC7515] Compact
// JSON Web Signature (JWS):
type SignedData struct {
}

// Specification: https://identity.foundation/decentralized-web-node/spec/#grant
// PermissionsGrant messages are JSON objects containing capabilities granted to parties that curtail the scope of permitted activities an invoker
// can perform. They are generated either in response to a PermissionsRequest message or optimistically by a user agent without an initiating PermissionsRequest.
type PermissionsGrant struct {

	// The message object MUST contain a descriptor property, and its value MUST be a JSON object composed as follows:
	// The object MUST contain a method property, and its value MUST be the string PermissionsGrant.
	MessageDescriptor

	// The object MUST contain an permissionGrantId property, and its value MUST be a [RFC4122] UUID Version 4 string representing the reply object.2
	ID string `json:"permissionGrantId"`

	// If the granted permission is in response to a PermissionRequest, the object MUST contain a permissionRequestId property, and its value MUST be
	// the [RFC4122] UUID Version 4 string of the PermissionRequest object the permission is being granted in relation to.
	PermissionRequestID string `json:"permissionRequestId"`

	// The object MUST contain a grantedBy property, and its value MUST be the DID URI string of the party that is granting the permission.
	GrantedBy string `json:"grantedBy"`

	// The object MUST contain a grantedTo property, and its value MUST be the DID URI string of the party that is being granted the permission.
	GrantedTo string `json:"grantedTo"`

	// The object MUST contain a expiry property, and its value MUST be a Unix epoch timestamp that can be used to trigger revocation activities.
	Expiry string `json:"expiry"`

	// The object MUST contain a scope property, and its value MUST be an object
	GrantScope Scope `json:"scope"`

	// The object MAY contain a conditions property, and its value MUST be an object
	GrantConditions Conditions `json:"conditions"`

	// The message object MUST contain an attestation property, which MUST be a JSON object as defined by the Signed Data section of this specification,
	// with the requirement that the kid and signature MUST match the DID of the requesting party.
	Attestation SignedData `json:"attestation"`

	// The message MUST contain a data payload, which is a JSON Web Token representation of the granted permission, as defined in the Capability Objects section.
	Data string `json:"data"`

	// The message MUST contain an encryptionKey property if the data transacted using the permission grant is to be encrypted, per the directives for encryption
	// under the encryption field of the permission’s conditions. If present, the value of the encryptionKey property MUST be a [RFC7516] JSON Web Encryption (JWE)
	// object that contains the encrypted key material required for an authorized party to decrypt the JWE represented by the dataCid value within the descriptor object.
	EncryptionKey string `json:"encryptionKey"` // TODO:  This is an object, I think.  Come back to this
}

// Specification: https://identity.foundation/decentralized-web-node/spec/#grant
// Implementations may extend this:
// The object MAY contain an identifier property that corresponds with the method
// specified (i.e. recordId for Collections, threadId for Threads), and its value
// MUST be a UUID 4 string reference to an object. If an identifier property is present the
// scope object MUST include a schema property.
type Scope struct {
	// The object MUST contain a method property, and its value MUST be the interface method the requesting party wants to invoke.
	Method interfaces.InterfaceMethod `json:"method"`

	// The object MAY contain a schema property, and its value MUST be a URI string that indicates the schema of the associated data.
	Schema string `json:"schema"`
}

// Specification: https://identity.foundation/decentralized-web-node/spec/#grant
type SigningConditions string

const (
	// the object MUST NOT be signed.
	Prohibited SigningConditions = "prohibited"

	// the object MAY be signed using a key linked to the DID of the owner of a Decentralized Web Node or authoring party (whichever
	// is relevant to the application-level use case), and the signature MUST be in the [RFC7515] JSON Web Signature (JWS) format.
	Optional SigningConditions = "optional"

	// the object MUST be signed using a key linked to the DID of the owner of a Decentralized Web Node or authoring party (whichever
	// is relevant to the application-level use case), and the signature MUST be in the [RFC7515] JSON Web Signature (JWS) format.
	Required SigningConditions = "required"
)

// Specification: https://identity.foundation/decentralized-web-node/spec/#grant
type EncryptionConditions SigningConditions

// Specification: https://identity.foundation/decentralized-web-node/spec/#grant
type Conditions struct {

	// The object MAY contain an attestation property, and if present its value MUST be a string representing the signing conditions detailed below. If the property is not present it MUST be evaluated as if it were set to the value optional.
	// prohibited - the object MUST NOT be signed.
	// ptional - the object MAY be signed using a key linked to the DID of the owner of a Decentralized Web Node or authoring party (whichever is relevant to the application-level use case), and the signature MUST be in the [RFC7515] JSON Web Signature (JWS) format.
	// required - the object MUST be signed using a key linked to the DID of the owner of a Decentralized Web Node or authoring party (whichever is relevant to the application-level use case), and the signature MUST be in the [RFC7515] JSON Web Signature (JWS) format.
	Attestation SigningConditions `json:"attestation"`

	// The object MAY contain an encryption property, and if present its value MUST be a string representing the encryption conditions detailed below. If the property is not present it MUST be evaluated as if it were set to the value optional.
	// optional - the object MAY be encrypted using the key provided by the owner of a Decentralized Web Node in the [RFC7516] JSON Web Encryption (JWE) format.
	// required - the object MUST be encrypted using the key provided by the owner of a Decentralized Web Node in the [RFC7516] JSON Web Encryption (JWE) format.
	Encryption EncryptionConditions `json:"encryption"`

	// The object MAY contain a delegation property, and its value MUST be a boolean, wherein true indicates the issuing party is allowing the grantee the ability to delegate the capability. A value of false or omission of the property MUST be evaluated as false, and
	// indicates the grantee MUST NOT be allowed to delegate the capability.
	Delegation bool `json:"delegation"`

	// The object MAY contain a publication property, and its value MUST be a boolean, wherein true indicates the issuing party is allowing the grantee the ability to publish data tied to methods that support the public boolean value in their descriptor field sets.
	// Conforming implementations MUST throw an error and fail to grant a permission if this property is present and the method does not support publication.
	Publication bool `json:"publication"`

	// The object MAY contain a sharedAccess property, and its value MUST be a boolean, wherein true indicates the requesting party wants the ability to use the permission against any object or data that aligns with the capability’s definition, regardless of which
	// entity created the object or data. A value of false or omission of the property MUST be evaluated as false, and indicates the requesting party only needs the ability to invoke the permission against objects or data it creates.
	SharedAccess bool `json:"sharedAccess"`
}
