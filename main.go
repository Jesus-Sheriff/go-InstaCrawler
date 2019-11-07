package main

import (
	"log"
	"os"
	"net/http"
	"encoding/json"
	"strconv"
	

	//"github.com/ahmdrz/goinsta"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Imagen struct {
    Status string `json:"status"`
    URI string `json:"URI"`
}
// Lista de struct Imagen
var imagenStatus []Imagen

// Funciones de route handlers
// func <your-function-name>(w http.ResponseWriter, r *http.Request)
func getStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	imagenStatus := Imagen{Status: "OK", URI: ""}
	json.NewEncoder(w).Encode(imagenStatus)
}
//implementar get latestimage que es el codigo del main
func getLatestImage() string{
	image := "https://scontent-mad1-1.cdninstagram.com/v/t51.2885-15/e35/74533385_166361284555809_7727768850258146020_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=109&bc=1571337657&oh=121bcbbc0ebee792f067f0d9cfcd5549&oe=5E2D09EB&ig_cache_key=MjE1ODI0Mjc3ODE0MDUwNjc2Mg%3D%3D.2"
	return image
}
func getImage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	uri:= getLatestImage()
	imagenStatus := Imagen{Status: "OK", URI: uri}
	json.NewEncoder(w).Encode(imagenStatus)
}
func getImageNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"]) // Se pasa a entero la entrada. Filtrado de datos de seguridad
	if id >=0 && id < len(imagenStatus) { // 'len' indica el número de elementos de un vector 'slice' 
		json.NewEncoder(w).Encode(imagenStatus[id])
	} else {
		imagenStatus := Imagen{Status: "NOT - OK: 416 Requested Range Not Satisfiable", URI: ""}
		json.NewEncoder(w).Encode(imagenStatus)
	}
}
func notFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	imagenStatus := Imagen{Status: "NOT - OK: 404", URI: ""}
	json.NewEncoder(w).Encode(imagenStatus)
}



func main() {
	router := mux.NewRouter()
	imagenStatus = append(imagenStatus, Imagen{Status: "OK", URI: "https://scontent-mad1-1.cdninstagram.com/v/t51.2885-15/e35/74533385_166361284555809_7727768850258146020_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=109&bc=1571337657&oh=121bcbbc0ebee792f067f0d9cfcd5549&oe=5E2D09EB&ig_cache_key=MjE1ODI0Mjc3ODE0MDUwNjc2Mg%3D%3D.2"})
	imagenStatus = append(imagenStatus, Imagen{Status: "OK", URI: "https://scontent-mad1-1.cdninstagram.com/v/t51.2885-15/e35/s1080x1080/72415541_145699950150997_7117851733228337833_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=101&bc=1571337657&oh=4a5b93972d502b2e13b4627fb112f319&oe=5E4A614E&ig_cache_key=MjE2NjMyNTg0MjY4NjA0MzY5Mw%3D%3D.2"})
	
	router.HandleFunc("/status", getStatus).Methods("GET")
	router.HandleFunc("/latest", getImage).Methods("GET")
	router.HandleFunc("/latest/{id}", getImageNumber).Methods("GET")
	// Función para capturar los errores 404
	router.NotFoundHandler = http.HandlerFunc(notFound)
	//Gestíon de las variables de entorno (puerto)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port, exists := os.LookupEnv("MYPORT")
	
	if exists == false {
		log.Println("Error al leer puerto. ¿Está definido? ")
		return
	}
	// insta := goinsta.New(
	// 	os.Getenv("INSTAGRAM_USERNAME"),
	// 	os.Getenv("INSTAGRAM_PASSWORD"),
	// )

	// if err := insta.Login(); err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// defer insta.Logout()

	// feedTag, err := insta.Feed.Tags("golang")
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	//log.Printf("ultima foto: ", feedTag.Images[0].User.Username)
	log.Printf("ultima foto:  %s", imagenStatus[len(imagenStatus)-1].URI)

	http.ListenAndServe(":"+port, router)
}
