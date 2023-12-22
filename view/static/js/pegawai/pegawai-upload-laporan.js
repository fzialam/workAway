class RequestPDF {
    constructor() {
        this.dok_laporan_name = '';
        this.dok_laporan_pdf = '';
    }

    setDokName(dokumen) {
        this.dok_laporan_name = dokumen;
    }

    setDokPdf(pdf) {
        this.dok_laporan_pdf = pdf;
    }
}

const fileAk = document.getElementById('file-ak');
const fileAng = document.getElementById('file-ang');

const buttAk = document.getElementById('butt-ak');
const buttAng = document.getElementById('butt-ang');

const url = window.location.href;
const idUser = url.split('/')[4];
const suratId = url.split('id=')[1];

var endPoint = '';

if ((fileAk != undefined) && (buttAk != undefined)) {
    buttAk.addEventListener('click',(e)=>{
        e.preventDefault();
        
        x = buttAk.firstElementChild.value;

        if (x == 'add'){
            endPoint = '/wp/'+idUser+'/laporan-ak?id='+ suratId
        } else if (x == 'set'){
            endPoint = '/wp/'+idUser+'/set-laporan-ak?id='+ suratId
        }
        
        x = buttAk.firstElementChild.value;
        if (fileAk.hidden == true){
            fileAk.hidden = false;
        } else {
            UploadFile(fileAk, endPoint);
        }
    })
} 

if ((fileAng != undefined) && (buttAng != undefined)){
    buttAng.addEventListener('click',(e)=>{
        e.preventDefault();
        
        x = buttAng.firstElementChild.value;

        if (x == 'add'){
            endPoint = '/wp/'+idUser+'/laporan-ang?id='+ suratId
        } else if (x == 'set'){
            endPoint = '/wp/'+idUser+'/set-laporan-ang?id='+ suratId
        }
        if (fileAng.hidden == true) {
            fileAng.hidden = false;
        } else {
            UploadFile(fileAng ,endPoint)
        }
    })
}



function UploadFile(file, endPoint) {
    const selectedFile = file.files[0];
    if (!selectedFile){
        alert('Pilih file terlebih dahulu');
        return;
    }

    const requestPDF = new RequestPDF();


    const fileReader = new FileReader();
    fileReader.onload = async function () {
        const fileData = fileReader.result.split(',')[1];
        
        requestPDF.setDokName(selectedFile.name);
        requestPDF.setDokPdf(fileData);

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
                if (!ok){
                    setTimeout(()=>{
                      window.location.reload();
                    },500)
                }
            }
            else{
                ok = alert(responseJson.status);
                if (!ok){
                    setTimeout(()=>{
                        window.location.reload();
                    },500)
                }
            }
        })
    };
    fileReader.readAsDataURL(selectedFile);
};

function searchGoogleKetua() {
    // Get the <a> element
    var linkElement = document.querySelector('div#img-bukti a');
    console.log(linkElement);

    // Update the href attribute with the coordinate value
    if (linkElement) {
        linkElement.click();
    }
}