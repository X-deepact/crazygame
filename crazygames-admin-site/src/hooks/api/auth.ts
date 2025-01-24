import { apiClient } from '@/utils/axios';
import { useMutation } from '@tanstack/react-query';
import { z } from 'zod';

export const loginSchema = z
  .object({
    email: z.string().email(),
    password: z.string().max(100).nonempty({ message: 'Password is required' }),
  })
  .required();

export const useLoginMutation = () => {
  return useMutation({
    mutationFn: (data) => {
      return apiClient.post('/auth/login', data);
    },
  });
};

export const registerSchema = loginSchema
  .extend({
    username: z
      .string()
      .max(100, { message: 'Username must be at most 100 characters long' })
      .nonempty({ message: 'Username is required' }),
    password: z
      .string()
      .min(8, { message: 'Password must be at least 8 characters long' })
      .max(100, { message: 'Password must be at most 100 characters long' })
      .nonempty({ message: 'Password is required' }),
  })
  .required();

export const useRegisterMutation = () => {
  return useMutation({
    mutationFn: (data) => {
      return apiClient.post('/auth/register', data);
    },
  });
};
