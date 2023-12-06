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

const IsValid = document.querySelector('#IsValid').value
// console.log(IsValid);

const url = window.location.pathname
const idSurat =  url.split('/')[2]

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
        const fileData = fileReader.result.split(',')[1];

        requestPDF.setDokName(selectedFile.name);
        requestPDF.setDokPdf(fileData);
        
        var endPoint = '/wk/'+idSurat+'/set-rincian';
        if (!IsValid) {
            endPoint = '/wk/'+idSurat+'/rincian-biaya';
        }
        await fetch(endPoint, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(requestPDF)
        })
        .then(response => response.json())
        .then(responseJson => {
            if (responseJson.code == 200){
                ok = alert(responseJson.data.message);
                if (!ok) {
                    setTimeout(()=>{
                        window.location.reload();
                        window.location.href = '/wk/rincian-biaya';
                    }, 300);
                }
            }
            else{
                alert(responseJson.status);
            }
        })

    }
    fileReader.readAsDataURL(selectedFile);
});
