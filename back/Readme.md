1. Получить список автомобилей 
GET hostname/Client/get
Accept: application/vnd.api+json
Content-Type: application/vnd.api+json
2. Получить информацию об одном авто
GET hostname/Client/get/{id}
Accept: application/vnd.api+json
Content-Type: application/vnd.api+json
3. Создать запись
POST hostname/Client/create
Accept: application/vnd.api+json
Content-Type: application/vnd.api+json
{
  "data":
    {
        "type":"Clients",
        "attributes":{"brand":"Kamaz","mileage":1000000,"model":"6969","price":15,"status":2}
    }
}
4. Обновить запись
PATCH hostname/Client/update
Accept: application/vnd.api+json
Content-Type: application/vnd.api+json
{
  "data":[
    {
      "type":"Clients",
      "id":"1",
      "attributes":{
        "brand":"Renault","mileage":200,"model":"TrucksT","price":100000,"status":2}
      }
  ]
}
5. Удалить запись
DELETE hostname/Client/delete
Accept: application/vnd.api+json
Content-Type: application/vnd.api+json
{
  "data":[
    {
      "type":"Clients",
      "id":"1",
      "attributes":{
        "brand":"Renault","mileage":200,"model":"TrucksT","price":100000,"status":2}
      }
  ]
}