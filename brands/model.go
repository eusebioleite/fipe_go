package brands

type Brand struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Fipe        string `json:"fipe"`
	Type_id     int    `json:"type_id"`
	Ref_id      int    `json:"ref_id"`
}
