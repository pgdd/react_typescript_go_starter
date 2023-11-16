package models

type Transfer struct {
	SourceAccount      Account
	DestinationAccount Account
	Amount             float64
	// Other transfer-related fields
}
