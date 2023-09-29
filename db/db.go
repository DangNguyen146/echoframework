package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
	_ "github.com/lib/pq"
)

type Sql struct {
	Db       *sqlx.DB
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
}

func (s *Sql) Connect() {
	dataSource := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", s.Host, s.Port, s.UserName, s.Password, s.DbName)
	s.Db = sqlx.MustConnect("postgres", dataSource)

	if err := s.Db.Ping(); err != nil {
		log.Errorf(err.Error())
		return
	}
	fmt.Println("Connect database ok")
}

func (s *Sql) Connect2() {
	dataSource := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", s.UserName, s.Password, s.Host, s.Port, s.DbName)
	s.Db, _ = sqlx.Connect("postgres", dataSource)

	if err := s.Db.Ping(); err != nil {
		log.Errorf(err.Error())
		return
	}
	fmt.Println("Connect database2 ok")
}

func (s *Sql) Close() {
	err := s.Db.Close()
	if err != nil {
		return
	}
}
