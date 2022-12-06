const express = require('express')
const { restart } = require('nodemon')
const app = express()
const port = 8383

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
    if(!parcel){
        return res.status(400).send({status: 'failed'})
    }
    console.log(parcel)
    //Should make a call to the Google API here
    res.status(200).send({status: 'received'})
})

app.listen(port, () => console.log(`Server has started on port: ${port}`))