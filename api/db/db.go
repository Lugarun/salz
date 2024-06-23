package db

import (
  "database/sql" 
  "github.com/Salz-Team/salz/api/models"
  _ "github.com/glebarez/go-sqlite"
)

type ApiDBHandler interface {
  GetDB() *sql.DB
  GetUser(int) (models.User, error)
  GetUserByLogin(string) (models.User, error)
  GetUserByIdentity(string, string) (models.User, error)
  CreateUser(models.User) (models.User, error)
  CreateBotFile(models.BotFile) (models.BotFile, error)
  ConfirmBotFile(models.BotFile) error
}

type AuthDBHandler interface {
  GetDB() *sql.DB
  GetToken(string) (models.AuthToken, error)
  DeleteTokenByUserId(int) error
  DeleteTokenByToken(string) error
  CreateToken(models.AuthToken) error
}

