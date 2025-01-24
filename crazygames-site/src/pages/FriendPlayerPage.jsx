import React from 'react';

import Gallery from "@/components/Gallery"

const FriendPlayerPage = () => {
  const games = Array.from({ length: 50 }, (_, index) => ({
    id: index + 1,
    title: `Game ${index + 1}`,
    tag: index%5===1 ? "NEW" : (index%5===0 ? "HOT" : (index%5===2 ? "UPDATED" : "")) , // Alternate between NEW and HOT tags
    img: `https://dummyimage.com/200x150/cccccc/000000&text=Game ${index+1}`, // Placeholder image
    video: "https://www.w3schools.com/html/mov_bbb.mp4", // Placeholder video
  }));

  return (
    <div className="bg-gray-800 min-h-screen p-4">
      <h1 className="text-xl font-bold text-white mb-2">Games to Play with Friends</h1>
      <p className="text-sm text-gray-500 mb-5">You and your best pal looking for some game time? This page is full of free online games to play with friends, so check some of them out and have a blast!</p>

      <Gallery games={games} noPagination={false} type="card" />
    </div>
  );
}

export default FriendPlayerPage;