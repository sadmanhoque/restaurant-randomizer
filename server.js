const express = require('express')
const { restart } = require('nodemon')
const randomResult = require('./test2')
const app = express()
const port = 8383
var searchResult = ''

app.use(express.static('public'))
app.use(express.json())

app.get('/info/:dynamic', (req, res) => {
    const {dynamic} = req.params
    const {key} = req.query
    console.log(dynamic, key)
    res.status(200).json({info: 'preset text'})
})

app.post('/', (req, res) => {
    const {parcel} = req.body
    console.log(req.body.inputCoordinates)
    if(!parcel){
        return res.status(400).send({status: 'failed'})
    }
    //console.log(parcel)
    randomResult(parcel, '44.666070, -63.657702', '1000').then(searchResult =>{
        //console.log(searchResult)
        res.status(200).json({status: 'received', 'searchResult': searchResult})
      })
})

app.listen(port, () => console.log(`Server has started on port: ${port}`))