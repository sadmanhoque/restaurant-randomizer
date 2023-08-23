import { useState, useEffect } from "react";
//import React from "react"

const URL = ` http://localhost:8000/books`

function Test() {
    const [temp, setTemp] = useState(0)

    useEffect(() => {
        const fetchData = async () => {
            const result = await fetch(URL)
            result.json().then(json => {
                setTemp(json[0].quantity)
            })
        }
        fetchData()
    }, [])

    return (
        <div className="Test">
            Book API return: {temp}
        </div>
    )
}

export default Test;