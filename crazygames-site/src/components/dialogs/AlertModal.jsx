import React from 'react';

const AlertModal = ({ isOpen, title, message, onClose }) => {
  if (!isOpen) return null; // Don't render anything if modal is not open

  return (
    <div className="fixed inset-0 bg-gray-800 bg-opacity-50 flex items-center justify-center z-50">
      <div className="bg-gray-900 p-6 rounded-lg shadow-lg max-w-sm w-full">
        <h2 className="text-xl text-white font-semibold mb-4">{title}</h2>
        <p className="text-gray-100">{message}</p>
        <div className="mt-4 flex justify-end">
          <button
            onClick={onClose}
            className="px-4 py-2 bg-purple-600 text-white rounded hover:bg-purple-500"
          >
            OK
          </button>
        </div>
      </div>
    </div>
  );
};

export default AlertModal;
