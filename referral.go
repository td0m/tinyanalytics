package model

type Referral struct {
	From  Page `json:"from"`
	Count int  `json:"count"`
}
