package models

import (
  "strconv"
  "path"
  "time"
)

// Go, y u no enum
const (
  BOT_STATUS_PENDING = "pending" // File has been created in the database, but not uploaded yet
  BOT_STATUS_UPLOADED = "uploaded" // File has been uploaded to S3
  BOT_STATUS_RUNNABLE = "runnable" // File has been uploaded and has participated in a game without crashing
  BOT_STATUS_FAILED = "failed" // File has been uploaded and has participated in a game, but crashed. Exclude from future games.
)

type BotFile struct {
  BotId int `json:"bot_id"`
  CreatedAt time.Time `json:"created_at"`
  UpdatedAt time.Time `json:"updated_at"`
  UploadPath string `json:"upload_path"`
  UserId int `json:"user_id"`
  Status string `json:"status"`
}

func (b BotFile) BotPathBuilder(extension string) string {
  return path.Join("bots", strconv.Itoa(b.UserId), strconv.Itoa(b.BotId)) + extension
}

func NewBotFile(userId int, extension string) BotFile {
  return BotFile{
    UserId: userId,
    Status: BOT_STATUS_PENDING,
  }
}
