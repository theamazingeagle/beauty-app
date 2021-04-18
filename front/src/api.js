export class API {
  baseURL = 'http://localhost:8080';

  getClients() {
    return fetch(
      this.baseURL + '/api/client/get', {
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
}
