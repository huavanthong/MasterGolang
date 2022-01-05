package main

import (
	"fmt"
	"net/http"
)

func BootApplication() {
	// Run Bootstrapping Application
	fmt.Printf("Bootstrapping the Application")

	http.HandleFunc("/employee", controllers.GetEmployee)

	err := http.ListenAndServe(":8080", nil)  
	if err != null {
		panic(err.Error())
	}
}