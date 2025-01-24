import React, { useEffect, useState } from 'react';

import SideModal from '@/components/dialogs/SideModal';
import PeerInput from '@/components/PeerInput';
import {
  useCheckEmailMutation,
  useLoginWithGoogleMutation,
} from '@/hooks/api/auth';
import useAuthStore from '@/hooks/store/useAuth';
import { useLoginError } from '@/hooks/store/useLoginError';

const LoginModal = ({
  isOpen,
  onClose,
  emailAddress,
  onSignup,
  onLogin,
  alert,
}) => {
  const [email, setEmail] = useState(emailAddress || '');
  const [password, setPassword] = useState('');
  const [alreadyRegistered, setAlreadyRegistered] = useState(false); // Control visibility
  const checkEmailMutation = useCheckEmailMutation();
  const { setUser } = useAuthStore();
  const { errors } = useLoginError();

  // Email validation using regex
  const emailRegex = /^[a-zA-Z0-9._-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,6}$/;
  const isEmailValid = emailRegex.test(email);

  useEffect(() => {
    if (isOpen) {
      // Reset the email input field when the modal is opened
      setEmail(emailAddress || '');
      setPassword('');
      setAlreadyRegistered(false);
    }
  }, [isOpen]); // Trigger when modal opens or emailAddress changes

  const handleContinue = async () => {
    // TODO - check email is already registered or not
    // if registered, show password field. otherwise go to "Sign Up" dialog
    const response = await checkEmailMutation.mutateAsync({ email });

    if (response?.message === 'email_exists') {
      setUser(response?.userInfo);
      setAlreadyRegistered(true);

      if (password !== '') {
        onLogin({ email, password });
      }
    } else {
      setAlreadyRegistered(false);
      onSignup(email);
    }
  };

  const googleLoginMutation = useLoginWithGoogleMutation();
  const handleGoogleLogin = async () => {
    const response = await googleLoginMutation.mutateAsync();
    window.location.href = response.redirect_url;
  };

  return (
    <SideModal title='' isOpen={isOpen} onClose={onClose}>
      <h3 className='text-2xl text-center mb-[2.5rem]'>Log in or Sign up</h3>

      <button
        className='w-full flex items-center justify-center bg-red-600 hover:bg-red-500 p-2 rounded-full mb-4'
        onClick={handleGoogleLogin}
      >
        <i className='fab fa-google mr-2'></i>
        Sign in with Google
      </button>
      <button className='w-full flex items-center justify-center bg-blue-600 hover:bg-blue-500 p-2 rounded-full mb-4'>
        <i className='fab fa-facebook mr-2'></i>
        Continue with Facebook
      </button>

      <div className='text-center my-4'>OR</div>

      <PeerInput
        id='email'
        label='Email'
        type='email'
        value={email}
        setValue={setEmail}
      />

      {alreadyRegistered && (
        <PeerInput
          id='password'
          label='Password'
          type='password'
          value={password}
          setValue={setPassword}
        />
      )}

      <button
        disabled={!isEmailValid} // Disable if email is not valid
        className={`w-full ${
          isEmailValid
            ? 'bg-purple-600 hover:bg-purple-500'
            : 'cursor-not-allowed bg-gray-700'
        } text-white font-semibold rounded-lg py-3`}
        onClick={() => handleContinue()}
      >
        Continue
      </button>

      {errors?.login && (
        <p className='text-red-700 text-center font-bold mt-2'>
          {errors?.login}
        </p>
      )}
    </SideModal>
  );
};

export default LoginModal;
