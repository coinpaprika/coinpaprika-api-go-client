package coinpaprika

// ICO represents an ICO project.
type ICO struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	IsNew  bool   `json:"is_new"`
}
