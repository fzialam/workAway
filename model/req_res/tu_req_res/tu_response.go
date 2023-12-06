package tureqres

type CreateSPPDResponse struct {
	DokumenName string `json:"dokumen_name"`
	Message     string `json:"message"`
}

type IndexTU struct {
	Iscreated int `json:"iscreated"`
	Approved  int `json:"approved"`
	Reject    int `json:"reject"`
}
