import { InputErrorMessage } from '@/components/InputErrorMessage';
import { registerSchema, useRegisterMutation } from '@/hooks/api/auth';
import { zodResolver } from '@hookform/resolvers/zod';
import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import toast from 'react-hot-toast';
import { useHistory } from 'react-router-dom';

const Register = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm({
    resolver: zodResolver(registerSchema),
  });
  const history = useHistory();
  const registerMutation = useRegisterMutation();

  const handleRegister = (data) => {
    registerMutation.mutate(data, {
      onSuccess: () => {
        toast.success('Account created successfully!');
        history.push('/login');
      },
      onError: (error) => {
        const message = error?.response?.data?.error;
        const duplicateEmailMsg = 'email already registered';
        if (message.includes(duplicateEmailMsg)) {
          toast.error('Invalid Email Address');
        }
      },
    });
  };

  return (
    <div className='min-h-screen flex items-center justify-center bg-gray-900'>
      <div className='w-full max-w-md bg-gray-800 p-8 rounded-lg'>
        <h2 className='text-2xl font-bold text-white mb-6 text-center'>
          Crazy Games Admin Register
        </h2>
        <form onSubmit={handleSubmit(handleRegister)}>
          <div className='mb-4'>
            <label className='block text-white mb-2'>Username</label>
            <input
              type='username'
              className='w-full p-2 rounded bg-gray-700 text-white'
              placeholder='Enter your username'
              {...register('username')}
            />
            <InputErrorMessage message={errors.username?.message} />
          </div>
          <div className='mb-4'>
            <label className='block text-white mb-2'>Email</label>
            <input
              type='email'
              className='w-full p-2 rounded bg-gray-700 text-white'
              placeholder='Enter your email'
              {...register('email')}
            />
            <InputErrorMessage message={errors.email?.message} />
          </div>
          <div className='mb-4'>
            <label className='block text-white mb-2'>Password</label>
            <input
              type='password'
              className='w-full p-2 rounded bg-gray-700 text-white'
              placeholder='Enter your password'
              {...register('password')}
            />
            <InputErrorMessage message={errors.password?.message} />
          </div>
          <button
            type='submit'
            className='w-full bg-blue-500 p-2 rounded text-white font-bold hover:bg-blue-600'
          >
            Register
          </button>
        </form>
        <p className='text-gray-400 text-sm mt-4 text-center'>
          Already have an account?{' '}
          <a href='/login' className='text-blue-400 hover:underline'>
            Login
          </a>
        </p>
      </div>
    </div>
  );
};

export default Register;
