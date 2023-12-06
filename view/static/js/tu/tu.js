class RequestPDF {
  constructor() {
    this.status = '';
    this.dokumen_name = '';
    this.dokumen_pdf = '';
    this.message = '';
  }
  
  setStatus(status) {
    this.status = status
  }
  
  setDokName(dokumen) {
    this.dokumen_name = dokumen;
  }
  setMessage(m) {
    this.message = m;
  }
  
  setDokPdf(pdf) {
    this.dokumen_pdf = pdf;
  }
}

const url = window.location.href
const idSurat =  url.split('/')[4]

// TTD Section
const submitTTDButton = document.getElementById("submit-file");

submitTTDButton.addEventListener("click", (e) => {
    e.preventDefault();
    
    const pendukung = document.querySelector('#file')

    const selectedFile = pendukung.files[0];
    if (!selectedFile) {
      alert('Pilih file terlebih dahulu.');
      return;
    }
    const requestPDF = new RequestPDF()

    const fileReader = new FileReader();
    fileReader.onload = async function () {
      const fileData = fileReader.result.split(',')[1]; // Ambil data base64 saja

      // Buat objek JSON dengan data formulir dan file base64
      requestPDF.setDokName(selectedFile.name);
      requestPDF.setDokPdf(fileData);
    
      // Kirim data ke server dalam format JSON
      await fetch('/wt/'+idSurat+'/sppd', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(requestPDF)
      })
      .then(response => response.json())
      .then(responseJson => {
        if (responseJson.code == 200){
          ok =  alert(responseJson.data.message);
          if (!ok){
            setTimeout(()=>{
              window.location.reload();
              window.location.href = '/wt/sppd' 
            },500)
          }
        }
        else{
          alert(responseJson.code);
        }
      })
    };

    fileReader.readAsDataURL(selectedFile);
    
});

