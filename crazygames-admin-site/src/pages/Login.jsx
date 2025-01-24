import { loginSchema, useLoginMutation } from '@/hooks/api/auth';
import useAuthStore from '@/hooks/store/useAuth';
import React, { useEffect } from 'react';
import { useForm } from 'react-hook-form';
import { useHistory } from 'react-router-dom';
import { zodResolver } from '@hookform/resolvers/zod';
import { InputErrorMessage } from '@/components/InputErrorMessage';
import toast from 'react-hot-toast';

const Login = ({}) => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(loginSchema),
  });
  const history = useHistory();
  const loginMutation = useLoginMutation();
  const { token, setToken } = useAuthStore();

  useEffect(() => {
    if (token) {
      history.push('/');
    }
  }, []);

  const handleLogin = (data) => {
    loginMutation.mutate(data, {
      onSuccess: (data) => {
        setToken(data?.data);
        toast.success('Successfully logged in!', {
          position: 'top-right',
        });
        history.push('/');
      },
      onError: (error) => {
        toast.error('Invalid Credentials', {
          position: 'top-right',
        });
      },
    });
  };

  return (
    <div className='min-h-screen flex items-center justify-center bg-gray-900'>
      <div className='w-full max-w-md bg-gray-800 p-8 rounded-lg'>
        <h2 className='text-2xl font-bold text-white mb-6 text-center'>
          Crazy Games Admin Login
        </h2>
        <form onSubmit={handleSubmit(handleLogin)}>
          <div className='mb-4'>
            <label className='block text-white mb-2'>Email</label>
            <input
              type='email'
              className='w-full p-2 rounded bg-gray-700 text-white'
              placeholder='Enter your email'
              {...register('email')}
            />
            <InputErrorMessage message={errors?.email?.message} />
          </div>
          <div className='mb-4'>
            <label className='block text-white mb-2'>Password</label>
            <input
              type='password'
              className='w-full p-2 rounded bg-gray-700 text-white'
              placeholder='Enter your password'
              {...register('password')}
            />
            <InputErrorMessage message={errors?.password?.message} />
          </div>
          <button
            type='submit'
            className='w-full bg-blue-500 p-2 rounded text-white font-bold hover:bg-blue-600'
          >
            Login
          </button>
        </form>
        <p className='text-gray-400 text-sm mt-4 text-center'>
          Don't have an account?{' '}
          <a href='/register' className='text-blue-400 hover:underline'>
            Register
          </a>
        </p>
      </div>
    </div>
  );
};

export default Login;
