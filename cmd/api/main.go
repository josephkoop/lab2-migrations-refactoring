package main

import(
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome! This is Joseph Koop's territory. The project I have chosen is #7 - Gym/Fitness Class Scheduling System. Why? It's easier to work on something that you're interested in, and I have an interest in gyms.     \n"))
}

func about(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("About me     \nName: Joseph Koop     \nField of Study: Information Technology     \nCurrent Residence: Spanish Lookout     \n"))
}

func contact(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Email: 2023159207@ub.edu.bz     \nGitHub username: josephkoop (main account) + Joseph-Koop (school email account)     \n"))
}

func hobby(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hobby: Chess     \nI doubt this one needs explaining     \n"))
}

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/about", about)
	mux.HandleFunc("/contact", contact)
	mux.HandleFunc("/hobby", hobby)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}