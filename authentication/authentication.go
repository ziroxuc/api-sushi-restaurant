package authentication

// se firman los token con llave privada
// se verifica con llave publica

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
     jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
     "../models"
	"time"
	"net/http"
	"encoding/json"
	utils "../utils"
	mo "../models"
	db "../dbConnection"
	"gopkg.in/mgo.v2/bson"
)

var cUsuario = db.GetCollectionUsuario()

var (
	privateKey *rsa.PrivateKey
	publicKey *rsa.PublicKey
)

func init()  {
	privateBytes, err := ioutil.ReadFile("./privateToken.rsa")
	if err != nil{
		log.Fatal("no se pudo leer clave privada")
	}

	publicBytes, err := ioutil.ReadFile("./public.rsa.pub")
	if err != nil{
		log.Fatal("no se pudo leer clave publica")
	}

	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil{
		log.Fatal("No se pudo tratar la privada")
	}

	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil{
		log.Fatal("No se pudo tratar la privada")
	}
}
//funcion para generar token

func GenerateJWT(user models.User) (string)  {
	claims := models.Claim{
		User:user,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 8).Unix(),
			Issuer: "restaurant",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodPS256, claims)
	result, err := token.SignedString(privateKey)
	if err != nil{
		log.Fatal("No se pudo firmar el token")
	}
	return result
}

func Login(w http.ResponseWriter, r *http.Request){
	var userUknown mo.User
	err := json.NewDecoder(r.Body).Decode(&userUknown)
	if(err != nil){
		utils.RespondWithError(w, http.StatusInternalServerError, "Error al leer el body.")
		return
	}

	var user mo.User
	err = cUsuario.Find(bson.M{"usuario":userUknown.Usuario}).One(&user)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "Usuario o clave incorrecta.")
		return
	}

	if user.Usuario == userUknown.Usuario && utils.ComparePassword(user.PasswordHash,userUknown.Password){
		userUknown.Password = ""
		userUknown.Usuario = ""
		userUknown.Role = "USER_ADMIN"

		token := GenerateJWT(userUknown)
		if err != nil{
			utils.RespondWithError(w, http.StatusInternalServerError, "Error al generar json de token.")
			return
		}
		//Se asigna token creado al objeto user
		user.Token = token
		user.PasswordHash = nil
		utils.RespondWithJSON(w,http.StatusOK,user)
		
	}else{
		utils.RespondWithError(w, http.StatusUnauthorized, "Usuario o clave incorrecta.")
	}
}

func ValidateToken(r *http.Request) string  {
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor,&models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey,nil
	})
	if err!= nil{
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors{
			case jwt.ValidationErrorExpired:
				return "Su token ha expirado."
			case jwt.ValidationErrorSignatureInvalid:
				return "Su token no coincide."
			default:
				return "Su token no es correcto."
			}
		default:
			return "Su token no es válido."
		}
	}
	if token.Valid{
		return ""
	}else{
		return "su token no es valido inautorizado"
	}
}


func IsValidToken(w http.ResponseWriter,r *http.Request)  {
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor,&models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey,nil
	})
	if err!= nil{
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors{
			case jwt.ValidationErrorExpired:
				utils.RespondWithError(w, http.StatusUnauthorized, "false")
				return
			case jwt.ValidationErrorSignatureInvalid:
				utils.RespondWithError(w, http.StatusUnauthorized, "false")
				return
			default:
				utils.RespondWithError(w, http.StatusUnauthorized, "false")
				return
			}
		default:
			utils.RespondWithError(w, http.StatusUnauthorized, "false")
			return
		}
	}
	if token.Valid{
		utils.RespondWithJSON(w ,http.StatusOK ,true)
	}else{
		utils.RespondWithError(w, http.StatusUnauthorized, "false")
		return
	}
}