# Golang Embed

## Agenda

- Pengenalan Embed Package
- Embed File ke String
- Embed File ke Byte[]
- Embed Multiple File
- Hasil Embed di Compile
- Dan lain-lain

## Pengenalan Embed Package

- Sejak Golang versi 1.16, terdapat package baru dengan nama `embed`
- Package embed adalah fitur baru untuk mempermudah membaca isi file pada saat compile time secara otomatis dimasukkan isi filenya ke dalam sebuah variable yang kita tuju

## Cara Embed File

- Untuk melakukan emded file ke variable, kita bisa mengimport package embed terlebih dahulu
- Selanjutnya kita bisa tambahkan komentar `//go:embed` diikuti dengan `nama filenya`, diatas variable yang kita tuju
- Variable yang dituju tersebut nanti secara otomatis akan berisi konten file yang kita inginkan secara otomatis ketika kode golang di compile
- Variable yang dituju `tidak bisa` disimpan didalam function (artinya harus diluar package nya)
