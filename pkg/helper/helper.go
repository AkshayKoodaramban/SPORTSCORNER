package helper

import (
	"errors"
	"fmt"
	cfg "sportscorner/pkg/config"
	helper "sportscorner/pkg/helper/interface"
	"sportscorner/pkg/utils/models"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"github.com/twilio/twilio-go"
	"golang.org/x/crypto/bcrypt"

	twilioApi "github.com/twilio/twilio-go/rest/verify/v2"
)

type Helper struct {
	Cfg cfg.Config
}

func NewHelper(Config cfg.Config) helper.Helper {
	return &Helper{
		Cfg: Config,
	}
}

var client *twilio.RestClient

func (h *Helper) GenerateTokenClients(user models.UserDetailsResponse) (string, error) {
    // Define the claims for the JWT
    claims := jwt.MapClaims{
        "id":    user.Id,
        "email": user.Email,
        "role":  "client",
        "exp":   time.Now().Add(time.Minute * 20).Unix(),
        "iat":   time.Now().Unix(),
    }

    // Create a new JWT with the specified claims and signing method
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign the JWT with the secret key "secret"
    tokenString, err := token.SignedString([]byte("secret"))

    // Check for errors during token signing
    if err != nil {
        return "", err
    }

    // Return the generated token string and nil error (indicating success)
    return tokenString, nil
}

func (h *Helper) GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, int64, error) {
	// accessTokenClaims := &AuthCustomClaims{
	// 	Id:    admin.ID,
	// 	Email: admin.Email,
	// 	Role:  "admin",
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
	// 		IssuedAt:  time.Now().Unix(),
	// 	},
	// }

	// refreshTokenClaims := &AuthCustomClaims{
	// 	Id:    admin.ID,
	// 	Email: admin.Email,
	// 	Role:  "admin",
	// 	StandardClaims: jwt.StandardClaims{
	// 		ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
	// 		IssuedAt:  time.Now().Unix(),
	// 	},
	// }

	// accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	// accessTokenString, err := accessToken.SignedString([]byte("accesssecret"))
	// if err != nil {
	// 	return "", "", err
	// }

	// refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	// refreshTokenString, err := refreshToken.SignedString([]byte("refreshsecret"))
	// if err != nil {
	// 	return "", "", err
	// }

	// return accessTokenString, refreshTokenString, nil
	accessTokenClaims := jwt.MapClaims{
		"id":    admin.ID,
		"email": admin.Email,
		"role":  "admin",
		"exp":   time.Now().Add(time.Minute * 20).Unix(),
		"iat":   time.Now().Unix(),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenString, err := accessToken.SignedString([]byte("secret")) // Use your secret key here

	if err != nil {
		return "", 0, err
	}

	// Retrieve the expiration time from the claims
	expirationTime, ok := accessTokenClaims["exp"].(int64)
	if !ok {
		return "", 0, fmt.Errorf("unable to retrieve expiration time from claims")
	}

	// Return both the access token and its expiration time as int64
	return accessTokenString, expirationTime, nil

}

func (h *Helper) TwilioSetup(username string, password string) {
	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: username,
		Password: password,
	})

}

func (h *Helper) TwilioSendOTP(phone string, serviceID string) (string, error) {
	fmt.Println("sid", serviceID)
	to := "+91" + phone
	fmt.Println("phone", to)
	params := &twilioApi.CreateVerificationParams{}
	params.SetTo(to)
	params.SetChannel("sms")
	resp, err := client.VerifyV2.CreateVerification(serviceID, params)
	if err != nil {
		return " ", err
	}
	return *resp.Sid, nil

}

func (h *Helper) TwilioVerifyOTP(serviceID string, code string, phone string) error {
	params := &twilioApi.CreateVerificationCheckParams{}
	params.SetTo("+91" + phone)
	params.SetCode(code)
	resp, err := client.VerifyV2.CreateVerificationCheck(serviceID, params)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if *resp.Status == "approved" {
		return nil
	}

	return errors.New("failed to validate otp")

}

func (h *Helper) PasswordHashing(password string) (string, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("internal server error")
	}

	hash := string(hashedPassword)
	return hash, nil
}

func (h *Helper) CompareHashAndPassword(a string, b string) error {
	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
	if err != nil {
		return err
	}
	return nil
}
func (h *Helper) Copy(a *models.UserDetailsResponse, b *models.UserSignInResponse) (models.UserDetailsResponse, error) {
	err := copier.Copy(a, b)
	if err != nil {
		return models.UserDetailsResponse{}, err
	}

	return *a, nil
}

func ExtractUserIDFromToken(tokenString, secretKey string) (int, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return 0, err
	}

	// Check if the token is valid
	if !token.Valid {
		return 0, errors.New("internal  error")
	}

	// Extract the user ID from the claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("internal  error")
	}

	userIDfloat64, ok := claims["id"].(float64) // Assuming "Id" is the key for user ID
	if !ok {
		return 0, errors.New("internal error")
	}
	userID := int(userIDfloat64)

	return userID, nil
}

// package helper

// import (
// 	"sportscorner/pkg/utils/models"
// 	"time"

// 	"github.com/golang-jwt/jwt"
// 	"golang.org/x/crypto/bcrypt"
// )

// type authCustomClaims struct {
// 	Id    int    `json:"id"`
// 	Email string `json:"email"`
// 	Role  string `json:"role"`
// 	jwt.StandardClaims
// }

// func GenerateTokenClients(user models.UserDetailsResponse) (string, error) {
// 	claims := &authCustomClaims{
// 		Id:    user.Id,
// 		Email: user.Email,
// 		Role:  "client",
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
// 			IssuedAt:  time.Now().Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte("secret"))

// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// func GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {
// 	claims := &authCustomClaims{
// 		Id:    admin.ID,
// 		Email: admin.Email,
// 		Role:  "admin",
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
// 			IssuedAt:  time.Now().Unix(),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString([]byte("secret"))

// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }
// func  CompareHashAndPassword(a string, b string) error {
// 	err := bcrypt.CompareHashAndPassword([]byte(a), []byte(b))
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
