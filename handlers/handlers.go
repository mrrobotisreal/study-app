package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"study-app/auth"
	"study-app/aws_services"
	"study-app/db_stuff"
	"time"
)

var Mux *http.ServeMux

type GetCollectionsRequest struct {
	User string `json:"user,omitempty" bson:"user,omitempty"`
}

func enableCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
	(*w).Header().Set("Content-type", "*")
	(*w).Header().Set("Access-Control-Allow-Credentials", "true")
}

func AddCollection(response http.ResponseWriter, request *http.Request) {
	enableCORS(&response)
	if request.Method != "POST" {
		var err db_stuff.ServerMessage
		err.Message = "Errors are yonder this way..."
		err.Type = "ERROR"
		json.NewEncoder(response).Encode(&err)
		return
	}
	var req db_stuff.AddCollectionRequest
	json.NewDecoder(request.Body).Decode(&req)
	fmt.Println(req.Description)
	result := db_stuff.AddCollection(req)
	if !result {
		var err db_stuff.ServerMessage
		err.Message = "There was a problem adding your collection, please try again."
		err.Type = "ERROR"
		json.NewEncoder(response).Encode(&err)
		return
	}
	var success db_stuff.ServerMessage
	success.Message = "Successfully added new collection."
	success.Type = "SUCCESS"
	json.NewEncoder(response).Encode(&success)
}

func GetCollections(response http.ResponseWriter, request *http.Request) {
	enableCORS(&response)
	//response.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	//response.Header().Set("Access-Control-Allow-Headers", "Content-type")
	//response.Header().Set("Access-Control-Allow-Methods", "GET, POST")
	response.Header().Add("content-type", "application/json")
	if request.Method != "POST" {
		var err db_stuff.ServerMessage
		err.Message = "Errors getting collections..."
		err.Type = "ERROR"
		json.NewEncoder(response).Encode(&err)
		return
	}

	var user GetCollectionsRequest
	json.NewDecoder(request.Body).Decode(&user)
	collections := db_stuff.GetCollections(user.User)
	json.NewEncoder(response).Encode(&collections)
}

func GetCollection(response http.ResponseWriter, request *http.Request) {
	enableCORS(&response)
	if request.Method != "GET" {
		var err db_stuff.ServerMessage
		err.Message = "Errors here yo..."
		err.Type = "ERROR"
		json.NewEncoder(response).Encode(&err)
		return
	}
	var req GetCollectionsRequest
	json.NewDecoder(request.Body).Decode(&req)
	params := strings.Split(request.URL.Path, "/")
	result := db_stuff.GetCollection(req.User, params[3])
	json.NewEncoder(response).Encode(&result)
}

func SetUser(response http.ResponseWriter, request *http.Request) {
	enableCORS(&response)
	if request.Method != "POST" {
		var err db_stuff.ServerMessage
		err.Message = "Errors be here matey...."
		err.Type = "ERROR"
		json.NewEncoder(response).Encode(&err)
		return
	}
	var user db_stuff.User
	type tempUser struct {
		Email    string `json:"email,omitempty" bson:"email,omitempty"`
		Password string `json:"password,omitempty" bson:"password,omitempty"`
		Username string `json:"username,omitempty" bson:"username,omitempty"`
	}
	var temp tempUser
	json.NewDecoder(request.Body).Decode(&temp)
	user.Username = temp.Username
	user.Email = temp.Email
	user.JWT = auth.GenerateJWT(user.Username, user.Email)
	password, _ := auth.HashPassword(temp.Password)
	user.Password = password
	user.Collections = []db_stuff.CardCollection{}
	result := db_stuff.SetUser(user)
	if !result {
		var err db_stuff.ServerMessage
		err.Message = "Something went wrong, please try again."
		err.Type = "ERROR"
		json.NewEncoder(response).Encode(&err)
	}
	json.NewEncoder(response).Encode(&user.JWT)
}

func GetUser(response http.ResponseWriter, request *http.Request) {
	enableCORS(&response)
	if request.Method != "GET" {
		var err db_stuff.ServerMessage
		err.Message = "Errors be here matey GET...."
		err.Type = "ERROR"
		return
	}
	params := strings.Split(request.URL.Path, "/")
	user := db_stuff.GetUser(params[2])
	json.NewEncoder(response).Encode(&user)
}

