import React from 'react';

import SideModal from '@/components/dialogs/SideModal'

const NotificationsModal = ({isOpen, onClose, onSettings}) => {
    return (
      <SideModal title="Notifications" isOpen={isOpen} onClose={onClose} onSettings={() => onSettings('notifications')} >
        <div className="flex flex-col items-center text-center">
            No notifications
        </div>
      </SideModal>
      );
  }

export default NotificationsModal;