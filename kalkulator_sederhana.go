/*
Created By Anwar Sanusi
*/

package main

import (
   "fmt";
   "math";
)

var (
   menuutama, menu int
   celcius,farenheit,reamur int
   a,b,mas,vol,f,s,w,t,vo float64
   kelvin,cel,far,j,luas,m,v,g,h,x  float64
   pi float64 = 3.14
)

func main() {
   fmt.Println("=================WELCOME================")
   fmt.Println("Menu Aplikasi Perhitungan")
   fmt.Println("1. Luas Persegi")
   fmt.Println("2. Luas Segitiga")
   fmt.Println("3. Luas Lingkaran")
   fmt.Println("4. Sudut Cosinus, Sinus, dan Tangen")
   fmt.Println("5. Gerak Lurus Beraturan")
   fmt.Println("6. Gerak Lurus Berubah Beraturan")
   fmt.Println("7. Energi Kinetik dan Energi Potensial")
   fmt.Println("8. Frekuensi dan Getaran")
   fmt.Println("9. Massa Jenis")
   fmt.Println("10. Gaya, Tekanan, Usaha dan Daya")
   fmt.Println("11. Konversi Suhu")
   fmt.Println("12. Exit")
   fmt.Println("=========================================")
   fmt.Print("Pilih salah satu menu diatas (masukkan nomor menu): ")
   fmt.Scanf("%d", &menuutama)

switch menuutama{
   case 1 :
      LuasPersegi()
   case 2 :
      LuasSegitiga()
   case 3 :
      LuasLingkaran()
   case 4 :
      SdtCoSinTan()
   case 5 :
      GLB()
   case 6 :
      GLBB()
   case 7 :
      KinetikPoten()
   case 8 :
      FrekuensiGet()
   case 9 :
      MassaJenis()
   case 10 :
      GTUD()
   case 11 :
      Suhu()
   default :
      exit()
   }
}

//=====================LUAS PERSEGI=======================//
func LuasPersegi(){
   fmt.Println("=====Mengukur Luas Persegi=====")
   fmt.Print("Masukkan panjang persegi : ")
   fmt.Scanf("%f\n", &a)
   fmt.Print("Masukkan lebar persegi : ")
   fmt.Scanf("%f\n", &b)
   luas := a*b
   fmt.Println("Luas Persegi= ", luas)
   main()
}

//=======================LUAS SEGITIGA====================//
func LuasSegitiga(){
   fmt.Println("=====Mengukur Luas Segitiga=====")
   fmt.Print("Masukkan alas segitiga : ")
   fmt.Scanf("%f", &a)
   fmt.Print("Masukkan tinggi segitiga : ")
   fmt.Scanf("%f", &b)
   luas := (a*b)/2
   fmt.Println("Luas Segitiga= \n", luas)
   main()
}

//===================LUAS LINGKARAN======================//
func LuasLingkaran(){
   fmt.Println("=====Mengukur Luas Lingkaran=====")
   fmt.Print("Masukkan jari-jari lingkaran : ")
   fmt.Scanf("%f", &j)
   luas = pi*(j*j)
   fmt.Println("Luas Lingkaran= \n", luas)
   main()
}

