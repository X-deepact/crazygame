import React, { useState, useEffect } from 'react';
import {useHistory} from "react-router-dom";

import SideModal from '@/components/dialogs/SideModal'
import ProfileGallery from "@/components/ProfileGallery";

const ChangeAvatarModal = ({isOpen, onClose, profile, onChange, onBack}) => {
  const [selectedItem, setSelectedItem] = useState("");

  const items = Array.from({ length: 112 }, (_, index) =>
    `https://imgs.crazygames.com/userportal/avatars/${index + 1}.png`
  );

  const handleSelect = (url) => {
    setSelectedItem(url);
    onChange([{avatarImage: url}]);
  };

  useEffect(() => {
    setSelectedItem(profile.avatarImage || "");
  }, [profile]);

  return (
    <SideModal title="Change Avatar" isOpen={isOpen} onClose={onClose} onBack={onBack}>
      <ProfileGallery type="avatar" items={items} selectedItem={selectedItem} onSelect={handleSelect} />
    </SideModal>
  );
}

export default ChangeAvatarModal;