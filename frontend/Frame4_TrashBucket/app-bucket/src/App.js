import React, { useState, useEffect } from 'react';
import './ProductCard.css';

function ProductCard({ title, description, price, image }) {
  const [count, setCount] = useState(0);

  useEffect(() => {
    if (window.Telegram?.WebApp) {
      window.Telegram.WebApp.ready();
      window.Telegram.WebApp.expand();
    }
  }, []);

  const handleDecrement = () => {
    if (count > 0) setCount(count - 1);
  };

  const handleIncrement = () => {
    setCount(count + 1);
  };

  const handleOrder = () => {
    if (count > 0) {
      const orderData = { title, count, price };
      window.Telegram.WebApp.sendData(JSON.stringify(orderData));
    }
  };

  return (
    <div className="product-card">
      <div
        className="product-image"
        style={{ backgroundImage: `url(${image})` }}
      />
      <div className="product-info">
        <div className="product-title">{title}</div>
        <div className="product-description">{description}</div>
        <div className="product-controls">
          <button onClick={handleDecrement} disabled={count === 0}>
            -
          </button>
          <span>{count}</span>
          <button onClick={handleIncrement}>+</button>
        </div>
        <div className="product-price">{price} ₽</div>
        <button className="order-button" onClick={handleOrder} disabled={count === 0}>
          Заказать
        </button>
      </div>
    </div>
  );
}

export default ProductCard;