//============SUDUT COSINUS, SINUS, DAN TANGEN=========//
func SdtCoSinTan(){
   fmt.Print("Masukkan nilai x = ")
   fmt.Scanf("%f", &x)
   sin := math.Sin(x)
   cos := math.Cos(x)
   tan := math.Tan(x)

   fmt.Println("Hasil hitung sudut Sin = ", sin)
   fmt.Println("Hasil hitung sudut Cos = ", cos)
   fmt.Println("Hasil hitung sudut Tan = ", tan)
   main()
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
   main()
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

//========================EK dan EP==================//
func KinetikPoten(){
   fmt.Println("1. Energi Kinetik")
   fmt.Println("2. Energi Potensial")
   fmt.Println("3. Menu Utama")
   fmt.Print("Pilih salah satu menu diatas ini (ketik nomor menu): ")
   fmt.Scanf("%d", &menu)

   if menu == 1 {
      kinetik()
   }else if menu == 2{
      potensial()
   }else{
      main()
   }
}

func kinetik(){
   fmt.Println("=====Mengukur Energi Kinetik=====")
   fmt.Print("Masukkan nilai massa = ")
   fmt.Scanf("%f", &m)
   fmt.Print("Masukkan nilai kecepatan = ")
   fmt.Scanf("%f", &v)
   kinetik := 0.5*(m*(v*v))
   fmt.Println("Besar Energi Kinetik = \n", kinetik)
   KinetikPoten()
}

func potensial(){
   fmt.Println("=====Mengukur Energi Potensial=====")
   fmt.Print("Masukkan nilai massa = ")
   fmt.Scanf("%f", &m)
   fmt.Print("Masukkan nilai gravitasi = ")
   fmt.Scanf("%f", &g)
   fmt.Print("Masukkan nilai tinggi = ")
   fmt.Scanf("%f", &h)
   potensial := m*g*h
   fmt.Println("Besar Energi Kinetik = \n", potensial)
   KinetikPoten()
}

//=================FREKUENSI atau GETARAN=================//
func FrekuensiGet() {
   fmt.Println("=====Mengukur Frekuensi atau Getaran=====")
   fmt.Println("Pilih menu sesuai keinginan")
   fmt.Println("1. Frekuensi")
   fmt.Println("2. Getaran")
   fmt.Println("3. Menu Utama")
   fmt.Println("Masukkan pilihan menu (masukkan nomor menu) ")
   fmt.Scanf("%d", &menu)

   switch menu{
      case 1 :
         frekuensi()
      case 2 :
         getaran()
      default :
         main()
   }
}

func frekuensi(){
   fmt.Println("=====Mengukur Frekuensi=====")
   fmt.Print("Masukkan periode = ")
   fmt.Scanf("%f", &g)
   frek := 1 / g
   fmt.Println("Nilai Massa Jenis = ", frek)
   main()
}

func getaran(){
   fmt.Println("=====Mengukur Getaran=====")
   fmt.Print("Masukkan jumlah getaran = ")
   fmt.Scanf("%f", &g)
   fmt.Print("Masukkan nilai t = ")
   fmt.Scanf("%f", &t)
   frek := g / t
   fmt.Println("Nilai Massa Jenis = ", frek)
   main()
}

//======================MASSA JENIS==================//
func MassaJenis(){
   fmt.Println("=====Mengukur Massa Jenis=====")
   fmt.Print("Masukkan nilai massa = ")
   fmt.Scanf("%f", &mas)
   fmt.Print("Masukkan nilai volume = ")
   fmt.Scanf("%f", &vol)
   massa := m*v
   fmt.Println("Nilai Massa Jenis = ", massa)
   main()
}

//===========GAYA, TEKANAN, USAHAN DAN DAYA==========//
func GTUD(){
   fmt.Println("Menu")
   fmt.Println("1. Gaya")
   fmt.Println("2. Tekanan")
   fmt.Println("3. Usaha")
   fmt.Println("4. Daya")
   fmt.Println("5. Memu Utama")
   fmt.Print("Pilih salah satu menu diatas (ketik nomor menu): ")
   fmt.Scanf("%d", &menu)

switch menu{
   case 1 :
      Gaya();
   case 2 :
      Tekanan()
   case 3 :
      Usaha()
   case 4 :
      Daya()
   default :
      main()
   }
}

func Gaya(){
   fmt.Println("=====Mengukur Besar Gaya=====")
   fmt.Print("Masukkan nilai massa = ")
   fmt.Scanf("%f", &m)
   fmt.Print("Masukkan nilai percepatan = ")
   fmt.Scanf("%f", &a)
   gaya := mas*a
   fmt.Print("Besar Gaya = ", gaya)
   GTUD()
}

func Tekanan(){
   fmt.Println("=====Mengukur Besar Tekanan=====")
   fmt.Print("Masukkan nilai gaya = ")
   fmt.Scanf("%f", &f)
   fmt.Print("Masukkan nilai percepatan = ")
   fmt.Scanf("%d", &a)
   tekanan := f*a
   fmt.Print("Besar Tekanan = ", tekanan)
   GTUD()
}

func Usaha(){
   fmt.Println("=====Mengukur Besar Usaha=====")
   fmt.Print("Masukkan nilai gaya = ")
   fmt.Scanf("%f", &f)
   fmt.Print("Masukkan jarak = ")
   fmt.Scanf("%f", &s)
   usaha := f*s
   fmt.Print("Besar Usaha = ", usaha)
   GTUD()
}

func Daya(){
   fmt.Println("=====Mengukur Besar Daya=====")
   fmt.Print("Masukkan nilai usaha = ")
   fmt.Scanf("%f", &w)
   fmt.Print("Masukkan waktu = ")
   fmt.Scanf("%f", &t)
   daya := w/t
   fmt.Print("Besar Daya = ", daya)
   GTUD()
}

//========================SUHU==================//
func Suhu(){
   fmt.Println("Konversi Suhu")
   fmt.Println("1. Celcius")
   fmt.Println("2. Farenheit")
   fmt.Println("3. Reamur")
   fmt.Println("4. Kelvin")
   fmt.Println("5. Menu Utama")
   fmt.Print("Pilih salah satu suhu yang ingin dikonversi pada menu diatas (ketik nomor menu): ")
   fmt.Scanf("%d", &menu)

switch menu{
   case 1 :
      fmt.Println("1. Celcius - Farenheit")
      fmt.Println("2. Celcius - Reamur")
      fmt.Println("3. Celcius - Kelvin")
      fmt.Println("4. Menu Utama")
      fmt.Print("Pilih salah satu menu diatas (ketik nomor menu): ")
      fmt.Scanf("%d", &celcius)

      switch celcius{
         case 1 :
            CF()
         case 2 :
            CR()
         case 3 :
            CK()
         default :
            Suhu()
      }
   case 2 :
      fmt.Println("1. Farenheit - Celcius")
      fmt.Println("2. Farenheit - Reamur")
      fmt.Println("3. Menu Utama")
      fmt.Print("Pilih salah satu menu diatas (ketik nomor menu): ")
      fmt.Scanf("%d", &farenheit)

      switch farenheit{
         case 1 :
            FC()
         case 2 :
            FR()
         default :
            Suhu()
      }
   case 3 :
      fmt.Println("1. Reamur - Celcius")
      fmt.Println("2. Reamur - Farenheit")
      fmt.Println("3. Menu Utama")
      fmt.Print("Pilih salah satu menu diatas (ketik nomor menu): ")
      fmt.Scanf("%d", &reamur)

      switch reamur{
         case 1 :
            RC()
         case 2 :
            RF()
         default :
            Suhu()
         }
   case 4 :
      fmt.Println("konversi Kelvin - Celcius")
      fmt.Print("Masukkan nilai suhu Kelvin = ")
      fmt.Scanf("%f", &kelvin)
      KC := kelvin-273
      fmt.Print("Nilai Suhu dalam Celcius = ", KC)
      Suhu()
   default :
      main()
}
}

//-----------------------------------------------------------------------------------------------//
//CELCIUS

func CF(){
   var cek int
   fmt.Println("konversi Celcius - Farenheit")
   fmt.Print("Masukkan nilai suhu Celcius = ")
   fmt.Scanf("%f", &cel)
   hasil := ((9*cel)+32)/5
   fmt.Println("Nilai Suhu dalam Fareinheit =  ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

func CR(){
   var cek int
   fmt.Println("konversi Celcius - Reamur")
   fmt.Print("Masukkan nilai suhu Celcius = ")
   fmt.Scanf("%f", &cel)
   hasil := cel*4/5
   fmt.Println("Nilai Suhu dalam Reamur = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
      exit()
   default :
      Suhu()
   }
}

func CK(){
   var cek int
   fmt.Println("konversi Celcius - Kelvin")
   fmt.Print("Masukkan nilai suhu Celcius = ")
   fmt.Scanf("%f", &cel)
   hasil := cel+273
   fmt.Println("Nilai Suhu dalam Kelvin = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

//--------------------------------------------------------------------------------------------//
// FARENHEIT

func FC(){
   var cek int
   fmt.Println("konversi Farenheit - Celcius")
   fmt.Print("Masukkan nilai suhu Farenheit = ")
   fmt.Scanf("%f", &far)
   hasil := 5/9*(far-32)
   fmt.Println("Nilai Suhu dalam Celcius = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
        exit()
      default :
         Suhu()
   }
}

func FR(){
   var cek int
   fmt.Println("konversi Farenheit - Reamur")
   fmt.Print("Masukkan nilai suhu Farenheit = ")
   fmt.Scanf("%f", &far)
   hasil := 4*(far-32)/9
   fmt.Println("Nilai Suhu dalam Reamur = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

//-------------------------------------------------------------------------------------------//
// REAMUR

func RC(){
   var r1 float64
   var cek int

   fmt.Println("konversi Reamur - Celcius")
   fmt.Print("Masukkan nilai suhu Reamur = ")
   fmt.Scanf("%f", &r1)
   hasil := (r1*5)/4
   fmt.Println("Nilai Suhu dalam Celcius = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

func RF(){
   var r2 float64
   var cek int

   fmt.Println("konversi Reamur - Farenheit")
   fmt.Print("Masukkan nilai suhu Reamur = ")
   fmt.Scanf("%f", &r2)
   hasil := (r2*9)/4+32
   fmt.Println("Nilai Suhu dalam Farenheit = ", hasil)

   fmt.Println("Apa anda ingin keluar ? (jika ya tekan 1, jika tidak tekan 0) ")
   fmt.Scanf("%d", &cek)

   switch cek{
      case 1 :
         exit()
      default :
         Suhu()
   }
}

//=========================EXIT=====================//
func exit(){
   fmt.Println("====Selesai===")
}