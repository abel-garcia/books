package books

import (
	"fmt"

	"github.com/wackGarcia/books/cmd/books/config"
)

func Run() error {
	fmt.Println("set up configurations")
	conf := config.Config{}
	if err := conf.LoadConfig(); err != nil {
		return err
	}

	fmt.Println("set up database")
	db, err := conf.Storage.Conn()
	if err != nil {
		return err
	}

	fmt.Println("runnning server", conf.HttpServer)
	if err := HttpServer(conf.HttpServer, db); err != nil {
		return err
	}
	return nil
}
