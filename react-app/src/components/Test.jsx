import { useState, useEffect } from "react";
//import React from "react"

const URL = `http://localhost:9000/books`
const id = "pizza";
const returnURL = `http://localhost:9000/search/:${id}`

function Test() {
    const [temp, setTemp] = useState(0)

    useEffect(() => {
        const fetchData = async () => {
            const result = await fetch(URL)
            result.json().then(json => {
                setTemp(json[0].quantity)
                //console.log(json)
            })
        }
        fetchData()
    }, [])

    const submitKey = async () => {
        const myData = {
            'key': 'pizza'
        }

        const result = await fetch(returnURL, {
            method: 'GET',
            //headers: {
            //    'Content-Type': 'application/json'
            //},
            params: JSON.stringify(myData)
        })

        const resultInJson = await result.json();
        console.log(resultInJson)
    }

    return (
        <div className="Test">
            Book API return: {temp}
            <button onClick={submitKey}>Return book</button>
        </div>
    )
}

export default Test;