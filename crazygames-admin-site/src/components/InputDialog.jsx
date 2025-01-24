import React from "react";

const InputDialog = ({isOpen, onClose, title, error, onOk, onCancel, children}) => {
    if (!isOpen) return null;

    return (
        <div className="fixed inset-0 bg-gray-600 bg-opacity-50 flex justify-center items-center">
            <div className="w-1/3 bg-white p-6 rounded-lg shadow-lg relative">
                <button
                    onClick={onClose}
                    className="absolute top-3 right-3 text-gray-500 hover:text-gray-700">
                    <i className="fa fa-times"></i> {/* Close Icon */}
                </button>
                <h3 className="text-lg font-semibold mb-4">{title}</h3>

                <form>
                    { error && (
                        <h6 className="bg-red-100 text-red-500 rounded-md px-4 py-4 mb-5">{error}</h6>
                    )}

                    {children}

                    <div className="flex justify-end space-x-4">
                        <button
                            type="button"
                            onClick={onOk}
                            className="px-4 py-2 bg-blue-700 hover:bg-blue-600 text-white rounded-md"
                        >
                            Save
                        </button>
                        <button
                            type="button"
                            onClick={onCancel}
                            className="px-4 py-2 bg-gray-300 hover:bg-gray-200 rounded-md"
                        >
                            Cancel
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default InputDialog;