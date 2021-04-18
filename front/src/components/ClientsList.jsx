import React from 'react';
import { AppContext } from '../context';

export class ClientsList extends React.Component {
  static contextType = AppContext;
  constructor(props) {
    super(props);
    this.state = { items: [] };
  }
  componentDidMount() {
    this.fetchClients();
  }
  fetchClients() {
    this.context.api.getClients()
      .then(response => response.json())
      .then(items => { this.setState({ items, errText: undefined }) })
      .catch(err => {
        let errText = "crticalError";
        if (err instanceof Error) {
          errText = err.message;
        }
        this.setState({ errText, items: [] });
      });
  }

  render() {
    return (
      <div>
        {this.state.errText}
        {!!this.state.errText && <button onClick={() => { this.fetchClients() }} >tryAgain</button>}
        <ul>
          {
            this.state.items.map((item) => {
              return <li onClick={() => {
                let userName = prompt("ololo");
                this.context.api.updateClient({ ...item, name: userName })
                  .then(() => { this.fetchClients() }).catch(err => {
                    let errText = "crticalError";
                    if (err instanceof Error) {
                      errText = err.message;
                    }
                    alert(errText);
                  })
              }} key={item.id}>
                {item.name}
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
              </li>
            })
          }
        </ul>
      </div>);
  }
}
