package app

type Task struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
}

type Endpoint struct {
	Path   string
	Method string

	Service *MainService
}

type Response struct {
	Message string
}
