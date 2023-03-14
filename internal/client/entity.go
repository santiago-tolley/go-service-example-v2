package client

type CustomGetRequest struct {
	PathVar   string
	QueryVar1 string
	QueryVar2 string
	HeaderVar string
}

type CustomGetResponse struct {
	StringValue string `json:"stringValue"`
	IntValue    int32  `json:"intValue"`
}

type CustomPostRequest struct {
	StringValue string `json:"stringValue"`
	IntValue    int32  `json:"intValue"`
}

type CustomPostResponse struct {
	Results []Result `json:"results"`
	Notes   string   `json:"notes"`
}

type Result struct {
	DataPoint1 int32 `json:"dataPoint1"`
	DataPoint2 int32 `json:"dataPoint2"`
	DataPoint3 int32 `json:"dataPoint3"`
}
