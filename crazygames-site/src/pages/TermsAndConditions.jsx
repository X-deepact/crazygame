import React from 'react';

const TermsAndConditions = ({}) => {

  return (
    <div className="bg-gray-800 text-white min-h-screen p-4">
      {/* Container */}
      <div className="w-full">
        {/* Header Section */}
        <div className="flex justify-between items-center mb-6">
          <h1 className="text-xl font-semibold">Terms & Conditions</h1>
        </div>

        {/* Content Section */}
        <div className="bg-gray-900 rounded-lg shadow-md p-6">
          {/* Last Updated */}
          <p className="text-gray-400 text-sm mb-6">Last updated: June 2024</p>

          {/* Section 1 */}
          <h2 className="text-lg font-semibold mb-4">1. Who we are</h2>
          <p className="text-gray-300 text-sm mb-4">
            We are CrazyGames. We offer a browser game platform where you can
            discover and experience many different web games in an easy and
            accessible way (<strong>“Platform”</strong>). You can access our
            Platform via an account (<strong>“Account”</strong>) or without. The
            Platform is made available via one of CrazyGames’ websites as visible
            in the address bar depending on your region (<strong>“Website”</strong>).
          </p>
          <p className="text-gray-300 text-sm mb-4">
            Legal information about CrazyGames (<strong>“CrazyGames”</strong>,
            <strong>“we”</strong>, <strong>“us”</strong>):
          </p>
          <ul className="list-disc list-inside text-sm text-gray-300 mb-4">
            <li>Maxflow BV</li>
            <li>Ketelmakerij 20, 3010 Kessel-Lo, Belgium</li>
            <li>KBO 0550.758.377.</li>
          </ul>
          <p className="text-gray-300 text-sm mb-4">
            We recognize the trust you place in us and take our responsibility to
            protect your privacy seriously. This Privacy Policy (this
            <strong>“Policy”</strong>) provides important details about how we
            collect, process, disclose, retain, and protect your personal data.
          </p>
          <p className="text-gray-300 text-sm mb-4">
            The Platform is intended for visitors and users who are thirteen (13)
            years of age (or the applicable minimum age in your country) or older.
            If you are under thirteen (13) years of age (or the applicable minimum
            age in your country), this Platform is not intended for you. We do not
            knowingly collect or solicit personal information from children under
            the age of thirteen (13) (or the applicable minimum age in your
            country) through our Platform. However, we have launched a website
            specifically for younger children below the applicable minimum age,{" "}
            <a
              href="https://kids.crazygames.com"
              target="_blank"
              rel="noopener noreferrer"
              className="text-teal-400 hover:underline"
            >
              https://kids.crazygames.com
            </a>{" "}
            (<strong>“Kids Site”</strong>) and have a separate privacy policy for
            the Kids Site, available on{" "}
            <a
              href="https://kids.crazygames.com/privacy-policy"
              target="_blank"
              rel="noopener noreferrer"
              className="text-teal-400 hover:underline"
            >
              https://kids.crazygames.com/privacy-policy
            </a>
            . CrazyGames does not allow personalized advertising on its Kids Site.
          </p>
          <p className="text-gray-300 text-sm mb-4">
            For information about the terms upon which we do business, you should
            also read our{" "}
            <a
              href="/terms-of-service"
              className="text-teal-400 hover:underline"
            >
              Terms of Service
            </a>
            .
          </p>

          {/* Section 2 */}
          <h2 className="text-lg font-semibold mb-4">2. What is Personal Data?</h2>
          <p className="text-gray-300 text-sm mb-4">
            Personal Data is any information about you that allows us to identify
            you. This could be, for example, your name or email address. But
            equally data about the games you played if we can link it to your
            account or location (<strong>“Account”</strong>).
          </p>

          {/* Section 3 */}
          <h2 className="text-lg font-semibold mb-4">3. From whom do we collect personal data?</h2>
          <p className="text-gray-300 text-sm mb-4">
            In order to operate our Platform, we may collect data from users and visitors of the Platform and Website, developers of web games, persons who otherwise provide us with their contact details, and persons who contact us by email or other means. As described above under article 1.3, we do collect personal data from children older than the age of thirteen (13) or the applicable minimum age in your country.
          </p>

          <h2 className="text-lg font-semibold mb-4">4. What personal data does CrazyGames process and why?</h2>
          <p className="text-gray-300 text-sm mb-4">
            CrazyGames may collect the following Personal Data of you:
          </p>
          <p className="text-gray-300 text-sm mb-4">
            To evaluate your job application
          </p>
          <ul className="list-disc list-inside text-sm text-gray-300 mb-4">
            <li>Which personal data: Email address, name, age, gender, CV</li>
            <li>On what basis: Necessary for the performance of a contract or to take steps at your request, before entering a contract</li>
          </ul>
          <p className="text-gray-300 text-sm mb-4">
            To register and authenticate your Account of the Platform
          </p>
          <ul className="list-disc list-inside text-sm text-gray-300 mb-4">
            <li>Which personal data: email address, name, age, gender external account, username and password</li>
            <li>On what basis: The prior, express, free, specific and informed consent of you</li>
          </ul>
          <p className="text-gray-300 text-sm mb-4">
            Personalisation of your game experience (Account)
          </p>
          <ul className="list-disc list-inside text-sm text-gray-300 mb-4">
            <li>Which personal data: username, device and connection data, selected interests, gaming behaviour and feedback</li>
            <li>On what basis: The prior, express, free, specific and informed consent of you</li>
          </ul>

        </div>
      </div>
    </div>
  );
};

export default TermsAndConditions;
