package dto

type SearchNewsReq struct {
	Keywords string `json:"keywords,omitempty"`
	Offset int `json:"offset,omitempty"`
	Limit int `json:"limit,omitempty"`
}
