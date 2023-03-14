package entity

type GetRequest struct {
	ID   string
	Info string
	Data GetData
}

type GetData struct {
	Value  string `json:"value"` // TODO add validations
	IValue int32  `json:"ivalue"`
}

type GetResponse struct {
	Str     string `json:"string"`
	Integer int32  `json:"int"`
}

type PostRequest struct {
	Str     string `json:"string"`
	Integer int32  `json:"int"`
}

type PostResponse struct {
	DataPoints map[int][]int32 `json:"dataPoints"`
}
