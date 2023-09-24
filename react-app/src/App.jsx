import { useState } from "react"
import SearchBar from './components/SearchBar';
import axios from 'axios';

export default function App() {
  const [searchResults, setSearchResults] = useState([]);
  const [isLoading, setIsLoading] = useState(false);

  const handleSearch = (query) => {
    // Clear previous search results
    setSearchResults([]);

    // Perform the GET request here using Axios or any other HTTP library
    setIsLoading(true);

    axios.get(`http://localhost:9000/search/:${query}`)
      .then((response) => {
        setSearchResults(response.data);
        setIsLoading(false);
      })
      .catch((error) => {
        console.error(error);
        setIsLoading(false);
      });
  };

  return (
    <div>
      <h1>Search App</h1>
      <SearchBar onSearch={handleSearch} />
      {isLoading ? (
        <p>Loading...</p>
      ) : searchResults.error ? (
        <p>No results found</p>
      ) : (
        <div>
          <h2>{searchResults.name}</h2>
          <p>{searchResults.storeAddress}</p>
        </div>
      )}
    </div>
  );
}