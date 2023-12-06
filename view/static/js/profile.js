class ChangePasswordRequest {
    constructor(password, new_password) {
      this.password = password;
      this.new_password = new_password;
    }

}

class UpdateProfiledRequest {
    constructor() {
        this.nip = '';
        this.nik = '';
        this.npwp = '';
        this.name = '';
        this.email = "";
        this.no_telp = '';
        this.tgl_lahir = '';
        this.alamat = '';
        this.gambar = '';
    }

    setNIP(klm) {
        this.nip = klm;
    }
    
    setNIK(ghj) {
        this.nik = ghj;
    }
    
    setNPWP(def) {
        this.npwp = def;
    }
    
    setName(n) {
        this.name = n;
    }

    setEmail(abc) {
        this.email = abc;
    }
    setNo(no) {
        this.no_telp = no;
    }
    
    setAlamat(kakak) {
        this.alamat = kakak;
    }

    setTglLahir(tanggal) {
        this.tgl_lahir = tanggal;
    }
    setGambar(g) {
        this.gambar = g;
    }
}

var namein =  document.getElementById('name');
var nik =  document.getElementById('nik');
var nip =  document.getElementById('nip');
var email =  document.getElementById('email');
var npwp =  document.getElementById('npwp');
var lahir =  document.getElementById('lahir');
var noTelp =  document.getElementById('noTelp');
var alamat =  document.getElementById('alamat');
var userImage = document.getElementById('userImage');

var gantiImage = document.getElementById('gantiImage');
var saveImage = document.getElementById('saveImage');
var ganCon = document.getElementById('ganCon');

gantiImage.addEventListener('click', ()=>{
    openFileInput();
    ganCon.style.marginLeft = "-50px";
    saveImage.classList.remove('d-none');
})

saveImage.addEventListener('click',()=>{
     var req = new UpdateProfiledRequest();

     req.setGambar(userImage.src)

    fetch('image', {
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
                window.location.href = 'profile'
                }, 500);
        } else if (responseJson.code == 400){
            alert(responseJson.status+' Isi Seluruh Form');
        } else{
            alert(responseJson.errorMessages);
        }
    })
})

function openFileInput() {
    document.getElementById('fileInput').click();
}

function displaySelectedImage() {
    var fileInput = document.getElementById('fileInput');
    
    var file = fileInput.files[0];
    
    if (file) {
        // Baca file sebagai URL data
        var reader = new FileReader();
        reader.onload = function (e) {
            // Tampilkan gambar yang dipilih
            userImage.src = e.target.result;
        };
        reader.readAsDataURL(file);
    }
    console.log(userImage.src);
    saveImage.classList.remove('d-none');
}

const editButtonSubmit = document.querySelector('#edit');
const editForm = document.querySelector('#edit-form');

const changeButtonSubmit = document.querySelector('#change');
const changeForm = document.querySelector('#change-form');
const toggleButton = document.querySelector('#changePass');
var a = 0;

editButtonSubmit.addEventListener('click', (e)=>{
    if (a==0 ) {
        namein.readOnly = false;
        nik.readOnly = false;
        nip.readOnly = false;
        npwp.readOnly = false;
        email.readOnly = false;
        lahir.readOnly = false;
        noTelp.readOnly = false;
        alamat.readOnly = false;

        // Menonaktifkan tombol setelah diklik
        editButtonSubmit.disabled = true;
        editButtonSubmit.classList.remove('bg-danger');
        editButtonSubmit.classList.add('bg-warning');
        a++;
    } else {

        const req = new UpdateProfiledRequest();
        req.setNIP(nip.value)
        req.setNIK(nik.value)
        req.setNPWP(npwp.value)
        req.setName(namein.value)
        req.setEmail(email.value)
        req.setNo(noTelp.value)
        req.setAlamat(alamat.value)
        req.setTglLahir(lahir.value)
        req.setGambar(userImage.src)

        fetch('profile',{
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
                    window.location.href = 'profile'
                    }, 500);
            } else if (responseJson.code == 400){
                alert(responseJson.status+' Isi Seluruh Form');
            } else{
                alert(responseJson.errorMessages);
            }
        })
    }
});

var initialValues = {
    name: document.getElementById('name').value,
    nik: document.getElementById('nik').value,
    nip: document.getElementById('nip').value,
    npwp: document.getElementById('npwp').value,
    email: document.getElementById('email').value,
    alamat: document.getElementById('alamat').value,
    lahir : document.getElementById('lahir').value,
    noTelp : document.getElementById('noTelp').value,
};

var inputElements = document.querySelectorAll('.form-text');
inputElements.forEach(function(inputElement) {
    inputElement.addEventListener('input', checkForChanges);
});


function checkForChanges() {
    // Cek apakah nilai input telah berubah dari nilai awal
    var isChanged = Object.keys(initialValues).some(function (key) {
        var element = document.getElementById(key);

        if (element) {
            return initialValues[key] !== element.value;
        } else {
            return false; // Consider it as not changed if the element is not found
        }
    });

    // Aktifkan atau nonaktifkan tombol "Edit Profile" berdasarkan hasil cek
    editButtonSubmit.disabled = !isChanged;
}


toggleButton.addEventListener('click', (e)=>{
    x = toggleButton.id
    if (x == 'changePass') {
        toggleButton.id = 'editPass';
        toggleButton.innerHTML = 'Edit Profile';
        
        editForm.classList.add('d-none');
        changeForm.classList.remove('d-none');
        
        namein.readOnly = true;
        nik.readOnly = true;
        nip.readOnly = true;
        npwp.readOnly = true;
        email.readOnly = true;
        lahir.readOnly = true;
        noTelp.readOnly = true;
        alamat.readOnly = true;

        // Menonaktifkan tombol setelah diklik
        editButtonSubmit.disabled = false;
        editButtonSubmit.classList.add('bg-danger');
        editButtonSubmit.classList.remove('bg-warning');
        a = 0;
    } else {
        toggleButton.id = 'changePass';
        toggleButton.innerHTML = 'Ganti Password';
        
        changeForm.classList.add('d-none');
        editForm.classList.remove('d-none');
    }
})


function submitForm() {
    const formElements = document.querySelector('#changeForm').elements;
    const errorMessages = [];

    for (let i = 0; i < formElements.length; i++) {
        const element = formElements[i];

        if (element.hasAttribute('required') && element.value.trim() === '') {
        errorMessages.push(i);
        }
    }

    if (errorMessages.length > 0) {
        // Menampilkan pesan kesalahan dalam bentuk alert
        alert("Isi Seluruh Form")
        return 1;
    }
}

changeButtonSubmit.addEventListener('click', (e)=>{
    e.preventDefault();
    x = submitForm();
    if (x) {
        return
    } else {
        var pswLama = document.querySelector('#passwordLama').value;
        var pswBaru = document.querySelector('#passwordBaru').value;

        if (pswLama === pswBaru) {
            alert("Password yang anda masukkan sama");
            return;
        }

        const req = new ChangePasswordRequest(pswLama, pswBaru);
        fetch('password',{
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
                    window.location.href = 'profile'
                    }, 500);
            } else if (responseJson.code == 400){
                alert(responseJson.status+' Isi Seluruh Form');
            } else{
                alert(responseJson.errorMessages);
            }
        })
    }
})