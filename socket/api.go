package socket

func Api() {
	AddRoute("sign_in/create", HandlerFunc(SignInCreate))
	AddRoute("level/get-info", HandlerFunc(GetLevelInfo))
	AddRoute("level/get-result", HandlerFunc(GetLevelResult))
	AddRoute("user", HandlerFunc(GetUserInfo))
	AddRoute("user/plan/change", HandlerFunc(UserPlanChange))
}
