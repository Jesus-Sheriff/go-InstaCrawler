package tests

import (
	// "os"
	"testing"
	// "net/url"
	"net/http"
	"io"
	"net/http/httptest"
	// "gopkg.in/ahmdrz/goinsta.v2"
	
)
func getStatusHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{Status: "OK", URI: ""}`)
}
func getImageHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{Status: "OK", URI: "https://scontent-mad1-1.cdninstagram.com/v/t51.2885-15/e35/74533385_166361284555809_7727768850258146020_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=109&bc=1571337657&oh=121bcbbc0ebee792f067f0d9cfcd5549&oe=5E2D09EB&ig_cache_key=MjE1ODI0Mjc3ODE0MDUwNjc2Mg%3D%3D.2"}`)
}
func getImageNumberHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{Status: "OK", URI: "https://scontent-mad1-1.cdninstagram.com/v/t51.2885-15/e35/74533385_166361284555809_7727768850258146020_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=109&bc=1571337657&oh=121bcbbc0ebee792f067f0d9cfcd5549&oe=5E2D09EB&ig_cache_key=MjE1ODI0Mjc3ODE0MDUwNjc2Mg%3D%3D.2"}`)
}
func notFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{Status: "NOT - OK: 404", URI: ""}`)
}

func TestGetStatus(t *testing.T){
	req, err := http.NewRequest("GET", "/status", nil)
    if err != nil {
        t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getStatusHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{Status: "OK", URI: ""}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
	t.Logf("getStatus correcto: '%s'",expected)
}
func TestGetImage(t *testing.T){
	req, err := http.NewRequest("GET", "/latest", nil)
    if err != nil {
        t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getImageHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{Status: "OK", URI: "https://scontent-mad1-1.cdninstagram.com/v/t51.2885-15/e35/74533385_166361284555809_7727768850258146020_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=109&bc=1571337657&oh=121bcbbc0ebee792f067f0d9cfcd5549&oe=5E2D09EB&ig_cache_key=MjE1ODI0Mjc3ODE0MDUwNjc2Mg%3D%3D.2"}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
	t.Logf("getImage correcto: '%s'",expected)
}
func TestGetImageNumber(t *testing.T){
	req, err := http.NewRequest("GET", "/latest/{id}", nil)
    if err != nil {
        t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getImageNumberHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	expected := `{Status: "OK", URI: "https://scontent-mad1-1.cdninstagram.com/v/t51.2885-15/e35/74533385_166361284555809_7727768850258146020_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=109&bc=1571337657&oh=121bcbbc0ebee792f067f0d9cfcd5549&oe=5E2D09EB&ig_cache_key=MjE1ODI0Mjc3ODE0MDUwNjc2Mg%3D%3D.2"}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
	t.Logf("getImageNumber correcto: '%s'",expected)
}

func TestNotFound(t *testing.T){
	req, err := http.NewRequest("GET", "/something", nil)
    if err != nil {
        t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(notFoundHandler)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusNotFound)
	}
	expected := `{Status: "NOT - OK: 404", URI: ""}`
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
	t.Logf("notFound correcto: '%s'",expected)
}


//OJO!!
//Antiguo test para obtener una imagen real de Instagram. (Parece que falla api que se usa: goinsta)

// func TestImportAccount(t *testing.T) {
// 	//obtenemos variables de entorno para login
// 	insta := goinsta.New(
// 		os.Getenv("INSTAGRAM_USERNAME"),
// 		os.Getenv("INSTAGRAM_PASSWORD"),
// 	)
// 	//logueamos...
// 	if err := insta.Login(); err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	//... o si hay un error deslogueamos
// 	defer insta.Logout()

// 	//llamamos a la función que busca un hashtag y recogemos errores si hubiere
// 	feedTag, err := insta.Feed.Tags("golang")
// 	if err != nil {
// 		t.Fatal(err)
// 		return
// 	}
// 	// URL de ejemplo:
// 	// https://scontent-mad1-1.cdninstagram.com/vp/7c8004a33e8ef83675e7c62a62c821d7/5E39388E/t51.2885-15/e35/70513351_167187977761265_1918517610523590583_n.jpg?_nc_ht=scontent-mad1-1.cdninstagram.com&_nc_cat=105&se=8&ig_cache_key=MjE1MDc2MzEwMzMzMTk0ODE0Mw%3D%3D.2

// 	//GetBest() es una función que es coge la imagen 
// 	//de mayor de resolución de las que se han devuelto en insta.Feed.Tags("golang")
// 	bestImageURL := feedTag.Images[0].Images.GetBest()

// 	//chequeo de que la URL devuelta es válida en estructura
// 	_, err = url.ParseRequestURI(bestImageURL)
//     if err != nil {
//         t.Fatal(err)
// 		return
//     }

// 	//mostramos información de log: url obtenida y usuario con el que nos hemos logueado
// 	t.Logf("URL is: %s", bestImageURL)

// 	t.Logf("logged into Instagram as user '%s'", insta.Account.Username)
// }
