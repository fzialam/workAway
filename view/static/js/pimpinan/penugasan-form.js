class PenugasanRequest {
  constructor() {
    this.tipe = 1;
    this.user_ketua_id = 0;
    this.lokasi_tujuan = '';
    this.jenis_program = '';
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

  setKetuaId(ketuaId) {
    this.user_ketua_id = ketuaId;
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

var today = new Date();
var dd = String(today.getDate()).padStart(2, '0');
var mm = String(today.getMonth() + 1).padStart(2, '0'); // January is 0!
var yyyy = today.getFullYear();

today = yyyy + '-' + mm + '-' + dd;
var awal = document.getElementById('awal');

awal.min = today;

awal.addEventListener('change',(e)=>{
  document.getElementById('akhir').min = document.getElementById('awal').value;
  })

const submitButton = document.getElementById("submit");
const soloRadio = document.getElementById("solo");
const groupRadio = document.getElementById("group");
const participanContainer = document.getElementById("participan-container");
const participanList = document.getElementById("partcipan-list");

var selectedOptions = [];

soloRadio.addEventListener("change", function() {
  if (soloRadio.checked) {
    selectedOptions = [];
    submitButton.disabled = true;
    
    while (participanContainer.firstChild) {
      participanContainer.removeChild(participanContainer.firstChild);
    }
    
    while (participanList.firstChild) {
      participanList.removeChild(participanList.firstChild);
    }
    
    createComboBox(0);
  }
  else if(!soloRadio.checked) {
    while (participanContainer.firstChild) {
      participanContainer.removeChild(participanContainer.firstChild);
    }
  }
});


groupRadio.addEventListener("change", function() {
  if (groupRadio.checked) {
    submitButton.disabled = true;
    
    while (participanContainer.firstChild) {
      participanContainer.removeChild(participanContainer.firstChild);
    }
    
    const divList = document.getElementById("partcipan-list");
    
    alert('Pilihan Anggota Pertama akan menjadi Ketua Kelompok')
    
    createComboBox(1);
    createAddResetButton(divList);    
  }
  else if(!groupRadio.checked) {
    
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
  const label = document.getElementById('id-label-par')
  const option = document.querySelector('#comboBox')
  addButton.addEventListener("click", (e) => {
    e.preventDefault();
    
    opNum--;
    resetButton.disabled = false;
    
    const selectedOption = comboBox.value;        
    
    selectedOptions.push(selectedOption);
    
    createList(selectedOption, divList);
    removeList(comboBox, options, opNum, addButton);
    
    if (selectedOptions.length > 1) {
      submitButton.disabled = false
    }
    else if (selectedOptions.length === 0){
      label.innerText = 'Pilih Ketua Kelompok';
      option.text = 'Pilih Ketua Kelompok';
    } 
    
    if (selectedOptions.length === 1){
      label.innerText = 'Pilih Anggota Kelompok';
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

function createComboBox(info) {
  if (info === 0){
    let label = document.createElement('label')
    label.innerText = 'Pilih Pelaksana';
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
    let option = new Option("Pilih Pelaksana");
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

    comboBox.addEventListener("change", () => {
      const selectedOption = comboBox.value;
      
      if (selectedOption) {
        submitButton.disabled = false;
      }
    });
  } else{
    let label = document.createElement('label')
    label.id = 'id-label-par'
    label.classList.add('col-md-4');
    label.classList.add('col-form-label');
    label.innerText = 'Pilih Ketua Kelompok';
    
    let div = document.createElement("div");
    div.classList.add('col-md-5');
    div.classList.add('d-flex');
    div.classList.add('w-50');
    div.id = 'comboBoxDiv'
    
    let comboBox = document.createElement("select");
    comboBox.id = "comboBox";
    comboBox.classList.add('form-select');
    comboBox.classList.add('form-control');
    comboBox.ariaLabel = ".form-select-sm example";
    
    // Create and append option elements
    let option = new Option('Pilih Anggota');
    
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

    if (info === 1){
      comboBox.addEventListener("change", () => {
        const selectedOption = comboBox.value;
        
        if (selectedOption) {
          addComboBoxButton.disabled = false;
        }
      });
    }
  }
  
}

submitButton.addEventListener("click", (e) => {
  e.preventDefault(); // Mencegah perilaku default formulir
  
  const lokasi = document.getElementById('lokasi').value;
  var ketuaId = document.getElementById('comboBox');
  if (!ketuaId) {
    alert('BAD REQUEST');
    return
  }
  const jenis = document.getElementById('jenis').value;
  const awal = document.getElementById('awal').value;
  const akhir = document.getElementById('akhir').value;
  var pNum = [];
  pNum = selectedOptions.map(str => parseInt(str));
  
  const penugasanRequest = new PenugasanRequest()
  
  if (selectedOptions.length == 0){
    penugasanRequest.setKetuaId(parseInt(ketuaId.value));
  }else if (selectedOptions.length > 0){
    ketuaId = pNum[0];
    penugasanRequest.setKetuaId(ketuaId);
    pNum.shift();
  }
  
  penugasanRequest.setLokasiTujuan(lokasi);
  penugasanRequest.setJenisProgram(jenis);
  penugasanRequest.setParticipansId(pNum);
  penugasanRequest.setTglAwal(awal);
  penugasanRequest.setTglAkhir(akhir);

  okForm = submitForm();
  
  if (okForm) {
    return;
  }else{
    fetch('/wpp/penugasan', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
        },
        body: JSON.stringify(penugasanRequest)
    })
    .then(response => response.json())
    .then(responseJson => {
      if (responseJson.code == 200){
          alert('Succes');
          window.location.href = '/wpp/sppd';
        }
        else{
          alert(responseJson.status);
          window.location.href = '/wpp/sppd';
      }
    })
  }
});

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