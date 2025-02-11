import { useState } from "react";

interface SearchResult {
  title: string;
  url: string;
  metaData?: string;
}

const LandingPage: React.FC = () => {
  const [isSearched, setIsSearched] = useState<boolean>(false);
  const [query, setQuery] = useState<string>("");
  const [results, setResults] = useState<SearchResult[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [error, setError] = useState<string>("");

  const fetchResults = async (query: string): Promise<void> => {
    setIsLoading(true);
    setError("");

    try {
      const response = await fetch(
        `http://localhost:8081/search?limit=100&query=${encodeURIComponent(query)}`,
        {
          headers: {
            "User-Agent": "my-reddit-app",
          },
        }
      );

      if (!response.ok) {
        throw new Error("Failed to fetch results");
      }

      const data: SearchResult[] = await response.json();

      setResults(Array.isArray(data) ? data : []);
      setIsSearched(true);
    } catch (err) {
      setError("An error occurred while fetching results. Please try again.");
      console.error(err);
    } finally {
      setIsLoading(false);
    }
  };

  const handleSearch = (): void => {
    if (query.trim()) {
      fetchResults(query);
    }
  };

  return (
    <div>
      <div className="flex flex-col mt-12 items-center bg-white">
        <h1 className="mb-12 text-4xl text-slate-800">IntelliSearch</h1>

        <div className="flex items-center">
          <input
            type="text"
            placeholder="Search the web..."
            className="w-96 px-4 py-2 border border-gray-300 rounded-l-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            style={{ height: "40px" }}
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            onKeyPress={(e) => e.key === "Enter" && handleSearch()}
          />
          <button
            className="px-6 py-2 bg-blue-500 text-white rounded-r-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
            style={{ height: "40px" }}
            onClick={handleSearch}
            disabled={isLoading}
          >
            {isLoading ? "Searching..." : "Search"}
          </button>
        </div>

        <div className="mt-6 space-x-4">
          <a href="https://sushantcodes.tech" className="text-blue-500 hover:underline">
            About
          </a>
          <a href="mailto:sushant.dhiman9812@gmail.com" className="text-blue-500 hover:underline">
            Contact
          </a>
          <a href="https://linkedin.com/in/sushant102004" className="text-blue-500 hover:underline">
            Follow
          </a>
        </div>
      </div>

      {isSearched ? (
        <div className="mt-12 mx-24">
          <h2 className="text-2xl text-slate-800 mb-6">Search Results for "{query}"</h2>
          {error ? (
            <p className="text-red-500">{error}</p>
          ) : isLoading ? (
            <p>Loading...</p>
          ) : results.length > 0 ? (
            <div className="space-y-4">
              {results.map((result, index) => (
                <div key={index} className="border-b border-gray-200 pb-4">
                  <a
                    href={result.url}
                    className="text-blue-600 hover:underline text-lg"
                    target="_blank"
                    rel="noopener noreferrer"
                  >
                    {result.title}
                  </a>
                  <p className="text-gray-600">{result.metaData || "No description available."}</p>
                </div>
              ))}
            </div>
          ) : (
            <p>No results found.</p>
          )}
        </div>
      ) : (
        <div className="ml-24 mt-12 mr-24">Please search something to see results.</div>
      )}
    </div>
  );
};

export default LandingPage;
