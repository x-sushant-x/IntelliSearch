import { useState } from "react";

const LandingPage = () => {
  const [isSearched, setIsSearched] = useState(false);
  const [query, setQuery] = useState("");

  const handleSearch = () => {
    if (query.trim()) {
      setIsSearched(true);
    }
  };

  const demoResults = [
    {
      title: "What is IntelliSearch?",
      link: "https://intellisearch.com",
      description: "IntelliSearch is a smart search engine designed to provide accurate and fast results.",
    },
    {
      title: "How to use IntelliSearch",
      link: "https://intellisearch.com/guide",
      description: "Learn how to make the most out of IntelliSearch with our comprehensive guide.",
    },
    {
      title: "Benefits of using IntelliSearch",
      link: "https://intellisearch.com/benefits",
      description: "Discover the advantages of using IntelliSearch over traditional search engines.",
    },
  ];

  return (
    <div>
      {/* Search Engine Box */}
      <div className="flex flex-col mt-12 items-center bg-white">
        <h1 className='mb-12 text-4xl text-slate-800'>IntelliSearch</h1>

        <div className="flex items-center">
          <input
            type="text"
            placeholder="Search the web..."
            className="w-96 px-4 py-2 border border-gray-300 rounded-l-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
            style={{ height: '40px' }}
            onChange={(e) => setQuery(e.target.value)}
            onKeyDown={(e) => e.key === "Enter" && handleSearch()}
          />
          <button
            className="px-6 py-2 bg-blue-500 text-white rounded-r-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
            style={{ height: '40px' }}
            onClick={handleSearch}
          >
            Search
          </button>
        </div>

        <div className="mt-6 space-x-4">
          <a href="https://sushantcodes.tech" className="text-blue-500 hover:underline">
            About
          </a>
          <a href="https://sushantcodes.tech" className="text-blue-500 hover:underline">
            Contact
          </a>
          <a href="https://sushantcodes.tech" className="text-blue-500 hover:underline">
            Follow
          </a>
        </div>
      </div>

      {/* Results Area */}

      {isSearched ? (
        <div className="mt-12 mx-24">
          <h2 className="text-2xl text-slate-800 mb-6">Search Results for "{query}"</h2>
          <div className="space-y-4">
            {demoResults.map((result, index) => (
              <div key={index} className="border-b border-gray-200 pb-4">
                <a
                  href={result.link}
                  className="text-blue-600 hover:underline text-lg"
                >
                  {result.title}
                </a>
                <p className="text-gray-600">{result.description}</p>
              </div>
            ))}
          </div>
        </div>
      ) : (
        <div className="ml-24 mt-12 mr-24">Please search something to see results.</div>
      )}
    </div>
  );
};



export default LandingPage;