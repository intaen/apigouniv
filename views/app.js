function searchForm(e) {
    var search = document.getElementById("search").value;
    if (search == "") {
        alert("Please enter the keyword")
        e.preventDefault();
        return
    }

    alert("Wait a minute...");
    e.preventDefault();

    fetch('/master/university', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        "name": search,
        "country": search,
      })
    }).then(function(response) {
      return response.json();
    }).then(function(data) {
    // Success code goes here
      let response = JSON.stringify(data)
      let result = JSON.parse(response)   
      let name, country, url = ""
      const urls = []

      // check each country name
      for (let i = 0; i < result.result.data.length; i++) {
        name = result.result.data[i].name;
        country = result.result.data[i].country_name;

        // add to array so we can join it with separator
        urls.push(result.result.data[i].url);
        url = urls.join(" , ");
      }

      alert(name +" is in "+country+", go to "+url+" for detail.")  
    }).catch(function(err) {
    // Failure
      alert('Error: '+err)
    });
}