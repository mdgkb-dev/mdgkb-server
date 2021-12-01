package auth

import (
	"github.com/gin-gonic/gin"
	"mdgkb/mdgkb-server/models"
	"net/http"
)

func (h *Handler) Register(c *gin.Context) {
	var user *models.User
	err := c.Bind(&user)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	err = user.GenerateHashPassword()

	item, err := h.service.Register(user)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *Handler) Login(c *gin.Context) {
	var item models.User
	err := c.Bind(&item)
	if h.helper.HTTP.HandleError(c, err, http.StatusInternalServerError) {
		return
	}
	res, err := h.service.Login(&item)

	c.JSON(http.StatusOK, res)
}

func (h *Handler) Logout(c *gin.Context) {
	//_, err := h.helper.Token.ExtractTokenMetadata(c.Request)
	//if err != nil {
	//	c.JSON(http.StatusUnauthorized, "unauthorized")
	//	return
	//}
	//delErr := helpers.DeleteTokens(metadata, h.redis)
	//if delErr != nil {
	//	c.JSON(http.StatusUnauthorized, delErr.Error())
	//	return
	//}
	c.JSON(http.StatusOK, "Successfully logged out")
}

// Parse, validate, and return a token.
// keyFunc will receive the parsed token and should return the key for validating.
//func VerifyToken(r *http.Request) (*jwt.Token, error) {
//	tokenString := helpers.ExtractToken(r)
//	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(os.Getenv("ACCESS_SECRET")), nil
//	})
//	if err != nil {
//		return nil, err
//	}
//	return token, nil
//}
//
//func (h *Handler) Refresh(c *gin.Context) {
//	mapToken := map[string]string{}
//	if err := c.ShouldBindJSON(&mapToken); err != nil {
//		c.JSON(http.StatusUnprocessableEntity, err.Error())
//		return
//	}
//	refreshToken := mapToken["refresh_token"]
//
//	//verify the token
//	os.Setenv("REFRESH_SECRET", "mcmvmkmsdnfsdmfdsjf") //this should be in an env file
//	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
//		//Make sure that the token method conform to "SigningMethodHMAC"
//		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//		}
//		return []byte(os.Getenv("REFRESH_SECRET")), nil
//	})
//	//if there is an error, the token must have expired
//	if err != nil {
//		fmt.Println("the error: ", err)
//		c.JSON(http.StatusUnauthorized, "Refresh token expired")
//		return
//	}
//	//is token valid?
//	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
//		c.JSON(http.StatusUnauthorized, err)
//		return
//	}
//	//Since token is valid, get the uuid:
//	claims, ok := token.Claims.(jwt.MapClaims) //the token claims should conform to MapClaims
//	if ok && token.Valid {
//		refreshUuid, ok := claims["refresh_uuid"].(string) //convert the interface to string
//		if !ok {
//			c.JSON(http.StatusUnprocessableEntity, err)
//			return
//		}
//		userId := claims["user_id"].(string)
//
//		//Delete the previous Refresh Token
//		deleted, delErr := helpers.DeleteAuth(refreshUuid, h.redis)
//		if delErr != nil || deleted == 0 { //if any goes wrong
//			c.JSON(http.StatusUnauthorized, "unauthorized")
//			return
//		}
//		//Create new pairs of refresh and access tokens
//		ts, createErr := helpers.CreateToken(userId)
//		if createErr != nil {
//			c.JSON(http.StatusForbidden, createErr.Error())
//			return
//		}
//		//save the tokens metadata to redis
//		saveErr := helpers.CreateAuth(userId, ts, h.redis)
//		if saveErr != nil {
//			c.JSON(http.StatusForbidden, saveErr.Error())
//			return
//		}
//		tokens := map[string]string{
//			"access_token":  ts.AccessToken,
//			"refresh_token": ts.RefreshToken,
//		}
//		c.JSON(http.StatusCreated, tokens)
//	} else {
//		c.JSON(http.StatusUnauthorized, "refresh expired")
//	}
//}
