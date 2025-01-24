import React from "react";

const FooterLinks = ({type}) => {
    const links = [
        {label: "About", url: "https://about.crazygames.com", target: "_blank"},
        {label: "Developers", url: "https://developer.crazygames.com", target: "_blank"},
        {label: "Kids site", url: "https://kids.crazygames.com", target: "_blank"},
        {label: "Jobs", url: "https://crazygames.recruitee.com", target: "_blank"},
//        {label: "Info for parents", url: "https://about.crazygames.com", target: "_blank"},
        {label: "Terms & conditions", url: "/terms-and-conditions", target: ""},
        {label: "Privacy & Policy", url: "/privacy-policy", target: ""},
//        {label: "All games", url: "https://about.crazygames.com", target: "_blank"},
//        {label: "Preferences", url: "https://about.crazygames.com", target: "_blank"},
//        {label: "Do not sell my data", url: "https://about.crazygames.com", target: "_blank"},
    ];

    return (
        <div className={`gap-2 px-4 ${type==='sidebar' ? "flex flex-col" : "grid grid-cols-2"} `}>
        {links.map((link, index) => (
            <a key={index}
               href={link.url}
               target={link.target}
               className={`text-gray-300 text-sm ${type==='sidebar' ? "px-4" : ""} `}>{link.label}</a>
        ))}
        </div>
    );
};

export default FooterLinks;