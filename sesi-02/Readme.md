# Scalable Web Service with Golang - DTS Kominfo

## Sesi 02

Pada Sesi dua mempelajari 
- Conditions <br>
    Kondisional pada pemrograman dapat digunakan untuk mengatur alur dari suatu program. <br>
    Pada bahasa Go, terdapat 2 macam kondisional yang dapat kita pakai yaitu if else dan switch
    - If Else <br>
    keyword if, else if dan else
    - Switch <br>
    keyword switch untuk membuat suatu kondisional. keyword default sama dengan
    keyword else yang dimana akan dieksekusi blocknya jika tidak ada kondisi yang terpenuhi.
    - Switch with relational operators <br>
    Kita juga dapat menggunakan switch dengan operator perbandingan seperti layaknya pada sebuah kondisional dengan keyword if, else if dan else. <br><br>
    Lalu kita juga bisa menggunakan kurung kurawal, pada scope block default. 
    Hal ini sangat bagus diterapkan jika kita ingin membuat lebih dari satu statement pada scope block default agar syntax kita lebih rapi dan mudah di maintain. <br>
    - Switch (fallthrough keyword) <br>
    Ketika kita ingin switch pada bahasa Go melanjutkan pengecekan kepada case selanjutnya walaupun suatu case telah terpenuhi kondisinya, maka kita bisa menggunakan keyword fallthrough. <br>
    - Nested Conditions <br>
     Kondisional bersarang merupakan sebuah proses kondisional yang didalamnya terdapat proses kondisional kembali. Kita dapat menggunakan if, else if , else, switch ataupun menggabungkan seluruhnya. <br><br>
     
- Loopings <br>
    Looping atau perulangan merupakan suatu proses berulang yang dimana proses tersebut akan berhenti jika memenuhi suatu kondisi.

- Array <br>
    Array pada bahasa Go merupakan sebuah tipe data untuk menyimpan sebuah kumpulan dari data-data. Data-data yang kita simpan pada sebuah array dalam bahasa Go harus memiliki tipe data yang sama, kecuali kita menyimpannya sebagai suatu interface kosong <br>
    Kita dapat membuat array dengan 2 macam cara yaitu dengan mendeklarasikannya terlebih dahulu tanpa memberi nilai apapun, dan kita juga bisa mendeklarasikan dan langsung menginisialisasikannya dengan memberikannya sebuah nilai. 

- Slice <br>
Slice merupakan suatu tipe data yang mirip dengan tipe data array, yang juga memiliki kegunaan untuk menyimpan satu atau lebih data. Namun tipe data slice dan array memiliki sifat yang berbeda. Slice tidak memiliki sifat ﬁxed-length  yang berarti panjang dari slice tidak tetap sehingga kita bisa dengan leluasa menentukan panjang dari slice nya. Slice termasuk dalam kategori reference type yang dimana jika kita melakukan copy terhadap suatu slice, dan kita ubah element dari yang kita copy tersebut, maka slice semulanya juga akan ikut terubah. <br> 
Cara membuat slice cukup mudah hampir mirip dengan jika kita membuat suatu array. Yang menjadi perbedaan adalah kita tidak perlu menuliskan panjang dari slice nya tidak seperti array. Contohnya seperti pada gambar dibawah ini. 
<br>
Slice (append function with ellipsis) <br>
Jika kita ingin memasukkan seluruh element-element pada suatu array ke dalam array lainnya, maka kita dapat menggunakan tanda ellipsis (...) atau tanda titik tiga berurut. 
<br>
Slice (copy function) <br>
Kita juga bisa menggunakan fungsi copy untuk meng-copy seluruh element pada sebuah slice ke dalam slice lainnya. Perlu diingat disini bahwa ketika kita mencoba untuk meng-copy sebuah slice kedalam slice lainnya, maka seluruh element pada slice lainnya tersebut akan ter-replace oleh  element-element yang di copy kannya. 


### Next Review materi tiap praktek Code


