import React from "react";

const ConfirmDialog = ({isOpen, onClose, title, onOk, onCancel}) => {
  if (!isOpen) return null;

    return (
        <div
          onClick={onClose}
          className="fixed inset-0 bg-gray-600 bg-opacity-50 flex justify-center items-center">
            <div
              onClick={(e) => e.stopPropagation()}
              className="bg-white p-6 rounded-lg shadow-lg">
              <h3 className="text-lg font-semibold mb-4">{title}</h3>
              <div className="flex justify-end space-x-4">
                <button
                  onClick={onOk}
                  className="px-4 py-2 bg-red-700 hover:bg-red-500 text-white rounded-md">
                  Yes
                </button>
                <button
                  onClick={onCancel}
                  className="px-4 py-2 bg-gray-300 hover:bg-gray-400 rounded-md">
                  No
                </button>
              </div>
            </div>
          </div>
    );
};

export default ConfirmDialog;