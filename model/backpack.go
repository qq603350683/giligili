package model

import (
	"encoding/json"
	"giligili/constbase"
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"time"
)

// 背包道具
type Backpack struct {
	BID int `json:"-" gorm:"column:b_id;type:int(10) unsigned auto_increment;not null;primary_key;comment:'背包物品主键ID'"`
	UID int `json:"u_id" gorm:"column:u_id;type:int(10) unsigned; not null;default:0;index:idx_u_id;comment:'用户ID 来自 users 表的 u_id'"`
	PID int `json:"p_id" gorm:"column:p_id;type:int(10) unsigned;not null;default:0;comment:'道具ID 来自 prop 表的 p_id'"`
	PropDetail *Prop `json:"prop_detail" comment:"道具详情"`
	Quantity int `json:"quantity" gorm:"-" comment:"数量"`
	//Type string `json:"type" gorm:"column:type;type:enum('bullet_enhancer','bullet_speed_enhancer', 'skill_enhancer', 'skill_speed_enhancer');not null;comment:'道具分类'"`
	//Title string `json:"title" gorm:"column:title;type:varchar(50);not null;comment:'标题'"`
	//Remark string `json:"-" gorm:"column:remark;type:varchar(50);not null;comment:'备注说明、领取途径等'"`
	IsUse int8 `json:"-" gorm:"column:is_use;type:int(1);not null;default:0;comment:'是否已使用 0 - 未使用 1 - 已使用'"`
	IsSell int8 `json:"-" gorm:"column:is_sell;type:int(1);not null;default:0;comment:'是否已出售 0 - 未出售 1 - 已出售'"`
	UseAt time.Time `json:"-" gorm:"type:datetime;not null;default:'1000-01-01 00:00:00'; comment:'使用时间'"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at; type:datetime; not null; comment:'创建时间'"`
}

type PropUse struct {
	PID int `json:"p_id" comment:"道具ID"`
	UpID int `json:"up_id" comment:"飞机ID"`
	BID int `json:"b_id" comment:"子弹ID"`
	SID int `json:"s_id" comment:"技能ID"`
}

func NewBackpack() *Backpack {
	return &Backpack{
		UseAt:      time.Time{},
		CreatedAt:  time.Time{},
	}
}

func NewPropUse() *PropUse {
	return &PropUse{}
}

func GetMyBackpackInfo(p_id int) *Backpack {
	backpack := &Backpack{}

	err := DB.Where("u_id = ? AND p_id = ? AND is_use = 0", UserInfo.UID, p_id).First(backpack).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Println(err.Error())
		}
		return nil
	}

	return backpack
}

func GetBackpacks(u_id int) []Backpack {
	var backpacks []Backpack

	err := DB.Where("u_id = ? AND is_use = 0", u_id).Limit(1).Find(&backpacks).Error
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	err = DB.Raw("SELECT *, COUNT(*) as quantity FROM backpacks WHERE u_id = ? AND is_use = 0 GROUP BY p_id", u_id).Find(&backpacks).Error
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if len(backpacks) == 0 {
		return nil
	}

	for index, backpack := range(backpacks) {
		backpack.PropDetail = GetPropInfo(backpack.PID)

		backpacks[index] = backpack
	}

	return backpacks
}


func (backpack *Backpack) Use() bool {
	if backpack.IsUse == constbase.YES {
		return false
	}

	if backpack.PropDetail == nil {
		backpack.PropDetail = GetPropInfo(backpack.PID)
	}

	backpack.IsUse = constbase.YES
	backpack.UseAt = time.Time{}

	err := DB.Save(backpack).Error
	if err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}

// 出售
func (backpack *Backpack) Sell() bool {
	if backpack.IsUse == constbase.YES {
		return false
	}

	if backpack.IsSell == constbase.YES {
		return false
	}

	res := DB.Model(backpack).Where("is_use = 0 AND is_sell = 0").UpdateColumns(map[string]interface{} {
		"is_use": constbase.YES,
		"is_sell": constbase.YES,
		"use_at": time.Now(),
	})
	if res.RowsAffected == 0 {
		return false
	}

	return true
}

// 打开金币礼包
func (backpack *Backpack) OpenGoldPack() (int, bool) {
	if backpack.IsUse == constbase.YES {
		return 0, false
	}

	var b bool
	var quantity int

	if backpack.PropDetail.MinQuantity == backpack.PropDetail.MaxQuantity {
		// 固定金币金额
		quantity = backpack.PropDetail.MinQuantity
		b = backpack.PropDetail.AddToUserGold(backpack.PropDetail.MinQuantity)
	} else {
		// 随机金币金额
		rand.Seed(time.Now().Unix())
		quantity = rand.Intn(backpack.PropDetail.MaxQuantity - backpack.PropDetail.MinQuantity) + backpack.PropDetail.MinQuantity

		b = backpack.PropDetail.AddToUserGold(quantity)
	}

	if b == false {
		return 0, false
	}

	backpack.IsUse = constbase.YES
	backpack.UseAt = time.Now()

	err := DB.Save(backpack).Error
	if err != nil {
		log.Println(err.Error())
		return 0, false
	}

	return quantity, true
}

