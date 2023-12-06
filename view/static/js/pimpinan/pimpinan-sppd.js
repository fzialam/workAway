class AppproveSPPDJOINAnggaran{
    constructor() {
        this.rincian_id = 0;
        this.status = '';
        this.dok_name = '';
        this.dok_pdf = '';
        this.message = '';
    }

    setStatus(status) {
        this.status = status
    }
    setRincianId(r_id) {
        this.rincian_id = r_id
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

const url = window.location.href;
const idSurat =  url.split('/')[4];
const rincianId  = document.getElementById('rincianId').value;

// TTD Section
const submitTTDButton = document.getElementById("submit-file");
if (!rincianId) {
    submitTTDButton.disabled = true;
}

submitTTDButton.addEventListener("click", (e) => {
    e.preventDefault(); // Mencegah perilaku default formulir
    const pendukung = document.querySelector('#file')
    const selectedFile = pendukung.files[0];
    
    console.log(selectedFile);
    if (!selectedFile) {
      alert('Pilih file terlebih dahulu.');
      return;
    }
    const requestPDF = new AppproveSPPDJOINAnggaran();
    requestPDF.setStatus('1')
    requestPDF.setMessage('OK')
    requestPDF.setRincianId(parseInt(rincianId))
    
    const fileReader = new FileReader();
    fileReader.onload = async function () {
      const fileData = fileReader.result.split(',')[1]; // Ambil data base64 saja

      // Buat objek JSON dengan data formulir dan file base64
      requestPDF.setDokName(selectedFile.name);
      requestPDF.setDokPdf(fileData);

      console.log(requestPDF);
    
      // Kirim data ke server dalam format JSON
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
            ok = alert('Success');
            if (!ok) {
              setTimeout(()=>{
                window.location.reload();
                window.location.href = '/wpp/sppd';
              }, 300)
            }
          } else {
              alert(responseJson.status);
          }
      })
    };

    fileReader.readAsDataURL(selectedFile);
    
});

// Tolak Section
const tolakButton = document.getElementById('tolak');

tolakButton.addEventListener('click', (e) => {
  e.preventDefault();
  sure = prompt('Beri Pesan Alasan');
  namePdf = document.getElementById('file-sppd').value;
  var pdfBase64 = document.getElementById('pdfViewer-anggaran').src;
  pdfBase64 = pdfBase64.split(',')[1];
  if (sure){    
    const requestPDF = new AppproveSPPDJOINAnggaran();
    requestPDF.setStatus('2');
    requestPDF.setDokName(namePdf);
    requestPDF.setDokPdf(pdfBase64);
    requestPDF.setRincianId(parseInt(rincianId));
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
            ok = alert('Success');
            if (!ok) {
              setTimeout(()=>{
                window.location.reload();
                window.location.href = '/wpp/sppd';
              }, 300)
            }
        } else {
            alert(responseJson.status);
        }
    })
  } else{
    window.location.reload();
  }
})