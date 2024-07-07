package utils

type Response struct {
	Message string `json:"message"`
}

type Result struct {
	Data interface{} `json:"data"`
}

type Errors struct {
	Errors []string `json:"errors"`
}

type Error struct {
	Error []string `json:"error"`
}

type DeleteStatus struct {
	Deleted bool `json:"deleted"`
}
