class PermohonanRequest {
  constructor() {
    this.tipe = 0;
    this.lokasi_tujuan = '';
    this.jenis_program = "";
    this.dok_pendukung_name = '';
    this.dok_pendukung_pdf = '';
    this.participans_id = [];
    this.tgl_awal = '';
    this.tgl_akhir = '';
  }

  setLokasiTujuan(lokasi) {
    this.lokasi_tujuan = lokasi;
  }

  setJenisProgram(jenis) {
    this.jenis_program = jenis;
  }

  setDokPendukungName(dokumen) {
    this.dok_pendukung_name = dokumen;
  }

  setDokPendukungPdf(pdf) {
    this.dok_pendukung_pdf = pdf;
  }

  setParticipansId(participans) {
    this.participans_id = participans;
  }

  setTglAwal(tanggal) {
    this.tgl_awal = tanggal;
  }

  setTglAkhir(tanggal) {
    this.tgl_akhir = tanggal;
  }
}

const backB = document.getElementById('back');

const url = window.location.pathname;
const userId = url.split('/')[2];
backB.setAttribute('href','/wp/'+userId+'/permohonan')

const submitButton = document.getElementById("submit");
const groupRadio = document.querySelector(".checkbox");
const participanContainer = document.getElementById("participan-container");
const participanList = document.getElementById("partcipan-list");

var selectedOptions = [];

groupRadio.addEventListener("change", function() {
    if (groupRadio.checked) {
      while (participanContainer.firstChild) {
        participanContainer.removeChild(participanContainer.firstChild);
      }
      createComboBox();
      createAddResetButton(participanList);
      if (selectedOptions.length === 0 ) {
        submitButton.disabled = true;
      }
    }
    else if(!groupRadio.checked) {
      submitButton.disabled = false;
      while (participanContainer.firstChild) {
        participanContainer.removeChild(participanContainer.firstChild);
      }
      while (participanList.firstChild) {
        participanList.removeChild(participanList.firstChild);
      }
    }
});

function createAddResetButton(divList){
  let addButton = document.createElement("button");
  addButton.id = "addComboBoxButton";
  addButton.disabled = true;
  addButton.type = "button";
  addButton.innerText = "Add";
  addButton.classList.add('btn');
  addButton.classList.add('btn-warning');
  addButton.classList.add('mx-2');
  addButton.classList.add('px-4');
  
  let resetButton = document.createElement("button");
  resetButton.id = "reset";
  resetButton.disabled = true;
  resetButton.type = "button";
  resetButton.innerText = "Reset";
  resetButton.classList.add('btn');
  resetButton.classList.add('btn-danger');
  resetButton.classList.add('mx-2');
  resetButton.classList.add('px-4');

  var div = document.querySelector('#comboBoxDiv')
  
  div.appendChild(addButton)
  div.appendChild(resetButton)
  
  
  const comboBox = document.getElementById("comboBox")
  const options = comboBox.getElementsByTagName("option");
  var opNum = options.length;
  comboBox.addEventListener("change", () => {
    if (comboBox.value) {
      addButton.disabled = false;
    }
    else if(comboBox.length === 0){
      selectedOptions =[]
      comboBox.selected = 0;
      addButton.disabled = true;
      resetButton.disabled = true;
    }
  });
  
  // Tambahkan event listener ke tombol + untuk membuat Combobox B atau yang baru
  const selectDisable = document.querySelector("#disabled")
  addButton.addEventListener("click", (e) => {
    resetButton.disabled = false;
    opNum--;
    e.preventDefault();
    const selectedOption = comboBox.value;        
    selectedOptions.push(selectedOption);
    createList(selectedOption, divList);
    removeList(comboBox, options, opNum, addButton);
    if (selectedOptions.length > 0) {
      submitButton.disabled = false
    }
  });

  resetButton.addEventListener("click", (e) => {
    e.preventDefault();
    selectedOptions =[];
    submitButton.disabled = true;
    while (participanContainer.firstChild) {
      participanContainer.removeChild(participanContainer.firstChild);
    }

    while (participanList.firstChild) {
      participanList.removeChild(participanList.firstChild);
    }

    while (divList.firstChild) {
      divList.removeChild(divList.firstChild);
    }
    createComboBox("");
    createAddResetButton(divList);
  });
}

