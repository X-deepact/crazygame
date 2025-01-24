import React, { useState, useEffect } from 'react';
import {useHistory} from "react-router-dom";

import SideModal from '@/components/dialogs/SideModal'
import SocialButtons from "@/components/SocialButtons";
import FooterLinks from "@/components/FooterLinks";

const ProfileModal = ({isOpen, onClose, profile, onShare, onNotificationPreferences, onPrivacyPreferences, onAccountSettings, onLogout}) => {
  const history = useHistory();
  const onViewProfile = () => {
    history.push("/profile");
    window.location.reload();
  };

  useEffect(() => {
  }, [profile]);

  return (
    <SideModal title="" isOpen={isOpen} onClose={onClose}>
      {/* User Name */}
        <div className="relative flex justify-center mb-2">
          <img alt=""
               className="object-cover w-[64px] h-[64px] border-[3px] border-white rounded-full"
               src="https://imgs.crazygames.com/userportal/avatars/78.png?auto=format%2Ccompress&q=45&cs=strip&ch=DPR&w=150&h=150" />
        </div>

        <div className="text-center text-white mb-4">
          <h2 className="text-xl font-bold">@{profile.username}</h2>
          <p className="text-sm text-gray-400">{profile.email}</p>
        </div>

        {/* Action Buttons */}
        <div className="flex flex-col mb-4 space-y-2">
          <button
            onClick={onViewProfile}
            className="bg-purple-600 text-white py-2 px-4 rounded-full hover:bg-purple-500">
            View profile
          </button>
          <button
            onClick={() => onShare('profile')}
            className="bg-blue-600 text-white py-2 px-4 rounded-full hover:bg-blue-500">
            Share profile
          </button>
        </div>

        {/* Profile Completion */}
        <div className="mb-2">
          <p className="text-white pl-2 mb-1">Your profile is 80% complete</p>
          <div className="h-2 bg-gray-600 rounded-full">
            <div className="h-2 bg-blue-500 rounded-full w-4/5"></div>
          </div>
        </div>

        <div className="space-y-1">
          <button
            onClick={() => onNotificationPreferences('profile')}
            className="flex items-center space-x-2 bg-transparent text-white py-2 px-4 rounded-md hover:bg-gray-700">
            <i className="fas fa-bell"></i> {/* Notification Icon */}
            <span>Notification preferences</span>
          </button>
          <button
            onClick={() => onPrivacyPreferences()}
            className="flex items-center space-x-2 bg-transparent text-white py-2 px-4 rounded-md hover:bg-gray-700">
            <i className="fas fa-shield-alt"></i> {/* Privacy Icon */}
            <span>Privacy preferences</span>
          </button>
          <button
            onClick={onAccountSettings} 
            className="flex items-center space-x-2 bg-transparent text-white py-2 px-4 rounded-md hover:bg-gray-700">
            <i className="fas fa-cogs"></i> {/* Account settings Icon */}
            <span>Account settings</span>
          </button>
          <button
            className="flex items-center space-x-2 bg-transparent text-white py-2 px-4 rounded-md hover:bg-gray-700"
            onClick={() => onLogout()}
          >
            <i className="fas fa-sign-out-alt"></i> {/* Log Out Icon */}
            <span>Log out</span>
          </button>
          <hr></hr>
          <button className="flex items-center space-x-2 bg-transparent text-white py-2 px-4 rounded-md hover:bg-gray-700">
            <i className="fas fa-envelope"></i> {/* Contact us Icon */}
            <span>Contact us</span>
          </button>
          <button className="flex items-center space-x-2 bg-transparent text-white py-2 px-4 rounded-md hover:bg-gray-700">
            <i className="fas fa-download"></i> {/* Install App Icon */}
            <span>Install app</span>
          </button>
          <hr></hr>
        </div>


        <div className="mt-3 mb-3">
          <FooterLinks type="profile" />
        </div>

        <hr className="border-gray-600" />

        <div className="mt-3">
          <SocialButtons type="profile" />
        </div>

    </SideModal>
  );
}

export default ProfileModal;