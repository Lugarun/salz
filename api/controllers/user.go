package controllers

import (
  "net/http"
  "strings"
  "github.com/gin-gonic/gin"
  "github.com/charmbracelet/log"

  "github.com/Salz-Team/salz/api/db"
)

func GetUser(c *gin.Context, dbHandler db.ApiDBHandler, authHandler db.AuthDBHandler) {
  // Get the user from the database
  
  authTokenHeader := c.GetHeader("Authorization")
  authToken := strings.TrimSpace(strings.Replace(authTokenHeader, "Bearer", "", 1))

  token, err := authHandler.GetToken(authToken)
  if err != nil {
    log.Error("Unable to get token", "error", err)
    c.AbortWithStatus(http.StatusInternalServerError)
    return
  }

  user, err := dbHandler.GetUser(token.UserId)
  if err != nil {
    log.Error("Unable to get user", "error", err)
    c.AbortWithStatus(http.StatusInternalServerError)
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "userid": user.Id,
    "username": user.UserName,
    "iconpath": user.IconPath,
  })
}

