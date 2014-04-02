package model

import (
)

type PlayerInfoVo struct {
	InstanceId	int64
	Id        	int64
	LastScene  	int64		`db:"last_scene"`
	X						int64		`db:"last_x"`
	Y						int64		`db:"last_y"`
}

var idMax int64 = 1000

func CreatePlayerVo() *PlayerInfoVo{
	idMax++
	return &PlayerInfoVo{InstanceId: idMax}
}
