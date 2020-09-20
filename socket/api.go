package socket

func Api() {
	AddRoute("sign_in/create", HandlerFunc(SignInCreate))
}
