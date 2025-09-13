package auths

type SignUpController struct {
	FirstName	string	`json:"first-name"`
	LastName  string `json:"last-name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type Response struct {
	Message  string      `json:"message"`
	Status   bool        `json:"status"`
	TryLater string      `json:"try_later,omitempty"`
	Data     interface{} `json:"data,omitempty"`
}

type ResponseUserData struct {
	FirstName	string	`json:"first-name"`
	LastName  string `json:"last-name"`
	Email     string `json:"email"`
}

type LoginController struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
}