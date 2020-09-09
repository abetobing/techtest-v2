# Recommendation

## 1. Return Standard HTTP Status (RFC 7231)
Contoh:
* 401 untuk unauthorized
* 404 untuk resource not found. Contoh untuk: `GET /customer/list`
* 201 untuk resource yg berhasil ter-create. Contoh untuk: `POST /customer/create`

## 2. Enkripsi Password
Password masih di simpan sebagai plain text di kolom password.<br>
Action items:
* Password di encrypt dengan standar `sha512`.
* Menambah field baru: `salt` di table administrators sebagai salt untuk password.
* Inisialisasi data user admin di `schema.sql` dipindah ke `bootstrap.go`.
* Mengganti definisi table dengan primary key `id`.<br>
  Diganti menjadi SERIAL (auto-increment).


## 3. Group Router
Group router, untuk guest dan protected URI, menggunakan middleware yg berbeda.<br>
Alasan: tidak semua URI membutuhkan proteksi autentikasi dengan token. 
Contoh: `POST /auth/login`.<br>
Action items:
* Grouping router untuk guest dan authenticated users
* Halt http handlers sampai ke titik unauthorized (proses tidak diteruskan ke middleware lain)


## 4. SQL Injection
Menghindari kemungkinan sql injection dengan (minimal) menggunakan argument placeholder saat meng-query ke db.
Action items:
* Implementasi sederhana menggunakan native database/sql parameterized query interface.


## 5. Return json sesuai standard rest: dengan header “Content-Type: application/json”
Header saat ini masih default: `text/plain`. Diganti menjadi `application/json`
Action items:
* Middleware custom untuk menambahkan response header


## 6. Production/Dev/Local Environment
Profiling aplikasi dengan menseparasi environment *dev* dan *production*.<br>
Dengan passing argument pertama saat start aplikasi:

    $ ./app prod   --> menggunakan config file `config-prod.yml`
    $ ./app dev    --> menggunakan config file `config-dev.yml`
    $ ./app        --> menggunakan default config file `config.yml`

Catatan : 
 * Default config untuk development adalah `config.yml`
 * Untuk aplikasi yang di start dengan `docker-compose up`, profile yg digunakan adalah `prod`


## 7. CORS
Implementasi standard browser security policy, CORS.<br>
Limitasi akses endpoint-endpoint di API khusus dari origin tertentu.<br>
Untuk saat ini masih dibuka untuk semua domain.<br>
> `AllowedOrigins: *`

## 8. JWT Token
Mengganti authentication token ke authentikasi standar industri, JSON Web Token.
Pertimbangan:
* Lebih panjang, mengurangi kemungkinan brute force
* Bisa menyimpan informasi tambahan didalam JWT [claims](https://auth0.com/docs/tokens/json-web-tokens/json-web-token-claims). Sehingga memungkinkan identifikasi/profiling user berdasarkan:
  * device_id (untuk firebease)
  * user_agent
  * msisdn

Sebelumnya, untuk authentikasi hanya mengirimkan query string `token=jE1OTk1NzU3NDIsInVzZXJfaWQiOjF9` di setiap request.<br>Dengan JWT, token di kirimkan di header:

> `Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTk1NzU3NDIsInVzZXJfaWQiOjF9.eUzAQfZYlctoyzHW4ZNN4muQH-NF-WfeblVvgsrW5VA`

## 9. Data integrity
* Saat deletion customer, data address di customer_address perlu di delete
* Implementasi soft delete *(Kendala waktu tidak cukup)*

## 10. API Compatibility
Terkait dengan poin nomor 8, yang mengganti metode authentication existing, maka perlu dipastikan agar API tetap men-serve dengan authentication versi yang lama. Pertimbangannya adalah agar client yang masih menggunakan authentication yg lama tetap masih bisa meng-consume API.
<br>
<br>
<br>

# Poin Tambahan & Belum Terimplementasi
Beberapa poin tambahan dibawah 

### 11. Standarisasi REST method untuk CRUD:
Poin ini lebih ke best practice REST dengan pattern endpoint yg lebih simple.

<table>
    <thead>
        <tr>
            <td><strong>METHOD</strong></td>
            <td><strong>BEFORE</strong></td>
            <td><strong>AFTER</strong></td>
        </tr>
    </thead>
    <tbody>
        <tr>
            <td><code>GET</code></td>
            <td><code>/customer/list</code></td>
            <td><code>/customers</code></td>
        </tr>
        <tr>
            <td><code>GET</code></td>
            <td><code>/customer/get/{id}</code></td>
            <td><code>/customer/{id}</code></td>
        </tr>
        <tr>
            <td><code>POST</code></td>
            <td><code>/customer/create</code></td>
            <td><code>/customer</code></td>
        </tr>
        <tr>
            <td><code>DELETE</code></td>
            <td><code>/customer/delete/{id}</code></td>
            <td><code>/customer/{id}</code></td>
        </tr>
    </tbody>
</table> 
<br>

### 12. Paging
Memecah output dari listing customer menjadi beberapa page.
### 13. Logger
Enhance logging output dari aplikasi sehingga bisa mendeteksi bug dengan mudah.
<br>
<br>
<br>

# Catatan
* Urutan commit di tidak sesuai dengan urutan poin-poin diatas.<br>
* Commit message menggunakan contoh format: <br>
  `[N] implementasi untuk poin ke-N` <br>
  dimana `N` mewakili implementasi poin-poin diatas.<br>
* Contoh, untuk poin kedua, commit messagenya adalah:<br>
  `[2] Return HTTP Status code sesuai dengan error` 

