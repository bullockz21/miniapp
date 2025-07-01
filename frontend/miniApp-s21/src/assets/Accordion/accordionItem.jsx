import React, { useRef, useEffect } from 'react';

export const AccordionItem = ({ profileItem, onClick, isOpen }) => {
  const itemRef = useRef(null);

  // Устанавливаем высоту после монтирования
  useEffect(() => {
    if (itemRef.current) {
      itemRef.current.style.height = isOpen ? `${itemRef.current.scrollHeight}px` : '0px';
    }
  }, [isOpen]);

  return (
    <li className="accordion-item">
      <button className="accordion-header" onClick={onClick}>
        {profileItem.q}
        <i className={`bi ${isOpen ? 'bi-chevron-up' : 'bi-chevron-down'}`}></i>
      </button>
      <div
        className="accordion-collapse"
        style={{ height: isOpen ? `${itemRef.current?.scrollHeight || 0}px` : '0px' }}
      >
        <div className="accordion-body" ref={itemRef}>
          {profileItem.a}
        </div>
      </div>
    </li>
  );
};