// 打开钻石大礼包
func (backpack *Backpack) OpenDiamondPack() (int, bool) {
	if backpack.IsUse == constbase.YES {
		return 0, false
	}

	var b bool
	var quantity int

	if backpack.PropDetail.MinQuantity == backpack.PropDetail.MaxQuantity {
		// 固定金币金额
		quantity = backpack.PropDetail.MinQuantity
		b = backpack.PropDetail.AddToUserDiamond(backpack.PropDetail.MinQuantity)
	} else {
		// 随机金币金额
		rand.Seed(time.Now().Unix())
		quantity = rand.Intn(backpack.PropDetail.MaxQuantity - backpack.PropDetail.MinQuantity) + backpack.PropDetail.MinQuantity

		b = backpack.PropDetail.AddToUserDiamond(quantity)
	}

	if b == false {
		return 0, false
	}

	backpack.IsUse = constbase.YES
	backpack.UseAt = time.Now()

	err := DB.Save(backpack).Error
	if err != nil {
		log.Println(err.Error())
		return 0, false
	}

	return quantity, true
}

// 子弹强化器
func (backpack *Backpack) UseBulletEnhancer(up_id int, id int) (bool, bool) {
	enhancer_result := false

	plan := GetUserPlanInfo(up_id)
	if plan == nil {
		return enhancer_result, false
	}

	if plan.UID != UserInfo.UID {
		return enhancer_result, false
	}

	index := -1
	bullet := Bullet{}

	for i, b := range(plan.Detail.Bullets) {
		if b.BID == id {
			index = i
			bullet = b
			break
		}
	}

	if bullet.BID == 0 {
		log.Println("请选择需要强化的 bullet: b_id 不能为0")
		return enhancer_result, false
	}

	if index == -1 {
		log.Println("请选择需要强化的 bullet: index 不能为 -1")
		return enhancer_result, false
	}

	// 已达到最高强化级别
	if plan.Detail.Bullets[index].MaxLevel <= plan.Detail.Bullets[index].Level {
		return enhancer_result, false
	}

	enhancer_result = GetBulletEnhancerIsSuccess(backpack.PropDetail.Type, bullet.Level)

	if enhancer_result == true {
		// 强化成功
		bullet.Level += 1

		plan.Detail.Bullets[index] = bullet

		str, err := json.Marshal(plan.Detail)
		if err != nil {
			log.Println(err.Error())
		}

		plan.DetailJson = string(str)

		err = DB.Save(plan).Error
		if err != nil {
			log.Println(err.Error())
			return enhancer_result, false
		}
	}

	return enhancer_result, true
}

// 子弹速度强化器
func (backpack *Backpack) UseBulletSpeedEnhancer(up_id int, id int) (bool, bool) {
	enhancer_result := false

	plan := GetUserPlanInfo(up_id)
	if plan == nil {
		return enhancer_result, false
	}

	if plan.UID != UserInfo.UID {
		return enhancer_result, false
	}

	index := -1
	bullet := Bullet{}

	for i, b := range(plan.Detail.Bullets) {
		if b.BID == id {
			index = i
			bullet = b
			break
		}
	}

	if bullet.BID == 0 {
		return enhancer_result, false
	}

	if index == -1 {
		return enhancer_result, false
	}

	enhancer_result = GetSpeedEnhancerIsSuccess(backpack.PropDetail.Type, bullet.Speed)

	// 已达到最高强化级别
	if plan.Detail.Bullets[index].MaxSpeed <= plan.Detail.Bullets[index].Speed {
		return enhancer_result, false
	}

	if enhancer_result == true {
		// 强化成功
		bullet.Speed += 1

		plan.Detail.Bullets[index] = bullet

		str, err := json.Marshal(plan.Detail)
		if err != nil {
			log.Println(err.Error())
		}

		plan.DetailJson = string(str)

		err = DB.Save(plan).Error
		if err != nil {
			log.Println(err.Error())
			return enhancer_result, false
		}
	}

	return enhancer_result, true
}

// 子弹频率强化器
func (backpack *Backpack) UseBulletRateEnhancer(up_id int, id int) (bool, bool) {
	enhancer_result := false

	plan := GetUserPlanInfo(up_id)
	if plan == nil {
		return enhancer_result, false
	}

	if plan.UID != UserInfo.UID {
		return enhancer_result, false
	}

	index := -1
	bullet := Bullet{}

	for i, b := range(plan.Detail.Bullets) {
		if b.BID == id {
			index = i
			bullet = b
			break
		}
	}

	if bullet.BID == 0 {
		return enhancer_result, false
	}

	if index == -1 {
		return enhancer_result, false
	}

	// 已达到最高强化级别
	if plan.Detail.Bullets[index].MaxRate >= plan.Detail.Bullets[index].Rate {
		return enhancer_result, false
	}

	enhancer_result = GetSpeedEnhancerIsSuccess(backpack.PropDetail.Type, bullet.Speed)

	if enhancer_result == true {
		// 强化成功
		bullet.Rate -= 1

		plan.Detail.Bullets[index] = bullet

		str, err := json.Marshal(plan.Detail)
		if err != nil {
			log.Println(err.Error())
		}

		plan.DetailJson = string(str)

		err = DB.Save(plan).Error
		if err != nil {
			log.Println(err.Error())
			return enhancer_result, false
		}
	}

	return enhancer_result, true
}

