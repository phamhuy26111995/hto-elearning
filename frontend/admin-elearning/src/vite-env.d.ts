/// <reference types="vite/client" />


interface ImportMetaEnv {
    readonly VITE_API_URL: string
    readonly VITE_TEACHER_ENDPOINT: string
    // add more custom keys here...
  }
  
  interface ImportMeta {
    readonly env: ImportMetaEnv
  }