/*
Created By Anwar Sanusi
*/

package main

import (
   "fmt";
)

var (
   menuutama, menu int
   celcius,farenheit,reamur int
   a,s,t,vo float64
)

func main() {
   fmt.Println("=================WELCOME================")
   fmt.Println("Menu Aplikasi Perhitungan")
   fmt.Println("1. Gerak Lurus Beraturan")
   fmt.Println("2. Gerak Lurus Berubah Beraturan")
   fmt.Println("3. Exit")
   fmt.Println("=========================================")
   fmt.Print("Pilih salah satu menu diatas (masukkan nomor menu): ")
   fmt.Scanf("%d", &menuutama)

switch menuutama{
   case 1 :
      GLB()
   case 2 :
      GLBB()
   default :
      exit()
   }
}


//======================GLB======================//
func GLB(){
   fmt.Println("=====Mengukur Gerak Lurus Beraturan=====")
   fmt.Print("Masukkan jarak = ")
   fmt.Scanf("%f", &s)
   fmt.Print("Masukkan waktu = ")
   fmt.Scanf("%f", &t)
   v := s*t
   fmt.Println("Nilai Gerak Lurus Beraturan = ", v)
}

//=======================GLBB=====================//
func GLBB() {
   fmt.Println("=====Mengukur Gerak Lurus Berubah Beraturan=====")
   fmt.Println("Pilih menu sesuai keinginan")
   fmt.Println("1. S = Vo + 1/2 * a*t*t")
   fmt.Println("2. Vt = Vo - a.t")
   fmt.Println("3. Vt2 = Vo2 + 2*a*S")
   fmt.Println("4. Menu Utama")
   fmt.Println("Masukkan pilihan menu (masukkan nomor menu) ")
   fmt.Scanf("%d", &menu)

   switch menu{
      case 1 :
         nomor1()
      case 2 :
          nomor2()
      case 3 :
         nomor3()
      default :
         main()
   }
}

func nomor1(){
   fmt.Println("=====Nomor 1=====")
   fmt.Print("Masukkan kecepatan awal = ")
   fmt.Scanf("%f", &vo)
   fmt.Print("Masukkan nilai percepatan = ")
   fmt.Scanf("%f", &a)
   fmt.Print("Masukkan nilai waktu= ")
   fmt.Scanf("%f", &t)
   GLBB := vo + 0.5 * a*t*t
   fmt.Println("Nilai Gerak Lurus Beraturan = ", GLBB)
}

func nomor2(){
   fmt.Println("=====Nomor 2=====")
   fmt.Print("Masukkan kecepatan awal = ")
   fmt.Scanf("%f", &vo)
   fmt.Print("Masukkan nilai percepatan = ")
   fmt.Scanf("%f", &a)
   fmt.Print("Masukkan nilai waktu= ")
   fmt.Scanf("%f", &t)
   GLBB := vo-a*t
   fmt.Println("Nilai Gerak Lurus Beraturan = ", GLBB)
}

func nomor3(){
   fmt.Println("=====Nomor 3=====")
   fmt.Print("Masukkan kecepatan awal = ")
   fmt.Scanf("%f", &vo)
   fmt.Print("Masukkan nilai percepatan = ")
   fmt.Scanf("%f", &a)
   fmt.Print("Masukkan nilai jarak= ")
   fmt.Scanf("%f", &s)
   GLBB := (vo*vo) + 2*a*s
   fmt.Println("Nilai Gerak Lurus Beraturan = ", GLBB)
}

//=========================EXIT=====================//
func exit(){
   fmt.Println("====Selesai===")
}