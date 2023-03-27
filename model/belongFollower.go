/**
  @author: cilang
  @qq: 1019383856
  @bili: https://space.bilibili.com/433915419
  @gitee: https://gitee.com/OpencvLZG
  @since: 3/25/2023
  @desc: //TODO
**/

package model

import "time"

type BelongFollower struct {
	Id          int64     `json:"id" xorm:"autoincr pk"`
	BelongId    int64     `json:"belong_id" xorm:"not null"`
	Follower    int64     `json:"follower" xorm:"not null"`
	CreatedTime time.Time `json:"created_time" xorm:"created"`
	UpdateTime  time.Time `json:"update_time" xorm:"updated"`
}