func CheckLogin(response http.ResponseWriter, request *http.Request) {
	enableCORS(&response)
	if request.Method != "POST" {
		var err db_stuff.ServerMessage
		err.Message = "Errors be here matey POST...."
		err.Type = "ERROR"
		json.NewEncoder(response).Encode(&err)
		return
	}
	type pwd struct {
		Password string `json:"Password,omitempty" bson:"Password,omitempty"`
	}
	var password pwd
	json.NewDecoder(request.Body).Decode(&password)
	username := strings.Split(request.URL.Path, "/")
	authResult := db_stuff.CheckLogin(username[2], password.Password)
	type responseBody struct {
		AuthResult bool
		JWT        string
		Username   string
	}
	var resBody responseBody
	resBody.AuthResult = authResult.AuthResult
	resBody.JWT = authResult.JWT
	resBody.Username = username[2]
	json.NewEncoder(response).Encode(&resBody)
}

func Logout(response http.ResponseWriter, request *http.Request) {
	enableCORS(&response)
	if request.Method != "GET" {
		var err db_stuff.ServerMessage
		err.Message = "Errors be here matey DELETE...."
		err.Type = "ERROR"
		return
	}
	username := strings.Split(request.URL.Path, "/")
	log.Println("Logout handler: username: " + username[2])
	logoutResult := db_stuff.Logout(username[2])
	if !logoutResult {
		json.NewEncoder(response).Encode("There was a problem logging you out, please try again.")
	} else {
		json.NewEncoder(response).Encode("Successfully logged out!")
	}
}

func CheckSession(response http.ResponseWriter, request *http.Request) {
	enableCORS(&response)
	if request.Method != "POST" {
		var err db_stuff.ServerMessage
		err.Message = "Errors be here matey POST...."
		err.Type = "ERROR"
		return
	}
	username := strings.Split(request.URL.Path, "/")
	type jwt struct {
		Token string `json:"Token,omitempty" bson:"Token,omitempty"`
	}
	var token jwt
	json.NewDecoder(request.Body).Decode(&token)
	log.Println("username-> " + username[2])
	log.Println("token-> " + token.Token)
	if err := auth.ValidateJWT(token.Token, username[2]); err != nil {
		var err db_stuff.ServerMessage
		err.Message = "Errors validating JWT"
		err.Type = "ERROR"
		json.NewEncoder(response).Encode(false)
	}
	json.NewEncoder(response).Encode(true)
}

func SetSession(response http.ResponseWriter, request *http.Request) {
	enableCORS(&response)
	if request.Method != "POST" {
		var err db_stuff.ServerMessage
		err.Message = "Errors be here matey POST...."
		err.Type = "ERROR"
		return
	}
	log.Println("SET SESSION REQUEST INCOMING")
	var sessionReq db_stuff.SessionRequest
	json.NewDecoder(request.Body).Decode(&sessionReq)
	log.Println(sessionReq.Name)
	log.Println(sessionReq.Email)
	log.Print("isSignup = ")
	log.Println(sessionReq.IsSignup)
	username := strings.Split(request.URL.Path, "/")
	log.Println("USERNAME: " + username[2])
	session := db_stuff.SetSession(username[2], sessionReq)
	json.NewEncoder(response).Encode(&session)
}

func Testy(response http.ResponseWriter, request *http.Request) {
	enableCORS(&response)
	fmt.Println("CHECKING FOR THE TIME")
	fmt.Println(time.Now().UTC().UnixMilli())
	fmt.Printf("millis = %T", time.Now().UTC().UnixMilli())
}

func AudioTest(response http.ResponseWriter, request *http.Request) {
	enableCORS(&response)
	fmt.Println("CHECKING AUDIO")

	parseErr := request.ParseMultipartForm(32 << 20)
	if parseErr != nil {
		fmt.Println("There was an error parsing the multipart form")
		return
	}

	for _, h := range request.MultipartForm.File["audio_data"] {
		fmt.Println("H filename is:")
		fmt.Println(h.Filename)
		file, _ := h.Open()
		tmpfile, _ := os.Create("./" + h.Filename)
		io.Copy(tmpfile, file)
		file.Close()
	}

}

func handleAudioUpload(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var audioName string
	json.NewDecoder(r.Body).Decode(&audioName)

	// Read the audio blob from the request body
	audioBlob, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error reading audio blob"))
		return
	}

	// Upload the audio blob to S3
	bucketName := "study-app-audio"
	keyName := "audio/" + audioName + ".mp3"
	err = aws_services.UploadAudioBlobToS3(audioBlob, bucketName, keyName)
	if err != nil {
		log.Printf("Error uploading audio blob to S3: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error uploading audio blob to S3"))
		return
	}

	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Audio blob uploaded successfully"))
}
