package nwayconnection

import (
	"database/sql"
	"fmt"
	//"go_fs/fs_control/conf"
	//"go_fs/fs_control/log"
	"nway/utils/log"

	_ "github.com/lib/pq"

	"sync"
)

type NwayConnect struct {
	NwayDb *sql.DB
}

var _init_db sync.Once
var _instance_ *NwayConnect

func NewDb() *NwayConnect {
	_init_db.Do(func() {
		_instance_ = new(NwayConnect)

	})
	return _instance_
}

func (this *NwayConnect) InitDB(dbengine, dbstring string) bool {
	if dbengine == "postgresql" {
		return this.Init(dbstring)
	} else if dbengine == "mysql" {
		var err error
		this.NwayDb, err = sql.Open("mysql", dbstring)
		this.NwayDb.SetMaxIdleConns(4)
		this.NwayDb.SetMaxOpenConns(10)
		if err != nil {
			//log.Fatalf()
			logger.Error("Got error when connect database, the error is '%v'", err)
			fmt.Println("Got error when connect database, the error is '%v'", err)
			return false
		}

		return true
	}
	return false
}
func (this *NwayConnect) Init(dbstring string) bool {
	//	result := true
	var err error
	//var conf *config.NwayConfig
	//var isload bool
	//conf, isload = config.LoadConf()
	//if isload == true {

	//dbstring := conf.GetConnStr()

	this.NwayDb, err = sql.Open("postgres", dbstring)
	this.NwayDb.SetMaxIdleConns(4)
	this.NwayDb.SetMaxOpenConns(50)
	if err != nil {
		//log.Fatalf()
		logger.Error("Got error when connect database, the error is '%v'", err)
		fmt.Println(dbstring)
		return false
	}

	return true
	//} else {
	//	return false
	//}
}

func (this *NwayConnect) GetConn() *sql.DB {
	//保证所有的连接都是有效的
	strsql := "select now();"
	rows, err := this.NwayDb.Query(strsql)
	if err != nil {
		logger.Error("failed to check_user：  ", err)

		return nil
	}
	rows.Close()
	return this.NwayDb
}
