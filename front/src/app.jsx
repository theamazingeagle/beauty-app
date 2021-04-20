import ReactDOM from 'react-dom';
import React from 'react';
import { ClientsList } from './components/ClientsList';
import { AppContext } from './context';
import { API } from './api.js';

const api = new API();

const TabContent = ({ title, element }) => (
  <div className="tabcontent">
    <h3>{title}</h3>
    <div>{element}</div>
  </div>
);

function Tabs({ items }) {
  const [active, setActive] = React.useState(null);

  const openTab = e => setActive(+e.target.dataset.index);

  return (
    <div>
      <div className="tab">
        {items.map((n, i) => (
          <button
            onClick={openTab}
            data-index={i}
          >{n.title}</button>
        ))}
      </div>
      {items[active] && <TabContent {...items[active]} />}
    </div>
  );
}

const items = [
  { title: "Список клиентов", element: React.createElement(ClientsList) },
  { title: "Список услуг", element: 'Paris' },
  { title: "Список заказов", element: 'Tokyo' },
];

ReactDOM.render(
  <AppContext.Provider value={{ api }}>
    <Tabs items={items} />
  </AppContext.Provider>,
  document.getElementById("app")
)