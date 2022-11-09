package salesforceclient

type CreateCaseRequest struct {
	Subject     string `json:"Subject"`
	Description string `json:"Description"`
	Priority    string `json:"Priority"`
}