function createList(selectedOption, divList){
  var ol = document.querySelector('ol');
  
  var li = document.createElement('li');
  li.classList.add('list-group-item');
  
  var inputHidden = document.createElement("input");
  inputHidden.value = selectedOption;
  inputHidden.name = "participan";
  inputHidden.type = "hidden";
  inputHidden.className = "participan";
  
  divList.appendChild(inputHidden);
  
  var idName = dataGOId.indexOf(selectedOption);
  li.innerText = dataGOName[idName];
  
  ol.appendChild(li);
  divList.appendChild(ol);
}

function removeList(comboBox, options, opNum, button){
    var selectedIndex = comboBox.selectedIndex;
    opNum--;

    if (selectedIndex !== -1) {
        comboBox.remove(selectedIndex);
    }
    
    var selectedOptions = Array.from(options).filter(option => option.selected);
    
    if ((selectedOptions.length === 0) && (opNum === 0)) {
        comboBox.setAttribute("disabled", true); // Menonaktifkan elemen <select>
        comboBox.selected = "disabled"; // Memilih opsi "disabled"
        button.disabled = true;
    }
}

submitButton.addEventListener("click", (e) => {
  e.preventDefault(); // Mencegah perilaku default formulir
  
  const lokasi = document.getElementById('lokasi').value;
  const jenis = document.getElementById('jenis').value;
  const pendukung = document.getElementById('pendukung');
  const awal = document.getElementById('awal').value;
  const akhir = document.getElementById('akhir').value;
  var pNum = [];
  pNum = selectedOptions.map(str => parseInt(str));

  const selectedFile = pendukung.files[0];
  
  const permohonanRequest = new PermohonanRequest()
  permohonanRequest.setLokasiTujuan(lokasi);
  permohonanRequest.setJenisProgram(jenis);
  if (groupRadio.checked){
    permohonanRequest.setParticipansId(pNum);
  }
  permohonanRequest.setTglAwal(awal);
  permohonanRequest.setTglAkhir(akhir);

  

  okForm = submitForm();
  if (okForm) {
    return
  } else {
    if (!selectedFile) {
      alert('Pilih file terlebih dahulu.');
      return;
    }
  
    const fileReader = new FileReader();
    fileReader.onload = async function () {
      const fileData = fileReader.result.split(',')[1]; // Ambil data base64 saja

      // Buat objek JSON dengan data formulir dan file base64
      permohonanRequest.setDokPendukungName(selectedFile.name);
      permohonanRequest.setDokPendukungPdf(fileData);

      
    
      // Kirim data ke server dalam format JSON
      await fetch('/wp/'+match[2]+'/permohonan', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(permohonanRequest)
      })
      .then(response => response.json())
      .then(responseJson => {
        if (responseJson.code == 200){
            alert('Success');
            setTimeout(function () {
              window.location.reload();
              window.location.href = '/wp/'+userId+'/permohonan'
            }, 500);
        }
        else if (responseJson.code == 400){
            alert(responseJson.status+' Isi Seluruh Form');
        }
        else{
          alert(responseJson.errorMessages);
        }
      })
    };

    fileReader.readAsDataURL(selectedFile);
  }
});

function createComboBox(info) {
  let label = document.createElement('label')
  label.innerText = 'Pilih Anggota';
  label.classList.add('col-md-4');
  label.classList.add('col-form-label');
  
  let div = document.createElement("div");
  div.classList.add('col-md-5');
  div.classList.add('d-flex');
  div.classList.add('w-50');
  div.id = 'comboBoxDiv'
  
  let comboBox = document.createElement("select");
  comboBox.id = "comboBox";
  comboBox.classList.add('form-select');
  comboBox.classList.add('form-control');
  // comboBox.classList.add('form-block');
  comboBox.ariaLabel = ".form-select-sm example";
  // Create and append option elements
  let option = new Option("Pilih Anggota");
  option.disabled = true;
  option.selected = true;
  option.id = "disabled";
  comboBox.appendChild(option);
  for(i = 0; i < dataGOId.length; i++){
    option = new Option(dataGOName[i], dataGOId[i]);
    comboBox.appendChild(option);
  }

  div.appendChild(comboBox);

  participanContainer.appendChild(label);
  participanContainer.appendChild(div);

  var ol = document.createElement('ol');
  ol.classList.add('list-group');
  ol.classList.add('list-group-numbered');
  participanList.appendChild(ol);

  
  // Tambahkan event listener ke combobox yang baru untuk memungkinkan pembuatan combobox berikutnya
  comboBox.addEventListener("change", () => {
    const selectedOption = comboBox.value;
    
    if (selectedOption) {
        addComboBoxButton.disabled = false;
    }
  });
}

function submitForm() {
  const formElements = document.querySelector('form').elements;
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