package handler

import (
	"context"
	"fmt"
	"main/models"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var ctx = context.TODO()

func (r routes) logout(rg *gin.RouterGroup) {
	rg.POST("/", Logout)
}

func (r routes) token(rg *gin.RouterGroup) {
	rg.POST("/refresh", Refresh)
	rg.GET("/valid", IsAccessTokenValid)
}

func IsAccessTokenValid(c *gin.Context) {
	r := c.Request
	tokenMetaData, err := ExtractTokenMetadata(r)
	if err != nil {
		c.JSON(http.StatusBadRequest, false)
		return
	}
	num, err2 := GetAuthentication(tokenMetaData)
	if err2 != nil {
		c.JSON(http.StatusUnauthorized, false)
		return
	} else if num == 0 {
		c.JSON(http.StatusOK, false)
		return
	}
	c.JSON(http.StatusOK, true)
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}

func ExtractToken(r *http.Request) string {
	bT := r.Header.Get("Authorization")
	newArr := strings.Split(bT, " ")
	if len(newArr) == 2 {
		return newArr[1]
	}
	return ""
}

func TokenVerify(r *http.Request) (*jwt.Token, error) {
	tS := ExtractToken(r)
	token, err := jwt.Parse(tS, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func TokenValid(r *http.Request) error {
	token, err := TokenVerify(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(r *http.Request) (*models.AccessDetails, error) {
	verifiedToken, err := TokenVerify(r)
	if err != nil {
		return nil, err
	}
	claims, ok := verifiedToken.Claims.(jwt.MapClaims)
	if ok && verifiedToken.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, err
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 64)
		if err != nil {
			return nil, err
		}
		return &models.AccessDetails{
			AccessUuid: accessUuid,
			UserId:     userId,
		}, nil
	}
	return nil, err
}

func GetAuthentication(authD *models.AccessDetails) (uint64, error) {
	userid, err := Client.Get(ctx, authD.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseUint(userid, 10, 64)
	return userID, nil
}

func DeleteAuthentication(givenUuid string) (int64, error) {
	deletedToken, err := Client.Del(ctx, givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deletedToken, nil
}

func Logout(c *gin.Context) {
	r := c.Request
	au, err := ExtractTokenMetadata(r)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "No valid authentication; unauthorized")
		return
	}
	deletedToken, delErr := DeleteAuthentication(au.AccessUuid)
	if delErr != nil || deletedToken == 0 {
		c.JSON(http.StatusUnauthorized, "No valid authentication; unauthorized")
		return
	}
	c.JSON(http.StatusOK, "Successfully logged out")
}

func Refresh(c *gin.Context) {
	mapToken := map[string]string{}
	if err := c.ShouldBindJSON(&mapToken); err != nil {
		c.JSON(http.StatusForbidden, err.Error())
		return
	}
	rT := mapToken["refresh_token"]

	tokenParsed, err := jwt.Parse(rT, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_SECRET")), nil
	})
	if err != nil {
		c.JSON(http.StatusForbidden, "Refresh token expired")
		return
	}
	if _, ok := tokenParsed.Claims.(jwt.Claims); !ok && !tokenParsed.Valid {
		c.JSON(http.StatusForbidden, err)
		return
	}
	parsedClaims, ok := tokenParsed.Claims.(jwt.MapClaims)
	if ok && tokenParsed.Valid {
		refreshUuid, ok := parsedClaims["refresh_uuid"].(string)
		if !ok {
			c.JSON(http.StatusForbidden, err)
			return
		}
		userId, err := strconv.ParseUint(fmt.Sprintf("%.f", parsedClaims["user_id"]), 10, 64)
		if err != nil {
			c.JSON(http.StatusForbidden, "Error occurred")
			return
		}
		deleted, delErr := DeleteAuthentication(refreshUuid)
		if delErr != nil || deleted == 0 {
			c.JSON(http.StatusForbidden, "No valid authentication; unauthorized")
			return
		}
		ts, createErr := NewToken(userId)
		if createErr != nil {
			c.JSON(http.StatusForbidden, createErr.Error())
			return
		}
		authErr := AuthFunc(userId, ts)
		if authErr != nil {
			c.JSON(http.StatusForbidden, authErr.Error())
			return
		}
		tokens := map[string]string{
			"access_token":  ts.AccessToken,
			"refresh_token": ts.RefreshToken,
		}
		c.JSON(http.StatusCreated, tokens)
	} else {
		c.JSON(http.StatusForbidden, "Refresh token expired.")
	}
}
