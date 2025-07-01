import React, { useReducer, useEffect, useState } from 'react';
import './App.css';
import './ProductCardMain.css';
import './PCAdaptMain.css';
import './assets/ProfileStyles/ProductCardProfile.css';
import Accordion from './assets/Accordion/accordionOrders.jsx';

const profileList = [
  {
    q: "Lorem upsum dolor sit amet",
    a: "Lorem upsum dolor sit amet consectetur adipisicing elit. Eius accusamus cumque"
  },
  {
    q: "Lorem upsum dolor sit amet",
    a: "Lorem upsum dolor sit amet consectetur adipisicing elit. Eius accusamus cumque"
  },
  {
    q: "Lorem upsum dolor sit amet",
    a: "Lorem upsum dolor sit amet consectetur adipisicing elit. Eius accusamus cumque"
  },
  {
    q: "Lorem upsum dolor sit amet",
    a: "Lorem upsum dolor sit amet consectetur adipisicing elit. Eius accusamus cumque"
  },
  {
    q: "Lorem upsum dolor sit amet",
    a: "Lorem upsum dolor sit amet consectetur adipisicing elit. Eius accusamus cumque"
  },
  {
    q: "Lorem upsum dolor sit amet",
    a: "Lorem upsum dolor sit amet consectetur adipisicing elit. Eius accusamus cumque"
  },
]

function ProductCard({ title, description, price, image }) {
  const [count, setCount] = useState(0);

  useEffect(() => {
  if (window.Telegram?.WebApp) {
    window.Telegram.WebApp.ready();
    window.Telegram.WebApp.expand();
    const viewportWidth = window.Telegram.WebApp.viewportWidth;
    document.documentElement.style.setProperty('--viewport-width', `${viewportWidth}px`);
    document.body.style.width = 'var(--viewport-width)';
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

  const [activeTab, setActiveTab] = useState('home'); 

  const handleTabClick = (tab) => {
    setActiveTab(tab);
  };

  const toggleOrders = () => {
    setIsOrdersOpen(!isOrdersOpen);
  };

  return (
    <div className="product-card">
      <div
        className="product-header"
        style={{ backgroundImage: `url(${image})` }}
    >
      <div className={`product-rectangle ${activeTab === 'home' ? 'active' : ''}`}
          onClick={() => handleTabClick('home')}>
      <i class="bi bi-house"></i>
        <p className="pro-rect-title">Главная</p> </div>
      <div className={`product-rectangle ${activeTab === 'profile' ? 'active' : ''}`}
          onClick={() => handleTabClick('profile')}> 
      <i class="bi bi-person"></i>
      <p className="pro-rect-title">Профиль</p></div>
       <div className="product-rectangle"> 
       <i class="bi bi-cart"></i>
       <p className="pro-rect-title">Корзина</p></div>
        <div className="product-rectangle">
        <i class="bi bi-telephone"></i>
         <p className="pro-rect-title">Контакты</p></div>
      </div>


{activeTab === 'home' && ( 
  <>
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
        </>
        )}

        {activeTab === 'profile' && (
       <>
          <div className="card-profile-data">

          <p>Здесь будет информация о профиле пользователя.</p>
          <div>        
            <button className='return-to-main'>
              <i class="bi bi-chevron-compact-left"></i>
              <p>Продолжить покупки</p>
            </button>
            </div>
          <h2>Профиль</h2>
        </div>
<div className="my-orders-data" onClick={toggleOrders}>
            <p>Мои заказы</p>
            <i className={`bi ${isOrdersOpen ? 'bi-chevron-up' : 'bi-chevron-down'}`}></i>
            {isOrdersOpen && <Accordion profileList={profileList} />}
          </div>
        </>
      )}
    </div>
  );
};

export default ProductCard;