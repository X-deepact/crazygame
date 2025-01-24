import React from 'react';

const SideModal = ({
  title,
  isOpen,
  onClose,
  onSettings,
  onBack,
  children,
}) => {
  if (!isOpen) return null;

  return (
    <div
      className='fixed inset-x-0 bottom-0 top-[4.6rem] bg-black bg-opacity-50 z-50 overflow-auto'
      onClick={onClose}
    >
      {/* Modal Content */}
      <div
        className='absolute top-[0.4rem] right-0 min-h-[calc(100vh-5.5rem)] w-[400px] bg-gray-800 text-white p-6 rounded-lg'
        onClick={(e) => e.stopPropagation()}
        style={{
          marginBottom: '1rem', // Adding bottom margin for better spacing
          marginRight: '0.4rem',
        }}
      >
        <h4 className='text-md text-center mb-[2.5rem]'>{title}</h4>

        {/* Back Button */}
        {onBack && (
          <button
            className='absolute top-5 left-5 text-white focus:outline-none'
            onClick={onBack}
          >
            <i className='fas fa-chevron-left'></i>
          </button>
        )}

        {onSettings && (
          <button
            className='absolute top-5 left-5 text-white focus:outline-none'
            onClick={onSettings}
          >
            <i className='fas fa-cog'></i>
          </button>
        )}

        {/* Close Button */}
        <button
          className='absolute top-5 right-5 text-white focus:outline-none'
          onClick={onClose}
        >
          <i className='fa text-xl fa-times'></i>
        </button>

        {children}
      </div>
    </div>
  );
};

export default SideModal;
