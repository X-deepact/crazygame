import { create } from 'zustand';

export const useLoginError = create((set) => {
  return {
    errors: { login: '', register: '' },
    setErrors: (key, value) =>
      set((state) => ({
        errors: { ...state.errors, [key]: value },
      })),
  };
});
