package entities

type Quote struct {
	Quote    string `json:"quote" binding:"required,max=1000"`
	Language string `json:"language" binding:"required,len=2" validate:"country-code"`
}
