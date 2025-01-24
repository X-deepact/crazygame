import React from 'react';

export function InputErrorMessage({ message }) {
  if (!message) return null;

  return <p className='text-red-500 text-sm'>{message}</p>;
}
