package albums

type Album struct {
    ID     string  `json:"id"`
    Title  string  `json:"Title"`
    Artist string  `json:"Artist"`
    Price  float64 `json:"Price"`
}
