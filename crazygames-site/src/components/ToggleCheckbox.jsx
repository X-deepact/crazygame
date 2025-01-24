import React from 'react';

const ToggleCheckbox = ({label, checked, onChange}) => {
    return (
        <div className="flex justify-between items-center">
            <span className="pr-4">{label}</span>

            <label className="relative inline-flex items-center cursor-pointer">
                <input
                    type="checkbox"
                    checked={checked}
                    onChange={onChange}
                    className="hidden"
                />
                <div
                    className={`w-12 h-6 bg-gray-400 rounded-full flex items-center px-1 ${
                        checked ? 'bg-green-400' : 'bg-gray-400'
                    }`}
                >
                    <div
                        className={`w-4 h-4 bg-white rounded-full shadow-md transform transition-transform ${
                            checked ? 'translate-x-6' : 'translate-x-0'
                        }`}
                    />
                </div>
            </label>
        </div>
    );
};

export default ToggleCheckbox;