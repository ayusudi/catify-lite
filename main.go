import (
	"github.com/labstack/echo/v4"
	"catify-lite/config"
	"catify-lite/routes"
)

func main() {
	e := echo.New()

	config.ConnectDB()
	config.Migrate()
	routes.InitRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
