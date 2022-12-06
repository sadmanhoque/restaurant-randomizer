require('dotenv').config()
var axios = require('axios');

//random number generator
function between(min, max) {  
  return Math.floor(
    Math.random() * (max - min) + min
  )
}

//Random place search function
const randomResult = async (searchText, coordinates, searchRadius) => {
    try{
        const response = await axios(`https://maps.googleapis.com/maps/api/place/textsearch/json?query=${searchText}&location=${coordinates}&radius=${searchRadius}&key=${process.env.GOOGLE_MAPS_API_KEY}`);
        jsonData = JSON.stringify(response.data)
        jsonParsed = JSON.parse(jsonData)
        const numberOfResults = Object.keys(jsonParsed.results).length
        const rando = between(0, numberOfResults)
        console.log(jsonParsed.results[rando].name)
    } catch(err){
        console.log(err);
    }
}

//Used for testing the generator function
console.log(randomResult('pie', '44.666070, -63.657702', '1000'));