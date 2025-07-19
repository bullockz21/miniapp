import React from 'react';
import './addressAccordionStyles.css';

const AddressAccordion = ({ addressList = [] }) => {
  return (
    <ul className="addresses-list">
      {addressList.map((addressItem, id) => (
        <li key={id} className="address-item">
          <div className="address-header">{addressItem.name || `Адрес ${id + 1}`}</div>
          <div className="address-details">
            <div className="address-info">{addressItem.fullAddress || 'Полный адрес не указан'}</div>
            <div className="address-footer">
              <div className="address-notes">{addressItem.notes || 'Заметки отсутствуют'}</div>
              <button className="select-address-btn">Выбрать адрес</button>
            </div>
          </div>
        </li>
      ))}
    </ul>
  );
};

export default AddressAccordion;