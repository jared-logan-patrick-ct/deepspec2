'use client';

import dynamic from 'next/dynamic';
import Header from './components/Header';

const Grid = dynamic(() => import('./components/Grid'), { ssr: false });

export default function Home() {
  return (
    <div className="h-screen flex flex-col bg-bg-dark text-text-white">
      <Header />
      <main className="flex-1 overflow-hidden">
        <Grid />
      </main>
    </div>
  );
}