// 技能强化器
func (backpack *Backpack) UseSkillEnhancer(up_id int, id int) (bool, bool) {
	enhancer_result := false

	plan := GetUserPlanInfo(up_id)
	if plan == nil {
		return enhancer_result, false
	}

	if plan.UID != UserInfo.UID {
		log.Printf("您(u_id: %d)没有权限操作(up_id: %d)", UserInfo.UID, up_id)
		return enhancer_result, false
	}

	index := -1
	skill := Skill{}

	for i, b := range(plan.Detail.Skills) {
		if b.SID == id {
			index = i
			skill = b
			break
		}
	}

	if skill.SID == 0 {
		log.Println("请选择需要强化的 skill: s_id 不能为0")
		return enhancer_result, false
	}

	if index == -1 {
		log.Println("请选择需要强化的 skill: index 不能为 -1")
		return enhancer_result, false
	}

	// 已达到最高强化级别
	if plan.Detail.Skills[index].MaxLevel <= plan.Detail.Skills[index].Level {
		return enhancer_result, false
	}

	enhancer_result = GetBulletEnhancerIsSuccess(backpack.PropDetail.Type, skill.Level)

	if enhancer_result == true {
		// 强化成功
		skill.Level += 1

		plan.Detail.Skills[index] = skill

		str, err := json.Marshal(plan.Detail)
		if err != nil {
			log.Println(err.Error())
		}

		plan.DetailJson = string(str)

		err = DB.Save(plan).Error
		if err != nil {
			log.Println(err.Error())
			return enhancer_result, false
		}
	}

	return enhancer_result, true
}

// 技能速度强化器
func (backpack *Backpack) UseSkillSpeedEnhancer(up_id int, id int) (bool, bool) {
	enhancer_result := false

	plan := GetUserPlanInfo(up_id)
	if plan == nil {
		return enhancer_result, false
	}

	if plan.UID != UserInfo.UID {
		return enhancer_result, false
	}

	index := -1
	skill := Skill{}

	for i, b := range(plan.Detail.Skills) {
		if b.SID == id {
			index = i
			skill = b
			break
		}
	}

	if skill.SID == 0 {
		return enhancer_result, false
	}

	if index == -1 {
		return enhancer_result, false
	}

	// 已达到最高强化级别
	if plan.Detail.Skills[index].MaxSpeed <= plan.Detail.Skills[index].Speed {
		return enhancer_result, false
	}

	enhancer_result = GetSpeedEnhancerIsSuccess(backpack.PropDetail.Type, skill.Speed)

	if enhancer_result == true {
		// 强化成功
		skill.Speed += 1

		plan.Detail.Skills[index] = skill

		str, err := json.Marshal(plan.Detail)
		if err != nil {
			log.Println(err.Error())
		}

		plan.DetailJson = string(str)

		err = DB.Save(plan).Error
		if err != nil {
			log.Println(err.Error())
			return enhancer_result, false
		}
	}

	return enhancer_result, true
}

// 子弹射频强化器
func (backpack *Backpack) UseSkillRateEnhancer(up_id int, id int) (bool, bool) {
	enhancer_result := false

	plan := GetUserPlanInfo(up_id)
	if plan == nil {
		return enhancer_result, false
	}

	if plan.UID != UserInfo.UID {
		return enhancer_result, false
	}

	index := -1
	skill := Skill{}

	for i, b := range(plan.Detail.Skills) {
		if b.SID == id {
			index = i
			skill = b
			break
		}
	}

	if skill.SID == 0 {
		return enhancer_result, false
	}

	if index == -1 {
		return enhancer_result, false
	}

	// 已达到最高强化级别
	if plan.Detail.Skills[index].MaxRate >= plan.Detail.Skills[index].Rate {
		return enhancer_result, false
	}

	enhancer_result = GetSpeedEnhancerIsSuccess(backpack.PropDetail.Type, skill.Speed)

	if enhancer_result == true {
		// 强化成功
		skill.Rate -= 1

		plan.Detail.Skills[index] = skill

		str, err := json.Marshal(plan.Detail)
		if err != nil {
			log.Println(err.Error())
		}

		plan.DetailJson = string(str)

		err = DB.Save(plan).Error
		if err != nil {
			log.Println(err.Error())
			return enhancer_result, false
		}
	}

	return enhancer_result, true
}