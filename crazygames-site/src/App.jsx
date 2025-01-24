import React, { useEffect, useState } from 'react';
import {
  BrowserRouter as Router,
  Route,
  Switch,
  Redirect,
  NavLink,
} from 'react-router-dom';

import SideBar from '@/pages/SideBar';
import Header from '@/pages/Header';
import HomePage from '@/pages/HomePage';
import RecentPage from '@/pages/RecentPage';
import NewPage from '@/pages/NewPage';
import TrendingPage from '@/pages/TrendingPage';
import UpdatedPage from '@/pages/UpdatedPage';
import OriginalsPage from '@/pages/OriginalsPage';
import MultiplayerPage from '@/pages/MultiplayerPage';
import CategoryPage from '@/pages/CategoryPage';
import GamePage from '@/pages/GamePage';
import ProfilePage from '@/pages/ProfilePage';
import LoginModal from '@/components/dialogs/LoginModal';
import SignupModal from '@/components/dialogs/SignupModal';
import AlertModal from '@/components/dialogs/AlertModal';
import ProfileModal from '@/components/dialogs/ProfileModal';
import FriendsModal from '@/components/dialogs/FriendsModal';
import NotificationsModal from '@/components/dialogs/NotificationsModal';
import MyGamesModal from '@/components/dialogs/MyGamesModal';
import EditProfileModal from '@/components/dialogs/EditProfileModal';
import ChangeUsernameModal from '@/components/dialogs/ChangeUsernameModal';
import ChangeBirthdayModal from '@/components/dialogs/ChangeBirthdayModal';
import ShareProfileModal from '@/components/dialogs/ShareProfileModal';
import FriendPlayerPage from '@/pages/FriendPlayerPage';
import NotificationPreferencesModal from '@/components/dialogs/NotificationPreferencesModal';
import PrivacyPreferencesModal from '@/components/dialogs/PrivacyPreferencesModal';
import PrivacyPage from '@/pages/PrivacyPage';
import AccountSettingsModal from '@/components/dialogs/AccountSettingsModal';
import ChangeAvatarModal from '@/components/dialogs/ChangeAvatarModal';
import ChangeCoverModal from '@/components/dialogs/ChangeCoverModal';
import { useLoginMutation, useRegisterMutation } from './hooks/api/auth';
import useAuthStore from './hooks/store/useAuth';
import { useLoginError } from './hooks/store/useLoginError';
import TermsAndConditions from "@/pages/TermsAndConditions";

