import React, { useState } from 'react';
import './accordionStyles.css';
import { AccordionItem } from './accordionItem.jsx';

const Accordion = ({ profileList = [] }) => {
  const [openID, setOpenID] = useState(null);

  return (
    <ul className="accordion">
      {profileList.map((profileItem, id) => (
        <AccordionItem
          key={id}
          profileItem={profileItem}
          isOpen={openID === id}
          onClick={() => setOpenID(openID === id ? null : id)}
        />
      ))}
    </ul>
  );
};

export default Accordion;