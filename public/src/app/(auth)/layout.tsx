'use client';

import React, { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import Cookies from 'universal-cookie';

export default function AuthLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();
  const cookies = new Cookies(); // Initialize universal-cookie instance.

  useEffect(() => {
    // Get the 'auth-storage' cookie.
    const authCookie = cookies.get('auth-storage');
    // console.log('Auth Cookie:', authCookie);

    // If the cookie exists, attempt to parse it and check if it's valid.
    if (authCookie) {
      try {
        const parsedAuth = typeof authCookie === 'string' ? JSON.parse(authCookie) : authCookie;

        // Check if the token exists in the parsed data
        if (parsedAuth?.token) {
          router.push('/'); // Redirect to home page if the user is already authenticated.
        }
      } catch (error) {
        console.error('Error parsing auth cookie:', error);
        // Optionally clear the cookie if it’s invalid
        cookies.remove('auth-storage');
      }
    }
  }, [router, cookies]);

  return (
    <main className="min-h-screen flex items-center justify-center p-4">
      {children}
    </main>
  );
}
