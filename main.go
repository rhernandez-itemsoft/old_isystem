package main

import (

	//"context"
	stdContext "context"

	"github.com/rhernandez-itemsoft/helpers/ilicense"
	"github.com/rhernandez-itemsoft/helpers/iresponse"
	"itemsoftmx/isystem/api/controllers/securityctrl"
	"itemsoftmx/isystem/api/controllers/usersctrl"
	"itemsoftmx/isystem/config"
	"log"
	"time"
	"github.com/kataras/iris/middleware/i18n"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"github.com/kataras/iris/mvc"
	"xorm.io/core"
)

//LanguageHandler permite manejar la traduccion (i18n)
var LanguageHandler context.Handler
var _response iresponse.Definition

func main() {
	fmt.Println(iris.Version)
	if !config.GralConf.ValidateLicense || ilicense.UIDAllowed() {

		//Crea una instancia de app con la configuración inicial
		app := initialice()

		////NewLanguageHandler incializa la configuración para manejar traducciones (i18n )
		LanguageHandler = i18n.New(
			i18n.Config{
				Default:      "es-MX",
				URLParameter: "language",
				Languages: map[string]string{
					"en-US": "./resources/languages/en-US.ini",
					"es-MX": "./resources/languages/es-MX.ini",
				},
			}
		)

		//Establece los handlers para el manejo de errores y para el manejo de traducciones
		app.Use(LanguageHandler)

		//conecta con la BD y mantiene la conexion abierta
		//db, err := sql.Open("mysql", "root:root@(localhost:3306)/isystem")
		xorm, err := xorm.NewEngine("mysql", "root:root@(localhost:3306)/isystem")
		if err != nil {
			log.Println(err.Error())
			return
		}
		if xorm.Ping() != nil {
			log.Println("No se pudo conectar a la Base de Datos.")
			xorm.Close()
			return
		}

		iris.RegisterOnInterrupt(func() {
			xorm.Close()
		})

		xorm.SetTableMapper(core.SameMapper{})
		//xorm.SetMapper(core.NewCacheMapper(new(core.SameMapper)))
		/*
			engine.SetTableMapper(core.SameMapper{})
			engine.SetColumnMapper(core.SnakeMapper{})


			SnakeMapper inserts an _ (underscore) between each word (Capital Case) to get the table name or column name.
			SameMapper uses the same name between the struct and table.
			GonicMapper is basically the same as SnakeMapper, but doesn't insert an underscore between commonly used acronyms. For example ID is converted to id in GonicMapper, but ID is converted to i_d in SnakeMapper.
		*/
		//engine.SetMapper(core.GonicMapper{})

		//OnAnyErrorCode se ejecuta cuando hay una petición a un servicio que no existe.

		//crea la tabla de ruteo
		createRoutes(app, xorm)

		app.OnAnyErrorCode(LanguageHandler, func(ctx iris.Context) {
			_response := iresponse.New(ctx, xorm)
			_response.JSON(iris.StatusUnauthorized, nil, "access_denied")
			return
		})

		//Arranca el servidor
		app.Run(iris.Addr(":8080"), iris.WithConfiguration(config.IrisConfiguration))

	} else {
		ilicense.RegisterUID()
	}
}

//Initialice Inicializa la configuración de la API
func initialice() *iris.Application {
	app := iris.New()
	/*
	* Prevenimos que se cierre el server, cuando un error fatal ocurre
	 */
	iris.RegisterOnInterrupt(func() {
		timeout := 10000 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		// close all hosts
		app.Shutdown(ctx)
	})
	return app
}

//createRoutes establece la tabla de ruteo
func createRoutes(app *iris.Application, xorm *xorm.Engine) {

	mvc.Configure(app.Party("/security", LanguageHandler), func(app *mvc.Application) {
		app.Register(app, xorm, &config.GralConf).Handle(new(securityctrl.Definition))
	})

	mvc.Configure(app.Party("/users", LanguageHandler), func(app *mvc.Application) {
		app.Register(app, xorm, &config.GralConf).Handle(new(usersctrl.Definition))
	})
}
