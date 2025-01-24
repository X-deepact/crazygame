import Axios from 'axios';
import { env } from './env';

function authRequestInterceptor(config) {
  if (config.headers) {
    config.headers.Accept = 'application/json';
  }
  const token = localStorage.getItem('token');
  if (token) {
    config.headers.Authorization = `Bearer ${token}`;
  }
  return config;
}

const apiClient = Axios.create({
  baseURL: env.API_URL,
});

apiClient.interceptors.request.use(authRequestInterceptor);

apiClient.interceptors.response.use(
  (response) => {
    return response.data;
  },
  async (error) => {
    return Promise.reject(error);
  }
);

export { apiClient };
