import { useCallback, useState } from "react";
import "./styles.css";

const SEARCH_URL = "https://openlibrary.org/search.json?q=";

async function searchBooks(query: string) {
  const fetchResults = await window.fetch(`${SEARCH_URL}${query}`);
  const results = await fetchResults.json();
  return results["docs"] as any[];
}

export default function App() {
  const [query, setQuery] = useState("");
  const [books, setBooks] = useState([] as any[]);

  const doSearch = useCallback(() => {
    searchBooks(query).then((books) => setBooks(books));
  }, [query, setBooks]);

  return (
    <div className="App">
      <input
        value={query}
        placeholder="Search for a book"
        onChange={(e) => setQuery(e.target.value)}
      />
      <button onClick={doSearch}>Search</button>
      <hr />

      <div>
        {books.length === 0 && <div>No Results</div>}
        {books.map((book, i) => (
          <div key={i} className="book">
            <div>
              "{book.title}" by {book.author_name && book.author_name[0]}
            </div>
            <div>First published in {book.first_publish_year}</div>
          </div>
        ))}
      </div>
    </div>
  );
}
