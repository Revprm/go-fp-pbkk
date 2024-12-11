'use client';

import {
  createContext,
  useContext,
  useCallback,
  ReactNode,
  useEffect,
  useState,
} from 'react';
import { useRouter } from 'next/navigation';
import { User } from '@/lib/definitions';
import Cookies from 'universal-cookie'; // Import universal-cookie

interface AuthState {
  user: User | null;
  token: string | null;
  isLoading: boolean;
}

interface AuthContextType extends AuthState {
  login: (user: User, token: string) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

const AUTH_STORAGE_KEY = 'auth-storage';

export function AuthProvider({ children }: { children: ReactNode }) {
  const router = useRouter();
  const [state, setState] = useState<AuthState>({
    user: null,
    token: null,
    isLoading: true,
  });

  const cookies = new Cookies(); // Initialize universal-cookie

  useEffect(() => {
    // Check authentication on mount
    const checkAuth = () => {
      try {
        const stored = localStorage.getItem(AUTH_STORAGE_KEY);
        if (stored) {
          const parsedAuth = JSON.parse(stored);
          setState({ ...parsedAuth, isLoading: false });
        } else {
          setState((prev) => ({ ...prev, isLoading: false }));
        }
      } catch (error) {
        setState((prev) => ({ ...prev, isLoading: false }));
      }
    };

    checkAuth();
  }, []);

  const login = useCallback((user: User, token: string) => {
    const newState = { user, token, isLoading: false };
    setState(newState);
    localStorage.setItem(AUTH_STORAGE_KEY, JSON.stringify({ user, token }));
    cookies.set('auth-storage', JSON.stringify({ user, token }), { path: '/' }); // Set cookie on login
    router.push('/');
  }, [router]);

  const logout = useCallback(() => {
    setState({ user: null, token: null, isLoading: false });
    localStorage.removeItem(AUTH_STORAGE_KEY);
    cookies.remove('auth-storage', { path: '/' }); // Remove cookie on logout
    router.push('/login');
  }, [router]);

  return (
    <AuthContext.Provider value={{ ...state, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error('useAuth must be used within an AuthProvider');
  }
  return context;
}
