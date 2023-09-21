import { useState } from "react"
import { SearchBar } from "./components/SearchBar";
import Test from "./components/Test";
import { SearchResults } from "./components/SearchResults";

export default function App() {
  const [results, setResults] = useState([])

  return (
    <>
      <h1 className="text-3xl font-bold underline">
        Hello world!
      </h1>

      Test Values:
      <Test />
      <div className="search-bar-container">
        <SearchBar setResults={setResults} />
        <SearchResults results={results} />
      </div>
    </>
  )
}