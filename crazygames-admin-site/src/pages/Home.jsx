import React, { useState, useEffect } from 'react';
import { useHistory, Switch, Route, Redirect } from 'react-router-dom';

import Header from '@/pages/Header'
import Sidebar from '@/pages/Sidebar'
import Dashboard from '@/pages/Dashboard'
import GameManagement from '@/pages/GameManagement'
import CategoryManagement from '@/pages/CategoryManagement'
import UserManagement from '@/pages/UserManagement'
import Statistics from "@/pages/Statistics";
import AdManagement from "@/pages/AdManagement";
import FeedbackManagement from "@/pages/FeedbackManagement";
import SystemConfiguration from "@/pages/SystemConfiguration";
import MenuManagement from "@/pages/MenuManagement";
import GameDetails from "@/pages/GameDetails";
import useAuthStore from '@/hooks/store/useAuth';

export default function Home({ onLogout }) {
  const [isSidebarCollapsed, setIsSidebarCollapsed] = useState(false); // Sidebar collapse state
  const history = useHistory();
  const { token } = useAuthStore();

  useEffect(() => {
    if (!token) history.push('/login');
  }, []);

  // Handle responsive sidebar
  useEffect(() => {
    const handleResize = () => {
      if (window.innerWidth < 1024) {
        setIsSidebarCollapsed(true);
      } else {
        setIsSidebarCollapsed(false);
      }
    };

    // Attach the event listener
    window.addEventListener('resize', handleResize);

    // Trigger once on component mount
    handleResize();

    // Cleanup the event listener
    return () => {
      window.removeEventListener('resize', handleResize);
    };
  }, []);

  return (
    <div className='flex flex-col h-screen'>
      <Header
        toggleSidebar={() => setIsSidebarCollapsed(!isSidebarCollapsed)}
        isSidebarCollapsed={isSidebarCollapsed}
        onLogout={onLogout}
      />

      <div className='flex flex-grow'>
        <Sidebar isCollapsed={isSidebarCollapsed} />

              <div
                  className={`flex-grow bg-gray-100 p-6 transition-all duration-300 mt-[4.5rem] ${
                    isSidebarCollapsed ? "ml-[4.5rem]" : "ml-64"
                  }`}
                >
                  <Switch>
                      <Route exact path="/">
                          <Redirect to="/dashboard"/>
                      </Route>
                      <Route path="/dashboard" component={Dashboard}/>
                      <Route exact path="/games" component={GameManagement}/>
                      <Route path="/games/:id" render={(props) => <GameDetails key={props.match.params.id} {...props} />}/>
                      <Route path="/categories" component={CategoryManagement}/>
                      <Route path="/users" component={UserManagement}/>
                      <Route path="/statistics" component={Statistics}/>
                      <Route path="/ad" component={AdManagement}/>
                      <Route path="/feedback" component={FeedbackManagement}/>
                      <Route path="/configuration" component={SystemConfiguration}/>
                  </Switch>
              </div>
      </div>
    </div>
  );
}
