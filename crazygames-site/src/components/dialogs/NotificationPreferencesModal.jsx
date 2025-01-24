import React, {useState} from 'react';

import ToggleCheckbox from "@/components/ToggleCheckbox";
import SideModal from '@/components/dialogs/SideModal'

const NotificationPreferencesModal = ({isOpen, onClose, onBack}) => {
  // States for each toggle switch
  const [gameUpdates, setGameUpdates] = useState(true);
  const [friendRequests, setFriendRequests] = useState(true);
  const [events, setEvents] = useState(true);
  const [achievements, setAchievements] = useState(true);
  const [platformUpdates, setPlatformUpdates] = useState(true);

  const [interactiveNotifications, setInteractiveNotifications] = useState(true);
  const [informationalNotifications, setInformationalNotifications] = useState(true);

  return (
    <SideModal title="Notification Preferences" isOpen={isOpen} onClose={onClose} onBack={onBack}>
      <section className="mb-6">
          <h3 className="text-xl text-white mb-2">Emails</h3>
          <p className="text-sm text-gray-300 mb-4">
            Receive emails from CrazyGames on your mailbox. We’ll promise not to spam you!
          </p>
          <div className="space-y-4 rounded-lg bg-gray-900 p-4">
            <ToggleCheckbox label="Game updates & recommendations" checked={gameUpdates} onChange={() => setGameUpdates(!gameUpdates)} />
            <ToggleCheckbox label="Friend requests" checked={friendRequests} onChange={() => setFriendRequests(!friendRequests)} />
            <ToggleCheckbox label="Events & competitions" checked={events} onChange={() => setEvents(!events)} />
            <ToggleCheckbox label="Achievements & leaderboards" checked={achievements} onChange={() => setAchievements(!achievements)} />
            <ToggleCheckbox label="Platform updates" checked={platformUpdates} onChange={() => setPlatformUpdates(!platformUpdates)} />
          </div>
        </section>

        {/* Site Notifications Section */}
        <section>
          <h3 className="text-xl text-white mb-2">Site notifications</h3>
          <p className="text-sm text-gray-300 mb-4">
            Receive site notifications directly on your device. Stay updated on your friends' activities, game invites, and much more!
          </p>
          <div className="space-y-4 rounded-lg bg-gray-900 p-4">
            <ToggleCheckbox label="Interactive (friend requests & invites, CrazyGames events, ...)" checked={interactiveNotifications} onChange={() => setInteractiveNotifications(!interactiveNotifications)} />
            <ToggleCheckbox label="Informational (friends activity and more)" checked={informationalNotifications} onChange={() => setInformationalNotifications(!informationalNotifications)} />
          </div>
        </section>
    </SideModal>
    );
  }

export default NotificationPreferencesModal;