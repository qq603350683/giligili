package socket

func Api() {
	AddRoute("sign_in/create", HandlerFunc(SignInCreate))
	AddRoute("level/get-info", HandlerFunc(GetLevelInfo))
}
