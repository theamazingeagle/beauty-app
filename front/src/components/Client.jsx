import React from 'react';
import { AppContext } from '../context';

export class Client extends React.Component {
  static contextType = AppContext;
  constructor(props) {
    super(props);
    this.state = { items: [], newname: '' };
    this.handleChange = this.handleChange.bind(this);
  }
  componentDidMount() {
    this.fetchClients();
  }
  fetchClients() {
    this.context.api.getClients()
      .then(response => response.json())
      .then(items => { this.setState({ items, errText: undefined }) })
      .catch(err => {
        let errText = "criticalError";
        if (err instanceof Error) {
          errText = err.message;
        }
        this.setState({ errText, items: [] });
      });
  }

  handleChange(event) {
    this.setState({ newname: event.target.value });
  }

  render() {
    return (
      <div>
        {this.state.errText}
        {!!this.state.errText && <button onClick={() => { this.fetchClients() }} >tryAgain</button>}
        <div>
          Имя нового клиента:<input type="text" value={this.state.newname} onChange={this.handleChange} />
          <button onClick={() => {
            this.context.api.createClient({ name: this.state.newname }).then(() => { this.fetchClients() }).catch(err => {
              let errText = "criticalError";
              if (err instanceof Error) {
                errText = err.message;
              }
              alert(errText);
            })
          }}>Добавить
          </button>
        </div>
        <ul>
          {
            this.state.items.map((item) => {
              return <div>
                <input onClick={() => {
                  let userName = prompt("Edit...", item.name);
                  this.context.api.updateClient({ ...item, name: userName })
                    .then(() => { this.fetchClients() }).catch(err => {
                      let errText = "crticalError";
                      if (err instanceof Error) {
                        errText = err.message;
                      }
                      alert(errText);
                    })
                }} key={item.id} value={item.name} />

                <button onClick={() => {
                  this.context.api.deleteClient(item.id).then(() => { this.fetchClients() }).catch(err => {
                    let errText = "crticalError";
                    if (err instanceof Error) {
                      errText = err.message;
                    }
                    alert(errText);
                  })
                }}>
                  x
                </button>
              </div>
            })
          }
        </ul>
      </div >);
  }
}
