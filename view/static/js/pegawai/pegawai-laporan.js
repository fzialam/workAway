const view = document.querySelectorAll('.view');
const suratIds = document.getElementById('suratIds').value;

const url = window.location.href

const userId = url.split('/')[4]
idSplit = suratIds.split(',')

for (i=0; i < view.length; i++){
    s = '/wp/'+userId+'/laporan?id='+idSplit[i]
    view[i].setAttribute('href', s)
}
