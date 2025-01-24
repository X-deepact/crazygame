import React from 'react';

import Gallery from "@/components/Gallery"

const ProfilePage = ({profile, onEditProfile, onChangeAvatar, onChangeCover}) => {
  const games = Array.from({ length: 12 }, (_, index) => ({
    id: index + 1,
    title: `Game ${index + 1}`,
    tag: index%3===1 ? "NEW" : (index%3===0 ? "HOT" : "") , // Alternate between NEW and HOT tags
    img: `https://dummyimage.com/200x150/cccccc/000000&text=Game ${index+1}`, // Placeholder image
    video: "https://www.w3schools.com/html/mov_bbb.mp4", // Placeholder video
  }));

  return (
    <div className="bg-gray-800 text-white min-h-screen">
      {/* Main Content Section */}
      <div className="flex justify-center mt-6 px-6 mb-10">
        {/* Container with max-width of 1200px */}
        <div className="w-full max-w-[1200px] flex flex-col space-y-6">

          {/* Top Row: Profile Info */}
          <div className="w-full relative mb-[75px]">
            {/* Profile Image */}
            <div className="relative group">
              <img
                src={profile.coverImage} // Placeholder for background image
                alt="background"
                className="w-full h-[320px] object-cover rounded-b-xl group-hover:opacity-70 transition duration-300"
              />
              <button
                onClick={onChangeCover}
                className="absolute bottom-2.5 right-2.5 bg-white text-gray-800 px-4 py-2 flex items-center rounded-full shadow-lg opacity-0 group-hover:opacity-100 transition duration-300"
              >
                <i className="fas fa-camera mr-1"></i> Edit
              </button>
            </div>

            <div className="absolute bottom-[-75px] left-4 group">
              {/* Profile Image (overlapping) */}
              <div className="relative w-full h-full rounded-xl overflow-hidden border-4 border-black">
                <img
                  src={profile.avatarImage} // Placeholder for profile image
                  alt="profile"
                  className="object-cover w-[150px] h-[150px] rounded-xl border-4 border-black group-hover:opacity-70 transition duration-300"
                />
                <div className="absolute inset-0 bg-black opacity-0 group-hover:opacity-30 transition-opacity duration-300"></div>
              </div>

              <button
                onClick={onChangeAvatar}
                className="absolute bottom-2.5 right-2.5 w-8 h-8 bg-white px-[0.55rem] flex items-center rounded-full shadow-lg opacity-0 group-hover:opacity-100 transition duration-300"
              >
                <i className="fas fa-camera text-gray-800"></i>
              </button>
            </div>


            <div className="absolute left-[180px] bottom-[-64px] hidden sm:block">
              <h1 className="text-xl font-semibold">{profile.username}</h1>
              <p className="text-gray-400">{profile.country}</p>
            </div>

            <div className="absolute right-4 bottom-[-70px] flex items-center">
                <button
                  onClick={onEditProfile}
                  className="bg-gray-700 px-4 py-4 rounded-full text-md hover:bg-gray-600 transition flex items-center">
                  <i className="fas fa-edit mr-2"></i> Edit my profile
                </button>
              </div>
          </div>

          {/* Middle Row: Stats, Friends, and Liked Games */}
          <div className="flex flex-col lg:flex-row space-y-6 lg:space-y-0 lg:space-x-6">

            {/* Left Section (1/3): Friends and Stats */}
            <div className="w-full lg:w-1/3 space-y-6">
              {/* Friends Card */}
              <div className="w-full bg-gray-700 rounded-lg shadow-lg p-6">
                <h2 className="text-2xl font-semibold">Friends <span className="ml-2 text-xl">0</span></h2>
              </div>

              {/* Stats Card */}
              <div className="w-full bg-gray-700 rounded-lg shadow-lg p-6">
                <h2 className="text-2xl font-semibold mb-6">Stats</h2>
                <div className="space-y-4 px-2">
                  <div className="flex items-center space-x-4">
                    <i className="fas fa-cogs text-2xl text-gray-400"></i>
                    <div>
                      <p className="text-sm text-gray-400">Games played</p>
                      <p className="text-xl font-semibold">5</p>
                    </div>
                  </div>

                  <div className="flex items-center space-x-4">
                    <i className="fas fa-calendar-day text-2xl text-gray-400"></i>
                    <div>
                      <p className="text-sm text-gray-400">Member for</p>
                      <p className="text-xl font-semibold">24 days</p>
                    </div>
                  </div>

                  <div className="flex items-center space-x-4">
                    <i className="fas fa-thumbs-up text-2xl text-gray-400"></i>
                    <div>
                      <p className="text-sm text-gray-400">Games liked</p>
                      <p className="text-xl font-semibold">1</p>
                    </div>
                  </div>

                  <div className="flex items-center space-x-4">
                    <i className="fas fa-fire text-2xl text-gray-400"></i>
                    <div>
                      <p className="text-sm text-gray-400">Playstreak</p>
                      <p className="text-xl font-semibold flex items-center mb-2">
                        2 days
                        <span className="text-purple-500 ml-2">
                          <i className="fas fa-question-circle"></i>
                        </span>
                      </p>
                      <p className="text-sm text-gray-400">Highest streak: 3 days</p>
                    </div>
                  </div>
                </div>
              </div>
            </div>

            {/* Right Section (2/3): My Liked Games */}
            <div className="w-full lg:w-2/3 bg-gray-700 rounded-lg shadow-lg p-6">
              <h2 className="text-2xl font-semibold mb-4">My liked games</h2>
              <div className="w-full">
                <Gallery games={games} noPagination={true} type="hover" />
              </div>
            </div>

          </div>
        </div>
      </div>
    </div>
  );
};

export default ProfilePage;
