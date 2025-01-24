import React, { useState, useEffect } from 'react';

const PeerInput = ({ id, type, label, value, setValue, children }) => {
  const [input, setInput] = useState(value || '');
  const handleOnChange = (e) => {
    setValue(e.target.value);
    setInput(e.target.value);
  };

  useEffect(() => {
    setInput(value);
  }, [value]);

  return (
    <div className='relative mb-4'>
      <input
        type={type}
        id={id}
        value={value}
        onChange={(e) => handleOnChange(e)}
        placeholder=' '
        className='peer w-full bg-gray-700 text-white rounded-lg p-3 focus:ring-2 focus:ring-blue-500 focus:outline-none'
      />
      <label
        htmlFor={id}
        className={`absolute left-3 top-3 text-gray-400 text-sm transform transition-all px-1 rounded
                peer-placeholder-shown:top-3 peer-placeholder-shown:text-gray-400 peer-placeholder-shown:text-base
                peer-focus:top-[-10px] peer-focus:left-2 peer-focus:text-blue-500 peer-focus:text-sm peer-focus:bg-gray-800
                ${
                  input
                    ? 'top-[-10px] left-2 text-blue-500 bg-gray-800 text-sm'
                    : ''
                }
              `}
      >
        {label}
      </label>

      {children}
    </div>
  );
};

export default PeerInput;
