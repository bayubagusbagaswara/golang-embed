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

## Embed File ke String

- Embed file bisa kita lakukan ke variable dengan tipe data String
- Secara otomatis isi file akan dibaca sebagai text dan di masukkan ke variable tersebut

## Embed File ke []byte atau Slice of Byte

- Selain ke tipe data String, embed file juga bisa dilakukan ke variable tipe data []byte
- Ini cocok sekali jika kita ingin melakukan embed file dalam bentuk binary, seperti gambar dan lain-lain

## Embed Multiple Files

- Kadang ada kebutuhan kita ingin melakukan embed beberapa file sekaligus
- Hal ini juga bisa dilakukan menggunakan embed package
- Kita bisa menambahkan komentar `//go:embed lebih dari satu baris`
- Selain itu variable nya bisa kita gunakan tipe data `embed.FS`

## Path Matcher

- Selain manual satu per satu, kita bisa menggunakan path matcher untuk membaca multiple file yang kita inginkan
- Ini sangat cocok ketika misal kita punya `pola jenis file` yang kita inginkan untuk kita baca
- Caranya, kita perlu menggunakan path matcher seperti pada package function `path.Match`

## Hasil Embed di Compile

- Perlu diketahui, bahwa hasil embed yang dilakukan oleh package embed adalah permanent dan data file yang dibaca akan disimpan dalam binary file golang nya
- Artinya bukan dilakukan secara realtime membaca file yang ada diluar
- Hal ini menjadikan jika binary file golang sudah di compile, kita tidak butuh lagi file external nya, dan bahkan jika diubah file external nya, isi variablenya tidak akan berubah lagi
- cara running file binary : `golang-embed.exe`
