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
	"fmt"
)

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
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer: "taller ql",
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
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if(err != nil){
		fmt.Fprintln(w, "Error al leer el usuario", err)
		return
	}
	
	if user.Name == "alexys" && user.Password == "alexys"{
		user.Password = ""
		user.Role = "admin"
		
		token := GenerateJWT(user)
		result := models.ResponseToken{token}
		jsonResult, err := json.Marshal(result)
		if err != nil{
			fmt.Fprintln(w,"Error al generar json")
			return 
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type","application/json")
		w.Write(jsonResult)
		
	}else{
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w,"Usuario o clave invalido")
	}
}

func ValidateToken(w http.ResponseWriter, r *http.Request)  {
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor,&models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey,nil
	})

	if err!= nil{
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors{
			case jwt.ValidationErrorExpired:
				fmt.Fprintln(w,"Su token ha expirado")
				return
			case jwt.ValidationErrorSignatureInvalid:
				fmt.Fprintln(w,"Su token no coincide")
				return
			default:
				fmt.Fprintln(w,"Su token exploto")
				return
			}
		default:
			fmt.Fprintln(w,"Su token exploto 2")
			return
		}
	}

	if token.Valid{
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w,"Bienvenido al sistema")
	}else{
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "su token no es valido inautorizado")
	}
}