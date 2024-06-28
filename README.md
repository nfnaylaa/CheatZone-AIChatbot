# CheatZone

CheatZone adalah aplikasi web yang dirancang untuk memberikan bantuan sehari-hari melalui model AI. Aplikasi ini mendukung tiga fitur utama:
1. **Table QA**: Menerima file CSV dari pengguna dan memungkinkan pengguna untuk mengajukan pertanyaan mengenai data tersebut. Fitur ini menggunakan model [google/tapas-base-finetuned-wtq](https://huggingface.co/google/tapas-base-finetuned-wtq).
2. **Summarization**: Merangkum teks yang ditulis di kolom chat oleh pengguna. Fitur ini menggunakan model [Falconsai/text_summarization](https://huggingface.co/Falconsai/text_summarization).
3. **Chat Me**: Menjawab pertanyaan pengguna mirip seperti ChatGPT. Fitur ini menggunakan model [meta-llama/Meta-Llama-3-8B-Instruct](https://huggingface.co/meta-llama/Meta-Llama-3-8B-Instruct)

## Komponen Utama

- **Frontend**: HTML, CSS, JavaScript
- **Backend**: Go (Gin framework)
- **AI Models**: Menggunakan [Hugging Face API](https://huggingface.co/)
- **Modul** : [The Hugging Face Inference Client](https://github.com/hupe1980/go-huggingface)

- ## Instalasi
  
1. **Clone Repository ini**
     ```bash
     git clone https://github.com/username/cheatzone.git
     cd cheatzone/final
     ```
2. Install Dependencies
     ```bash
     go mod tidy
     ```
3. Buat Token di Hugging Face
   - Buka [Hugging Face](https://huggingface.co/)
   - Daftar atau masuk ke akun Anda.
   - Klik pada gambar profil Anda di sudut kanan atas dan pilih "Settings".
   - Di menu sebelah kiri, pilih "Access Tokens".
   - Klik tombol "New token" untuk membuat token baru.
   - Beri nama token Anda, pilih "Read" permissions, dan klik "Generate".
   - Salin token yang dihasilkan dan tempelkan ke dalam file .env Anda di HUGGING_FACE_TOKEN.
4. Masukan Token ke file .env
    ```bash
     HUGGING_FACE_TOKEN="YOUR_TOKEN_HERE"
     ```
5. Jalankan Aplikasi
     ```bash
     go run main.go
     ```
6. Buka browser dan akses http://localhost:8080

## Penggunaan

**1. Table QA**
- Pilih model "Table QA" di sisi kiri.
- Unggah file CSV dengan menggunakan form yang tersedia.
- Ketik pertanyaan tentang data yang ada di dalam file CSV tersebut.
- Klik tombol "Send" untuk mendapatkan jawaban dari model.

  
**2. Summarization**
- Pilih model "Summarization" di sisi kiri.
- Ketik teks yang ingin dirangkum di kolom chat.
- Klik tombol "Send" untuk mendapatkan ringkasan dari teks yang diberikan.
  
**3. Chat Me**
- Pilih model "Chat Me" di sisi kiri.
- Ketik pertanyaan umum di kolom chat.
- Klik tombol "Send" untuk mendapatkan jawaban dari model.



