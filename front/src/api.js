export class API {
    baseURL = 'http://localhost:8080';

    getClients() {
        return fetch(
            this.baseURL + '/api/client/get', {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        });
    }

    createClient(client) {
        return fetch(
            this.baseURL + '/api/client/create', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(client)
        });
    }

    getClient(id) {
        return fetch(
            this.baseURL + '/api/client/get/' + id, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        });
    }

    updateClient(client) {
        return fetch(
            this.baseURL + '/api/client/update', {
            method: 'PATCH',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(client)
        });
    }

    deleteClient(id) {
        return fetch(
            this.baseURL + '/api/client/delete/' + id, {
            method: 'DELETE',
            headers: { 'Content-Type': 'application/json' },
        });
    }

    getServices() {
        return fetch(
            this.baseURL + '/api/service/get', {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        });
    }

    createService(service) {
        return fetch(
            this.baseURL + '/api/service/create', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(service)
        });
    }

    getService(id) {
        return fetch(
            this.baseURL + '/api/service/get/' + id, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        });
    }

    updateService(service) {
        return fetch(
            this.baseURL + '/api/service/update', {
            method: 'PATCH',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(service)
        });
    }

    deleteService(id) {
        return fetch(
            this.baseURL + '/api/service/delete/' + id, {
            method: 'DELETE',
            headers: { 'Content-Type': 'application/json' },
        });
    }

    getOrders() {
        return fetch(
            this.baseURL + '/api/order/get', {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        });
    }

    createOrder(order) {
        return fetch(
            this.baseURL + '/api/order/create', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(order)
        });
    }

    getOrder(id) {
        return fetch(
            this.baseURL + '/api/order/get/' + id, {
            method: 'GET',
            headers: { 'Content-Type': 'application/json' },
        });
    }

    updateOrder(order) {
        return fetch(
            this.baseURL + '/api/order/update', {
            method: 'PATCH',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(order)
        });
    }

    deleteOrder(id) {
        return fetch(
            this.baseURL + '/api/order/delete/' + id, {
            method: 'DELETE',
            headers: { 'Content-Type': 'application/json' },
        });
    }

}
