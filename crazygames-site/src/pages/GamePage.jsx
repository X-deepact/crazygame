import React, {useEffect, useState} from 'react';

import Gallery from "@/components/Gallery"

const GamePage = ({match}) => {
    const { id } = match.params;

    const games = Array.from({ length: 36 }, (_, index) => ({
      id: index + 1,
      title: `Game ${index + 1}`,
      tag: index%3===1 ? "NEW" : (index%3===0 ? "HOT" : "") , // Alternate between NEW and HOT tags
      img: `https://dummyimage.com/200x150/cccccc/000000&text=Game ${index+1}`, // Placeholder image
      video: "https://www.w3schools.com/html/mov_bbb.mp4", // Placeholder video
    }));

    const tags = [
      { id: 1, name: "Casual", icon: "🌈" },
      { id: 2, name: "Defense", icon: "🛡️" },
      { id: 3, name: "Tower Defense", icon: "🏰" },
      { id: 4, name: "Strategy", icon: "♟️" },
      { id: 5, name: "2D", icon: "🖼️" },
      { id: 6, name: "Mouse", icon: "🖱️" },
    ];

    const embedUrls = [
    "https://www.crazygames.com/embed/aod---art-of-defense",
    "https://www.crazygames.com/embed/cursed-treasure-2",
    "https://www.crazygames.com/embed/tanks-battlefield-desert",
    "https://www.crazygames.com/embed/sniper-master",
    "https://www.crazygames.com/embed/epic-empire-tower-defense",
    "https://www.crazygames.com/embed/defender-idle-2",
    "https://www.crazygames.com/embed/space-waves",
    "https://www.crazygames.com/embed/man-runner-2048",
    "https://www.crazygames.com/embed/jump-guys",
    "https://www.crazygames.com/embed/mini-golf-club",
  ];

    const iframeHTML = `
        <iframe
          src="${embedUrls[Math.floor(Math.random() * embedUrls.length)]}"
          style="width: 100%; height: 100%;"
          frameborder="0"
          allow="gamepad *;"
        ></iframe>
      `;

    return (
        <div className="flex w-full bg-gray-800 min-h-screen">
            {/* Main Game Section */}
            <div className="flex-1 flex flex-col p-4">
                <div
                    style={{ width: "100%", height: "60vh" }}
                    dangerouslySetInnerHTML={{ __html: iframeHTML }}
                  />

                  <div className="mt-5 flex flex-col lg:flex-row items-start lg:items-center justify-between mb-8">
                    <h1 className="text-4xl text-white font-bold mb-4 lg:mb-0">AOD - Art Of Defense</h1>
                    <div className="flex items-center gap-4">
                      <button className="bg-gray-800 px-4 py-2 rounded-full text-white bg-gray-700 hover:bg-gray-600">
                        <i className="fas fa-share"></i> Share
                      </button>
                      <button className="bg-gray-800 px-4 py-2 rounded-full text-white bg-gray-700 hover:bg-gray-600">
                        <i className="fas fa-code"></i> Embed
                      </button>
                    </div>
                  </div>

                  {/* Game Statistics Section */}
                  <table className="text-left">
                    <tbody>
                      {/* Rating */}
                      <tr className="">
                        <td className="py-2 text-sm text-gray-500">Rating:</td>
                        <td className="py-2">
                          <span className="text-white font-bold">
                            8.5 <span className="text-sm text-gray-400">(18,556 votes)</span>
                          </span>
                        </td>
                      </tr>

                      {/* Developer */}
                      <tr className="">
                        <td className="py-2 text-sm text-gray-500">Developer:</td>
                        <td className="py-2">
                          <a href="#" className="text-blue-500 hover:underline">
                            Sateda
                          </a>
                        </td>
                      </tr>

                      {/* Released */}
                      <tr className="">
                        <td className="py-2 text-sm text-gray-500">Released:</td>
                        <td className="py-2 text-white">March 2024</td>
                      </tr>

                      {/* Last Updated */}
                      <tr className="">
                        <td className="py-2 text-sm text-gray-500">Last Updated:</td>
                        <td className="py-2 text-white">November 2024</td>
                      </tr>

                      {/* Technology */}
                      <tr className="">
                        <td className="py-2 text-sm text-gray-500">Technology:</td>
                        <td className="py-2 text-white">HTML5 (Unity WebGL)</td>
                      </tr>

                      {/* Platform */}
                      <tr className="">
                        <td className="py-2 text-sm text-gray-500">Platform:</td>
                        <td className="py-2 text-white">Browser (desktop, mobile, tablet)</td>
                      </tr>

                      {/* Classification */}
                      <tr>
                        <td className="py-2 text-sm text-gray-500">Classification:</td>
                        <td className="py-2">
                          <a href="#" className="text-blue-500 hover:underline">
                            Games
                          </a>
                          <span className="text-gray-400"> » </span>
                          <a href="#" className="text-blue-500 hover:underline">
                            Casual
                          </a>
                          <span className="text-gray-400"> » </span>
                          <a href="#" className="text-blue-500 hover:underline">
                            Defense
                          </a>
                          <span className="text-gray-400"> » </span>
                          <a href="#" className="text-blue-500 hover:underline">
                            Tower Defense
                          </a>
                        </td>
                      </tr>
                    </tbody>
                  </table>

                  <p className="mt-4 text-md border-t border-gray-600 py-4 text-gray-100">AOD - Art of Defence is a casual tower defense game set in a post-apocalyptic world. As Commander of the A.O.D squad, you lead the charge against Mr. Ivil's ruthless thugs, determined to find the lost Inola project. With stunning isometric graphics, build your technological kingdoms and thwart the enemy's destructive ambitions. Will you save mankind or succumb to darkness?</p>

                  <h2 className="mt-5 text-2xl font-bold mb-4 text-white">How to Play AOD - Art of Defence</h2>
                  <p className="mt-4 text-sm text-gray-100">Your mission in AOD Art of Defense begins with an urgent message from Mr. Ivil—find the blueprints of the fabled Inola Project at all costs—even if it means reducing the city to rubble. Time is of the essence. Meanwhile, Ich Nevel warns of an imminent large-scale assault on your bases. The stakes couldn’t be higher, and it’s up to you to defend humanity’s last hope. You have one goal - to protect your base at all costs.</p>
                  <p className="mt-4 text-sm text-gray-100">Dynamic battles ensure that every moment is filled with action. Combining elements of Tower Defense, RPG, and tactical strategy, the game features over 500 sectors for real-time warfare, offering a variety of environments—from small villages to massive abandoned megapolises. You’ll face challenges in exciting new Tower Defense modes like escape, fog, and survival. </p>
                  <p className="mt-4 text-sm text-gray-100">The aim is strategically deploying siege houses, tanks, and armed trucks to stop enemy forces before they breach your defenses. With over 1,000 tower upgrades, including tanks, miniguns, and air defense systems, your arsenal is as vast as powerful. Along the way, you’ll unlock six unique heroes, each with their own abilities to level up and master. From nuclear bombs to ionic satellite strikes and ballistic barrages, a range of devastating boosters will turn the tide in your favor.</p>
                  <p className="mt-4 text-sm text-gray-100">Every successful defense wave rewards you with crystals, essential for upgrading your arsenal. The longer you survive, the greater the rewards, but enemies will become stronger with each attack. Hundreds of card upgrades allow you to customize your towers for maximum efficiency, keeping your strategies fresh and adaptable.</p>

                  <h5 className="mt-[2rem] text-xl text-white">More Games Like This</h5>
                  <p className="mt-4 text-sm text-gray-100">There are plenty of casual games where you can show off your game skills. Want to stick to defense games? Try our Bloons Tower Defense, a casual game where your mission is to build and fortify your defenses to stop an ever-growing wave of colorful balloons from reaching the end of the track. Another fun game is Doodle Jump, a thrilling game that has you jumping upwards to infinity, grabbing boosts, and avoiding beasts on your way up.</p>

                  <h5 className="mt-[2rem] text-xl text-white">Release Date</h5>
                  <ul className="list-disc list-inside text-gray-100 p-2 text-sm space-y-2">
                    <li>
                      <span className="text-gray-100">July 2019</span> <span className="text-gray-400">(Android)</span>
                    </li>
                    <li>
                      <span className="text-gray-100">December 2020</span> <span className="text-gray-400">(Steam)</span>
                    </li>
                    <li>
                      <span className="text-gray-100">March 2024</span>
                    </li>
                  </ul>

                  <h5 className="mt-[2rem] text-xl text-white">Platforms</h5>
                  <ul className="list-disc list-inside text-gray-100 p-2 text-sm space-y-2">
                    <li>
                      <span className="text-gray-100">Web browser</span> <span className="text-gray-400">(desktop and mobile)</span>
                    </li>
                    <li>
                      <span className="text-gray-100">Android</span> <span className="text-gray-400"></span>
                    </li>
                    <li>
                      <span className="text-gray-100">Steam</span>
                    </li>
                  </ul>

                  <h5 className="mt-[2rem] text-xl text-white">Last Updated</h5>
                  <p className="mt-4 text-sm text-gray-100">Nov 18, 2024</p>

                  <h5 className="mt-[2rem] text-xl text-white">Controls</h5>
                  <ul className="list-disc list-inside text-gray-100 p-2 text-sm space-y-2">
                    <li>
                      <span className="text-gray-100">Use the left mouse button to place a tower</span>
                    </li>
                    <li>
                      <span className="text-gray-100">Use the scroll wheel/pinch the touchpad to zoom in/out</span>
                    </li>
                  </ul>

                  <h4 className="mt-[2rem] text-2xl text-white">FAQ</h4>

                  <h5 className="mt-[2rem] text-lg text-white">What does a katana do in Art of Defense?</h5>
                  <p className="mt-4 text-sm text-gray-100">The katana is a rocket launcher that strikes ground targets in the Art of Defence?</p>

                  <h5 className="mt-[2rem] text-lg text-white">How do you change the tower for the hero in Art of Defense?</h5>
                  <p className="mt-4 text-sm text-gray-100">In the Art of Defense, the hero can move between towers whenever needed. Simply choose a tower and click the hero icon to send them upward.</p>

                  <h4 className="mt-[2rem] text-2xl text-white">Gameplay Video</h4>
                  <div className="mt-4 flex justify-center items-center bg-black p-4 rounded-md">
                    <iframe
                      className="w-full max-w-3xl h-[400px] rounded-lg"
                      src="https://www.youtube.com/embed/your-video-id"
                      frameBorder="0"
                      allow="accelerometer; autoplay; clipboard-write; encrypted-media; gyroscope; picture-in-picture"
                      allowFullScreen
                    ></iframe>
                  </div>

                  <ul className="flex space-x-4 p-4 rounded-lg">
                    {tags.map((tag) => (
                      <li
                        key={tag.id}
                        className="flex items-center space-x-2 px-4 py-2 bg-gray-700 hover:bg-gray-600 text-white rounded-full shadow-md cursor-pointer"
                      >
                        {/* Icon */}
                        <span className="text-lg">{tag.icon}</span>
                        {/* Name */}
                        <span className="text-sm font-medium">{tag.name}</span>
                      </li>
                    ))}
                  </ul>
            </div>

            {/* Ad Section */}
            <div className="w-[448px] flex flex-col gap-4 p-4 bg-gray-700">
              <div className="w-full h-[250px]">
                <h1>AD</h1>
              </div>

              <Gallery games={games} noPagination={true} type="hover" />
            </div>
        </div>
    );
};

export default GamePage;