package main

import (
	"bytes"
	log "github.com/Sirupsen/logrus"
	"math/rand"

	"github.com/alligrader/gradebook-api/src/app"
	_ "github.com/go-sql-driver/mysql"
	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"

	"github.com/alligrader/gradebook-api/src/models"
)

// UserController implements the user resource.
type UserController struct {
	*goa.Controller
	db *models.UserDB
}

// NewUserController creates a user controller.
func NewUserController(service *goa.Service) *UserController {
	log.Debug("Opening connection to DB")
	// Open the connection to the DB
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/alligrader")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	log.Debug("Automigrating user model")
	db.AutoMigrate(&models.User{})

	userDB := models.NewUserDB(db)

	return &UserController{Controller: service.NewController("UserController"), db: userDB}
}

func (c *UserController) Read(ctx *app.ReadUserContext) error {
	return ctx.OK(&app.UserMt{})
}

func (c *UserController) Update(ctx *app.UpdateUserContext) error {
	return ctx.OK(&app.UserMt{})
}

// Create runs the create action.
func (c *UserController) Create(ctx *app.CreateUserContext) error {
	// UserController_Create: start_implement

	// Fetch the user out of the request context
	log.Info("Converting payload")
	u := models.UserFromUserCreate(ctx.Payload)

	// Generate the salt
	log.Info("Getting the salt")
	u.Salt = getSalt()

	// Take the password out of the media type,
	// Salt and hash it

	log.Info("Generating the hash")
	var saltedPass bytes.Buffer
	if u.Password == nil {
		return ctx.InternalServerError()
	}
	saltedPass.Write([]byte(*u.Password))
	saltedPass.Write([]byte(u.Salt))
	pass, err := bcrypt.GenerateFromPassword(saltedPass.Bytes(), bcrypt.DefaultCost)
	if err != nil {
		log.Warn(err)
	}
	log.Info("Hash generated. Writing the user to the DB")
	// Store it back in.
	t := string(pass)
	u.Password = &t

	// Add that user to the database
	c.db.Add(ctx, u)
	mediaUser := u.UserToUserMt()
	log.Info("Success. Returning result.")

	// Rip out the password before returning it.
	u.Password = nil

	// UserController_Create: end_implement
	return ctx.OK(mediaUser)
}

func getSalt() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const length = 64

	randString := make([]byte, length)
	for i := range randString {
		randString[i] = letters[rand.Intn(len(letters))]
	}
	return string(randString)
}
