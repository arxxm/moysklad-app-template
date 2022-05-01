package main

import (
	"github.com/MaLowBar/moysklad-app-template"
	"github.com/MaLowBar/moysklad-app-template/storage"
	"github.com/labstack/echo/v4"
	"log"
	"os"
	"os/signal"
)

func main() {
	// Задаём информацию необходимую для работы приложения
	var info = moyskladapptemplate.AppConfig{
		ID:           "fb08e3f3-8f1a-488e-a609-1baa389cc546",
		SecretKey:    "8iv6RbvFlQsiDMqQz4ECczLjiwEZRfBkVKa2cMBmsHnzIg2ELuqdbQNXvloY65nQD1crmxdbCVXbx1CvnjY1Th9sUebNXOYnULPtZ40N2ujjv7EzbE6F5SEM9xucnEAL",
		VendorAPIURL: "/echo/api/moysklad/vendor/1.0/apps/:appId/:accountId",
	}

	// Можно использовать БД PostgreSQL
	//myStorage, err := storage.NewPostgreStorage("postgres://msgo:pswd@localhost/msgo_db")
	//if err != nil {
	//	log.Fatal(fmt.Errorf("cannot create storage: %w", err))
	//	return
	//}

	// Инициализируем файловое хранилище
	myStorage := storage.NewFileStorage("./")

	// Определяем простейший обработчик для HTML-документа
	var iframeHandler = moyskladapptemplate.AppHandler{
		Method: "GET",
		Path:   "/echo/iframe/purchases-report-go.sorochinsky",
		HandlerFunc: func(c echo.Context) error {
			return c.HTML(200, `<html>
    <head>
    </head>
    <body>
        <center>
            <h1> Hello, Malik! </h1>
        </center>    
    </body>
</html>
`)
		},
	}

	// Создаем приложение
	app := moyskladapptemplate.NewApp(info, myStorage, iframeHandler)

	e := make(chan error)
	go func() {
		e <- app.Run("0.0.0.0:8002") // Запускаем
	}()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt)

	select {
	case err := <-e:
		log.Printf("Server returned error: %s", err)
	case <-c:
		app.Stop(5)
		log.Println("Stop signal received")
	}
}
