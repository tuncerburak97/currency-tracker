package currency

type GoldResponse struct {
	ID                int     `json:"Id"`
	Code              string  `json:"Kod"`
	Description       string  `json:"Aciklama"`
	BuyPrice          string  `json:"Alis"`
	SellPrice         string  `json:"Satis"`
	LastUpdateTime    string  `json:"GuncellenmeZamani"`
	Status            *string `json:"Durum"`
	Main              bool    `json:"Main"`
	DataGroup         int     `json:"DataGroup"`
	Change            float64 `json:"Change"`
	MobileDescription string  `json:"MobilAciklama"`
	WebGroup          *int    `json:"WebGroup"`
	WidgetDescription string  `json:"WidgetAciklama"`
}

type GetCurrencyResponse struct {
	ID                int     `json:"Id"`
	Code              string  `json:"Kod"`
	Description       string  `json:"Aciklama"`
	BuyPrice          string  `json:"Alis"`
	SellPrice         string  `json:"Satis"`
	LastUpdateTime    string  `json:"GuncellenmeZamani"`
	Status            *string `json:"Durum"`
	Main              bool    `json:"Main"`
	DataGroup         int     `json:"DataGroup"`
	Change            float64 `json:"Change"`
	MobileDescription string  `json:"MobilAciklama"`
	WebGroup          *int    `json:"WebGroup"`
	WidgetDescription string  `json:"WidgetAciklama"`
}
