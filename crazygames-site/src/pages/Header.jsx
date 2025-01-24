import React from 'react';

const Header = ({ isAuthenticated, toggleSidebar, profile, openLoginModal, openFriendsModal, openNotificationsModal, openMyGamesModal, openProfileModal }) => {
  const handleUser = () => {
      if (isAuthenticated)
        openProfileModal();
      else
        openLoginModal({});
  };

    return (
      <header className="bg-gray-900 text-white p-4 flex items-center fixed top-0 left-0 z-50 w-full">
        {/* Menu Button */}
        <button
          className="text-white mr-4 focus:outline-none"
          onClick={toggleSidebar}
        >
          <i className="fas fa-bars"></i>
        </button>

        {/* Website Label */}
        <div className="text-xl font-bold flex items-center">
          Crazy Games
        </div>

        {/* Search Input */}
        <div className="relative mx-auto flex items-center w-[300px] hidden lg:block">
          <input
            type="text"
            placeholder="Search"
            className="w-full p-2 pl-4 bg-gray-700 text-white rounded-full focus:outline-none"
          />
          <button className="absolute right-2 top-1/2 transform -translate-y-1/2 text-white p-2 rounded-full">
            <i className="fas fa-search"></i>
          </button>
        </div>

        {/* Right-Side Buttons */}
        <div className="flex items-center space-x-4 ml-auto">
          {/* Round Buttons */}
          <button className="w-10 h-10 rounded-full bg-gray-700 hover:bg-gray-600" onClick={openFriendsModal}>
            <i className="fas fa-user-friends"></i>
          </button>
          { isAuthenticated && (
            <button className="w-10 h-10 rounded-full bg-gray-700 hover:bg-gray-600 hidden md:block" onClick={openNotificationsModal}>
              <i className="fas fa-bell"></i>
            </button>
          )}
          <button className="w-10 h-10 rounded-full bg-gray-700 hover:bg-gray-600 hidden md:block" onClick={openMyGamesModal}>
            <i className="fas fa-heart"></i>
          </button>

          {/* Log In Button */}
          <button
            style={
              isAuthenticated && profile?.avatarImage
                ? { backgroundImage: `url(${profile.avatarImage})`, backgroundSize: 'cover', backgroundPosition: 'center' }
                : {}
            }
            className={`h-10 ${isAuthenticated ? "w-10" : ""} bg-purple-600 px-2 py-2 rounded-full hover:bg-purple-500`} onClick={handleUser}>
            { isAuthenticated ? (
                !profile?.avatarImage && <i className="fas fa-user"></i>
              ) : (
                "Log in"
              )
            }
          </button>
        </div>
      </header>
    );
  };

export default Header;