package salesforceclient

type CreateCaseResponse struct {
	Id      string        `json:"id"`
	Success bool          `json:"success"`
	Errors  []interface{} `json:"errors"`
}
