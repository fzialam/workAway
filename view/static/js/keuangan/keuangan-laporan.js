class Konfirmasi{
    constructor() {
        this.id = 0;
        this.status = '';
        this.message = '';
    }
    setID(ID){
        this.id = ID;
    }
    setStatus(S){
        this.status = S;
    }
    setMessage(M){
        this.message = M;
    }
}
const konfirmasiButton = document.querySelectorAll('.konfirmasi');
const tolakButton = document.querySelectorAll('.tolak');

konfirmasiButton.forEach(element => {
    element.addEventListener('click', (e) =>{
        ok = confirm('Apakah anda yakin');
        if (!ok) {
            return
        } else {
            const hiddenKonfirmasi = element.firstElementChild.value;
            const req = new Konfirmasi();
            req.setID(parseInt(hiddenKonfirmasi));
            req.setStatus('1')
            req.setMessage('OK')

            fetch('/wk/laporan', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(req)
            })
            .then(response => response.json())
            .then(responseJson => {
                if (responseJson.code == 200){
                    ok = alert(responseJson.data.message);
                    if (!ok) {
                        setTimeout(()=>{
                            window.location.reload();
                        }, 300)
                    }
                }
                else{
                    alert(responseJson.status);
                }
            })
        }
    });
});

tolakButton.forEach(element => {
    element.addEventListener('click', (e) =>{
        ok = confirm('Apakah anda yakin');
        if (!ok) {
            return
        } else {
            x = prompt('Beri alasan anda')
            const hiddenKonfirmasi = element.firstElementChild.value;
            const req = new Konfirmasi();
            req.setID(parseInt(hiddenKonfirmasi));
            req.setStatus('2');
            req.setMessage(x);

            fetch('/wk/laporan', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(req)
            })
            .then(response => response.json())
            .then(responseJson => {
                if (responseJson.code == 200){
                    ok = alert(responseJson.data.message);
                    if (!ok) {
                        setTimeout(()=>{
                            window.location.reload();
                        }, 300)
                    }
                }
                else{
                    alert(responseJson.status);
                }
            })
        }
    });
});