package main

import (
	"keuangan/controller"
	"keuangan/config"
	"net/http"
)

func main() {
	db := config.ConnectToDB()

	http.HandleFunc("/", controller.ShowHtml(db))
	http.HandleFunc("/create", controller.CreateTransaksi(db))
	http.HandleFunc("/edit", controller.UpdateTransaksi(db))
	http.HandleFunc("/delete", controller.DeleteTransaksi(db))
	http.Handle("/view/", http.StripPrefix("/view/", http.FileServer(http.Dir("view"))))
	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		panic(err)
	}
}