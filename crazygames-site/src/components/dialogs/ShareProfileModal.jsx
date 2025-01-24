import React, {useState, useEffect} from 'react';
import QRCode from 'qrcode';

import SideModal from '@/components/dialogs/SideModal'

const ShareProfileModal = ({isOpen, onClose, profile, alert, onBack}) => {
    const [qrCodeUrl, setQrCodeUrl] = useState('');

    const copyProfileLink = () => {
        navigator.clipboard.writeText('https://example.com/profile'); // Replace with actual profile link
        alert('Profile link copied!');
      };

    useEffect(async () => {
        const generateQrCode = async () => {
            try {
                const url = 'https://example.com/profile'; // Replace with the actual profile URL
                const qrCode = await QRCode.toDataURL(url); // Generate QR code as a data URL
                setQrCodeUrl(qrCode); // Save the generated QR code URL to state
            } catch (error) {
                console.error('Failed to generate QR code:', error);
            }
        };

        await generateQrCode();
    }, []);

    return (
        <SideModal title="Share Profile" isOpen={isOpen} onClose={onClose} onBack={onBack}>
            <p className="text-sm text-gray-300 text-center px-4 mb-4">
                Show this QR code to your friends so they can scan it with their camera and send you a friend request
            </p>

            {qrCodeUrl ? (
                    <img src={qrCodeUrl} alt="QR Code" className="w-[200px] h-[200px] object-contain mx-auto mb-5" />
                  ) : (
                    <p className="text-gray-500">Generating QR Code...</p>
                  )}

            <div className="flex flex-col gap-2">
                <button
                    className="bg-gray-700 px-4 py-3 w-full rounded-full text-white hover:bg-gray-600">
                    <i className="fas fa-share-alt mr-2"></i> Share Profile
                </button>
                <button
                    onClick={copyProfileLink}
                    className="bg-gray-700 px-4 py-3 w-full rounded-full text-white hover:bg-gray-600">
                    <i className="fas fa-link mr-2"></i> Copy My Profile link
                </button>
            </div>
        </SideModal>
    );
}

export default ShareProfileModal;