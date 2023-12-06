class ChangeRole{
    constructor() {
        this.rank = 0;
    }
}

const gantiRank = document.getElementById('gantiRank');
const rank = document.getElementById('rank');
const saveImage = document.getElementById('saveImage');

const url = document.location.pathname
const userId = url.split('/')[2]

var initialValues = {
    rank: rank.value,
};

gantiRank.addEventListener('click', (e)=>{
    rank.disabled = false;
    saveImage.classList.remove('d-none');
})

rank.addEventListener('input', ()=>{
    
    var isChanged = Object.keys(initialValues).some(function (key) {
        var element = document.getElementById(key);

        if (element) {
            return initialValues[key] !== element.value;
        } else {
            return false;
        }
    })
    saveImage.disabled = !isChanged;
})

saveImage.addEventListener('click',(e)=>{
    e.preventDefault();

    var req = new  ChangeRole();
    req.rank = parseInt(rank.value);
    
    fetch('/wa/'+userId+'/user',{
        method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(req)
    })
    .then(response => response.json())
    .then(responseJson => {
        if (responseJson.code == 200){
            alert('Success');
            setTimeout(function () {
                window.location.reload();
                window.location.href = url
                }, 500);
        } else if (responseJson.code == 400){
            alert(responseJson.status+' Isi Seluruh Form');
        } else{
            alert(responseJson.errorMessages);
        }
    })
    
})