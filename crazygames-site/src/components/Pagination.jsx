import React from 'react';

const Pagination = ({ currentPage, totalPages, handlePageChange }) => {
    return (
        <div className="flex justify-center items-center mt-8 space-x-4">
            {/* Previous Button */}
            <button
              onClick={() => handlePageChange(currentPage - 1)}
              disabled={currentPage === 1}
              className={`px-4 py-2 rounded-full ${
                currentPage === 1
                  ? "bg-gray-700 text-gray-500 cursor-not-allowed"
                  : "bg-purple-600 hover:bg-purple-500 text-white"
              }`}
            >
              <i className="fas fa-arrow-left"></i>
            </button>

            {/* Page Numbers */}
            {[...Array(totalPages)].map((_, index) => (
              <button
                key={index}
                onClick={() => handlePageChange(index + 1)}
                className={`w-10 h-10 flex items-center justify-center rounded-full ${
                  currentPage === index + 1
                    ? "bg-purple-600 text-white"
                    : "bg-gray-800 hover:bg-purple-500 text-white"
                }`}
              >
                {index + 1}
              </button>
            ))}
            {/* Next Button */}
            <button
              onClick={() => handlePageChange(currentPage + 1)}
              disabled={currentPage === totalPages}
              className={`px-4 py-2 rounded-full ${
                currentPage === totalPages
                  ? "bg-gray-700 text-gray-500 cursor-not-allowed"
                  : "bg-purple-600 hover:bg-purple-500 text-white"
              }`}
            >
              <i className="fas fa-arrow-right"></i>
            </button>
          </div>
    );
};

export default Pagination;
