class ApprovedLaporan{
    constructor() {
        this.id = 0;
        this.status = '';
        this.message = '';
    }

    setID(ID){
        this.id = ID
    }
    setMewssage(m){
        this.message = m
    }
    setStatus(status){
        this.status = status
    }
}
const url = window.location.href

const idLap = url.split('/')[4]

const sButt = document.querySelector('#s')
const tButt = document.querySelector('#t')
const laporanId = document.querySelector('#laporanId').value

const req = new ApprovedLaporan();

sButt.addEventListener('click', (e)=> {
    e.preventDefault();

    x = confirm('Apakah anda yakin');

    if (x){
        req.setMewssage("OK");
        req.setStatus('1');
        req.setID(parseInt(laporanId));
        
        fetch('/wpp/'+idLap+'/laporan', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(req)
        })
        .then(response => response.json())
        .then(responseJson => {
            if (responseJson.code == 200){
                setTimeout(()=>{
                        window.location.href = '/wpp/laporan'
                    }, 300)
            } else {
                alert(responseJson.status);
            }
        })
        return
    }
})

tButt.addEventListener('click', (e)=> {
    e.preventDefault();

    x = prompt('Beri alasan anda');

    if (x){
        req.setMewssage(x);
        req.setStatus('2');
        req.setID(parseInt(laporanId));
        fetch('/wpp/'+idLap+'/laporan', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(req)
        })
        .then(response => response.json())
        .then(responseJson => {
            if (responseJson.code == 200){
                setTimeout(()=>{
                    window.location.href = '/wpp/laporan'
                }, 300)
            } else {
                alert(responseJson.status);
            }
        })
        return
    }
})