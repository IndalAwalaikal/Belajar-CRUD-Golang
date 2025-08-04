package model

import (
	"database/sql"
	"fmt"
)

type Transaction struct {
	IdTransaksi int
	Jumlah      int   
	Kategori    string
	Tanggal     string
	Jenis       string
	Catatan     string 
}

func (t Transaction) CreateTransaksi(db *sql.DB) error {
	_, err := db.Exec("INSERT INTO transaksi (jumlah, kategori, tanggal, jenis, catatan) VALUES (?, ?, ?, ?, ?)", t.Jumlah, t.Kategori, t.Tanggal, t.Jenis, t.Catatan)
	if err != nil {
		return err
	}
	return nil
}

func (t Transaction) ReadAllTransaksi(db *sql.DB) ([]Transaction, error) {
	rows, err := db.Query("SELECT id_transaksi, jumlah, kategori, tanggal, jenis, catatan FROM transaksi")
	if err != nil {
		return nil, err
	}
	transaksi := []Transaction{}
	for rows.Next() {
		err := rows.Scan(&t.IdTransaksi, &t.Jumlah, &t.Kategori, &t.Tanggal, &t.Jenis, &t.Catatan)
		if err != nil {
			panic(err)
		}
		transaksi = append(transaksi, t)
	}
	fmt.Print(t)
	return transaksi, nil
}

func (t Transaction) UpdateTransaksi(db *sql.DB) error {
	_, err := db.Exec("UPDATE transaksi SET jumlah = ?, kategori = ?, tanggal = ?, jenis = ?, catatan = ? WHERE id_transaksi = ?", t.Jumlah, t.Kategori, t.Tanggal, t.Jenis, t.Catatan, t.IdTransaksi)
	if err != nil {
		panic(err)
	}
	return nil
}

func (t Transaction) DeleteTransaksi(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM transaksi WHERE id_transaksi = ?", id)
	if err != nil {
		panic(err)
	}
	fmt.Print(t)
	return nil
}

func (t Transaction) ReadTransaksiById(db *sql.DB, id int) (Transaction, error) {
	row := db.QueryRow("SELECT id_transaksi, jumlah, kategori, tanggal, jenis, catatan FROM transaksi WHERE id_transaksi = ?", id)
	err := row.Scan(&t.IdTransaksi, &t.Jumlah, &t.Kategori, &t.Tanggal, &t.Jenis, &t.Catatan)
	if err != nil {
		panic(err)
	}
	return t, nil
}