package handlers

type ConsigneeHandler struct {
	ConsigneeRepo models.ConsigneeInterface
}

func NewConsigneeHandler(ConsigneeRepo models.ConsigneeInterface) *ConsigneeHandler {
	return &ConsigneeHandler{
		ConsigneeRepo,
	}
}
func (ch *ConsigneeHandler) CreateConsigneeHandler(w http.ResponseWriter, r *http.Request) {
	res, err := ch.ConsigneeRepo.CreateConsignee(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "consignee created successfully"})
}

func (ch *ConsigneeHandler) DeleteConsigneeHandler(w http.ResponseWriter, r *http.Request) {
	res, err := ch.ConsigneeRepo.DeleteConsignee(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Consignee deleted successfully"})
}

func (ch *ConsigneeHandler) GetConsigneeHandler(w http.ResponseWriter, r *http.Request) {
	res, err := ch.ConsigneeRepo.GetConsignee(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"message": err.Error()})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Consignee fetched successfully"})
}
