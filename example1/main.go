package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var logger = log.New(os.Stdout, "fgophers", log.LstdFlags)

	userDB := NewUserDB(logger)
	service := NewComplianceService(userDB, logger)
	handler := NewHandler(service, logger)

	handler.Compliance.UserDB.GetUser()
}

type Handler struct {
	Compliance *ComplianceService
	Logger     *log.Logger
}

func NewHandler(compliance *ComplianceService, logger *log.Logger) *Handler {
	return &Handler{Compliance: compliance, Logger: logger}
}

type ComplianceService struct {
	UserDB *UserDB
	Logger *log.Logger
}

func NewComplianceService(userDB *UserDB, logger *log.Logger) *ComplianceService {
	return &ComplianceService{UserDB: userDB, Logger: logger}
}

type UserDB struct {
	logger *log.Logger
}

func NewUserDB(logger *log.Logger) *UserDB {
	return &UserDB{
		logger: logger,
	}
}

func (u UserDB) GetUser() {
	fmt.Println("called UserDB.GetUser()")
}
