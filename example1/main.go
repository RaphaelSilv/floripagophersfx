package main

import (
	"fmt"
	"log"
	"os"
)

var Logger = log.New(os.Stdout, "fgophers", log.LstdFlags)

func main() {

	userDB := NewUserDB()
	service := NewComplianceService(userDB)
	handler := NewHandler(service)

	handler.Compliance.UserDB.GetUser()
}

type Handler struct {
	Compliance *ComplianceService
}

func NewHandler(compliance *ComplianceService) *Handler {
	return &Handler{Compliance: compliance}
}

type ComplianceService struct {
	UserDB *UserDB
}

func NewComplianceService(userDB *UserDB) *ComplianceService {
	return &ComplianceService{UserDB: userDB}
}

type UserDB struct{}

func NewUserDB() *UserDB {
	return &UserDB{}
}

func (u UserDB) GetUser() {
	Logger.Print("logging stuff")
	fmt.Println("called UserDB.GetUser()")
}
