package socket

import (
	"giligili/constbase"
	"giligili/model"
	"giligili/serializer"
	"giligili/util"
	"net/http"
)

func UserPlanChange(params Params) {
	if _, ok := params["up_id"]; !ok {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "参数异常", nil, "up_id 为 0"))
		return
	}

	up_id := util.StringToInt(params["up_id"])
	if up_id == 0 {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(http.StatusInternalServerError, "系统繁忙", nil, "up_id 为 0"))
		return
	}

	if model.UserInfo.UpID == up_id {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.USER_PLAN_CHANGE_SUCCESS, "更换成功", model.UserInfo.Plan, ""))
		return
	}

	current_user_plan := model.NewUserPlan()
	if (model.UserInfo.UpID > 0) {
		current_user_plan = model.GetUserPlanInfo(model.UserInfo.UpID)
		if current_user_plan == nil {
			SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.USER_PLAN_CHANGE_FAIL, "战斗机更换失败", nil, "up_id 异常"))
			return
		}
	}

	user_plan := model.GetUserPlanInfo(up_id)
	if user_plan == nil {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.USER_PLAN_CHANGE_FAIL, "战斗机更换失败", nil, "up_id 异常"))
		return
	}

	if user_plan.UID != model.UserInfo.UID {
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.USER_PLAN_CHANGE_FAIL, "您没有权限更换此战斗机", nil, "up_id 异常"))
		return
	}

	db := model.DBBegin()

	defer model.CancelDB()

	res := db.Model(user_plan).Update("is_put_on", constbase.YES)
	if res.RowsAffected == 0 {
		db.Rollback()
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.USER_PLAN_CHANGE_FAIL, "更换失败", nil, "UserPlanChange 更新数据失败001"))
		return
	}

	if (model.UserInfo.UpID > 0) {
		res = db.Model(current_user_plan).Update("is_put_on", constbase.NO)
		if res.RowsAffected == 0 {
			db.Rollback()
			SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.USER_PLAN_CHANGE_FAIL, "更换失败", nil, "UserPlanChange 更新数据失败002"))
			return
		}
	}

	res = db.Model(model.UserInfo).Update("up_id", up_id)
	if res.RowsAffected == 0 {
		db.Rollback()
		SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.USER_PLAN_CHANGE_FAIL, "更换失败", nil, "UserPlanChange 更新数据失败003"))
		return
	}

	db.Commit()

	model.UserInfo.Plan = user_plan

	SendMessage(model.UserInfo.UID, serializer.JsonByte(constbase.USER_PLAN_CHANGE_SUCCESS, "更换成功", model.UserInfo.Plan, ""))
}
