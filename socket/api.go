package socket

func Api() {
	AddRoute("watch_adv", HandlerFunc(WatchAdv))
	AddRoute("auto_commission", HandlerFunc(AutoCommission))
	AddRoute("withdraw", HandlerFunc(Withdraw))

	AddRoute("sign_in/create", HandlerFunc(SignInCreate))
	AddRoute("share/create", HandlerFunc(ShareCreate))
	AddRoute("level/get-info", HandlerFunc(GetLevelInfo))
	AddRoute("level/get-result", HandlerFunc(GetLevelResult))
	AddRoute("user", HandlerFunc(GetUserInfo))
	AddRoute("user/plan/change", HandlerFunc(UserPlanChange))

	AddRoute("user/plan/upgrade/bullets/level", HandlerFunc(UserPlanUpgradeBulletsLevel))
	AddRoute("user/plan/upgrade/bullets/speed", HandlerFunc(UserPlanUpgradeBulletsSpeed))
	AddRoute("user/plan/upgrade/bullets/rate", HandlerFunc(UserPlanUpgradeBulletsRate))

	AddRoute("user/plan/upgrade/bullet/level", HandlerFunc(UserPlanUpgradeBulletLevel))
	AddRoute("user/plan/upgrade/bullet/speed", HandlerFunc(UserPlanUpgradeBulletSpeed))
	AddRoute("user/plan/upgrade/bullet/rate", HandlerFunc(UserPlanUpgradeBulletRate))

	AddRoute("user/plan/upgrade/skill/level", HandlerFunc(UserPlanUpgradeSkillLevel))
	AddRoute("user/plan/upgrade/skill/speed", HandlerFunc(UserPlanUpgradeSkillSpeed))
	AddRoute("user/plan/upgrade/skill/rate", HandlerFunc(UserPlanUpgradeSkillRate))

	AddRoute("user/plan/upgrade/skills/level", HandlerFunc(UserPlanUpgradeSkillsLevel))
	AddRoute("user/plan/upgrade/skills/speed", HandlerFunc(UserPlanUpgradeSkillsSpeed))
	AddRoute("user/plan/upgrade/skills/rate", HandlerFunc(UserPlanUpgradeSkillsRate))

	AddRoute("backpack/get-list", HandlerFunc(GetBackpackList))
	AddRoute("backpack/prop/use", HandlerFunc(BackpackPropUse))

	AddRoute("backpack/prop/sell", HandlerFunc(BackpackPropSell))

	AddRoute("store/get-list", HandlerFunc(GetStoreList))
	AddRoute("store/buy", HandlerFunc(StoreBuy))


}
