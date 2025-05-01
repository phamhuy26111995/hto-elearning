import React from 'react'
import { useDarkMode } from '../hooks/useDarkMode';

type Props = {}

export default function Home({}: Props) {
    const [isDark, setIsDark] = useDarkMode();
    return (
      <button
        onClick={() => setIsDark(!isDark)}
        className="px-4 py-2 bg-primary text-primary-foreground rounded"
      >
        Switch to {isDark ? 'Light' : 'Dark'} Mode
      </button>
    );
}