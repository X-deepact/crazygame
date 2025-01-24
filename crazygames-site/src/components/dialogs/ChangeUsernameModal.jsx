import React, {useState, useEffect} from 'react';

import SideModal from '@/components/dialogs/SideModal'
import ProfileClickableLabel from "@/components/ProfileClickableLabel";

const ChangeUsernameModal = ({isOpen, onClose, profile, onBack, onSave}) => {
    const [username, setUsername] = useState("");

    const validate = () => {
        // Regular expression to match valid characters: letters, numbers, periods, and underscores
        const regex = /^[A-Za-z0-9._]{6,20}$/;

        // Check if the username matches the regex and its length is between 6 and 20 characters
        return regex.test(username);
    };

    const adjectives = ['Speedy', 'Mighty', 'Quick', 'Brave', 'Silent', 'Clever', 'Happy', 'Fierce', 'Loyal', 'Bold'];
    const animals = ['Tiger', 'Lion', 'Eagle', 'Shark', 'Wolf', 'Elephant', 'Panther', 'Cheetah', 'Falcon', 'Bear'];
    const numbers = ['123', '456', '789', '101', '202', '303', '404', '505'];

    const generateRandomUsername = () => {
        const adjective = adjectives[Math.floor(Math.random() * adjectives.length)];
        const animal = animals[Math.floor(Math.random() * animals.length)];
        const number = numbers[Math.floor(Math.random() * numbers.length)];

        // Combine the selected words and numbers to create a username
        const randomUsername = `${adjective}-${animal}.${number}`;

        // Ensure the username is between 6 and 20 characters
        if (randomUsername.length >= 6 && randomUsername.length <= 20) {
            setUsername(randomUsername);
        } else {
            generateRandomUsername(); // Recursively generate a new one if the length is not valid
        }
    };

    const onGenerateUsername = () => {
        generateRandomUsername();
    };

    const handleSave = () => {
        onSave([{username}]);
        onBack();
    };

    useEffect(() => {
        // country dropdown is loaded programmatically so need to select previous country after loaded.
        setUsername(profile.username);
    }, [profile]);

    return (
        <SideModal title="Change Username" isOpen={isOpen} onClose={onClose} onBack={onBack}>
            <h2 className="text-[17px] leading-[1.7] text-gray-300 text-center mb-3">Tired of your current username, huh? No problem, pick a new one here:</h2>

            <input
                type="text"
                value={username}
                onChange={(e) => setUsername(e.target.value)}
                className="w-full px-4 py-2 h-[64px] mb-3 bg-gray-900 border border-gray-600 hover:border-white rounded-lg text-white focus:outline-none focus:ring-0 focus:ring-white focus:border-white" />

            <div className="flex w-full justify-center mb-4">
                <span className="text-white">No idea?</span>
                <a  onClick={onGenerateUsername}
                    className="ml-1 cursor-pointer text-purple-600 hover:text-purple-500">Generate one</a>
            </div>

            <button
                disabled={username===profile.username || !validate()}
                onClick={handleSave}
                className="w-full h-[64px] mb-3 bg-purple-600 hover:bg-purple-500 disabled:bg-gray-700 disabled:cursor-not-allowed text-lg rounded-full ">Save Username</button>

            <p className="text-[13px] mt-2 mb-4 text-gray-500">Usernames must contain only letters, numbers, periods and underscores and have 6 - 20 characters. We will permanently ban accounts with toxic & inappropriate usernames without any notice. So let’s keep the community nice and safe for everybody!</p>

        </SideModal>
    );
}

export default ChangeUsernameModal;