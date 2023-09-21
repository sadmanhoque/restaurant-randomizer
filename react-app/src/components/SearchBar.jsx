import { useState } from 'react'
import { FaSearch } from "react-icons/fa"
import "./SearchBar.css"

// eslint-disable-next-line react/prop-types
export const SearchBar = ({ setResults }) => {
    const [input, setInput] = useState("")

    const fetchData = (value) => {
        fetch(`http://localhost:9000/search/:${value}`)
            .then((response) => response.json())
            .then((json) => {
                //const results = json.filter((user) => {
                //    return results;
                //});
                console.log(json)
                setResults(json);
                return json

            });
    }

    const handleChange = (value) => {
        setInput(value)
        fetchData(value)
    }

    return (
        <div className='input-wrapper'>
            <FaSearch id="search-icon" />
            <input placeholder='Type to search...' value={input} onChange={(e) => handleChange(e.target.value)} />
        </div>
    )
}
