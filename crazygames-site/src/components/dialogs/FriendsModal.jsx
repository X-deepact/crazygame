import React from 'react';
import {useHistory} from "react-router-dom";

import SideModal from '@/components/dialogs/SideModal'

const FriendsModal = ({isOpen, onClose, isAuthenticated, onLogin, onShare}) => {
  const history = useHistory();

  const handlePlayWithFriend = () => {
    history.push("/with-friend");
    window.location.reload();
  };

  return (
    <SideModal title="Friends" isOpen={isOpen} onClose={onClose}>
      <div className="flex flex-col items-center text-center">
        { isAuthenticated && (
          <div className="flex flex-col w-full space-y-4 mb-5">
            {/* Search Input */}
            <div className="flex items-center bg-gray-600 border border-gray-500 focus:border-purple-500 hover:border-purple-500 rounded-full py-2 px-4">
              <i className="fas fa-search text-white mr-2"></i>
              <input
                type="text"
                placeholder="Search new or existing friends"
                className="bg-transparent text-white outline-none flex-1 border-none focus:outline-none focus:ring-0"
              />
            </div>

            {/* "Games for friends" Button */}
            <button
              onClick={handlePlayWithFriend}
              className="bg-green-500 text-white py-4 px-4 rounded-full hover:bg-green-600">
              <i className="fas fa-users mr-1"></i>
              <span>Games for friends</span>
            </button>

            {/* "Share profile" Button */}
            <button
              onClick={() => onShare('friends')}
              className="bg-gray-600 text-white py-4 px-4 rounded-full hover:bg-gray-700">
              <i className="fas fa-share-alt mr-1"></i>
              <span>Share profile</span>
            </button>
          </div>
        )}

        {/* Image */}
        <img
          src="https://imgs.crazygames.com/crazygames/friends/BringYourFriends2.svg?auto=format%2Ccompress&q=45&cs=strip&ch=DPR" // Replace with actual image URL
          alt="Bring your friends"
          className="mb-4"
        />

        {/* Title */}
        <h3 className="text-lg font-bold mb-2">{ isAuthenticated ? "Invite your friends" : "Bring your friends!" }</h3>

        {/* Description */}
        { isAuthenticated && (
          <p className="text-gray-400 mb-6">
            Find friends by searching for their usernames, or <a className="text-purple-500 cursor-pointer" onClick={() => onShare('friends')}>share your QR code / invite link</a>
          </p>
        )}

        { !isAuthenticated && (
          <p className="text-gray-400 mb-6">
            Create a CrazyGames account to start inviting your friends.
          </p>
        )}

        {/* Button */}
        { !isAuthenticated && (
          <button
            onClick={onLogin}
            className="bg-purple-600 text-white px-6 py-2 rounded-full hover:bg-purple-500">
              Log in / Register
            </button>
        )}
        </div>
      </SideModal>
      );
  }

export default FriendsModal;