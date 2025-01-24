import React, { useState } from 'react';
import { useHistory } from "react-router-dom";
import {gridAutoColumns} from "tailwindcss/lib/plugins";

const GalleryItem = ({ game, type, columnWidth, columnHeight }) => {
  const [hoveredGameId, setHoveredGameId] = useState(null); // Track which game is hovered

  let tagClass = "bg-red-500";
  if (game.tag === "HOT")
    tagClass = "bg-yellow-500";
  else if (game.tag === "UPDATED")
    tagClass = "bg-green-500";

  const properties = game.properties ? game.properties.join(" • ") : "";

  const history = useHistory();
  const handleItemClick = (id) => {
    history.push(`/games/${id}`); // Navigate to the game page with the ID
    window.location.reload();
  };

  return (
        <div
              key={game.id}
              className={`relative rounded-lg overflow-hidden shadow-md group cursor-pointer border border-transparent hover:border-purple-600 transition-all duration-300
              ${ type==="card" ? "bg-gray-900" : "" }
              `}
              style={{ width: `${columnWidth}px` }}
              onClick={() => handleItemClick(game.id)}
              onMouseEnter={() => setHoveredGameId(game.id)} // Trigger video load on hover
              onMouseLeave={() => setHoveredGameId(null)}    // Remove hover state on mouse leave
            >
              {/* Game Image */}
              { type === 'card' ? (
                  <>
                    <img
                      src={game.img}
                      alt={game.title}
                      style={{height: `${columnHeight}px`}}
                      className="w-full object-cover"
                    />
                    <div className="mt-1 mb-2 px-2">
                      <h3 className="text-lg font-bold text-white">{game.title}</h3>
                      <p className="text-sm text-gray-400">
                        {properties}
                      </p>
                    </div>
                  </>
                ) : (
                  <>
                    <img
                      src={game.img}
                      alt={game.title}
                      style={{height: `${columnHeight}px`}}
                      className="w-full object-cover"
                    />
                  </>
                )}

              {/* Tag */}
              { game.tag!=="" && (
                <span className={`absolute top-0 left-0 ${tagClass} text-white text-xs font-bold px-2 py-1 rounded`}>
                  {game.tag}
                </span>
              )}

              {/* Video (Hidden by default, shown on hover) */}
              {hoveredGameId === game.id && (
                <video
                  src={game.video}
                  style={{height: `${columnHeight}px`}}
                  className={`absolute top-0 left-0 w-full object-cover opacity-0 group-hover:opacity-100 transition-opacity duration-300`}
                  autoPlay
                  loop
                  muted
                  playsInline
                />
              )}

{/*              { game.tag!="" && (
                <span className={`absolute top-0 left-0 ${tagClass} text-white bg-opacity-75 text-xs font-bold px-2 py-1 rounded`}>
                  {game.tag}
                </span>
              )}*/}

              {/* Game Title (Hidden by default, visible on hover) */}
              { type !== "card" &&  (
                <div className="absolute bottom-0 left-0 right-0 bg-gray-900 bg-opacity-75 text-white text-center py-2 opacity-0 group-hover:opacity-100 transition-opacity duration-300">
                  {game.title}
                </div>
              )}

            </div>
    );
}

export default GalleryItem;