import React, {useState, useEffect} from 'react';

import Constants from '@/components/Constants'

const DataTable = ({ columns, data, currentPage, totalPages, rowsPerPage, onPageChange, addLabel, onAdd, onEdit, onDelete, onPassword }) => {
  const [rows, setRows] = useState(rowsPerPage || 10); // Default rows per page
  const [searchTerm, setSearchTerm] = useState("");

  const getGenderValue = (value) => {
    const record = Constants.GENDERS.find((r) => r.value===value);
    return record?.name;
  };

  const handleRowsChange = (event) => {
    const r = Number(event.target.value);
    setRows(r);
    onPageChange(1, r); // Reset to first page whenever rows per page changes
  };

  const handleKeyDown = (event) => {
    if (event.key === 'Enter') {
      // Trigger search logic when Enter is pressed
      setSearchTerm(event.target.value);
      onPageChange(currentPage, rows, event.target.value);
    }
  };

  // Generate page numbers with ellipses when necessary, always keeping the count to 5 buttons
  const generatePageNumbers = () => {
    const pageNumbers = [];
    const totalButtons = 5;

    // Always include 1st and last page
    const firstPage = 1;
    const lastPage = totalPages;

    // Calculate range of buttons to show
    let startPage = Math.max(firstPage, currentPage - 2);
    let endPage = Math.min(lastPage, currentPage + 2);

    // Adjust the range to always show 5 buttons
    const totalRange = endPage - startPage + 1;
    if (totalRange < totalButtons) {
      if (startPage === firstPage) {
        endPage = Math.min(startPage + totalButtons - 1, lastPage);
      } else {
        startPage = Math.max(endPage - totalButtons + 1, firstPage);
      }
    }

    // Add the page numbers to the list
    if (startPage > firstPage) {
      pageNumbers.push(firstPage);
      if (startPage > firstPage + 1) pageNumbers.push('...');
    }

    for (let i = startPage; i <= endPage; i++) {
      pageNumbers.push(i);
    }

    if (endPage < lastPage) {
      if (endPage < lastPage - 1) pageNumbers.push('...');
      pageNumbers.push(lastPage);
    }

    return pageNumbers;
  };

  const handlePrev = () => {
    if (currentPage > 1) onPageChange(currentPage - 1, rows, searchTerm);
  };

  const handleNext = () => {
    if (currentPage < totalPages) onPageChange(currentPage + 1, rows, searchTerm);
  };

  const handleRefresh = () => {
    onPageChange(1, rows, searchTerm);
  };

  // initialize-------
  useEffect(() => {
  }, []);


  const paginate = (pageNumber) => onPageChange(pageNumber, rows, searchTerm);

  return (
    <div className="bg-white rounded-lg p-4 shadow-lg">
      <div className="flex justify-between mb-4 mt-1 px-1">
        <div className="flex items-center gap-4">
          <div className="flex items-center space-x-2">
            <label htmlFor="rowsPerPage" className="text-sm text-gray-600">Rows:</label>
            <select
              id="rowsPerPage"
              value={rows}
              onChange={handleRowsChange}
              className="p-2 rounded-lg border border-gray-300 w-20"
            >
              <option value={5}>5</option>
              <option value={10}>10</option>
              <option value={20}>20</option>
              <option value={50}>50</option>
              <option value={100}>100</option>
            </select>
          </div>

          <div className="relative w-64">
            <input
              type="text"
              className="border rounded-lg p-2 pl-10 w-full"
              placeholder="Search ....."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              onKeyDown={handleKeyDown}
            />
            <i className="fa fa-search absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-400"></i>
          </div>
        </div>

        <div className="flex items-center gap-4">
          { addLabel && (
            <button
              onClick={onAdd}
              className="bg-green-700 hover:bg-green-600 text-white px-4 py-2 rounded-lg flex items-center">
              <i className="fa fa-plus mr-2"></i> {addLabel}
            </button>
          )}
          <button
            onClick={handleRefresh}
            className="bg-blue-700 hover:bg-blue-600 text-white px-4 py-2 rounded-lg flex items-center">
            <i className="fa fa-sync-alt mr-2"></i> Refresh
          </button>
        </div>
      </div>

      {/* Table */}
      <table className="min-w-full px-1 border-collapse">
        <thead>
          <tr>
            {columns.map((col, index) => (
              <th key={`${col.field}_${index}`} className="text-left border-b border-gray-300 p-2">{col.header}</th>
            ))}
          </tr>
        </thead>
        <tbody>
          {data.map((row, i) => (
            <tr key={i} className="">
              {columns.map((col, index) => (
                <td key={`${col.field}_${index}`} className="text-sm border-b border-gray-300 p-2 ">
                  { col.type==='action' && (
                    <div>
                      { onPassword && (
                        <button title="Reset Password"
                          onClick={() => onPassword(row)}
                          className="text-yellow-500 hover:text-yellow-600 mr-2">
                          <i className="fa fa-key"></i>
                        </button>
                      )}

                      { onEdit && (
                        <button title="Edit"
                          onClick={() => onEdit(row)}
                          className="text-blue-700 hover:text-blue-600 mr-2">
                          <i className="fa fa-edit"></i>
                        </button>
                      )}

                      { onDelete && (
                        <button title="Delete"
                          onClick={() => onDelete(row)}
                          className="text-red-500 hover:text-red-700">
                          <i className="fa fa-trash"></i>
                        </button>
                      )}
                    </div>
                  )}
                  { col.type === 'text' && (
                    <span>{row[col.field]}</span>
                  )}
                  { col.type === 'gender' && (
                    <span>{getGenderValue(row[col.field])}</span>
                  )}

                  { col.type === 'image' && (
                    <img alt="" className="" style={{maxWidth: "120px"}} src={row[col.field]} />
                  )}

                  { col.type === 'video' && (
                      <video
                        src={row[col.field]}
                        style={{maxWidth: "120px"}}
                        className={`absolute top-0 left-0 w-full object-cover opacity-0 group-hover:opacity-100 transition-opacity duration-300`}
                        autoPlay
                        loop
                        muted
                        playsInline
                      />
                  )}
                </td>
              ))}
            </tr>
          ))}
        </tbody>
      </table>

      {/* Pagination */}
      <div className="flex justify-between items-center mt-6">
        <div className="flex-none">
          <span className="text-sm text-gray-900">
            Page {currentPage} of {totalPages}
          </span>
        </div>
        <div className="flex items-center w-1/2 justify-between">
          <button
            disabled={currentPage<=1}
            className="px-3.5 disabled:bg-gray-400 disabled:cursor-not-allowed bg-blue-700 hover:bg-blue-600 text-white rounded-full w-10 h-10 flex items-center"
            onClick={() => handlePrev()}
          >
            <i className="fa fa-arrow-left"></i>
          </button>

          <div className="flex items-center">
            {generatePageNumbers().map((page, index) => (
              <button
                key={index}
                disabled={page==="..."}
                onClick={() => paginate(page)}
                className={`text-sm w-8 h-8 mx-1 rounded-full disabled:text-gray-500
                    ${currentPage!==page && page!=="..." ? "bg-gray-200" : ""}
                    ${page==="..." ? "cursor-not-allowed bg-transparent hover:bg-transparent" : ""}
                    ${currentPage === page ? 'bg-blue-700 text-white cursor-not-allowed' : 'hover:bg-blue-600 hover:text-white'} `}
              >
                {page}
              </button>
            ))}
          </div>

          <button
            disabled={currentPage>=totalPages}
            className="px-3.5 disabled:bg-gray-400 disabled:cursor-not-allowed bg-blue-700 hover:bg-blue-600 text-white rounded-full w-10 h-10 flex items-center"
            onClick={() => handleNext()}
          >
            <i className="fa fa-arrow-right"></i>
          </button>
        </div>
      </div>

    </div>
  );
};

export default DataTable;
