package socket

func Api() {
	AddRoute("sign_in/create", HandlerFunc(SignInCreate))
	AddRoute("share/create", HandlerFunc(ShareCreate))
	AddRoute("level/get-info", HandlerFunc(GetLevelInfo))
	AddRoute("level/get-result", HandlerFunc(GetLevelResult))
	AddRoute("user", HandlerFunc(GetUserInfo))
	AddRoute("user/plan/change", HandlerFunc(UserPlanChange))

	AddRoute("backpack/get-list", HandlerFunc(GetBackpackList))
	AddRoute("backpack/prop/use", HandlerFunc(BackpackPropUse))

	AddRoute("backpack/prop/sell", HandlerFunc(BackpackPropSell))

	AddRoute("store/get-list", HandlerFunc(GetStoreList))
	AddRoute("store/buy", HandlerFunc(StoreBuy))


}
