package interfaces

type DataFormat string

const (
	JSON DataFormat = "application/json"
	JWT  DataFormat = "application/vc+jwt"
	LDP  DataFormat = "application/vc+ldp"
)
