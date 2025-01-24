import React from 'react';

import Gallery from "@/components/Gallery"

const CategoryPage = ({title, match}) => {
  const { id } = match.params;

  const games = Array.from({ length: 130 }, (_, index) => ({
    id: index + 1,
    title: `Game ${index + 1}`,
    tag: index%5===1 ? "NEW" : (index%5===0 ? "HOT" : (index%5===2 ? "UPDATED" : "")) , // Alternate between NEW and HOT tags
    img: `https://dummyimage.com/200x150/cccccc/000000&text=Game ${index+1}`, // Placeholder image
    video: "https://www.w3schools.com/html/mov_bbb.mp4", // Placeholder video
    properties: [`Up to ${index%20+1} players`, ".io"],
  }));

  return (
    <div className="bg-gray-800 min-h-screen p-4">
      <h1 className="text-xl font-bold text-white mb-4">{title} Games</h1>
      <p>Showing details for category with ID: {id}</p>

      <Gallery games={games} noPagination={false} type="card" />
    </div>
  );
}

export default CategoryPage;