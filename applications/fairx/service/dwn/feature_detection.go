package dwn

type FeatureDetection struct {
	Type       string           `json:"type"`
	Interfaces FeatureInterface `json:"interfaces"`
}

type FeatureInterface struct {
	Collections CollectionsFeatures `json:"collections,omitempty"`
	Actions     ActionsFeatures     `json:"actions,omitempty"`
	Permissions PermissionsFeatures `json:"permissions,omitempty"`
	Messaging   MessagingFeatures   `json:"messaging,omitempty"`
	FairX       FairXFeatures       `json:"fairx,omitempty"`
}

type CollectionsFeatures struct {
	CollectionsQuery  bool `json:"CollectionsQuery"`
	CollectionsWrite  bool `json:"CollectionsWrite"`
	CollectionsCommit bool `json:"CollectionsCommit"`
	CollectionsDelete bool `json:"CollectionsDelete"`
}

type ActionsFeatures struct {
	ThreadsQuery  bool `json:"ThreadsQuery"`
	ThreadsCreate bool `json:"ThreadsCreate"`
	ThreadsReply  bool `json:"ThreadsReply"`
	ThreadsClose  bool `json:"ThreadsClose"`
	ThreadsDelete bool `json:"ThreadsDelete"`
}

type PermissionsFeatures struct {
	PermissionsRequest bool `json:"PermissionsRequest"`
	PermissionsGrant   bool `json:"PermissionsGrant"`
	PermissionsRevoke  bool `json:"PermissionsRevoke"`
}

type MessagingFeatures struct {
	Batching bool `json:"batching"`
}

type FairXFeatures struct {
	SessionsEstablish       bool `json:"SessionsEstablish"`
	SessionsExecuteFunction bool `json:"SessionsExecuteFunction"`
}

var CurrentFeatureDetection FeatureDetection = FeatureDetection{
	Type: "FeatureDetection",
	Interfaces: FeatureInterface{
		Collections: CollectionsFeatures{
			CollectionsQuery:  false,
			CollectionsWrite:  false,
			CollectionsCommit: false,
			CollectionsDelete: false,
		},
		Actions: ActionsFeatures{
			ThreadsQuery:  false,
			ThreadsCreate: false,
			ThreadsReply:  false,
			ThreadsClose:  false,
			ThreadsDelete: false,
		},
		Permissions: PermissionsFeatures{
			PermissionsRequest: false,
			PermissionsGrant:   false,
			PermissionsRevoke:  false,
		},
		Messaging: MessagingFeatures{
			Batching: true,
		},
		FairX: FairXFeatures{
			SessionsEstablish:       true,
			SessionsExecuteFunction: true,
		},
	},
}
