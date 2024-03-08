package people

import (
	"os"

	"github.com/dpolimeni/fiber_app/common"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) fiber.Router {
	v1 := app.Group("/api/v1")
	postgres := common.Postgres{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: os.Getenv("password"),
		Dbname:   "gotest",
	}
	mongo := common.Mongo{
		Username: "diegopolimeni",
		Password: os.Getenv("mongo_password"),
		Host:     os.Getenv("mongo_host"),
	}
	mongoClient, err := mongo.DataBase()
	if err != nil {
		panic(err)
	}
	postgresClient, err := postgres.DataBase()
	if err != nil {
		panic(err)
	}
	handler := &PeopleHandler{
		DbClient:    postgresClient,
		MongoClient: mongoClient,
	}
	v1.Post("/person/new", handler.AddPerson)
	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_KEY"))},
	}))
	v1.Get("/people", handler.GetPeople)
	v1.Get("/person/:id<int>", handler.GetPerson)
	v1.Delete("/person/:id", handler.DeletePerson)
	v1.Put("/person/:username", handler.UpdatePerson)
	//defer DbClient.Close()
	return v1
}
