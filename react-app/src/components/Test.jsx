import { useState, useEffect } from "react";
//import React from "react"

const URL = `http://localhost:8000/books`
const id = 2;
const returnURL = `http://localhost:8000/return?id=${id}`

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

    const submitId = async () => {
        const myData = {
            'id': '2'
        }

        const result = await fetch(returnURL, {
            method: 'PATCH',
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
            <button onClick={submitId}>Return book</button>
        </div>
    )
}

export default Test;