import React, {useState, useEffect} from 'react';
import {useHistory} from "react-router-dom";

import SideModal from '@/components/dialogs/SideModal'
import ProfileClickableLabel from "@/components/ProfileClickableLabel";

const EditProfileModal = ({isOpen, onClose, onBack, profile, onEditUsername, onEditBirthday, onChangeAvatar, onChangeCover, onSave}) => {
    const countries = [
        'United States',
        'Canada',
        'United Kingdom',
        'Australia',
        'India',
        'South Africa',
        'Germany',
        'France',
        'Brazil',
        'Japan',
        'Mexico',
        'Italy',
        'Russia',
        'China',
    ];

    const [country, setCountry] = useState("");
    const [gender, setGender] = useState('');

    const onChangeSelection = (type, value) => {
        if (type === "country") {
            setCountry(value);
            onSave([{country: value}]);
        } else if (type === "gender") {
            setGender(value);
            onSave([{gender: value}])
        }
    }

    const handleUsername = () => {

    };

    useEffect(() => {
        // country dropdown is loaded programmatically so need to select previous country after loaded.
        setCountry(profile.country);
        setGender(profile.gender);
    }, [profile]);


    return (
        <SideModal title="Edit your Profile" isOpen={isOpen} onClose={onClose} onBack={onBack}>
            <div className="relative flex justify-center mb-[4rem]">
                <img
                    src={profile.coverImage} // Placeholder for profile image
                    alt="profile"
                    className="object-cover"
                />

                <button
                    onClick={onChangeCover}
                    className="absolute bottom-1 right-1 w-8 h-8 bg-purple-600 hover:bg-purple-500 rounded-full text-white">
                    <i className="fas fa-camera text-sm"></i> {/* Edit Icon */}
                </button>

                {/* Profile Image (overlapping) */}
                <img
                    src={profile.avatarImage} // Placeholder for profile image
                    alt="profile"
                    className="object-cover w-[90px] h-[90px] rounded-full absolute bottom-[-45px] left-0 right-0 m-auto z-10"
                />

                <button
                    onClick={onChangeAvatar}
                    className="absolute bottom-[-45px] z-10 left-[55px] right-0 m-auto w-8 h-8 bg-purple-600 hover:bg-purple-500 rounded-full text-white">
                    <i className="fas fa-camera text-sm"></i> {/* Edit Icon */}
                </button>
            </div>

            <div className="mt-8">
                <ProfileClickableLabel label="Username" value={profile.username} onClick={onEditUsername} />

                <div className="mt-4 mb-8">
                    <select
                        value={country}
                        onChange={(e) => onChangeSelection('country', e.target.value)}
                        className="w-full px-4 py-2 h-[64px] bg-gray-900 border border-gray-600 hover:border-white rounded-lg text-white focus:outline-none focus:ring-0 focus:ring-white focus:border-white">
                        <option disabled key={-1} value="">Location</option>
                        {countries.map((country, index) => (
                            <option key={index} value={country}>{country}</option>
                        ))}
                    </select>
                </div>
            </div>

            <hr/>

            {/* Personal Details Section */}
            <div className="mt-8">
                <h2 className="text-lg font-semibold text-white">Personal Details</h2>
                <p className="text-sm mt-2 mb-4 text-gray-300">This information is shared exclusively with CrazyGames
                    and will not be visible on your profile.</p>

                <ProfileClickableLabel label="Birthday" value={profile.birthday} onClick={onEditBirthday} />

                <div className="mt-4">
                    <select
                        value={gender}
                        onChange={(e) => onChangeSelection('gender', e.target.value)}
                        className="w-full px-4 py-2 h-[64px] bg-gray-900 border border-gray-600 hover:border-white rounded-lg text-white focus:outline-none focus:ring-0 focus:ring-white focus:border-white">
                        <option disabled value="">Gender</option>
                        <option value="M">Male</option>
                        <option value="F">Female</option>
                        <option value="O">Other</option>
                        <option value="Z">I prefer not to say</option>
                    </select>
                </div>
            </div>
        </SideModal>
    );
}

export default EditProfileModal;