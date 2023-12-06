// Popup Section
const active = document.querySelector('#active');
const closeButt = document.querySelector('#close');
const overlay = document.querySelector('#overlay');
const popupContainer = document.querySelector('#popup');
const popupOverlay = document.querySelector('.popup-overlay');


const nodes = overlay.getElementsByTagName('*');


active.addEventListener('click', (e)=>{
  overlay.classList.add('active-popup');
  popupContainer.style.display = 'block';
  popupOverlay.style.display = 'block';
  for(var i = 0; i < nodes.length; i++){
      nodes[i].disabled = true;
  }
});

closeButt.addEventListener('click', (e)=>{
  overlay.classList.remove('active-popup');
  popupContainer.style.display = 'none';
  popupOverlay.style.display = 'none';
  for(var i = 0; i < nodes.length; i++){
    nodes[i].disabled = false;
  }
});
