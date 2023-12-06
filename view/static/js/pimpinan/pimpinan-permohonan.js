class Konfirmasi{
  constructor() {
      this.status = '';
      this.message = '';
  }
  setStatus(S){
      this.status = S;
  }
  setMessage(M){
      this.message = M;
  }
}

const setujuButt = document.querySelector("#s");
const tolakButt = document.querySelector("#t");


const url = window.location.pathname;
urlSplit = url.split("/")[2]


setujuButt.addEventListener('click', (e)=>{
  ok = confirm('Apakah anda yakin');
  if (!ok) {
      return
  } else {
      const req = new Konfirmasi();
      req.setStatus('1')
      req.setMessage('OK')

      fetch("/wpp/"+ urlSplit+"/permohonan", {
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
                  window.location.reload()
                  window.location.href = '/wpp/permohonan'
                }, 300)
          }
          else{
            ok = alert(responseJson.status);
            if (!ok) {
              setTimeout(()=>{
                window.location.reload()
                window.location.href = '/wpp/permohonan'
              }, 300)
            }
          }
        }
      })
    }
  }
);

tolakButt.addEventListener('click', (e)=>{
  ok = confirm('Apakah anda yakin');
  if (!ok) {
    return
  } else {
    x = prompt('Beri alasan anda')
    const req = new Konfirmasi();
    req.setStatus('2')
    req.setMessage(x)

    fetch("/wpp/"+ urlSplit+"/permohonan", {
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
            window.location.reload()
            window.location.href = '/wpp/permohonan'
          }, 300)
        }
      }
      else{
        ok = alert(responseJson.status);
        if (!ok) {
          setTimeout(()=>{
            window.location.reload()
            window.location.href = '/wpp/permohonan'
          }, 300)
        }
      }
    });
  }
});