require('dotenv').config()
var axios = require('axios');

//const {Client} = require("@googlemaps/google-maps-services-js");
// const client = new Client({});
// client
//   .ElevationService({
//     params: {
//       locations: [{ lat: 45, lng: -110 }],
//       key: process.env.GOOGLE_MAPS_API_KEY,
//     },
//     timeout: 1000, // milliseconds
//   })
//   .then((r) => {
//     console.log(r.data.results[0].elevation);
//   })
//   .catch((e) => {
//     console.log(e.response.data.error_message);
//   });

var input = 'cake'
var location = '44.666070, -63.657702'
var radius = '1000'

var config = {
    method: 'get',
    url: `https://maps.googleapis.com/maps/api/place/textsearch/json?query=${input}&location=${location}&radius=${radius}&key=${process.env.GOOGLE_MAPS_API_KEY}`,
    headers: { }
  };
  
  axios(config)
  .then(function (response) {
    console.log(JSON.stringify(response.data));
  })
  .catch(function (error) {
    console.log(error);
  });