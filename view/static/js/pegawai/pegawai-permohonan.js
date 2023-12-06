const url = window.location.pathname;

const userId = url.split('/')[2]

const viewPermohonan = document.querySelectorAll('.view');
const suratIds = document.getElementById('suratIds').value;

idSplit = suratIds.split(',')

for (i=0; i < viewPermohonan.length; i++){
    s = '/wp/'+userId+'/permohonan?v='+idSplit[i]
    viewPermohonan[i].setAttribute('href', s)
}
