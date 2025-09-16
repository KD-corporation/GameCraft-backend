package auths


type SignUpController struct {
	Username string `json:"Username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
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
	Username string `json:"Username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Email     string `json:"email"`
}

type LoginController struct {
	Id    string `json:"Id"`
	Password string `json:"password"`
}

type OtpController struct {
	Email    		string 		`json:"email"`
	Otp    			string 		`json:"otp"`
}