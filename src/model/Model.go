package model

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/ziutek/mymysql/godrv"
	"database/sql"
	"github.com/coopernurse/gorp"
	"fmt"
)

type PlayerModel struct{
	playerInfoMap		map[int64]*PlayerInfoVo
	dbmap						*gorp.DbMap
}

var playerModelInstance *PlayerModel

func SharedPlayerModel() *PlayerModel {
	if playerModelInstance == nil {
		playerModelInstance = &PlayerModel{}
		playerModelInstance.init()
	}
	return playerModelInstance
}

func (this *PlayerModel) init() {
	db, _ := sql.Open("mymysql", "tcp:localhost:3306*muiltcube/root/root")
	this.dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	this.playerInfoMap = make(map[int64]*PlayerInfoVo)
}

func (this *PlayerModel) GerPlayerInfo(id int64)*PlayerInfoVo {
	playerInfo := this.playerInfoMap[id]
	if(playerInfo == nil){
		playerInfo = CreatePlayerVo()
		this.dbmap.SelectOne(playerInfo, "select * from player_info where id=?", id)
		this.playerInfoMap[id] = playerInfo;
	}

	fmt.Println("playerInfoMap:", playerInfo)

	return playerInfo
}

func (this *PlayerModel) Login(id int64, pwd string) bool{
	count, _ := this.dbmap.SelectInt("select count(*) from password where player_id=? and password =?", id, pwd)

	return count > 0
}

func (this *PlayerModel) HadSkin(id int64) bool{
	count, _ := this.dbmap.SelectInt("select count(*) from skin where player_id=?", id)

	return count > 0
}
