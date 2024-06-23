package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/charmbracelet/log"
  "github.com/Salz-Team/salz/api/models"
  "strings"
)

// TODO find a nice home for this
const (
  MAX_FILE_SIZE_BYTES = 25 << 20 // ~ 25 MB
)

func (ctrl *Controller) BotUploadHandler(c *gin.Context) {
  // TODO get user ID in middleware and pass it around via gin context?
  authTokenHeader := c.GetHeader("Authorization")
  authToken := strings.TrimSpace(strings.Replace(authTokenHeader, "Bearer", "", 1))

  // Get the user id from the auth token
  userToken, err := ctrl.cfg.AuthDBHandler.GetToken(authToken)
  if err != nil {
    log.Error("Error getting user id from token", "error", err)
    c.AbortWithStatus(http.StatusUnauthorized)
    return
  }
  userId := userToken.UserId

  // Note that file is a multipart.File which implements io.Reader
  file, fileHeader, err := c.Request.FormFile("data")
  if err != nil {
    log.Error("Error getting file from form", "error", err)
    c.AbortWithStatus(http.StatusInternalServerError)
    return
  }

  // TODO filesize check, content type, encoding, etc

  // Get the extension of the file
  extension := fileHeader.Filename[strings.LastIndex(fileHeader.Filename, "."):]

  // Transaction might not be a good idea if the upload takes a while
  // Create a new file in the database, note that upload path is null until file is uploaded.
  bot, err := ctrl.cfg.ApiDBHandler.CreateBotFile(models.BotFile{
    UserId: userId,
  })

  bot.UploadPath = bot.BotPathBuilder(extension)

  // Upload the file to S3
  err = ctrl.cfg.ObjectStoreHandler.UploadBotFile(bot.UploadPath, file)
  if err != nil {
    log.Error("Error uploading bot file", "error", err)
    c.AbortWithStatus(http.StatusInternalServerError)
    return
  }

  // Update the file in the database
  err = ctrl.cfg.ApiDBHandler.ConfirmBotFile(bot)
  if err != nil {
    log.Error("Error confirming bot file", "error", err)
    c.AbortWithStatus(http.StatusInternalServerError)
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "bot_id": bot.BotId,
  })
}
