function Get(theUrl) {
  var xmlHttp = new XMLHttpRequest();
  xmlHttp.open("GET", theUrl, false); // false for synchronous request
  xmlHttp.send();
  return JSON.parse(xmlHttp.responseText);
}

function numberWithCommas(x) {
  return x.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ",");
}

// append an entry to the table
// input is what the user had given as a option
// json is the returning value from the rest server
function appendTable(input, json) {
  var table = document.getElementById("output");
  let row = table.insertRow(1);
  let cell1 = row.insertCell(0);
  let cell2 = row.insertCell(1);
  let cell3 = row.insertCell(2);
  cell1.innerText = numberWithCommas(input);
  if (json.Index > 0) {
    cell2.innerText = numberWithCommas(json.Index);
    cell3.innerText = numberWithCommas(json.Surrounding);
  } else {
    cell2.innerText = "Could not find result";
    cell3.innerText = "No neighbours because no match was found";
  }
}

function update() {
  // getting dom elements
  let input = document.getElementById("input");
  let result = document.getElementById("result");
  // only update if a value was entered
  if (input.value == "") return;
  // do the api call
  let out = Get("http://localhost:8080/search/" + input.value);
  appendTable(input.value, out);
}

function clearTable() {
  var table = document.getElementById("output");
  console.log(table.rows.length);
  // length is up here because the table always gets updated thus each time we loop over the for our endpoint is different
  var length = table.rows.length;
  for (let i = 1; i < length; i++) {
    table.deleteRow(-1);
  }
}
