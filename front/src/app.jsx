import ReactDOM from 'react-dom';
import React from 'react';
import { ClientsList } from './components/ClientsList';
import { AppContext } from './context';
import { API } from './api.js';

const api = new API();

ReactDOM.render(
    <AppContext.Provider value={{ api }}>
        <ClientsList />
    </AppContext.Provider>,
    document.getElementById("app")
)