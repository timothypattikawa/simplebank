package main

import "github.com/timothypattikawa/simplebank/internal/config"

func main() {

	v := config.LoadViper()

	conf := config.NewConfiguration(v)

	_ = conf.DBConf.NewDbConn()

}
