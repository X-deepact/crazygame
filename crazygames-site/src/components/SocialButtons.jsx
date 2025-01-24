import React from "react";

const SocialButtons = ({type}) => {
  const links = [
    { icon: "fab fa-tiktok", url: "https://www.tiktok.com" },
    { icon: "fab fa-discord", url: "https://www.discord.com" },
    { icon: "fab fa-linkedin", url: "https://www.linkedin.com" },
    { icon: "fab fa-twitter", url: "https://www.twitter.com" }, // Using X as Twitter
    { icon: "fab fa-youtube", url: "https://www.youtube.com" },
    { icon: "fab fa-google-play", url: "https://play.google.com" },
    { icon: "fab fa-apple", url: "https://www.apple.com" },
  ];

  return (
    <div className={`flex flex-wrap space-x-2 p-2 justify-center`}>
      {links.map((link, index) => (
        <a
          key={index}
          href={link.url}
          target="_blank"
          rel="noopener noreferrer"
          className={`flex items-center justify-center w-9 h-9 mb-2 rounded-full border border-gray-500 text-gray-400 hover:text-white hover:border-purple-500 transition-all`}
        >
          <i className={link.icon + " text-md"}></i>
        </a>
      ))}
    </div>
  );
};

export default SocialButtons;