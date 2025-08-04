package controller

import (
	"database/sql"
	"html/template"
	"keuangan/model"
	"keuangan/view"
	"net/http"
	"strconv"
)

func ShowHtml(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		transaksi := model.Transaction{}

		transaction, err := transaksi.ReadAllTransaksi(db)
		if err != nil {
			panic(err)
		}
		tmpl, err := template.ParseFiles("view/index.html")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, transaction)
	}
}

func CreateTransaksi(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if http.MethodPost == r.Method {
			jumlahStr := r.FormValue("jumlah")
			kategori := r.FormValue("kategori")
			tanggal := r.FormValue("tanggal")
			jenis := r.FormValue("jenis")
			catatan := r.FormValue("catatan")

			jumlah, err := strconv.Atoi(jumlahStr)
			if err != nil {
				panic(err)
			}

			transaksi := model.Transaction{
				Jumlah: jumlah,
				Kategori: kategori,
				Tanggal: tanggal,
				Jenis: jenis,
				Catatan: catatan,
			}
			err = transaksi.CreateTransaksi(db)
			if err != nil {
				panic(err)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		tmpl, err := template.ParseFiles("view/create.html")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, nil)
	}
}

func ReadAllTransaksi(db *sql.DB) {
	transaksi := model.Transaction{}
	transaction, err := transaksi.ReadAllTransaksi(db)
	if err != nil {
		panic(err)
	}

	view.Output("Data Transaksi: ", transaction)
}

func UpdateTransaksi(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id_string := r.URL.Query().Get("id_transaksi")
		id_int, err := strconv.Atoi(id_string)
		if err != nil {
			panic(err)
		}

		if r.Method == http.MethodPost {
			jumlahStr := r.FormValue("jumlah")
			kategori := r.FormValue("kategori")
			tanggal := r.FormValue("tanggal")
			jenis := r.FormValue("jenis")
			catatan := r.FormValue("catatan")

			jumlah, err := strconv.Atoi(jumlahStr)
			if err != nil {
				panic(err)
			}

			transaksi := model.Transaction{
				IdTransaksi: id_int,
				Jumlah:      jumlah,
				Kategori:    kategori,
				Tanggal:     tanggal,
				Jenis:       jenis,
				Catatan:     catatan,
			}
			err = transaksi.UpdateTransaksi(db)
			if err != nil {
				panic(err)
			}
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}

		transaksi := model.Transaction{IdTransaksi: id_int}
		transaction, err := transaksi.ReadTransaksiById(db, id_int)
		if err != nil {
			panic(err)
		}
		tmpl, err := template.ParseFiles("view/edit.html")
		if err != nil {
			panic(err)
		}
		tmpl.Execute(w, transaction)
	}
}

func DeleteTransaksi(db *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		id_string := r.URL.Query().Get("id_transaksi")
		id_int, err := strconv.Atoi(id_string)
		if err != nil {
			panic(err)
		}

		transaksi := model.Transaction{
			IdTransaksi: id_int,
		}
		u, err := transaksi.ReadTransaksiById(db, id_int)
		if err != nil {
			panic(err)
		}

		err = transaksi.DeleteTransaksi(db, u.IdTransaksi)
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
