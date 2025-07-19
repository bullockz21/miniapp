import React from 'react';
import './accordionStyles.css';

const Accordion = ({ profileList = [] }) => {
  return (
    <ul className="orders-list">
      {profileList.map((profileItem, id) => (
        <li key={id} className="order-item">
          <div className="order-header">{profileItem.q}</div>
          <div className="order-details">
            <ul className="order-items">
              {profileItem.details.items.map((item, index) => (
                <li key={index}>{item.name} - {item.price} ₽</li>
              ))}
            </ul>
            <div className="order-footer">
              <div className="order-total">Итого: {profileItem.details.total} ₽</div>
              <button className="repeat-order-btn">Повторить заказ</button>
            </div>
          </div>
        </li>
      ))}
    </ul>
  );
};

export default Accordion;