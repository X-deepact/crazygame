import React, { useState, useRef, useEffect } from 'react';

import Gallery from "@/components/Gallery"

const NewPage = () => {
  const games = Array.from({ length: 120 }, (_, index) => ({
    id: index + 1,
    title: `Game ${index + 1}`,
    tag: index%3===1 ? "NEW" : "" , // Alternate between NEW and HOT tags
    img: `https://dummyimage.com/200x150/cccccc/000000&text=Game ${index+1}`, // Placeholder image
    video: "https://www.w3schools.com/html/mov_bbb.mp4", // Placeholder video
  }));

  return (
    <div className="bg-gray-800 min-h-screen p-4">
      <h1 className="text-xl font-bold text-white mb-4">New Games</h1>

      <Gallery games={games} noPagination={false} type="hover" />
    </div>
  );
}

export default NewPage;