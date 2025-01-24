import React, { useState, useEffect } from 'react';
import { useHistory } from "react-router-dom";

import Slider from "react-slick";
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";


const createSubArrays = (items) => {
  const result = [];
  let i = 0;

  while (i < items.length) {
    if (result.length % 2 === 0) {
      // Add 1 item for odd groups
      result.push(items.slice(i, i + 1));
      i += 1;
    } else {
      // Add 4 items for even groups
      let sub = items.slice(i, i + 4);
      if (sub.length===4)
        result.push(sub);
      i += 4;
    }
  }

  return result;
};

const Carousel = ({isMinimized, title, type, items}) => {
  const calculateSlidesToShow = () => {
    let width = window.innerWidth - (isMinimized ? 85 : 235) - 32;
    let slides = Math.floor(width / (type==="special" ? 416 : 216))-1;
    if (slides<1)
      slides = 1;

    return slides; // Extra small screens
  }

  // Update slidesToShow on window resize
  useEffect(() => {
    const handleResize = () => {
      setSlidesToShow(calculateSlidesToShow());
    };

    window.addEventListener("resize", handleResize);

    // Cleanup event listener
    return () => {
      window.removeEventListener("resize", handleResize);
    };
  }, []);

  const [hoveredGameId, setHoveredGameId] = useState(null); // Track which game is hovered
  const [slidesToShow, setSlidesToShow] = useState(calculateSlidesToShow());

  const history = useHistory();
  const handleItemClick = (id) => {
    history.push(`/games/${id}`); // Navigate to the game page with the ID
  };

  const settings = {
    dots: false,
    infinite: true,
    speed: 500,
    slidesToShow: slidesToShow,
    slidesToScroll: 1,
    arrows: true,
    initialSlide: 0,
    variableWidth: true,
    centerMode: true,
    swipeToSlide: true,
    autoPlay: true,
    draggable: false,
    lazyload: "ondemand",
    responsive: [
      {
        breakpoint: 1024, // For tablets and smaller screens
        settings: {
          slidesToShow: 3, // Show 3 items
        },
      },
      {
        breakpoint: 768, // For mobile devices
        settings: {
          slidesToShow: 2, // Show 2 items
        },
      },
      {
        breakpoint: 480, // Smaller mobile devices
        settings: {
          slidesToShow: 1, // Show 1 item
        },
      },
    ],
  };

  const groupedItems = type  ===  "special" ? createSubArrays(items) : null;

  return (
    <div className={`w-full ${isMinimized ? "w-[calc(100vw-170px)]" : "w-[calc(100vw-350px)]"}`}>
      <h2 className="text-xl font-bold text-white mb-4">{title}</h2>
      <Slider {...settings}>
        {type === "normal" &&
          items.map((item, index) => (
            <div key={item.id}
                 className={`px-2 w-[200px]`}>
              <div className="relative group rounded-lg overflow-hidden w-[200px] h-[150px] cursor-pointer border border-transparent hover:border-purple-600 transition-all duration-300"
                   onClick={() => handleItemClick(item.id)}
                   onMouseEnter={() => setHoveredGameId(item.id)} // Trigger video load on hover
                   onMouseLeave={() => setHoveredGameId(null)}    // Remove hover state on mouse leave
                 >
                <img
                  src={item.img}
                  alt={item.title}
                  className="w-full h-full object-cover"
                />

                {hoveredGameId === item.id && (
                  <video
                    src="https://www.w3schools.com/html/mov_bbb.mp4"
                    className={`absolute top-0 left-0 w-full h-full object-cover opacity-0 group-hover:opacity-100 transition-opacity duration-300`}
                    autoPlay
                    loop
                    muted
                    playsInline
                  />
                )}
              </div>
            </div>
          ))
        }

        {type === "special" &&
          groupedItems.map((group, index) => (
            <div key={`g-${index}`} className="px-2 w-[416px]">
              {index  %  2  ===  0 && (
                <div className={`relative group rounded-lg overflow-hidden w-[416px] h-[416px] cursor-pointer border border-transparent hover:border-purple-600 transition-all duration-300`}
                     onClick={() => handleItemClick(group[0].id)}
                     onMouseEnter={() => setHoveredGameId(group[0].id)} // Trigger video load on hover
                     onMouseLeave={() => setHoveredGameId(null)}    // Remove hover state on mouse leave
                  >
                  <img
                    src={group[0].img}
                    alt={`Game ${group[0].id}`}
                    className="w-full h-full object-cover"
                  />

                  {hoveredGameId === group[0].id && (
                    <video
                      src="https://www.w3schools.com/html/mov_bbb.mp4"
                      className={`absolute top-0 left-0 w-full h-full object-cover opacity-0 group-hover:opacity-100 transition-opacity duration-300`}
                      autoPlay
                      loop
                      muted
                    />
                  )}
                </div>
              )}

              {index  %  2  ===  1 && (
                <div className={`grid grid-rows-2 grid-cols-2 gap-4 w[-416px] h-[416px]`}>
                  <div className={`relative group rounded-lg overflow-hidden cursor-pointer border border-transparent hover:border-purple-600 transition-all duration-300`}
                       onClick={() => handleItemClick(group[0].id)}
                       onMouseEnter={() => setHoveredGameId(group[0].id)} // Trigger video load on hover
                       onMouseLeave={() => setHoveredGameId(null)}    // Remove hover state on mouse leave
                    >
                    <img
                      src={group[0].img}
                      alt={`Game ${group[0].id}`}
                      className="w-[200px] h-[200px] object-cover rounded-lg"
                    />

                    {hoveredGameId === group[0].id && (
                      <video
                        src="https://www.w3schools.com/html/mov_bbb.mp4"
                        className={`absolute top-0 left-0 w-full h-full object-cover opacity-0 group-hover:opacity-100 transition-opacity duration-300`}
                        autoPlay
                        loop
                        muted
                      />
                    )}
                  </div>

                  {group.length  >  1 && (
                      <div className={`relative group rounded-lg overflow-hidden cursor-pointer border border-transparent hover:border-purple-600 transition-all duration-300`}
                           onClick={() => handleItemClick(group[1].id)}
                           onMouseEnter={() => setHoveredGameId(group[1].id)} // Trigger video load on hover
                           onMouseLeave={() => setHoveredGameId(null)}    // Remove hover state on mouse leave
                            >
                        <img
                          src={group[1].img}
                          alt={`Game ${group[1].id}`}
                          className="w-[200px] h-[200px] object-cover"
                        />

                        {hoveredGameId === group[1].id && (
                          <video
                            src="https://www.w3schools.com/html/mov_bbb.mp4"
                            className={`absolute top-0 left-0 w-full h-full object-cover opacity-0 group-hover:opacity-100 transition-opacity duration-300`}
                            autoPlay
                            loop
                            muted
                          />
                        )}
                      </div>
                  )}

                  {group.length  >  2 && (
                      <div className={`relative group rounded-lg overflow-hidden cursor-pointer border border-transparent hover:border-purple-600 transition-all duration-300`}
                           onClick={() => handleItemClick(group[2].id)}
                           onMouseEnter={() => setHoveredGameId(group[2].id)} // Trigger video load on hover
                           onMouseLeave={() => setHoveredGameId(null)}    // Remove hover state on mouse leave
                            >
                        <img
                          src={group[2].img}
                          alt={`Game ${group[2].id}`}
                          className="w-[200px] h-[200px] object-cover"
                        />

                      {hoveredGameId === group[2].id && (
                        <video
                          src="https://www.w3schools.com/html/mov_bbb.mp4"
                          className={`absolute top-0 left-0 w-full h-full object-cover opacity-0 group-hover:opacity-100 transition-opacity duration-300`}
                          autoPlay
                          loop
                          muted
                        />
                      )}
                      </div>
                  )}

                  {group.length  >  3 && (
                      <div className={`relative group rounded-lg overflow-hidden cursor-pointer border border-transparent hover:border-purple-600 transition-all duration-300`}
                           onClick={() => handleItemClick(group[3].id)}
                           onMouseEnter={() => setHoveredGameId(group[3].id)} // Trigger video load on hover
                           onMouseLeave={() => setHoveredGameId(null)}    // Remove hover state on mouse leave
                            >
                        <img
                          src={group[3].img}
                          alt={`Game ${group[3].id}`}
                          className="w-[200px] h-[200px] object-cover rounded-lg"
                        />
                        <video
                          src="https://www.w3schools.com/html/mov_bbb.mp4"
                          className={`absolute top-0 left-0 w-full h-full object-cover opacity-0 group-hover:opacity-100 transition-opacity duration-300`}
                          autoPlay
                          loop
                          muted
                        />
                      </div>
                  )}
                </div>
              )}
            </div>
          ))
        }
      </Slider>
    </div>
  );
};

export default Carousel;