import React from 'react';

const LandingPage: React.FC = () => {
  return (
    <div className="flex flex-col mt-12 items-center min-h-screen bg-white">
      <h1 className='mb-12 text-4xl text-slate-800'>IntelliSearch</h1>

      <div className="flex items-center">
        <input
          type="text"
          placeholder="Search the web..."
          className="w-96 px-4 py-2 border border-gray-300 rounded-l-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
          style={{ height: '40px' }} 
        />
        <button
          className="px-6 py-2 bg-blue-500 text-white rounded-r-lg hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500"
          style={{ height: '40px' }} 
        >
          Search
        </button>
      </div>

      <div className="mt-6 space-x-4">
        <a href="https://example.com" className="text-blue-500 hover:underline">
          About
        </a>
        <a href="https://example.com" className="text-blue-500 hover:underline">
          Contact
        </a>
        <a href="https://example.com" className="text-blue-500 hover:underline">
          Follow
        </a>
      </div>
    </div>
  );
};

export default LandingPage;