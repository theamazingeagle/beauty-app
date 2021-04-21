var React = require('react');
import { AppContext } from '../context';

export class Service extends React.Component {
  static contextType = AppContext;

  constructor(props) {
    super(props);
    this.state = {};
    this.state.filterText = "";
    this.state.items = [];
  }

  componentDidMount() {
    this.fetchServices();
  }

  fetchServices() {
    this.context.api.getServices()
      .then(response => response.json())
      .then(items => { this.setState({ items }) })
  }

  handleUserInput(filterText) {
    this.setState({ filterText: filterText });
  }

  handleRowDel(service) {
    var index = this.state.items.indexOf(service);
    this.state.items.splice(index, 1);
    this.setState(this.state.items);
    //console.log(service)
    this.context.api.deleteService(service.id);
  }

  handleAddEvent(evt) {
    var id = 0;//(+ new Date() + Math.floor(Math.random() * 999999)).toString(36);
    var service = {
      id: id,
      name: "",
      cost: ""
    }
    this.state.items.push(service);
    this.setState(this.state.items);
  }

  handleServiceTable(evt) {
    var item = {
      id: evt.target.id,
      name: evt.target.name,
      value: evt.target.value
    };
    //console.log(item);
    var items = this.state.items;
    let newServiceData;
    var newitems = items.map(function (service) {
      for (var key in service) {
        if (key == item.name && service.id == item.id) {
          service[key] = item.value;
          newServiceData = service;
        }
      }
      return service;
    });
    this.setState(newitems);
    //console.log(newServiceData);
    // if (newServiceData.id !== 0) {
    //   this.context.api.updateService({ id: newServiceData.id, name: newServiceData.name, cost: parseInt(newServiceData.cost, 10) });
    // }
    // if (newServiceData.id === 0) {
    //   this.context.api.createService({ id: newServiceData.id, name: newServiceData.name, cost: parseInt(newServiceData.cost, 10) });
    // }
  };
  render() {
    return (
      <div>
        <SearchBar filterText={this.state.filterText} onUserInput={this.handleUserInput.bind(this)} />
        <ServiceTable onServiceTableUpdate={this.handleServiceTable.bind(this)} onRowAdd={this.handleAddEvent.bind(this)} onRowDel={this.handleRowDel.bind(this)} items={this.state.items} filterText={this.state.filterText} />
      </div>
    );
  }
}

class SearchBar extends React.Component {
  handleChange() {
    this.props.onUserInput(this.refs.filterTextInput.value);
  }
  render() {
    return (
      <div>
        <input type="text" placeholder="Search..." value={this.props.filterText} ref="filterTextInput" onChange={this.handleChange.bind(this)} />
      </div>
    );
  }
}

class ServiceTable extends React.Component {
  render() {
    var onServiceTableUpdate = this.props.onServiceTableUpdate;
    var rowDel = this.props.onRowDel;
    var filterText = this.props.filterText;
    var service = this.props.items.map(function (service) {
      if (service.name.indexOf(filterText) === -1) {
        return;
      }
      return (<ServiceRow onServiceTableUpdate={onServiceTableUpdate} service={service} onDelEvent={rowDel.bind(this)} key={service.id} />)
    });
    return (
      <div>
        <button type="button" onClick={this.props.onRowAdd} className="btn btn-success pull-right">Add</button>
        <table className="table table-bordered">
          <thead>
            <tr>
              <th>Название</th>
              <th>Стоимость</th>
            </tr>
          </thead>
          <tbody>
            {service}
          </tbody>
        </table>
      </div>
    );
  }
}

class ServiceRow extends React.Component {
  onDelEvent() {
    this.props.onDelEvent(this.props.service);
  }
  render() {
    return (
      <tr className="eachRow">
        <EditableCell onServiceTableUpdate={this.props.onServiceTableUpdate} cellData={{
          "type": "name",
          value: this.props.service.name,
          id: this.props.service.id
        }} />
        <EditableCell onServiceTableUpdate={this.props.onServiceTableUpdate} cellData={{
          type: "cost",
          value: this.props.service.cost,
          id: this.props.service.id
        }} />
        <td className="del-cell">
          <input type="button" onClick={this.onDelEvent.bind(this)} value="X" className="del-btn" />
        </td>
      </tr>
    );
  }
}

class EditableCell extends React.Component {
  render() {
    return (
      <td>
        <input
          type='text'
          name={this.props.cellData.type}
          id={this.props.cellData.id}
          value={this.props.cellData.value}
          onChange={this.props.onServiceTableUpdate}
        />
      </td>
    );
  }
}
