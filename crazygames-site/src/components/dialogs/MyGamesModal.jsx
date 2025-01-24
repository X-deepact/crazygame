import React, {useState} from 'react';

import SideModal from '@/components/dialogs/SideModal'
import Gallery from "@/components/Gallery"

const shuffleArray = (array) => {
  // Create a copy of the array to avoid mutating the original
  const shuffled = [...array];

  for (let i = shuffled.length - 1; i > 0; i--) {
    // Pick a random index from 0 to i
    const randomIndex = Math.floor(Math.random() * (i + 1));

    // Swap elements at i and randomIndex
    [shuffled[i], shuffled[randomIndex]] = [shuffled[randomIndex], shuffled[i]];
  }

  return shuffled;
}

const MyGamesModal = ({isOpen, onClose}) => {
  const [activeTab, setActiveTab] = useState("recent");

  const tabs = [
    { id: "recent", label: "Recent" },
    { id: "favorites", label: "Favorites" },
    { id: "liked", label: "Liked" },
  ];

    const games = Array.from({ length: 16 }, (_, index) => ({
      id: index + 1,
      title: `Game ${index + 1}`,
      tag: index%3===1 ? "NEW" : (index%3===0 ? "HOT" : "") , // Alternate between NEW and HOT tags
      img: `https://dummyimage.com/160x120/cccccc/000000&text=Game ${index+1}`, // Placeholder image
      video: "https://www.w3schools.com/html/mov_bbb.mp4", // Placeholder video
    }));

    return (
      <SideModal title="My Games" isOpen={isOpen} onClose={onClose}>
        <div className="border-b border-gray-700">
          <div className="flex justify-around">
            {tabs.map((tab) => (
              <button
                key={tab.id}
                onClick={() => setActiveTab(tab.id)}
                className={`relative py-2 px-4 text-sm font-medium ${
                  activeTab === tab.id
                    ? "text-white"
                    : "text-gray-400 hover:text-white"
                }`}
              >
                {tab.label}
                {activeTab === tab.id && (
                  <span className="absolute bottom-0 left-0 w-full h-1 bg-purple-500 rounded"></span>
                )}
              </button>
            ))}
          </div>
        </div>

        {/* Tab Content */}
        <div className="mt-4 text-gray-300 bg-gray-800 rounded-lg h-[calc(100vh-16rem)] overflow-y-auto">
          { activeTab === 'recent' && (
            <div>
              <Gallery games={shuffleArray(games)} noPagination={true} type="hover" width="160" />
            </div>
          )}

          { activeTab === 'favorites' && (
            <div>
              <Gallery games={shuffleArray(games)} noPagination={true} type="hover" width="160" />
            </div>
          )}

          { activeTab === 'liked' && (
            <div>
              <Gallery games={shuffleArray(games)} noPagination={true} type="hover" width="160" />
            </div>
          )}
        </div>
      </SideModal>
      );
  }

export default MyGamesModal;