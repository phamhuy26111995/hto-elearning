import { StrictMode } from 'react';
import { createRoot } from 'react-dom/client';
import App from './App.tsx';
import './index.css';
import GlobalModal from './components/project/common/GlobalModal.tsx';


createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <GlobalModal />
    <App />
  </StrictMode>,
)
