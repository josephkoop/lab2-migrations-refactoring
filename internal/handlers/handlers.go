package handlers

import(
	"net/http"
)

//Handler function for home route
func Home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome! This is Joseph Koop's territory. The project I have chosen is #7 - Gym/Fitness Class Scheduling System. Why? It's easier to work on something that you're interested in, and I have an interest in gyms.     \n"))
}

//Handler function for about route
func About(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("About me     \nName: Joseph Koop     \nField of Study: Information Technology     \nCurrent Residence: Spanish Lookout     \n"))
}

//Handler function for contact route
func Contact(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Email: 2023159207@ub.edu.bz     \nGitHub username: josephkoop (main account) + Joseph-Koop (school email account)     \n"))
}

//Handler function for hobby route
func Hobby(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hobby: Chess     \nI doubt this one needs explaining     \n"))
}