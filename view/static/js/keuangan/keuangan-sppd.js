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

            fetch('/wk/sppd', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(req)
            })
            .then(response => response.json())
            .then(responseJson => {
                if (responseJson.code == 200){
                    window.location.reload();
                    alert(responseJson.data.message);
                }
                else{
                    window.location.reload();
                    alert(responseJson.status);
                }
            })
        }
    });
});