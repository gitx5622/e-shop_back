package Forms


type StkPushForm struct {
	PartyA   string `gorm:"not_null" binding:"required" json:"party_a"`
	Amount   string `gorm:"not_null" binding:"required" json:"amount"`
	AuthorID uint32 `gorm:"not null" json:"author_id"`
}


//var StkPushForm Forms.StkPushForm
//err = c.BindJSON(&StkPushForm)
//if err != nil {
//	appError := err
//	appError.Error()
//	return
//}
//
//m := &StkPushForm