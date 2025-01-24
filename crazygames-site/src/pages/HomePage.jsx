import React, { useState, useRef, useEffect } from 'react';

import Carousel from "@/components/Carousel";

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

const HomePage = ({ isMinimized }) => {
  const games = [
    { id: 1, title: "Basket Battle", img: "https://dummyimage.com/200x150/cccccc/000000&text=Basket Battle" },
    { id: 2, title: "Survey.io", img: "https://dummyimage.com/200x150/cccccc/000000&text=Survey.io" },
    { id: 3, title: "Food Truck Chef", img: "https://dummyimage.com/200x150/cccccc/000000&text=Food Truck Chef" },
    { id: 4, title: "Connect", img: "https://dummyimage.com/200x150/cccccc/000000&text=Connect" },
    { id: 5, title: "Revolution Idle X", img: "https://dummyimage.com/200x150/cccccc/000000&text=Revolution Idle X" },
    { id: 7, title: "Art of Defense", img: "https://dummyimage.com/200x150/cccccc/000000&text=Art of Defense" },
    { id: 9, title: "Bank Heist", img: "https://dummyimage.com/200x150/cccccc/000000&text=Bank Heist" },
  ];


  return (
    <div className="bg-gray-800 min-h-screen p-4">
      <h1 className="text-xl font-bold text-white mb-4">Welcome to Crazy Games!</h1>

      <div className="rounded-lg p-3">
        <Carousel title="" isMinimized={isMinimized} type="special" items={shuffleArray(games)} />
      </div>

      <div className="rounded-lg p-3">
        <Carousel title="Featured Games" isMinimized={isMinimized} type="normal" items={shuffleArray(games)} />
      </div>

      <div className="rounded-lg p-3">
        <Carousel title="New Games" isMinimized={isMinimized} type="normal" items={shuffleArray(games)} />
      </div>

      <div className="rounded-lg p-3">
        <Carousel title="Crazy Games Originals" isMinimized={isMinimized} type="normal" items={shuffleArray(games)} />
      </div>

      <div className="rounded-lg p-3">
        <Carousel title="Games to Play with Friends" isMinimized={isMinimized} type="normal" items={shuffleArray(games)} />
      </div>

      <div className="rounded-lg p-3">
        <Carousel title=".io Games" isMinimized={isMinimized} type="normal" items={shuffleArray(games)} />
      </div>

      <div className="rounded-lg p-3">
        <Carousel title="Casual Games" isMinimized={isMinimized} type="normal" items={shuffleArray(games)} />
      </div>

      <div className="rounded-lg p-3">
        <Carousel title="Driving Games" isMinimized={isMinimized} type="normal" items={shuffleArray(games)} />
      </div>

      <div className="rounded-lg p-3">
        <Carousel title="Action Games" isMinimized={isMinimized} type="normal" items={shuffleArray(games)} />
      </div>

      <div className="rounded-lg p-3">
        <Carousel title="Puzzle Games" isMinimized={isMinimized} type="normal" items={shuffleArray(games)} />
      </div>

      <div className="rounded-lg p-3">
        <Carousel title="Adventure Games" isMinimized={isMinimized} type="normal" items={shuffleArray(games)} />
      </div>

      <div className="rounded-lg p-3">
        <Carousel title="Clicker Games" isMinimized={isMinimized} type="normal" items={shuffleArray(games)} />
      </div>

      <div className="rounded-lg p-3">
        <Carousel title="Shooting Games" isMinimized={isMinimized} type="normal" items={shuffleArray(games)} />
      </div>

      <div className="rounded-lg p-3">
        <Carousel title="Sports Games" isMinimized={isMinimized} type="normal" items={shuffleArray(games)} />
      </div>
    </div>
  );
}

export default HomePage;