export default function App() {
  const categories = [
    {
      id: 1,
      label: '2 Player',
      icon: 'https://imgs.crazygames.com/icon/2players.svg',
    },
    {
      id: 2,
      label: 'Action',
      icon: 'https://imgs.crazygames.com/icon/Action.svg',
    },
    {
      id: 3,
      label: 'Adventure',
      icon: 'https://imgs.crazygames.com/icon/Adventure.svg',
    },
    {
      id: 4,
      label: 'Basketball',
      icon: 'https://imgs.crazygames.com/icon/Basketball.svg',
    },
    {
      id: 5,
      label: 'Beauty',
      icon: 'https://imgs.crazygames.com/icon/Beauty.svg',
    },
    { id: 6, label: 'Bike', icon: 'https://imgs.crazygames.com/icon/Bike.svg' },
    { id: 7, label: 'Car', icon: 'https://imgs.crazygames.com/icon/Car.svg' },
    { id: 8, label: 'Card', icon: 'https://imgs.crazygames.com/icon/Card.svg' },
    {
      id: 9,
      label: 'Casual',
      icon: 'https://imgs.crazygames.com/icon/Casual.svg',
    },
    {
      id: 10,
      label: 'Clicker',
      icon: 'https://imgs.crazygames.com/icon/Clicker.svg',
    },
    {
      id: 11,
      label: 'Controller',
      icon: 'https://imgs.crazygames.com/icon/Controller.svg',
    },
    {
      id: 12,
      label: 'Dress Up',
      icon: 'https://imgs.crazygames.com/icon/DressUp.svg',
    },
    {
      id: 13,
      label: 'Driving',
      icon: 'https://imgs.crazygames.com/icon/Driving.svg',
    },
    {
      id: 14,
      label: 'Escape',
      icon: 'https://imgs.crazygames.com/icon/Escape.svg',
    },
    {
      id: 15,
      label: 'Flash',
      icon: 'https://imgs.crazygames.com/icon/Flash.svg',
    },
    { id: 16, label: 'FPS', icon: 'https://imgs.crazygames.com/icon/FPS.svg' },
    {
      id: 17,
      label: 'Horror',
      icon: 'https://imgs.crazygames.com/icon/Horror.svg',
    },
    { id: 18, label: '.io', icon: 'https://imgs.crazygames.com/icon/io.svg' },
    {
      id: 19,
      label: 'Mahjong',
      icon: 'https://imgs.crazygames.com/icon/Mahjong.svg',
    },
    {
      id: 20,
      label: 'Minecraft',
      icon: 'https://imgs.crazygames.com/icon/Minecraft.svg',
    },
    {
      id: 21,
      label: 'Pool',
      icon: 'https://imgs.crazygames.com/icon/Pool.svg',
    },
    {
      id: 22,
      label: 'Puzzle',
      icon: 'https://imgs.crazygames.com/icon/Puzzle.svg',
    },
    {
      id: 23,
      label: 'Shooting',
      icon: 'https://imgs.crazygames.com/icon/Shooting.svg',
    },
    {
      id: 24,
      label: 'Soccer',
      icon: 'https://imgs.crazygames.com/icon/Soccer.svg',
    },
    {
      id: 25,
      label: 'Sports',
      icon: 'https://imgs.crazygames.com/icon/Sports.svg',
    },
    {
      id: 26,
      label: 'Stickman',
      icon: 'https://imgs.crazygames.com/icon/Stickman.svg',
    },
    {
      id: 27,
      label: 'Tower Defense',
      icon: 'https://imgs.crazygames.com/icon/TowerDefense.svg',
    },
  ];
  const getCategoryTitle = (id) => {
    const record = categories.find((row) => row.id === id);
    return record ? record.label : '';
  };

  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [isMinimized, setIsMinimized] = useState(false);
  const toggleSidebar = () => {
    setIsMinimized(!isMinimized);
  };

  const [isAlertOpen, setIsAlertOpen] = useState(false);
  const [alertContent, setAlertContent] = useState({
    title: '',
    message: '',
    callback: null,
  });

  const openAlert = (title, message, callback) => {
    setAlertContent({ title, message, callback });
    setIsAlertOpen(true);
  };
  const closeAlert = () => {
    setIsAlertOpen(false);
    if (alertContent.callback != null) alertContent.callback();
  };

  const [sideModalFrom, setSideModalFrom] = useState('');

  const [isProfileModalOpen, setIsProfileModalOpen] = useState(false);
  const openProfileModal = () => {
    setIsProfileModalOpen(true);
    closeFriendsModal();
    closeNotificationsModal();
    closeMyGamesModal();
    closeLoginModal();
    closeSignupModal();
    closeEditProfileModal();
    closeEditUsernameModal();
    closeEditBirthdayModal();
    closeShareProfileModal();
    closeNotificationPreferencesModal();
    closePrivacyPreferencesModal();
    closeAccountSettingsModal();
    closeChangeAvatarModal();
    closeChangeCoverModal();
  };
  const closeProfileModal = () => setIsProfileModalOpen(false);

  const [isLoginModalOpen, setIsLoginModalOpen] = useState(false);
  const openLoginModal = ({ email }) => {
    setEmail(email || '');
    setIsLoginModalOpen(true);
    closeSignupModal();
    closeFriendsModal();
    closeNotificationsModal();
    closeMyGamesModal();
    closeProfileModal();
    closeEditProfileModal();
    closeEditUsernameModal();
    closeEditBirthdayModal();
    closeShareProfileModal();
    closeNotificationPreferencesModal();
    closePrivacyPreferencesModal();
    closeAccountSettingsModal();
    closeChangeAvatarModal();
    closeChangeCoverModal();
  };
  const closeLoginModal = () => setIsLoginModalOpen(false);

  const [isSignupModalOpen, setIsSignupModalOpen] = useState(false);
  const openSignupModal = () => {
    setIsSignupModalOpen(true);
    closeFriendsModal();
    closeNotificationsModal();
    closeMyGamesModal();
    closeLoginModal();
    closeProfileModal();
    closeEditProfileModal();
    closeEditUsernameModal();
    closeEditBirthdayModal();
    closeShareProfileModal();
    closeNotificationPreferencesModal();
    closePrivacyPreferencesModal();
    closeAccountSettingsModal();
    closeChangeAvatarModal();
    closeChangeCoverModal();
  };
  const closeSignupModal = () => setIsSignupModalOpen(false);

  const [isFriendsModalOpen, setIsFriendsModalOpen] = useState(false);
  const openFriendsModal = () => {
    setIsFriendsModalOpen(true);
    closeLoginModal();
    closeSignupModal();
    closeNotificationsModal();
    closeMyGamesModal();
    closeProfileModal();
    closeEditProfileModal();
    closeEditUsernameModal();
    closeEditBirthdayModal();
    closeShareProfileModal();
    closeNotificationPreferencesModal();
    closePrivacyPreferencesModal();
    closeAccountSettingsModal();
    closeChangeAvatarModal();
    closeChangeCoverModal();
  };
  const closeFriendsModal = () => setIsFriendsModalOpen(false);

  const [isNotificationsModalOpen, setIsNotificationsModalOpen] =
    useState(false);
  const openNotificationsModal = () => {
    setIsNotificationsModalOpen(true);
    closeLoginModal();
    closeSignupModal();
    closeFriendsModal();
    closeMyGamesModal();
    closeProfileModal();
    closeEditProfileModal();
    closeEditUsernameModal();
    closeEditBirthdayModal();
    closeShareProfileModal();
    closeNotificationPreferencesModal();
    closePrivacyPreferencesModal();
    closeAccountSettingsModal();
    closeChangeAvatarModal();
    closeChangeCoverModal();
  };
  const closeNotificationsModal = () => setIsNotificationsModalOpen(false);

  const [isMyGamesModalOpen, setIsMyGamesModalOpen] = useState(false);
  const openMyGamesModal = () => {
    setIsMyGamesModalOpen(true);
    closeLoginModal();
    closeSignupModal();
    closeFriendsModal();
    closeNotificationsModal();
    closeProfileModal();
    closeEditProfileModal();
    closeEditUsernameModal();
    closeEditBirthdayModal();
    closeShareProfileModal();
    closeNotificationPreferencesModal();
    closePrivacyPreferencesModal();
    closeAccountSettingsModal();
    closeChangeAvatarModal();
    closeChangeCoverModal();
  };
  const closeMyGamesModal = () => setIsMyGamesModalOpen(false);

  const [isEditProfileModalOpen, setIsEditProfileModalOpen] = useState(false);
  const openEditProfileModal = () => {
    setIsEditProfileModalOpen(true);

    closeFriendsModal();
    closeNotificationsModal();
    closeMyGamesModal();
    closeLoginModal();
    closeSignupModal();
    closeProfileModal();
    closeEditUsernameModal();
    closeEditBirthdayModal();
    closeShareProfileModal();
    closeNotificationPreferencesModal();
    closePrivacyPreferencesModal();
    closeAccountSettingsModal();
    closeChangeAvatarModal();
    closeChangeCoverModal();
  };
  const closeEditProfileModal = () => setIsEditProfileModalOpen(false);

  const [isEditUsernameModalOpen, setIsEditUsernameModalOpen] = useState(false);
  const openEditUsernameModal = () => {
    setIsEditUsernameModalOpen(true);

    closeFriendsModal();
    closeNotificationsModal();
    closeMyGamesModal();
    closeLoginModal();
    closeSignupModal();
    closeProfileModal();
    closeEditProfileModal();
    closeEditBirthdayModal();
    closeShareProfileModal();
    closeNotificationPreferencesModal();
    closePrivacyPreferencesModal();
    closeAccountSettingsModal();
    closeChangeAvatarModal();
    closeChangeCoverModal();
  };
  const closeEditUsernameModal = () => setIsEditUsernameModalOpen(false);

  const [isEditBirthdayModalOpen, setIsEditBirthdayModalOpen] = useState(false);
  const openEditBirthdayModal = () => {
    setIsEditBirthdayModalOpen(true);

    closeFriendsModal();
    closeNotificationsModal();
    closeMyGamesModal();
    closeLoginModal();
    closeSignupModal();
    closeProfileModal();
    closeEditProfileModal();
    closeEditUsernameModal();
    closeShareProfileModal();
    closeNotificationPreferencesModal();
    closePrivacyPreferencesModal();
    closeAccountSettingsModal();
    closeChangeAvatarModal();
    closeChangeCoverModal();
  };
  const closeEditBirthdayModal = () => setIsEditBirthdayModalOpen(false);

  const [isShareProfileModalOpen, setIsShareProfileModalOpen] = useState(false);
  const openShareProfileModal = (from) => {
    setSideModalFrom(from !== undefined ? from : '');
    setIsShareProfileModalOpen(true);

    closeFriendsModal();
    closeNotificationsModal();
    closeMyGamesModal();
    closeLoginModal();
    closeSignupModal();
    closeProfileModal();
    closeEditProfileModal();
    closeEditUsernameModal();
    closeEditBirthdayModal();
    closeNotificationPreferencesModal();
    closePrivacyPreferencesModal();
    closeAccountSettingsModal();
    closeChangeAvatarModal();
    closeChangeCoverModal();
  };
  const closeShareProfileModal = () => setIsShareProfileModalOpen(false);

  const backShareProfileModal = () => {
    if (sideModalFrom === 'profile') openProfileModal();
    else if (sideModalFrom === 'friends') openFriendsModal();
  };

  const [
    isNotificationPreferencesModalOpen,
    setIsNotificationPreferencesModalOpen,
  ] = useState(false);
  const openNotificationPreferencesModal = (from) => {
    setSideModalFrom(from !== undefined ? from : '');
    setIsNotificationPreferencesModalOpen(true);

    closeFriendsModal();
    closeNotificationsModal();
    closeMyGamesModal();
    closeLoginModal();
    closeSignupModal();
    closeProfileModal();
    closeEditProfileModal();
    closeEditUsernameModal();
    closeEditBirthdayModal();
    closeShareProfileModal();
    closePrivacyPreferencesModal();
    closeAccountSettingsModal();
    closeChangeAvatarModal();
    closeChangeCoverModal();
  };
  const closeNotificationPreferencesModal = () =>
    setIsNotificationPreferencesModalOpen(false);

  const backNotificationPreferencesModal = () => {
    if (sideModalFrom === 'notifications') openNotificationsModal();
    else if (sideModalFrom === 'profile') openProfileModal();
  };

  const [isPrivacyPreferencesModalOpen, setIsPrivacyPreferencesModalOpen] =
    useState(false);
  const openPrivacyPreferencesModal = (from) => {
    setIsPrivacyPreferencesModalOpen(true);

    closeFriendsModal();
    closeNotificationsModal();
    closeMyGamesModal();
    closeLoginModal();
    closeSignupModal();
    closeProfileModal();
    closeEditProfileModal();
    closeEditUsernameModal();
    closeEditBirthdayModal();
    closeShareProfileModal();
    closeNotificationPreferencesModal();
    closeAccountSettingsModal();
    closeChangeAvatarModal();
    closeChangeCoverModal();
  };
  const closePrivacyPreferencesModal = () =>
    setIsPrivacyPreferencesModalOpen(false);

  const [isAccountSettingsModalOpen, setIsAccountSettingsModalOpen] =
    useState(false);
  const openAccountSettingsModal = (from) => {
    setIsAccountSettingsModalOpen(true);

    closeFriendsModal();
    closeNotificationsModal();
    closeMyGamesModal();
    closeLoginModal();
    closeSignupModal();
    closeProfileModal();
    closeEditProfileModal();
    closeEditUsernameModal();
    closeEditBirthdayModal();
    closeShareProfileModal();
    closeNotificationPreferencesModal();
    closePrivacyPreferencesModal();
    closeChangeAvatarModal();
    closeChangeCoverModal();
  };
  const closeAccountSettingsModal = () => setIsAccountSettingsModalOpen(false);

  const [isChangeAvatarModalOpen, setIsChangeAvatarModalOpen] = useState(false);
  const openChangeAvatarModal = (from) => {
    setIsChangeAvatarModalOpen(true);

    closeFriendsModal();
    closeNotificationsModal();
    closeMyGamesModal();
    closeLoginModal();
    closeSignupModal();
    closeProfileModal();
    closeEditProfileModal();
    closeEditUsernameModal();
    closeEditBirthdayModal();
    closeShareProfileModal();
    closeNotificationPreferencesModal();
    closePrivacyPreferencesModal();
    closeAccountSettingsModal();
    closeChangeCoverModal();
  };
  const closeChangeAvatarModal = () => setIsChangeAvatarModalOpen(false);

  const [isChangeCoverModalOpen, setIsChangeCoverModalOpen] = useState(false);
  const openChangeCoverModal = (from) => {
    setIsChangeCoverModalOpen(true);

    closeFriendsModal();
    closeNotificationsModal();
    closeMyGamesModal();
    closeLoginModal();
    closeSignupModal();
    closeProfileModal();
    closeEditProfileModal();
    closeEditUsernameModal();
    closeEditBirthdayModal();
    closeShareProfileModal();
    closeNotificationPreferencesModal();
    closePrivacyPreferencesModal();
    closeAccountSettingsModal();
    closeChangeAvatarModal();
  };
  const closeChangeCoverModal = () => setIsChangeCoverModalOpen(false);

  // share email for Log in, Sign up
  const [email, setEmail] = useState('');
  const [profile, setProfile] = useState({
    email: '',
    username: '',
    birthday: '',
    gender: '',
    country: '',
    avatarImage: '',
    coverImage: '',
  });

  const onSignup = (address) => {
    setEmail(address);
    openSignupModal();
  };

  const registerMutation = useRegisterMutation();
  const onRegister = (address, password) => {
    // TODO - call api to register email, password
    registerMutation.mutate(
      { email: address, username: address, password },
      {
        onSuccess: () => {
          openAlert(
            'Sign Up',
            'Your account was successfully registered. You can login now!',
            () => {
              openLoginModal({ email: address });
            }
          );
        },
        onError: (error) => {
          openAlert('Sign Up', 'Invalid Credentials', null);
        },
      }
    );
  };

  const loginMutation = useLoginMutation();
  const { token, setToken, user, setUser, removeAuth } = useAuthStore();
  const { setErrors } = useLoginError();
  const onLogin = ({ email, password }) => {
    loginMutation.mutate(
      { email, password },
      {
        onError: () => {
          console.log('errors');
          setErrors('login', 'User disabled');
        },
        onSuccess: (data) => {
          setToken(data?.token);
          closeLoginModal();

          setIsAuthenticated(true);

          setEmail(email);
          setProfile((prevState) => ({
            ...prevState,
            ...user,
            birthday: '',
            gender: 'M',
            country: 'United States',
            coverImage: 'https://imgs.crazygames.com/userportal/covers/War.jpg',
            avatarImage:
              'https://imgs.crazygames.com/userportal/avatars/78.png',
          }));
        },
      }
    );
  };

  const onLogout = () => {
    closeProfileModal();
    setIsAuthenticated(false);
    removeAuth();
    window.location.href = '/home';
  };

  const handleProfileSave = (records) => {
    records.forEach((row) => {
      setProfile((prevState) => ({
        ...prevState,
        ...row,
      }));
    });

    console.log('-------profile save', records, profile);
  };

  // Handle Oauth
  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const token = params.get('token');
    const oauthEmail = params.get('uid');
    const oauthName = params.get('uname');

    if (!token || !oauthEmail) return;
    setToken(token);
    setIsAuthenticated(true);
    setEmail(oauthEmail);
    setUser({ email: oauthEmail, username: oauthName });
  }, []);

  // Handle initialize
  useEffect(() => {
    if (token) {
      if (user.email) {
        setIsAuthenticated(true);
        setEmail(user.email);

        // TODO - get user details by email
        setProfile((prevState) => ({
          ...prevState,
          ...user,
          birthday: '2000-01-22',
          gender: 'M',
          country: 'United States',
          coverImage: 'https://imgs.crazygames.com/userportal/covers/War.jpg',
          avatarImage: 'https://imgs.crazygames.com/userportal/avatars/78.png',
        }));
      }
    }
  }, [token]);

  // Handle responsive sidebar
  useEffect(() => {
    const handleResize = () => {
      if (window.innerWidth < 1024) {
        setIsMinimized(true);
      } else {
        setIsMinimized(false);
      }
    };

    // Attach the event listener
    window.addEventListener('resize', handleResize);

    // Trigger once on component mount
    handleResize();

    // Cleanup the event listener
    return () => {
      window.removeEventListener('resize', handleResize);
    };
  }, []);

  return (
    <Router>
      <div className='flex flex-col overflow-hidden'>
        {/* Header */}
        <Header
          isAuthenticated={isAuthenticated}
          toggleSidebar={toggleSidebar}
          profile={profile}
          openLoginModal={openLoginModal}
          openFriendsModal={openFriendsModal}
          openNotificationsModal={openNotificationsModal}
          openMyGamesModal={openMyGamesModal}
          openProfileModal={openProfileModal}
        />

        <LoginModal
          isOpen={isLoginModalOpen}
          onClose={closeLoginModal}
          emailAddress={email}
          onSignup={onSignup}
          alert={openAlert}
          onLogin={onLogin}
        />
        <SignupModal
          isOpen={isSignupModalOpen}
          onClose={closeSignupModal}
          emailAddress={email}
          onLogin={openLoginModal}
          onRegister={onRegister}
        />

        <ProfileModal
          isOpen={isProfileModalOpen}
          onClose={closeProfileModal}
          profile={profile}
          onShare={openShareProfileModal}
          onNotificationPreferences={openNotificationPreferencesModal}
          onPrivacyPreferences={openPrivacyPreferencesModal}
          onAccountSettings={openAccountSettingsModal}
          onLogout={onLogout}
        />
        <EditProfileModal
          isOpen={isEditProfileModalOpen}
          onClose={closeEditProfileModal}
          onBack={openProfileModal}
          profile={profile}
          onEditUsername={openEditUsernameModal}
          onEditBirthday={openEditBirthdayModal}
          onChangeAvatar={openChangeAvatarModal}
          onChangeCover={openChangeCoverModal}
          onSave={handleProfileSave}
        />
        <ChangeUsernameModal
          isOpen={isEditUsernameModalOpen}
          onClose={closeEditUsernameModal}
          profile={profile}
          onSave={handleProfileSave}
          onBack={openEditProfileModal}
        />
        <ChangeBirthdayModal
          isOpen={isEditBirthdayModalOpen}
          onClose={closeEditBirthdayModal}
          profile={profile}
          onSave={handleProfileSave}
          onBack={openEditProfileModal}
        />
        <ChangeAvatarModal
          isOpen={isChangeAvatarModalOpen}
          onClose={closeChangeAvatarModal}
          profile={profile}
          onChange={handleProfileSave}
          onBack={openEditProfileModal}
        />
        <ChangeCoverModal
          isOpen={isChangeCoverModalOpen}
          onClose={closeChangeCoverModal}
          profile={profile}
          onChange={handleProfileSave}
          onBack={openEditProfileModal}
        />

        <FriendsModal
          isAuthenticated={isAuthenticated}
          isOpen={isFriendsModalOpen}
          onClose={closeFriendsModal}
          onLogin={openLoginModal}
          onShare={openShareProfileModal}
        />
        <ShareProfileModal
          isOpen={isShareProfileModalOpen}
          onClose={closeShareProfileModal}
          profile={profile}
          alert={openAlert}
          onBack={backShareProfileModal}
        />
        <NotificationsModal
          isOpen={isNotificationsModalOpen}
          onClose={closeNotificationsModal}
          onSettings={openNotificationPreferencesModal}
        />
        <NotificationPreferencesModal
          isOpen={isNotificationPreferencesModalOpen}
          onClose={closeNotificationPreferencesModal}
          onBack={backNotificationPreferencesModal}
        />
        <PrivacyPreferencesModal
          isOpen={isPrivacyPreferencesModalOpen}
          onClose={closePrivacyPreferencesModal}
          onBack={openProfileModal}
        />
        <AccountSettingsModal
          isOpen={isAccountSettingsModalOpen}
          onClose={closeAccountSettingsModal}
          onBack={openProfileModal}
        />
        <MyGamesModal isOpen={isMyGamesModalOpen} onClose={closeMyGamesModal} />

        <AlertModal
          isOpen={isAlertOpen}
          title={alertContent.title}
          message={alertContent.message}
          onClose={closeAlert}
        />

        {/* Main Content Area */}
        <div className='flex flex-grow'>
          {/* Sidebar */}
          <SideBar isMinimized={isMinimized} categories={categories} />

          {/* Content */}
          <div
            className={`flex-grow bg-gray-800 mt-[4.5rem] min-h-screen ${
              isMinimized ? 'ml-[85px]' : 'ml-[235px]'
            } transition-all duration-300`}
          >
            <Switch>
              <Route exact path='/' render={() => <Redirect to='/home' />} />
              <Route
                path='/home'
                render={(props) => (
                  <HomePage {...props} isMinimized={isMinimized} />
                )}
              />
              <Route path='/new' component={NewPage} />
              <Route path='/recent' component={RecentPage} />
              <Route path='/trending' component={TrendingPage} />
              <Route path='/updated' component={UpdatedPage} />
              <Route path='/originals' component={OriginalsPage} />
              <Route path='/multiplayer' component={MultiplayerPage} />
              <Route path='/with-friend' component={FriendPlayerPage} />
              <Route
                path='/categories/:id'
                render={(props) => (
                  <CategoryPage
                    {...props}
                    title={getCategoryTitle(props.match.params.id)}
                  />
                )}
              />
              <Route
                path='/games/:id'
                render={(props) => (
                  <GamePage key={props.match.params.id} {...props} />
                )}
              />

              {isAuthenticated && (
                <Route
                  path='/profile'
                  render={(props) => (
                    <ProfilePage
                      {...props}
                      profile={profile}
                      onEditProfile={openEditProfileModal}
                      onChangeAvatar={openChangeAvatarModal}
                      onChangeCover={openChangeCoverModal}
                    />
                  )}
                />
              )}

              <Route path='/privacy-policy' component={PrivacyPage} />
              <Route path='/terms-and-conditions' component={TermsAndConditions} />
            </Switch>
          </div>
        </div>
      </div>
    </Router>
  );
}
