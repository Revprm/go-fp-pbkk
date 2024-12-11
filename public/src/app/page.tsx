'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import Cookies from 'universal-cookie';
import { TodoList } from '@/components/todo-list';

export default function HomePage() {
  const router = useRouter();
  const cookies = new Cookies();

  useEffect(() => {
    const authCookie = cookies.get('auth-storage');
    console.log('Auth cookie:', authCookie);

    if (!authCookie) {
      router.push('/login'); // Redirect to login if no cookie is found
      return;
    }

    try {
      // Check if the authCookie is already an object or needs to be parsed
      let authData;
      if (typeof authCookie === 'string') {
        authData = JSON.parse(authCookie);  // Parse only if it's a string
      } else {
        authData = authCookie;  // It's already an object, no need to parse
      }

      console.log('Parsed Auth Data:', authData);

      // Check if the token is valid
      if (!authData?.token) {
        router.push('/login'); // Redirect to login if no token
      }
    } catch (error) {
      console.error('Error parsing auth cookie:', error);
      router.push('/login'); // Redirect to login if cookie is malformed
    }
  }, [router, cookies]);

  return (
    <main className="container mx-auto p-4">
      <TodoList />
    </main>
  );
}
