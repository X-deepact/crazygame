import React, { useState, useEffect, useMemo } from 'react';

import SideModal from '@/components/dialogs/SideModal';
import PeerInput from '@/components/PeerInput';
import { registerSchema } from '@/hooks/api/auth';

const SignupModal = ({
  isOpen,
  onClose,
  emailAddress,
  onLogin,
  onRegister,
}) => {
  const [email, setEmail] = useState(emailAddress || '');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const isValid = useMemo(() => {
    const result = registerSchema.safeParse({
      email,
      password,
      confirmPassword,
      username: email,
    });

    return result.success;
  }, [email, password, confirmPassword]);

  const handleContinue = () => {
    // TODO - validation
    onRegister(email, password);
  };

  useEffect(() => {
    setEmail(emailAddress);
  }, [emailAddress]);

  return (
    <SideModal
      title='Sign up'
      isOpen={isOpen}
      onClose={onClose}
      onBack={onLogin}
    >
      <h3 className='text-2xl text-center mb-[2rem]'>Create a free account</h3>

      <PeerInput
        id='email'
        label='Email'
        type='email'
        value={email}
        setValue={setEmail}
      />
      <PeerInput
        id='password'
        label='Password'
        type='password'
        value={password}
        setValue={setPassword}
      />
      <PeerInput
        id='confirmPassword'
        label='Confirm Password'
        type='password'
        value={confirmPassword}
        setValue={setConfirmPassword}
        children={
          <p className='text-xs text-gray-400 mt-1 block'>
            Passwords must be at least 6 characters and include both letters and
            numbers
          </p>
        }
      />

      <button
        disabled={!isValid} // Disable if email is not valid
        className={`w-full ${
          isValid
            ? 'bg-purple-600 hover:bg-purple-500'
            : 'cursor-not-allowed bg-gray-700'
        } text-white font-semibold rounded-lg py-3`}
        onClick={() => handleContinue()}
      >
        Continue
      </button>
    </SideModal>
  );
};

export default SignupModal;
