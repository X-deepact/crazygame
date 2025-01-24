import React, { useState, useEffect } from 'react';
import {
  BrowserRouter as Router,
  Switch,
  Route,
  Redirect,
} from 'react-router-dom';

import Login from '@/pages/Login';
import Register from '@/pages/Register';
import Home from '@/pages/Home';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';
import { env } from '@/utils/env';
import { Toaster } from 'react-hot-toast';

const queryClient = new QueryClient();

export default function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Router>
        <Switch>
          <Route exact path='/login'>
            <Login />
          </Route>
          <Route exact path='/register' component={Register} />

          <Route path='/' render={(props) => <Home {...props} />} />
        </Switch>
      </Router>

      <Toaster />

      {env.APP_ENV !== 'production' && <ReactQueryDevtools />}
    </QueryClientProvider>
  );
}
