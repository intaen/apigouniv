function searchForm(e) {
  var search = document.getElementById("search").value;
  if (search == "") {
      alert("Please enter the keyword")
      e.preventDefault();
      return
  }

  alert("Wait a minute...");
  e.preventDefault();

  fetch('/master/university?'+ new URLSearchParams({
    "name": search,
   // "country": search,
    }),
  {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'}
  }).then(function(response) {
    return response.json();
  }).then(function(data) {
    clearResult()

    // Success code goes here
    let response = JSON.stringify(data)
    let result = JSON.parse(response)   
    let name, country, url = ""
    let urls, items = []

    if (result.result.count == 0) {
      alert("Sorry, university not found. Please check again the name.")
      return
    }

    // check each country name
    for (let i = 0; i < result.result.data.length; i++) {
      name = result.result.data[i].name;
      country = result.result.data[i].country_name;

      // add to array so we can join it with separator
      for (let j = 0; j < result.result.data[i].url.length; j++) {
        if (result.result.data[i].url.length == 1) {
          url = result.result.data[i].url[j]
        } else if (result.result.data[i].url.length > 1) {
          urls.push(result.result.data[i].url[j]);
          url = urls.join(" , ");
        }
      }

      // Show result in list
      items.push(name +" is in "+country+", go to "+url+" for detail.")
      showListResult(items)
    }
  }).catch(function(err) {
  // Failure
    alert('Error: '+err)
  });
}

function showListResult(texts) {
var ul = "<ul>"
texts.forEach(function(text) {
  ul += '<li>'+ text + '</li>';
  }); 
ul += '</ul>';

var dom = document.getElementById("result");
dom.setAttribute('style','text-align:justify;');
dom.innerHTML = ul;
}

function clearResult() {
document.getElementById('result').innerHTML = "";
}