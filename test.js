require('dotenv').config()
var axios = require('axios');

//random number generator
function between(min, max) {  
  return Math.floor(
    Math.random() * (max - min) + min
  )
}
var rando = between(0, 21)
//console.log(rando)

//Inputs for testing the API
var input = 'sandwich'
var location = '44.666070, -63.657702'
var radius = '1000'

//config for the axios function
var config = {
    method: 'get',
    url: `https://maps.googleapis.com/maps/api/place/textsearch/json?query=${input}&location=${location}&radius=${radius}&key=${process.env.GOOGLE_MAPS_API_KEY}`,
    headers: { }
  };

//this is where the real stuff happens
  axios(config)
  .then((response) =>{
    jsonData = JSON.stringify(response.data)
    jsonParsed = JSON.parse(jsonData)
    console.log(jsonParsed.results[rando].name)

  })
  .catch((error) => {
    console.log(error);
  });
  