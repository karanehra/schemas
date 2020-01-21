package schemas

//Test defines a test struct
type Test struct {
	Field string
}

//Process defines a task to be sent to a processor
type Process struct {
	Name   string `json:"processName"`
	Status string `json:"status"`
}
