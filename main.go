package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//PageVariables struct
type PageVariables struct {
	Title   string
	Name    string
	Data    string
	Degrees string
}

const degrees string = "℃"

//HomePage func
func HomePage(w http.ResponseWriter, r *http.Request) {

	HomePageVars := PageVariables{ //store the date and time in a struct
		Title: "Weather | Home",
		Name:  "WeatherAPI",
	}
	t, err := template.ParseFiles("./assets/index.html") //parse the html file homepage.html
	if err != nil {                                      // if there is an error
		log.Print("template parsing error: ", err) // log it
	}
	err = t.Execute(w, HomePageVars) //execute the template and pass it the HomePageVars struct to fill in the gaps
	if err != nil {                  // if there is an error
		log.Print("template executing error: ", err) //log it
	}
}

//ContactPage func
func ContactPage(w http.ResponseWriter, r *http.Request) {
	HomePageVars := PageVariables{
		Title: "WeatherAPI | Contact",
		Name:  "WeatherAPI",
	}
	t, err := template.ParseFiles("./assets/contact.html")
	if err != nil {
		log.Print("template executing error: ", err)
	}
	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}

//FormPage func
func FormPage(w http.ResponseWriter, r *http.Request) {
	pageVar := PageVariables{
		Title: "WeatherAPI | Weather",
		Name:  "WeatherAPI",
	}
	if r.Method == "GET" {
		t, err := template.ParseFiles("./assets/form.html")
		if err != nil {
			log.Print("Error loading the form page: ", err)
		}
		t.Execute(w, pageVar)
	} else {
		r.ParseForm()
		city := r.Form["city"]
		c := strings.ToLower(strings.Join(city, ","))
		data := Cities(c)
		//fix following code up
		if len(data.Name) == 0 {
			fmt.Println("Data entered is not valid")
		}
		fmt.Println(FloatToString(data.Main.Temp))
		pageVars := PageVariables{
			Title:   "WeatherAPI | Weather",
			Name:    "WeatherAPI",
			Data:    FloatToString(ToCelcius(data.Main.Temp)),
			Degrees: degrees,
		}
		fmt.Println(pageVars)
		t, err := template.ParseFiles("./assets/data.html")
		if err != nil {
			fmt.Println(err)
			log.Print("Error loading the data.html file :", err)
		}
		err = t.Execute(w, pageVars)
		if err != nil {
			log.Print("Error parsing data to file", err)
		}
	}
}

func newRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", HomePage).Methods("GET")
	r.HandleFunc("/Contact", ContactPage).Methods("GET")
	r.HandleFunc("/Weather", FormPage)
	// Declare the static file directory and point it to the
	// directory we just made
	staticFileDirectory := http.Dir("./assets/")
	// Declare the handler, that routes requests to their respective filename.
	// The fileserver is wrapped in the `stripPrefix` method, because we want to
	// remove the "/assets/" prefix when looking for files.
	// For example, if we type "/assets/index.html" in our browser, the file server
	// will look for only "index.html" inside the directory declared above.
	// If we did not strip the prefix, the file server would look for
	// "./assets/assets/index.html", and yield an error
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	// The "PathPrefix" method acts as a matcher, and matches all routes starting
	// with "/assets/", instead of the absolute route itself
	r.PathPrefix("/assets/").Handler(staticFileHandler).Methods("GET")
	return r

}

func main() {
	// declare a new router
	r := newRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}
