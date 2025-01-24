import React, {useState, useEffect} from 'react';

import SideModal from '@/components/dialogs/SideModal'
import ProfileClickableLabel from "@/components/ProfileClickableLabel";

const ChangeBirthdayModal = ({isOpen, onClose, profile, onBack, onSave}) => {
    const [birthday, setBirthday] = useState("");
    const [error, setError] = useState("");

    const validate = () => {
        if (birthday === "") {
//            setError("Please select a birthday");
            return false;
        }

//        if (birthday === profile.birthday) {
//            setError("Please select a birthday");
//            return false;
//        }

        if (!birthday) return false;
        const today = new Date();
        const birthDate = new Date(birthday);
        const age = today.getFullYear() - birthDate.getFullYear();
        const monthDiff = today.getMonth() - birthDate.getMonth();
        const dayDiff = today.getDate() - birthDate.getDate();

        // Adjust age calculation if the birthdate hasn't occurred yet this year
        if (monthDiff < 0 || (monthDiff === 0 && dayDiff < 0)) {
          return age - 1 >= 13; // Return true if user is at least 13
        }

        return age >= 13;
    };

    const handleSave = () => {
        onSave([{birthday}]);
        onBack();
    };

    useEffect(() => {
        // country dropdown is loaded programmatically so need to select previous country after loaded.
        setBirthday(profile.birthday);
    }, [profile]);

    return (
        <SideModal title="Update Birthday" isOpen={isOpen} onClose={onClose} onBack={onBack}>
            <input
                type="date"
                value={birthday}
                onChange={(e) => setBirthday(e.target.value)}
                className="w-full px-4 py-2 h-[64px] mb-3 bg-gray-900 border border-gray-600 hover:border-white rounded-lg text-white focus:outline-none focus:ring-0 focus:ring-white focus:border-white" />

            { !validate() && (
                <p className="text-[14px] mt-2 mb-4 px-4 text-red-500">You need to be 13 years or older to play on CrazyGames!</p>
            )}

            <button
                disabled={birthday===profile.birthday || !validate()}
                onClick={handleSave}
                className="w-full h-[64px] mt-3 mb-3 bg-purple-600 hover:bg-purple-500 disabled:bg-gray-700 disabled:cursor-not-allowed text-lg rounded-full ">Confirm Change</button>

        </SideModal>
    );
}

export default ChangeBirthdayModal;