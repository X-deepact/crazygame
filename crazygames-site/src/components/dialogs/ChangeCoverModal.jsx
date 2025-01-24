import React, { useState, useEffect } from 'react';
import {useHistory} from "react-router-dom";

import SideModal from '@/components/dialogs/SideModal'
import ProfileGallery from "@/components/ProfileGallery";

const ChangeCoverModal = ({isOpen, onClose, profile, onChange, onBack}) => {
  const [selectedItem, setSelectedItem] = useState("");

  const items = [
    "https://imgs.crazygames.com/userportal/covers/Abstract.jpg",
    "https://imgs.crazygames.com/userportal/covers/Beauty.jpg",
    "https://imgs.crazygames.com/userportal/covers/Beauty2.jpg",
    "https://imgs.crazygames.com/userportal/covers/Car.jpg",
    "https://imgs.crazygames.com/userportal/covers/Car2.jpg",
    "https://imgs.crazygames.com/userportal/covers/Gamer.jpg",
    "https://imgs.crazygames.com/userportal/covers/Landscape.jpg",
    "https://imgs.crazygames.com/userportal/covers/Landscape2.jpg",
    "https://imgs.crazygames.com/userportal/covers/Landscape3.jpg",
    "https://imgs.crazygames.com/userportal/covers/Landscape4.jpg",
    "https://imgs.crazygames.com/userportal/covers/Landscape5.jpg",
    "https://imgs.crazygames.com/userportal/covers/Landscape6.jpg",
    "https://imgs.crazygames.com/userportal/covers/War.jpg",
    "https://imgs.crazygames.com/userportal/covers/War2.jpg",
    "https://imgs.crazygames.com/userportal/covers/War3.jpg",
    "https://imgs.crazygames.com/userportal/covers/Space.jpg",
    "https://imgs.crazygames.com/userportal/covers/Space2.jpg",
    "https://imgs.crazygames.com/userportal/covers/Space3.jpg",
    "https://imgs.crazygames.com/userportal/covers/Space4.jpg",
    "https://imgs.crazygames.com/userportal/covers/Space5.jpg",
  ];

  const handleSelect = (url) => {
    setSelectedItem(url);
    onChange([{coverImage: url}]);
  };

  useEffect(() => {
    setSelectedItem(profile.coverImage || "");
  }, [profile]);

  return (
    <SideModal title="Change Cover" isOpen={isOpen} onClose={onClose} onBack={onBack}>
      <ProfileGallery type="cover" items={items} selectedItem={selectedItem} onSelect={handleSelect} />
    </SideModal>
  );
}

export default ChangeCoverModal;