import React from 'react';
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'
import '@fontsource/inter/400.css';
import '@fontsource/inter/500.css';
import '@fontsource/inter/700.css';

import { init, miniApp } from '@telegram-apps/sdk';


const initializeTelegramSDK = async () => {
  try {
    await init();


    if (miniApp.ready.isAvailable()) {
      await miniApp.ready();
      console.log('Mini App готово');
    }


  } catch (error) {
    console.error('Ошибка инициализации:', error);
  }
};


try {
  initializeTelegramSDK();
} catch (error) {
  console.warn("Ошибка инициализации Telegram SDK:", error);
}

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <App />
  </StrictMode>,
)