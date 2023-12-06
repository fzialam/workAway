class RequestPDF {
  constructor() {
    this.status = '';
    this.dok_name = '';
    this.dok_pdf = '';
    this.message = '';
  }

  setStatus(status) {
    this.status = status
  }

  setDokName(dokumen) {
    this.dok_name = dokumen;
  }
  setMessage(m) {
    this.message = m;
  }

  setDokPdf(pdf) {
    this.dok_pdf = pdf;
  }
}

const url = window.location.href
const idSurat =  url.split('/')[4]

// TTD Section
const submitTTDButton = document.getElementById("submit-file");

submitTTDButton.addEventListener("click", (e) => {
    e.preventDefault(); // Mencegah perilaku default formulir
    
    const pendukung = document.querySelector('#file')

    const selectedFile = pendukung.files[0];
    if (!selectedFile) {
      alert('Pilih file terlebih dahulu.');
      return;
    }
    const requestPDF = new RequestPDF();
    requestPDF.setStatus('1')
    requestPDF.setMessage('OK')
    
    const fileReader = new FileReader();
    fileReader.onload = async function () {
      const fileData = fileReader.result.split(',')[1]; // Ambil data base64 saja

      // Buat objek JSON dengan data formulir dan file base64
      requestPDF.setDokName(selectedFile.name);
      requestPDF.setDokPdf(fileData);
    
      // Kirim data ke server dalam format JSON
      try {
        const response = await fetch('/wpp/'+idSurat+'/sppd', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify(requestPDF)
        });

        const responseData = await response.json();
        if (responseData['code'] == 400){
          alert("400 : BAD REQUEST");
        }
        else if(responseData['code'] == 500){
          alert(responseData['status']);
        }
      } catch (error) {
        // Tangani kesalahan jika ada
        console.error(error);
      }
    };

    fileReader.readAsDataURL(selectedFile);
    
});

// Tolak Section
const tolakButton = document.getElementById('tolak');

tolakButton.addEventListener('click', (e) => {
  e.preventDefault();
  sure = prompt('Beri Pesan Alasan');
  if (sure){
    const requestPDF = new RequestPDF();
    requestPDF.setStatus('2');
    requestPDF.setDokName(pdfName);
    requestPDF.setDokPdf(pdfBase64);
    requestPDF.setMessage(sure);

    fetch('/wpp/'+idSurat+'/sppd', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(requestPDF)
    })
    .then(response => response.json())
    .then(responseJson => {
        if (responseJson.code == 200){
            alert(responseJson.data.message);
        } else {
            alert(responseJson.status);
        }
    })
    return
  } else{
    window.location.reload();
  }
})