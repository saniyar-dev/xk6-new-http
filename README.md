# xk6-new-http

This extension goal is to change the way users used to make HTTP connection with k6.

As you can see in [k6 repo](https://github.com/grafana/k6) there is [lots of issues](https://github.com/grafana/k6/issues?q=is%3Aissue%20state%3Aopen%20label%3Anew-http) with old/standard [HTTP API](https://t.me/saniyar_krmi) implemented in k6 originally.
So they came with an idea to design a complete new HTTP API (you can see [the design document here](https://github.com/grafana/k6/blob/master/docs/design/018-new-http-api.md)).

### What i'm doing?
I couldn't find an implementation for [the design document here](https://github.com/grafana/k6/blob/master/docs/design/018-new-http-api.md), so i decided to implement it myself and contribute to the great k6.
I need you to know that this project is a working progress and also my hubby as a side project, so don't expect something unordinary great from it. I hope some day it will be a great implementation of [phase 1](https://github.com/grafana/k6/issues/3038), but today is not the day.
Hope you enjoy it!

## Requirements

- [Goland 1.20+](https://go.dev/)
- [Git](https://git-scm.com/)
- [xk6](https://github.com/grafana/xk6) (`go install go.k6.io/xk6/cmd/xk6@latest`)

## Getting started

1. Build the k6 binary:
`make build`

2. Run an example:
`./k6 run ./examples/test.js`

## Usage/Examples

- Using a client with default transport settings, and making a GET request:
```javascript
import { Client } from 'k6/x/net/http';

export default async function () {
  const client = new Client();
  const response = await client.get('https://httpbin.test.k6.io/get');
  const jsonData = await response.json();
  console.log(jsonData);
}
```
***
- Creating a client with custom transport settings, some HTTP options, and making a POST request:
  
  > This example is on todo list and doesn't work now
```javascript
import { TCP } from 'k6/x/net';
import { Client } from 'k6/x/net/http';

export default async function () {
  const client = new Client({
    dial: async address => {
      return await TCP.open(address, { keepAlive: true });
    },
    proxy: 'https://myproxy',
    headers: { 'User-Agent': 'k6' },  // set some global headers
  });
  await client.post('http://10.0.0.10/post', {
    json: { name: 'k6' }, // automatically adds 'Content-Type: application/json' header
  });
}
```
***
- see `examples` dir for more examples
  
  > Some examples are on todo list but i will deliver them very fast don't wory.

## Contributing

Contributions are always welcome!

You can fork this project, create your branch, work on it, document your work and reasons behind it, contact me on telegram and do a pull request.

## Support

For support, email saniyar.dev@gmail.com or join my [telegram group]().


## 🔗 Links
[![linkedin](https://img.shields.io/badge/linkedin-0A66C2?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/saniyar-karami-818771231/)
[![twitter](https://img.shields.io/badge/twitter-1DA1F2?style=for-the-badge&logo=twitter&logoColor=white)](https://twitter.com/)

