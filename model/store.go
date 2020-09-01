package model

import (
	"giligili/constbase"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

type Store struct {
	SID int `json:"s_id" gorm:"column:s_id;type:int(10) unsigned auto_increment;not null;primary_key;comment:'商店道具ID'"`
	Title string `json:"title" gorm:"column:title;type:varchar(50);default:'';not null;comment:'标题'"`
	PID int `json:"p_id" gorm:"column:p_id;type:int(10) unsigned;not null;default:0;comment:'道具ID 来自 prop 表的 p_id'"`
	Gold int `json:"gold" gorm:"column:gold;type:int(10) unsigned;not null;default:0;comment:'道具所需的金币价值'"`
	Diamond int `json:"diamond" gorm:"column:diamond;type:int(10) unsigned;not null;default:0;comment:'道具所需的钻石价值'"`
	PorpDetail *Prop `json:"prop" comment:"道具详情"`
	Sell int `json:"-" gorm:"column:sell;type:int(10) unsigned;not null;default:0;comment:'卖出数量'"`
	Sort int `json:"-" gorm:"column:sort;type:int(10) unsigned;not null;default:0;comment:'排序'"`
	IsShelf int8 `json:"-" gorm:"column:is_shelf;type:tinyint(1) unsigned;not null;default:0;comment:'是否上架 0 - 下架 1 - 上架'"`
	DelAt time.Time `json:"-" gorm:"type:datetime;not null;default:'1000-01-01 00:00:00'; comment:'删除时间'"`
	CreatedAt time.Time `json:"-" gorm:"type:datetime;not null; comment:'创建时间'"`
}

// 获取商店列表
func GetStoreList() []Store {
	var stores []Store

	err := DB.Where("del_at = ? AND is_shelf = ?", DelAtDefault, constbase.YES).Order("sort desc, s_id desc").Find(&stores).Error
	if err != nil {
		log.Println(err.Error())
		return nil
	}

	if len(stores) == 0 {
		return nil
	}

	for index, store := range(stores) {
		store.PorpDetail = GetPropInfo(store.PID)

		stores[index] = store
	}

	return stores
}


func GetSroteInfo(s_id int) *Store {
	if s_id == 0 {
		return nil
	}

	store := &Store{}

	err := DB.Where("s_id = ?", s_id).First(store).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("商店ID(%d)找不到记录: ", s_id)
		} else {
			log.Println(err.Error())
		}

		return nil
	}

	store.PorpDetail = GetPropInfo(store.PID)

	return store
}

// 卖出个数 + 1
func (store Store) SellIncr() bool{
	res := DB.Model(store).Where("sell = ?", store.Sell).Update("sell", store.Sell + 1)
	if res.RowsAffected == 0 {
		return false
	}

	return true
}

// 兑换道具 有金币优先用金币兑换，没有则用钻石兑换
func (store Store) Change() bool {
	if store.IsShelf != constbase.YES {
		log.Printf("商店(%d) 已下架", store.SID)
		return false
	}

	if store.DelAt != DelAtDefaultTime {
		log.Printf("商店(%d) 已删除", store.SID)
		return false
	}

	user := GetUserInfo(UserInfo.UID)

	if store.Gold > 0 {
		if user.Gold < store.Gold {
			log.Printf("用户(%d) 兑换商店物品(%d) 金币不足", UserInfo.UID, store.SID)
			return false
		}

		res := DB.Model(&user).Where("gold = ?", user.Gold).Update("gold", user.Gold - store.Gold)
		if res.RowsAffected == 0 {
			log.Println("更新数据失败")
			return false
		}

	} else if store.Diamond > 0 {
		if user.Diamond < store.Diamond {
			log.Printf("用户(%d) 兑换商店物品(%d) 钻石不足", UserInfo.UID, store.SID)
			return false
		}

		res := DB.Model(&user).Where("diamond = ?", user.Diamond).Update("diamond", user.Diamond - store.Diamond)
		if res.RowsAffected == 0 {
			log.Println("更新数据失败")
			return false
		}
	} else {
		log.Printf("商店(%d) 金币和钻石兑换不能为0", store.SID)
		return false
	}

	b := store.SellIncr()
	if b == false {
		return false
	}

	b = store.PorpDetail.AddToBackpack()
	if b == false {
		return false
	}

	return true
}
