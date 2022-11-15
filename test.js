require('dotenv').config()
var axios = require('axios');

//Inputs for testing the API
var input = 'pizza'
var location = '44.666070, -63.657702'
var radius = '1000'

var config = {
    method: 'get',
    url: `https://maps.googleapis.com/maps/api/place/textsearch/json?query=${input}&location=${location}&radius=${radius}&key=${process.env.GOOGLE_MAPS_API_KEY}`,
    headers: { }
  };

  axios(config)
  .then((response) =>{
    jsonData = JSON.stringify(response.data)
    jsonParsed = JSON.parse(jsonData)
    //console.log(JSON.stringify(response.data));
    console.log(jsonParsed.results[0].name)

  })
  .catch((error) => {
    console.log(error);
  });
  