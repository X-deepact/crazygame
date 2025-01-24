import React from 'react';
import ReactDOM from 'react-dom';
import { Helmet } from 'react-helmet';

import '@/index.css';
import App from '@/App';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { Toaster } from 'react-hot-toast';
import { env } from '@/utils/env';
import { ReactQueryDevtools } from '@tanstack/react-query-devtools';

const queryClient = new QueryClient();

ReactDOM.render(
  <React.StrictMode>
    <QueryClientProvider client={queryClient}>
      <Helmet defaultTitle='Crazy Games' titleTemplate='%s | Crazy Games'>
        <meta charSet='utf-8' />
        <html lang='id' amp />
      </Helmet>
      <App />

      <Toaster />

      {env.APP_ENV !== 'production' && <ReactQueryDevtools />}
    </QueryClientProvider>
  </React.StrictMode>,
  document.getElementById('root')
);
