import React from 'react';
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.jsx'
import '@fontsource/inter/400.css';
import '@fontsource/inter/500.css';
import '@fontsource/inter/700.css';

import { init, miniApp } from '@telegram-apps/sdk';
import { retrieveLaunchParams } from '@telegram-apps/sdk';
import { isTMA } from '@telegram-apps/bridge';


if (await isTMA()) {
  console.log("Приложение запущено в Telegram");
} else {
  console.log("Приложение запущено вне Telegram (например, в браузере)");
  // Выполнить безопасный фолбек
}

let params;
try {
  params = retrieveLaunchParams();  // пытаемся получить initData
} catch (error) {
  console.error('Нет параметров запуска Telegram:', error);
  // Фолбек: показываем сообщение пользователю или моковые данные
}


const initializeTelegramSDK = async () => {
  try {
    const insideTelegram = await isTMA();

    if (!insideTelegram) {
      console.warn("Приложение не запущено в Telegram. Пропускаем инициализацию.");
      return;
    }

    await init();

    if (miniApp.ready.isAvailable()) {
      await miniApp.ready();
      console.log("Mini App готово");
    }

  } catch (error) {
    console.error("Ошибка инициализации Telegram SDK:", error);
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