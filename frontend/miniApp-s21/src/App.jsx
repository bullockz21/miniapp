import React, { useReducer, useEffect, useState } from 'react';
import './App.css';
import './ProductCardMain.css';
import './PCAdaptMain.css';
import './assets/ProfileStyles/ProductCardProfile.css';
import './assets/TrashBox/trashboxstyles.css';
import './assets/Contacts/contactsStyles.css';
import Accordion from './assets/Accordion/accordionOrders.jsx';
import AddressAccordion from './assets/addressAccordion/addressAccordion.jsx';
/* import TrashBox from './assets/TrashBox/trashbox.jsx'; */

const profileList = [
  {
    q: "12 апреля, 14:35",
    details: {
      items: [
        { name: "Пицца 'Маргарита' x1", price: 1 },
        { name: "Кола 0.5 л x2", price: 2 },
      ],
      total: 3,
    },
  },
  {
    q: "13 апреля, 15:30",
    details: {
      items: [
        { name: "Суши сет x1", price: 5 },
        { name: "Чай x1", price: 1 },
      ],
      total: 6,
    },
  },
];

const addressList = [
  { name: "Дом", fullAddress: "ул. Федора Попова, 1, кв. 5", notes: "Основной адрес" },
  { name: "Работа", fullAddress: "пр. Ленина, 10, офис 305", notes: "Доставка до 18:00" },
];

function ProductCard({ title, description, price, image }) {
  const [isOrdersOpen, setIsOrdersOpen] = useState(false);
  const [isAddressOpen, setIsAddressOpen] = useState(false);
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

  const toggleAddress = () => {
    setIsAddressOpen(!isAddressOpen);
  };

  return (
    <div className="product-card">
      <div
        className="product-header"
        style={{ backgroundImage: `url(${image})` }}
      >
        <div className={`product-rectangle ${activeTab === 'home' ? 'active' : ''}`}
          onClick={() => handleTabClick('home')}>
          <i className="bi bi-house"></i>
          <p className="pro-rect-title">Главная</p>
        </div>
        <div className={`product-rectangle ${activeTab === 'profile' ? 'active' : ''}`}
          onClick={() => handleTabClick('profile')}>
          <i className="bi bi-person"></i>
          <p className="pro-rect-title">Профиль</p>
        </div>
        <div className={`product-rectangle ${activeTab === 'trash' ? 'active' : ''}`}
          onClick={() => handleTabClick('trash')}>
          <i className="bi bi-cart"></i>
          <p className="pro-rect-title">Корзина</p>
        </div>
        <div className={`product-rectangle ${activeTab === 'contacts' ? 'active' : ''}`}
          onClick={() => handleTabClick('contacts')}>
          <i className="bi bi-telephone"></i>
          <p className="pro-rect-title">Контакты</p>
        </div>
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
                <i className="bi bi-chevron-right"></i>
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
                <p>Название товара</p>
              </div>
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
                <p>Название товара</p>
              </div>
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
            <div>
              <button className="return-to-main" onClick={() => handleTabClick('home')}>
                <i className="bi bi-chevron-compact-left"></i>
      
                <p>Продолжить покупки</p>
              </button>
            </div>
            <h2>Профиль</h2>
          </div>
          <div className="my-orders-data" onClick={toggleOrders}>
            <p>Мои заказы</p>
            <i className={`bi ${isOrdersOpen ? 'bi-chevron-down' : 'bi-chevron-right'}`}></i>
          </div>
          {isOrdersOpen && <Accordion profileList={profileList} />}
          <div className="my-addresses-data" onClick={toggleAddress}>
            <p>Адрес</p>
            <i className={`bi ${isAddressOpen ? 'bi-chevron-down' : 'bi-chevron-right'}`}></i>
          </div>
          {isAddressOpen && <AddressAccordion addressList={addressList} />}
        </>
      )}

         {activeTab === 'trash' && ( 
        <>
          <div className="trashbox-data-container">


              <div className="hey">
              <button className="return-to-main" onClick={() => handleTabClick('home')}>
                <i className="bi bi-chevron-compact-left"></i>
                <p>Продолжить покупки</p>
              </button>
            </div>
                        
        <div className="trashbox-added-item">
          <div className="item-imagecard">

          </div>
          <div className="column-name-price">
          <div className="item-header">
            <p>Название товара</p>
            <h3>Цена</h3>
          </div>

           <div className="product-controls-trash">
                <button onClick={handleDecrement} disabled={count === 0}>
                  -
                </button>
                <span>{count}</span>
                <div className="product-price">{price}</div>
                <button onClick={handleIncrement}>+</button>
              </div>
              </div>
          </div>

             <div className="trashbox-added-item">
          <div className="item-imagecard">

          </div>
          <div className="column-name-price">
          <div className="item-header">
            <p>Название товара</p>
            <h3>Цена</h3>
          </div>

           <div className="product-controls-trash">
                <button onClick={handleDecrement} disabled={count === 0}>
                  -
                </button>
                <span>{count}</span>
                <div className="product-price">{price}</div>
                <button onClick={handleIncrement}>+</button>
              </div>
              </div>
          </div>


             <div className="trashbox-added-item">
          <div className="item-imagecard">

          </div>
          <div className="column-name-price">
          <div className="item-header">
            <p>Название товара</p>
            <h3>Цена</h3>
          </div>

           <div className="product-controls-trash">
                <button onClick={handleDecrement} disabled={count === 0}>
                  -
                </button>
                <span>{count}</span>
                <div className="product-price">{price}</div>
                <button onClick={handleIncrement}>+</button>
              </div>
              </div>
          </div>

          <div className="summary-price-count">
            <p>В корзине 3 товара</p>
            <h3>Итого 3 ₽</h3>
            <p>Минимальная сумма заказа 1 ₽</p>
          </div>
          <button className="launch-to-delivery">
            <p>К оформлению</p>
          </button>
          </div>
      </>
      )}

      {activeTab === 'contacts' && ( 
        <>

          <div className="container-contacts">

            <div className="hey">
              <button className="return-to-main" onClick={() => handleTabClick('home')}>
                <i className="bi bi-chevron-compact-left"></i>
                <p>Продолжить покупки</p>
              </button>
            </div>

              <div className="contact-item">
                <button className="Delivery">
                  <p>Контакты доставки</p>
                </button>
  </div>

     <div className="contact-item">
                <button className="Telegram">
                  <p>Telegram</p>
                </button>
</div>

  <div className="contact-item">
                <button className="Whatsapp">
                  <p>Whatsapp</p>
                </button>
</div>

            
          </div>
        </>
      )}
    </div>

  );
};

export default ProductCard;