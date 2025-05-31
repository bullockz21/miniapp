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
      if (!title) {
        title = "Описание товара";
      }
      const orderData = { title, count, price };
      
      window.Telegram.WebApp.sendData(JSON.stringify(orderData));
    }
  };

  return (
    <div className="product-card">
      <div
        className="product-header"
        style={{ backgroundImage: `url(${image})` }}
    >
    {/* rectangle 20 21 22 23 */}
      <div className="product-rectangle">
      <i class="bi bi-house"></i>
        <p className="pro-rect-title">Главная</p> </div>
      <div className="product-rectangle"> 
      <i class="bi bi-person"></i>
      <p className="pro-rect-title">Профиль</p></div>
       <div className="product-rectangle"> 
       <i class="bi bi-cart"></i>
       <p className="pro-rect-title">Корзина</p></div>
        <div className="product-rectangle">
        <i class="bi bi-telephone"></i>
         <p className="pro-rect-title">Контакты</p></div>
      </div>
      <div className="product-address-arrow">
      <div className="product-rectangle18">
        <h3 className="cityaddress">Федора Попова, 1</h3>
        <h4 className="deliverytime">Сегодня, 14 апр., 14:00 - 16:30</h4>
      </div>
      <div className="product-rectangle19">
        <button className="arrow19">
        <i class="bi bi-chevron-right"></i>
        </button>
      </div>
      </div>

    <div className="product-rectangle3">
      <div className="search-input-container">
        <span className="material-icons search-icon">search</span>
        <input
          type="text"
          placeholder="Поиск"
          className="search-input"
        />
      </div>
      <button className="menu-button">
        <span className="material-icons">menu</span>
      </button>
    </div>

    <div className="product-rectangle5">
      <div className="product-list-row">
        <ul className="listFoodElement">
        <li>Бургеры</li>
        <li>Пицца</li>
        <li>Суши</li>
        <li>Дабстеп</li>
        <li>Марихуана</li>
    </ul>
      </div>
    </div>

<div className="checklist-products">
      <div className="product-info">
      <div className="product-picture"></div>
        <div className="product-title">{title}
        <p>Название товара</p></div>
        <div className="product-description">{description}
        <p>Описание товара</p>
        </div>
        <div className="product-controls">
          <button onClick={handleDecrement} disabled={count === 0}>
            -
          </button>
          <span>{count}</span>
            <div className="product-price">{price} ₽</div>
          <button onClick={handleIncrement}>+</button>
        </div>
      </div>

            <div className="product-info">
      <div className="product-picture"></div>
        <div className="product-title">{title}
        <p>Название товара</p></div>
        <div className="product-description">{description}
        <p>Описание товара</p>
        </div>
        <div className="product-controls">
          <button onClick={handleDecrement} disabled={count === 0}>
            -
          </button>
          <span>{count}</span>
            <div className="product-price">{price} ₽</div>
          <button onClick={handleIncrement}>+</button>
        </div>
      </div>
      </div>
        <button className="order-button" onClick={handleOrder} disabled={count === 0}> 
         {count} {price} ₽
        </button> 
    </div>
  );
}

export default ProductCard;