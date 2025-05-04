import React from 'react'
import { LoginForm } from './LoginForm'
import { Toaster } from '@/components/ui/sonner'

export default function Login() {
  return (
    <div className="flex min-h-svh flex-col items-center justify-center bg-muted p-6 md:p-10">
    <div className="w-full max-w-sm md:max-w-3xl">
      <LoginForm />
      <Toaster position="top-right" richColors />
    </div>
  </div>
  )
}
