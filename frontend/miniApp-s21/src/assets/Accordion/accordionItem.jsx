import React, { useRef, useEffect } from 'react';

export const AccordionItem = ({ profileItem, onClick, isOpen }) => {
  const itemRef = useRef(null);

  useEffect(() => {
    if (itemRef.current) {
    
      const timeout = setTimeout(() => {
        console.log(`Setting height for item ${profileItem.q}: ${isOpen ? itemRef.current.scrollHeight : 0}px`);
        itemRef.current.style.height = isOpen ? `${itemRef.current.scrollHeight}px` : '0px';
      }, 0);
      return () => clearTimeout(timeout);
    }
  }, [isOpen, profileItem.q]);

  return (
    <li className="accordion-item">
      <button className={`accordion-header ${isOpen ? 'active' : ''}`} onClick={onClick}>
        {profileItem.q}
        <i className={`bi ${isOpen ? 'bi-chevron-up' : 'bi-chevron-down'}`}></i>
      </button>
      <div
        className={`accordion-collapse ${isOpen ? 'open' : ''}`}
        style={{ height: isOpen ? 'auto' : '0px' }} 
      >
        <div className="accordion-body" ref={itemRef}>
          {isOpen && profileItem.details && (
            <div className="order-details">
              <ul className="order-items">
                {profileItem.details.items.map((item, index) => (
                  <li key={index}>{item.name} - {item.price} ₽</li>
                ))}
              </ul>
              <div className="order-total">Итого: {profileItem.details.total} ₽</div>
              <button className="repeat-order-btn">Повторить заказ</button>
            </div>
          )}
        </div>
      </div>
    </li>
  );
};