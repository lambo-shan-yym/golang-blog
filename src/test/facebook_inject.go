package main

import (
	"fmt"
	"github.com/facebookgo/inject"
	"github.com/xormplus/xorm"
	"golang-blog/src/common/datasource"
	"log"
)

type DBEngine struct {
	Name string
}

type UserDB struct {
	Db *DBEngine `inject:""`
}

type UserService struct {
	Db *UserDB `inject:""`
}

type App struct {
	Name string
	User *UserService `inject:""`
}

func (a *App) Create() string {
	return "create app, in db name:" + a.User.Db.Db.Name + " app name :" + a.Name
}

type Object struct {
	App *App
}

func Init() *Object {
	db := DBEngine{Name: "db1"}
	var g inject.Graph
	app := App{Name: "go-app"}

	_ = g.Provide(
		&inject.Object{Value: &app},
		&inject.Object{Value: &db},
	)
	_ = g.Populate()
	return &Object{
		App: &app,
	}

}

func main() {
	//obj := Init()
	//fmt.Println(obj.App.Create())
	var index Index
	//inject declare
	db := datasource.InstanceDbMaster()
	//Injection
	var injector inject.Graph
	err := injector.Provide(
		&inject.Object{Value: &index},
		&inject.Object{Value: db},
		&inject.Object{Value: StartRepo{}},
		&inject.Object{Value: StartService{}},
	)
	if err != nil {
		log.Fatal("inject fatal: ", err)
	}
	if err := injector.Populate(); err != nil {
		log.Fatal("inject fatal: ", err)
	}
}

type IStartRepo interface {
	Speak(message string) string
}
//StartRepo 注入数据库
type StartRepo struct {
	engine *xorm.Engine `inject:""`
}

//Speak 实现Speak方法
func (s *StartRepo) Speak(message string) string {
	//使用注入的IDb访问数据库
	//s.Source.DB().Where("name = ?", "jinzhu").First(&user)
	return fmt.Sprintf("[Repository] speak: %s", message)
}

//IStartService 定义IStartService接口
type IStartService interface {
	Say(message string) string
}

//StartService 注入IStartRepo
type StartService struct {
	Repo IStartRepo `inject:""`
}

//Say 实现Say方法
func (s *StartService) Say(message string) string {
	return s.Repo.Speak(message)
}

//Index 注入IStartService
type Index struct {
	Service IStartService `inject:""`
}

//GetName 调用IStartService的Say方法
func (i *Index) GetName() {

	fmt.Println(i.Service.Say("sss"))
}