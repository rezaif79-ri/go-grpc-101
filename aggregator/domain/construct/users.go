package construct

type GreetUserReq struct {
	Name       string `json:"name"`
	Salutation string `json:"salutation"`
}

type GreetUserRes struct {
	Greetings string `json:"greetings"`
}
