package model

type Referral struct {
	From  Page `json:"from"`
	To    Page `json:"to"`
	Count int  `json:"count"`
}
