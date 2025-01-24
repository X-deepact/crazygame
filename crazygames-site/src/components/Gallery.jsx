import React, { useState, useRef, useEffect } from 'react';

import GalleryItem from "./GalleryItem"
import Pagination from "./Pagination"

const Gallery = ({ games, noPagination, type, width, height }) => {
    const [currentPage, setCurrentPage] = useState(1); // Track the current page
    const [itemsPerPage, setItemsPerPage] = useState(21); // Default items per page
    const galleryRef = useRef(null); // Reference to the gallery container
    const [columnWidth, setColumnWidth] = useState(width || 200); // Reference to the gallery container
    const [columnHeight, setColumnHeight] = useState(height || 150); // Reference to the gallery container

    // Adjust items per page dynamically based on container width
    useEffect(() => {
      const handleResize = () => {
        if (galleryRef.current) {
          const containerWidth = galleryRef.current.offsetWidth; // Get container width
          const itemWidth = 216; // Each item's width (200px + gap)
          const itemsPerRow = Math.floor(containerWidth / itemWidth); // Calculate items per row
          setItemsPerPage(itemsPerRow * 5); // Update items per page
        }
      };

      // Initial calculation
      handleResize();

      // Recalculate on window resize
      window.addEventListener("resize", handleResize);
      return () => window.removeEventListener("resize", handleResize);
    }, [noPagination, games.length, type]);

    const totalPages = noPagination ? 1 : Math.ceil(games.length / itemsPerPage);

    const currentGames = noPagination ? games : games.slice(
      (currentPage - 1) * itemsPerPage,
      currentPage * itemsPerPage
    );

    const handlePageChange = (page) => {
      if (page >= 1 && page <= totalPages) {
        setCurrentPage(page);
      }
    };

    return (
      <div>
        {/* Gallery */}
        <div
          ref={galleryRef}
          className="grid"
          style={{
            gridTemplateColumns: `repeat(auto-fit, minmax(${columnWidth}px, 1fr))`,
            gap: "16px", // Ensures row and column gaps are equal
          }}
        >
          {currentGames.map((game) => (
            <GalleryItem key={`gallery-item-${game.id}`} game={game} type={type} columnWidth={columnWidth} columnHeight={columnHeight} />
          ))}
        </div>

        {/* Pagination */}
        { !noPagination && (
          <Pagination currentPage={currentPage} totalPages={totalPages} handlePageChange={handlePageChange} />
        )}
      </div>
    );
};

export default Gallery;