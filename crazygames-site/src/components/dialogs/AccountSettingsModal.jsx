import React, {useState} from 'react';
import {useHistory} from "react-router-dom";

import SideModal from '@/components/dialogs/SideModal'

const AccountSettingsModal = ({isOpen, onClose, onBack}) => {
  const history = useHistory();

  // States for each toggle switch
  const [publicProfile, setPublicProfile] = useState(true);
  const [friendRequests, setFriendRequests] = useState(true);
  const [shareFriends, setShareFriends] = useState(true);
  const [inviteFriends, setInviteFriends] = useState(true);

  const onDelete = () => {
//    history.push("/privacy-policy");
//    window.location.reload();
  };

  return (
    <SideModal title="Privacy Preferences" isOpen={isOpen} onClose={onClose} onBack={onBack}>
        <section>
          <h3 className="text-sm text-gray-600 mb-2">Delete your account</h3>
          <div className="space-y-4 py-2 mb-2">
            <div
              onClick={onDelete}
              className="flex justify-between font-semibold text-md items-center cursor-pointer">
              <span>Delete your CrazyGames Account</span>
              <i className="fa fa-arrow-right"></i>
            </div>
          </div>
          <hr />
        </section>
    </SideModal>
    );
  }

export default AccountSettingsModal;