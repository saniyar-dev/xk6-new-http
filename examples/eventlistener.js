import { Client } from 'k6/x/net/http';

export default async function () {
  const client = new Client();
  let requestID;  // used to correlate a specific response with the request that initiated it

  client.on('requestToBeSent', event => {
    const request = event.data;
    if (!requestID && request.url == 'https://httpbin.test.k6.io/get?name=k6'
        && request.method == 'GET') {
      // The request ID is a UUIDv4 string that uniquely identifies a single request.
      // This is a contrived check and example, but you can imagine that in a complex
      // script there would be many similar requests.
      requestID = request.id;
    }
  });

  client.on('responseReceived', event => {
    const response = event.data;
    if (requestID && response.request.id == requestID) {
      // Change the request duration metric to any value
      response.metrics['http_req_duration'].value = 3.1415;
      // Consider the request successful regardless of its response
      response.metrics['http_req_failed'].value = false;
      // Or drop a single metric
      delete response.metrics['http_req_duration'];
      // Or drop all metrics
      response.metrics = {};
    }
  });

  await client.get('https://httpbin.test.k6.io/get', { query: { name: 'k6' } });
}
