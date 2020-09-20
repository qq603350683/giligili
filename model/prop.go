package model

import (
	"github.com/jinzhu/gorm"
	"log"
	"math/rand"
	"time"
)

type Prop struct {
	PID int `json:"p_id" gorm:"column:p_id;type:int(10) unsigned auto_increment;not null;primary_key;comment:'道具ID'"`
	Type string `json:"type" gorm:"column:type;type:enum('gold','diamond','stone_enhancer_material','stone_speed_enhancer_material','bullet_enhancer','bullet_speed_enhancer','bullet_rate_enhancer','skill_enhancer','skill_speed_enhancer','skill_rate_enhancer');not null;comment:'道具分类'"`
	Image string `json:"image" gorm:"column:image;type:varchar(100);default:'';not null;comment:'图标'"`
	Title string `json:"title" gorm:"column:title;type:varchar(50);default:'';not null;comment:'标题'"`
	Remark string `json:"-" gorm:"column:remark;type:varchar(50);default:'';not null;comment:'备注说明、领取途径等'"`
	MinQuantity int `json:"min_quantity" gorm:"column:min_quantity;type:int(10) unsigned;not null;default:0;comment:'随机最低数量'"`
	MaxQuantity int `json:"max_quantity" gorm:"column:max_quantity;type:int(10) unsigned;not null;default:0;comment:'随机最多数量'"`
	GoldValue int `json:"gold_value" gorm:"column:gold_value;type:int(10) unsigned;not null;default:0;comment:'售出所值金币'"`
	DiamondValue int `json:"gold_value" gorm:"column:diamond_value;type:int(10) unsigned;not null;default:0;comment:'售出所值钻石'"`
	CreatedAt time.Time `json:"-" gorm:"column:created_at;type:datetime;not null;comment:'添加时间'"`
}

// 展示在前端的格式
type PropShow struct {
	*Prop `comment:"道具详情"`
	Quantity int `json:"quantity" comment:"数量"`
}

// 获取道具详情
func GetPropInfo(p_id int) *Prop {
	if p_id == 0 {
		return nil
	}

	prop := &Prop{}

	err := DB.Where("p_id = ?", p_id).First(prop).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("道具ID(%d)找不到记录: ", p_id)
		} else {
			log.Println(err.Error())
		}

		return nil
	}

	return prop
}

// 根据类型获取道具详情
func GetPropInfoByType(_type string) *Prop {
	if _type == "" {
		return nil
	}

	prop := &Prop{}

	err := DB.Where("type = ?", _type).First(prop).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("道具Type(%s)找不到记录: ", _type)
		} else {
			log.Println(err.Error())
		}

		return nil
	}

	return prop
}

func GetBulletEnhancerIsSuccess(t string, level int) bool {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(100)

	if t == "" {
		return false
	}

	if level <= 2 {
		return true
	} else if level > 2 && level <= 5 {
		if i <= 80 {
			return true
		} else {
			return false
		}
 	} else if level > 5 && level <= 7 {
 		if i <= 60 {
 			return true
		} else {
			return false
		}
	} else if level == 8 {
		if i <= 50 {
			return true
		} else {
			return false
		}
	} else if level == 9 {
		if i <= 40 {
			return true
		} else {
			return false
		}
	} else if level == 10 {
		if i <= 30 {
			return true
		} else {
			return false
		}
	} else if level == 11 {
		if i <= 25 {
			return true
		} else {
			return false
		}
	} else if level == 12 {
		if i <= 10 {
			return true
		} else {
			return false
		}
	} else if level > 13 && level <= 15 {
		if i <= 5 {
			return true
		} else {
			return false
		}
	} else if level > 16 && level < 20 {
		if i <= 3 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func GetSpeedEnhancerIsSuccess(t string, speed int) bool {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(100)

	if t == "" {
		return false
	}

	if speed <= 5 {
		return true
	} else if speed > 5 && speed <= 10 {
		if i <= 80 {
			return true
		} else {
			return false
		}
	} else if speed > 10 && speed <= 15 {
		if i <= 60 {
			return true
		} else {
			return false
		}
	} else if speed > 15 && speed <= 20 {
		if i <= 40 {
			return true
		} else {
			return false
		}
	} else if speed > 20 && speed <= 23 {
		if i <= 30 {
			return true
		} else {
			return false
		}
	} else if speed == 24 {
		if i <= 20 {
			return true
		} else {
			return false
		}
	} else if speed == 25 {
		if i <= 17 {
			return true
		} else {
			return false
		}
	} else if speed == 26 {
		if i <= 15 {
			return true
		} else {
			return false
		}
	} else if speed == 27 {
		if i <= 13 {
			return true
		} else {
			return false
		}
	} else if speed == 28 {
		if i <= 12 {
			return true
		} else {
			return false
		}
	} else if speed == 29 {
		if i <= 10 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func (prop *Prop) UseDB(db *gorm.DB) *Prop {
	DBTransaction = db
	return prop
}

func CancelDB() {
	if DBTransaction != nil {
		DBTransaction = nil
	}
}

// 道具加入到背包
func (prop *Prop) AddToBackpack() bool {
	backpack := NewBackpack()

	backpack.UID = UserInfo.UID
	backpack.PID = prop.PID

	var db *gorm.DB
	var err error

	if DBTransaction == nil {
		// 这里是普通业务逻辑
		db = DB
	} else{
		// 这里是事务
		db = DBTransaction
	}

	defer CancelDB()

	if err = db.Create(backpack).Error; err != nil {
		log.Printf("用户ID(%d)道具领取失败", UserInfo.UID)
		return false
	}

	return true
}

// 增加金币金额
func (prop *Prop) AddToUserGold(quantity int) bool {
	var db *gorm.DB

	if DBTransaction == nil {
		// 这里是普通业务逻辑
		db = DB
	} else{
		// 这里是事务
		db = DBTransaction
	}

	defer CancelDB()

	res := db.Model(UserInfo).Where("gold = ?", UserInfo.Gold).Update("gold", UserInfo.Gold + quantity)
	if res.RowsAffected == 0 {
		log.Println("更新数据失败")
		return false
	}

	return true
}

// 增加钻石
func (prop *Prop) AddToUserDiamond(quantity int) bool {
	var db *gorm.DB

	if DBTransaction == nil {
		// 这里是普通业务逻辑
		db = DB
	} else{
		// 这里是事务
		db = DBTransaction
	}

	defer CancelDB()

	res := db.Model(UserInfo).Where("diamond = ?", UserInfo.Diamond).Update("diamond", UserInfo.Diamond + quantity)
	if res.RowsAffected == 0 {
		log.Println("更新数据失败")
		return false
	}

	return true
}