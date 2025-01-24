import React from 'react';

import Gallery from "@/components/Gallery"

const RecentPage = () => {
  const games = Array.from({ length: 30 }, (_, index) => ({
    id: index + 1,
    title: `Game ${index + 1}`,
    tag: index%3===1 ? "NEW" : (index%3===0 ? "HOT" : "") , // Alternate between NEW and HOT tags
    img: `https://dummyimage.com/200x150/cccccc/000000&text=Game ${index+1}`, // Placeholder image
    video: "https://www.w3schools.com/html/mov_bbb.mp4", // Placeholder video
  }));

  return (
    <div className="bg-gray-800 min-h-screen p-4">
      <h1 className="text-xl font-bold text-white mb-4">Recently Playered Games</h1>

      <Gallery games={games} noPagination={true} type="hover" />
    </div>
  );
}

export default RecentPage;