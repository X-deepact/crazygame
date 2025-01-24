import React from 'react';

const ProfileClickableLabel = ({label, value, onClick}) => {
    return (
        <div
            onClick={onClick}
            className="w-full px-4 py-1 h-[64px] bg-gray-900 border border-gray-600 hover:border-white rounded-lg text-white">
            <p className="text-gray-400 text-[12px] pt-[2px]">{label}</p>
            <p className="text-white text-[18px]">{value}</p>
        </div>
    );
};

export default ProfileClickableLabel;