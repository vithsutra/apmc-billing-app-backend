package models

type Consignee struct {
	ConsigneeId          string `json:"consignee_id"`
	ConsigneeName        string `json:"consignee_name" validate:"required,max =50"`
	ConsigneeAddress     string `json:"consignee_address" validate:"required, max=100"`
	ConsigneeGstin       string `json:"consignee_gstin" validate:"require,max=50"`
	ConsigneePhoneNumber string `json:"consignee_phoone_number" validate:"required,max=50"`
	ConsigneeState       string `json:"consignee_state" validate:"required,max=50"`
	ConsigneeStateCode   string `json:"consignee_state_code" validate:"required,max=50"`
	UserId               string `json:"user_id" validate:"required,max=100"`
}
