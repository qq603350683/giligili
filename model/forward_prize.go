package model

import (
	"giligili/constbase"
	"github.com/jinzhu/gorm"
	"log"
)

type ForwardPrize struct {
	FpID int `json:"fp_id" gorm:"column:fp_id;type:int(10) unsigned auto_increment;not null;primary_key;comment:'转发奖品ID'"`
	PID int `json:"p_id" gorm:"column:p_id;type:int(10) unsigned;not null;default:0;comment:'道具ID 来自 prop 表的 p_id'"`
	PropDetail *Prop `json:"prop" comment:"道具详情"`
	IsOpen int8 `json:"is_open" gorm:"column:is_open;type:tinyint(3) unsigned;not null;default:0;comment'是否开启分享奖励 0 - 不开启 1 - 开启'"`
}

// 获取转发奖励
func GetForwardPrize() *Prop {
	fp := &ForwardPrize{}
	err := DB.Where("fp_id = 1").First(fp).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("forward_prize 转发奖品记录找不到")
		} else {
			log.Printf(err.Error())
		}
		return nil
	}

	if fp.IsOpen == constbase.NO {
		return nil
	}

	prop := &Prop{}
	err = DB.Where("p_id = ?", fp.PID).First(prop).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("prop 道具记录找不到")
		} else {
			log.Printf(err.Error())
		}
		return nil
	}

	fp.PropDetail = prop

	return prop
}



