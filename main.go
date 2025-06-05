package main

func main() {}

type LogInterface interface{}
type DatabaseInterface interface{}

type UserService struct {
	logger   LogInterface
	database DatabaseInterface
}

func NewUserService(logger LogInterface, database DatabaseInterface) *UserService {
	return &UserService{
		logger:   logger,
		database: database,
	}
}
