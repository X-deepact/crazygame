import React, {useState} from 'react';
import {useHistory} from "react-router-dom";

import ToggleCheckbox from "@/components/ToggleCheckbox";
import SideModal from '@/components/dialogs/SideModal'

const PrivacyPreferencesModal = ({isOpen, onClose, onBack}) => {
  const history = useHistory();

  // States for each toggle switch
  const [publicProfile, setPublicProfile] = useState(true);
  const [friendRequests, setFriendRequests] = useState(true);
  const [shareFriends, setShareFriends] = useState(true);
  const [inviteFriends, setInviteFriends] = useState(true);

  const openPolicy = () => {
    history.push("/privacy-policy");
    window.location.reload();
  };

  return (
    <SideModal title="Privacy Preferences" isOpen={isOpen} onClose={onClose} onBack={onBack}>
        <section className="mb-6">
          <h3 className="text-xl text-white mb-2">Community</h3>
          <div className="space-y-4 rounded-lg bg-gray-900 p-4">
            <ToggleCheckbox label="The content of my profile is public" checked={publicProfile} onChange={() => setPublicProfile(!publicProfile)} />
            <ToggleCheckbox label="Users can send me friend requests" checked={friendRequests} onChange={() => setFriendRequests(!friendRequests)} />
          </div>
        </section>

        {/* Site Notifications Section */}
        <section className="mb-6">
          <h3 className="text-xl text-white mb-2">Friends</h3>
          <div className="space-y-4 rounded-lg bg-gray-900 p-4">
            <ToggleCheckbox label="Share my game activity with friends" checked={shareFriends} onChange={() => setShareFriends(!shareFriends)} />
            <ToggleCheckbox label="Friends can send me game invites" checked={inviteFriends} onChange={() => setInviteFriends(!inviteFriends)} />
          </div>
        </section>

        <section>
          <h3 className="text-xl text-white mb-2">Your data at CrazyGames</h3>
          <div className="space-y-4 rounded-lg bg-gray-900 p-4">
            <div
              onClick={openPolicy}
              className="flex justify-between items-center cursor-pointer">
              <span>Privacy Policy & Cookie Preferences</span>
              <i className="fa fa-external-link-alt"></i>
            </div>
          </div>
        </section>
    </SideModal>
    );
  }

export default PrivacyPreferencesModal;