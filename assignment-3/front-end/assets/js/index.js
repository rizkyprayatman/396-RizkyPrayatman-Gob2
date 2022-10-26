window.onload = function () {
  document.body.classList.remove("is-preload");
};
window.ontouchmove = function () {
  return false;
};
window.onorientationchange = function () {
  document.body.scrollTop = 0;
};
const api_url = "http://localhost:8080";

const get_data = async (a) => {
  const request = await fetch("http://localhost:8080/weather", {
    method: "POST",
  });

  const response = await fetch(a);
  var data = await response.json();
  console.log(data);
  var water = data["Water"];
  var wind = data["Wind"];
  var msgWater = data["waterMessage"];
  var msgWind = data["windMessage"];

  document.getElementById("water").innerHTML = 
  `<p>${water} meter</p>`;;
  document.getElementById("wind").innerHTML = 
  `<p>${wind} m/s</p>`;
  setTimeout(() => get_data(a), 1000 * 15);

  const alertWater = (message, type) => {
    const alertWater = document.getElementById("alertMsgWater");
    alertWater.innerHTML =
      `<div class="alert alert-${type} text-center" role="alert">
         <h3 class="text-center">${message}</h3>
      </div>`
  };

  if (msgWater === "Bahaya") {
    alertWater(msgWater, "danger");
  } else if (msgWater === "Siaga") {
    alertWater(msgWater, "warning");
  } else {
    console.log(msgWater);
    alertWater(msgWater, "info");
  }

  const alertWind = (message, type) => {
    const alertWind = document.getElementById("alertMsgWind");
    alertWind.innerHTML =
      `<div class="alert alert-${type} text-center" role="alert">
         <h3 class="text-center">${message}</h3>
      </div>`
  };

  if (msgWind === "Bahaya") {
    alertWind("Bahaya", "danger");
  } else if (msgWind === "Siaga") {
    alertWind("Siaga", "warning");
  } else {
    console.log(msgWind);
    alertWind("Aman", "info");
  }
};

get_data(api_url);
