package controllers

import (
  "net/http"
  "github.com/gin-gonic/gin"
  "github.com/charmbracelet/log"
  "github.com/Salz-Team/salz/api/models"
  "strconv"

  ghlib "github.com/google/go-github/v62/github"
)

// TODO find a good place for this
const (
  DEFAULT_ELO = 1000.0
)

func (ctrl *Controller) OAuthLoginHandler(c *gin.Context) {
  url := ctrl.cfg.OAuth2Config.AuthCodeURL("ain't got no state yet")
  log.Info("Login redirect", "url", url)
  c.Redirect(http.StatusTemporaryRedirect, url)
}

func (ctrl *Controller) OAuthCallbackHandler(c *gin.Context) {
  // TODO Check the state to prevent CSRF attacks
  code := c.Query("code")
  token, err := ctrl.cfg.OAuth2Config.Exchange(c, code)
  if err != nil {
    log.Error("oauth2: Exchange failed", "error", err)
    c.AbortWithStatus(http.StatusUnauthorized)
    return
  }

  // Instead of pulling in this massive library, we could have done a basic query against the API.
  // But I didn't want to set up the models.
  ghc := ghlib.NewClient(nil).WithAuthToken(token.AccessToken)
  user, _, err := ghc.Users.Get(c, "")
  if err != nil {
    log.Error("Unable to get GH user info", "error", err)
    c.AbortWithStatus(http.StatusInternalServerError)
    return
  }

  // Check the db for the user. If it doesn't exist, create it.
  userId := strconv.FormatInt(*user.ID, 10)
  idp := "github"
  u, err := ctrl.cfg.ApiDBHandler.GetUserByIdentity(idp, userId)
  if err != nil {
    log.Warn("Unable to get user by identity", "error", err)
    // Create the user
    u = models.User{
      UserName: *user.Login,
      IconPath: *user.AvatarURL,
      IdentityProvider: idp,
      IdentityProviderId: userId,
      Elo: DEFAULT_ELO,
    }
    u, err = ctrl.cfg.ApiDBHandler.CreateUser(u)
    if err != nil {
      log.Error("Unable to create user", "error", err)
      c.AbortWithStatus(http.StatusInternalServerError)
      return
    }
    log.Info("User created", "user", u)
  }

  // Create a token for the user
  userAuthToken := models.NewAuthTokenForUser(u.Id)
  err = ctrl.cfg.AuthDBHandler.CreateToken(*userAuthToken)
  if err != nil {
    log.Error("Unable to create token", "error", err)
    c.AbortWithStatus(http.StatusInternalServerError)
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "access_token": userAuthToken.Token,
    "expires_at": userAuthToken.ExpiresAt,
  })
}
