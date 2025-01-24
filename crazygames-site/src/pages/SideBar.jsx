import React from "react";
import { NavLink } from "react-router-dom";

import SocialButtons from "@/components/SocialButtons";
import FooterLinks from "@/components/FooterLinks";

const SideBar = ({ isMinimized, categories }) => {
    const menuItems = [
      { label: "Home", path: '/home', icon: "https://imgs.crazygames.com/icon/Home.svg" },
      { label: "Recently played", path: '/recent', icon: "https://imgs.crazygames.com/icon/Recent.svg" },
      { label: "New", path: '/new', icon: "https://imgs.crazygames.com/icon/New.svg" },
      { label: "Trending now", path: '/trending', icon: "https://imgs.crazygames.com/icon/Trending.svg" },
      { label: "Updated", path: "/updated", icon: "https://imgs.crazygames.com/icon/Updated.svg" },
      { label: "Originals", path: "/originals", icon: "https://imgs.crazygames.com/icon/Originals.svg" },
      { label: "Multiplayer", path: "/multiplayer", icon: "https://imgs.crazygames.com/icon/Multiplayer.svg" },
    ];

    return (
      <div
        className={`scrollbar bg-gray-800 text-white border-r border-gray-900 ${
          isMinimized ? "w-[85px]" : "w-[235px]"
        } h-[calc(100vh-64px)] overflow-y-auto transition-all duration-300 fixed top-16 left-0`}
      >
        <ul className="p-4 mt-3">
          {menuItems.map((item, index) => (
            <li
              key={index}
              className="flex items-center mb-1 space-x-2 cursor-pointer whitespace-nowrap"
            >
              <NavLink
                to={item.path}
                activeClassName="bg-gray-600 "
                className={`flex items-center w-full p-2 rounded hover:bg-gray-700 hover:text-gray-300 ${isMinimized ? "text-center" : ""}`}
                exact
              >
                <img src={item.icon} alt="" />
                {/*<i className={`${item.icon} text-xl`}></i>*/}
                {!isMinimized && <span className="px-2 text-sm">{item.label}</span>}
              </NavLink>
            </li>
          ))}
        </ul>

        <ul className="px-4 py-2 border-t border-gray-900">
          {categories.map((item, index) => (
            <li
              key={index}
              className="flex items-center mb-1 space-x-2 cursor-pointer whitespace-nowrap"
            >
              <NavLink
                to={`/categories/${item.id}`}
                activeClassName="bg-gray-600"
                className={`flex items-center w-full p-2 rounded hover:bg-gray-700 hover:text-gray-300 ${isMinimized ? "text-center" : ""}`}
                exact
              >
                <img src={item.icon} alt="" />
                {/*<i className={`${item.icon} text-xl`}></i>*/}
                {!isMinimized && <span className="px-1 text-sm">{item.label}</span>}
              </NavLink>
            </li>
          ))}
        </ul>

        <ul className="px-4 py-2 border-t border-gray-900">
          <li
            key="tags"
            className="flex items-center mb-1 space-x-2 cursor-pointer whitespace-nowrap"
          >
            <NavLink
              to={`/tags`}
              activeClassName="bg-gray-600"
              className={`flex items-center w-full p-2 rounded hover:bg-gray-700 hover:text-gray-300 ${isMinimized ? "text-center" : ""}`}
              exact
            >
              <img src="https://imgs.crazygames.com/icon/Tags.svg" alt="" />
              {/*<i className={`${item.icon} text-xl`}></i>*/}
              {!isMinimized && <span className="px-1 text-sm">Tags</span>}
            </NavLink>
          </li>
        </ul>

        <div className="mb-3 w-full px-4">
          <button className="w-full text-white bg-purple-600 hover:bg-purple-500 rounded-full px-4 py-2">
            <i className="fa fa-inbox"> Contact Us</i>
          </button>
        </div>

        <div className="mt-4">
          <FooterLinks type="sidebar" />
        </div>

        <div className="mt-4">
          <SocialButtons type="sidebar" />
        </div>

      </div>
    );
  };

export default SideBar;