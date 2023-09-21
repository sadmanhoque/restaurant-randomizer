//import React from 'react'
import "./SearchResults.css"

// eslint-disable-next-line react/prop-types
export const SearchResults = ({ results }) => {
    return (
        <div className='results-list'>
            {
                // eslint-disable-next-line react/prop-types
                results.map((result, id) => {
                    return <div key={id}>{result.name}</div>
                })
            }
        </div>

    )
}
