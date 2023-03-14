package mocks

import "api-go-example/internal/client"

var (
	ClientHttpGETResponseOK = `{
	"stringValue": "sunny day",
	"intValue": 10
}`

	ClientHttpGETResponseError = `<Body>
	<stringValue> "sunny day",</stringValue>
	<intValue> 10</intValue>
</Body>`

	ClientGetRequestOK = &client.CustomGetRequest{
		PathVar:   "123456789",
		QueryVar1: "yes",
		QueryVar2: "10",
		HeaderVar: "nonerror",
	}
)

var (
	ClientHttpPOSTResponseOK = `{
		"results": [
			{
				"dataPoint1": 22,
				"dataPoint2": 33,
				"dataPoint3": 44
			},
			{
				"dataPoint1": 52,
				"dataPoint2": 63,
				"dataPoint3": 74
			}
		],
		"notes": "no errors found"
	}`

	ClientHttpPOSTResponseError = `{
		error json
		`

	ClientPostRequestOK = &client.CustomPostRequest{
		StringValue: "numbers",
		IntValue:    16,
	}
)
