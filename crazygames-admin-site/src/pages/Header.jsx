import useAuthStore from '@/hooks/store/useAuth';
import React from 'react';
import toast from 'react-hot-toast';
import { useHistory } from 'react-router-dom';

const Header = ({ toggleSidebar, isSidebarCollapsed }) => {
  const history = useHistory();
  const { clearToken } = useAuthStore();

  const onLogout = () => {
    clearToken();
    toast.success('Successfully logged out!');
    history.push('/login');
  };

  return (
    <header className='fixed top-0 left-0 right-0 h-[4.5rem] bg-gray-900 text-white flex items-center justify-between px-6 shadow-md z-10'>
      {/* Toggle Button */}
      <button
        className='p-2 bg-gray-800 rounded hover:bg-gray-700'
        onClick={toggleSidebar}
      >
        <i className={`fas ${isSidebarCollapsed ? 'fa-bars' : 'fa-bars'}`}></i>
      </button>

      {/* Title */}
      <div className='text-xl font-bold'>Crazy Games</div>

      {/* Actions */}
      <div className='flex items-center space-x-4'>
        {/* Notifications Button */}
        <button className='p-2 w-10 h-10 flex items-center justify-center bg-gray-800 rounded-full hover:bg-gray-700'>
          <i className='fas fa-bell'></i>
        </button>

        {/* Profile Button */}
        <button className='p-2 w-10 h-10 flex items-center justify-center bg-gray-800 rounded-full hover:bg-gray-700'>
          <i className='fas fa-user'></i>
        </button>

        {/* Logout Button */}
        <button
          onClick={onLogout}
          className='p-2 w-10 h-10 flex items-center justify-center bg-red-600 rounded-full hover:bg-red-500'
        >
          <i className='fas fa-sign-out-alt'></i>
        </button>
      </div>
    </header>
  );
};

export default Header;
