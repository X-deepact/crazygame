import React, {useState, useEffect} from 'react';

const ProfileGallery = ({type, items, selectedItem, onSelect}) => {
    return (
        <div className="flex justify-center p-4">
            <div
              style={{maxHeight: "calc(100vh - 215px)",}}
              className={`grid ${type==='avatar' ? "grid-cols-2" : "grid-cols-1 w-full"} gap-4 overflow-y-auto`}>
              {items.map((url) => (
                <div
                  key={url}
                  className={`relative ${type==="avatar" ? "w-32" : "w-full"} h-32 rounded-lg overflow-hidden cursor-pointer transition-all duration-300 ${
                    selectedItem === url
                      ? "border-4 border-green-600"
                      : "border-4 border-transparent"
                  }`}
                  onClick={() => onSelect(url)}
                >
                  <img
                    src={url}
                    alt="Avatar"
                    className="object-cover w-full h-full"
                  />
                  {selectedItem === url && (
                    <div className="absolute top-[-16px] right-[-16px] bg-green-600 w-10 h-10 flex items-end pb-1 px-2 rounded-full">
                      <i className="fas fa-check text-white text-sm"></i>
                    </div>
                  )}
                </div>
              ))}
            </div>
          </div>
    );
};

export default ProfileGallery;