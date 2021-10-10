package outbond

type FileUpload struct {
	Files []string `json:"files"`
}

type FileUploadResponse struct {
	Code int `json:"code"`
	Status string `json:"status"`
	Data []string `json:"data"`
}