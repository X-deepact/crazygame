import React from "react";
import { NavLink } from "react-router-dom";

const Sidebar = ({ isCollapsed }) => {
  // Define menu items with Font Awesome icons
  const menuItems = [
    {
      name: "Dashboard",
      path: "/dashboard",
      icon: "fas fa-tachometer-alt",
      description: "Admin dashboard",
    },
    {
      name: "Game Management",
      path: "/games",
      icon: "fas fa-gamepad",
      description: "Add, edit or delete games from the list.",
    },
    {
      name: "Category Management",
      path: "/categories",
      icon: "fas fa-th-large",
      description: "Add New, Update, Delete",
    },
    {
      name: "User Management",
      path: "/users",
      icon: "fas fa-users",
      description: "View and manage registered user information.",
    },
    {
      name: "Statistics",
      path: "/statistics",
      icon: "fas fa-chart-line",
      description: "View statistics on game plays, visits, and ratings.",
    },
    {
      name: "Ad Management",
      path: "/ad",
      icon: "fas fa-bullhorn",
      description: "Add new ads, Deliver ads, Track performance.",
    },
    {
      name: "Feedback Management",
      path: "/feedback",
      icon: "fas fa-comments",
      description:
        "View game reviews, Delete or flag feedback, Reply to feedback.",
    },
    {
      name: "System Configuration",
      path: "/configuration",
      icon: "fas fa-cogs",
      description: "Website settings management.",
    },
  ];

  return (
    <div
      className={`bg-gray-800 text-white ${
        isCollapsed ? "w-[4.5rem]" : "w-64"
      } fixed top-[4.5rem] left-0 bottom-0 transition-all duration-300 z-10`}
    >
      {/* Menu Items */}
      <ul className="mt-1">
        {menuItems.map((item) => (
          <li key={item.name} className="mb-4">
            <NavLink
              to={item.path}
              className={`flex items-center p-2 hover:bg-gray-700 ${
                isCollapsed ? "justify-center" : "space-x-4"
              }`}
              activeClassName="bg-gray-900"
            >
              {/* Font Awesome Icon */}
              <i className={`${item.icon} text-lg`}></i>
              {/* Menu Label (Hidden when collapsed) */}
              {!isCollapsed && <span>{item.name}</span>}
            </NavLink>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Sidebar;