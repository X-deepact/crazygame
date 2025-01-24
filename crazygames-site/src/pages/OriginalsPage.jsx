import React from 'react';

import Gallery from "@/components/Gallery"

const OriginalsPage = () => {
  const games = Array.from({ length: 30 }, (_, index) => ({
    id: index + 1,
    title: `Game ${index + 1}`,
    tag: index%5===1 ? "NEW" : (index%5===0 ? "HOT" : (index%5===2 ? "UPDATED" : "")) , // Alternate between NEW and HOT tags
    img: `https://dummyimage.com/200x350/cccccc/000000&text=Game ${index+1}`, // Placeholder image
    video: "https://www.w3schools.com/html/mov_bbb.mp4", // Placeholder video
    properties: [`Up to ${index%20+1} players`, ".io"],
  }));

  return (
    <div className="bg-gray-800 min-h-screen p-4">
      <h1 className="text-xl font-bold text-white mb-4">Original Games</h1>

      <Gallery games={games} noPagination={false} type="card" width="200" height="350" />
    </div>
  );
}

export default OriginalsPage;