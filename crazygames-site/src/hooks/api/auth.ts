import { apiClient } from '@/utils/axios';
import { useMutation, useQuery } from '@tanstack/react-query';
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

const passwordSchema = z
  .string()
  .min(6, 'Password must be at least 6 characters long')
  .regex(/[a-zA-Z]/, 'Password must contain at least one letter')
  .regex(/[0-9]/, 'Password must contain at least one number');

export const registerSchema = loginSchema
  .extend({
    username: z
      .string()
      .max(100, { message: 'Username must be at most 100 characters long' })
      .nonempty({ message: 'Username is required' }),
    password: passwordSchema,
    confirmPassword: z.string(),
  })
  .required()
  .refine((data) => data.password === data.confirmPassword, {
    message: 'Passwords must match',
    path: ['confirmPassword'],
  });

export const useRegisterMutation = () => {
  return useMutation({
    mutationFn: (data) => {
      return apiClient.post('/auth/register', data);
    },
  });
};

export const checkEmailSchema = z
  .object({
    email: z.string().email(),
  })
  .required();

export const useCheckEmailMutation = () => {
  return useMutation({
    mutationFn: (data) => {
      return apiClient.post('/auth/check-email', data);
    },
  });
};

export const useLoginWithGoogleMutation = () => {
  return useMutation({
    mutationFn: () => {
      return apiClient.get('/Oauth/google/login');
    },
  });
};
