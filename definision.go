package main

type FormatResp struct {
	Status     int    `json:"status"`
	ErrMessage string `json:"errorMessage"`
	Result     Result `json:"result"`
}

type Result struct {
	StdOut string `json:"stdOut"`
	StdErr string `json:"stdErr"`
}

func initializeResult() Result {
	r := Result {
		StdOut: "",
		StdErr: "",
	}
	return r
}