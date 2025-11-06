// curl -X POST http://localhost:42069/coffee -H 'Content-Type: application/json' -d '{"type": "dark mode", "size": "medium"}'
//this is for testing purposes only, idk why i use js

import net from "net";
import http from "http";
const body = {
  type: "dark mode",
  size: "medium",
  from: "javascript",
};

function makeRequest(path = "/coffee", body) {
  const url = new URL("http://localhost:42069/");
  url.pathname = path;

  const data = fetch(url.toString(), {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  });
  return data;
}

// makeRequest("/httpbin/stream/100", body).then((response) => {
//   console.log("response", response);
// });

// const request = `GET /httpbin/stream/100 HTTP/1.1\r\nHost: localhost:42069\r\nConnection: close\r\n\r\n`;

// const client = net.createConnection({ port: 42069, host: "localhost" }, () => {
//   console.log("Connected to server");
//   client.write(request);
// });

// client.on("data", (data) => {
//   process.stdout.write(data.toString()); // Stream response chunks
// });

// client.on("end", () => {
//   console.log("\nConnection closed by server");
// });

// client.on("error", (err) => {
//   console.error("Error:", err.message);
// });

// console.log(client);

const options = {
  hostname: "localhost",
  port: 42069,
  path: "/httpbin/stream/100",
  method: "GET",
  headers: {
    Host: "localhost:42069",
    Connection: "close",
  },
};

const req = http.request(options, (res) => {
  res.setEncoding("utf8");
  res.on("data", (chunk) => {
    process.stdout.write(chunk);
  });
  res.on("end", () => {
    console.log("\nStream ended");
  });
});

req.on("error", (err) => {
  console.error("Request error:", err.message);
});

req.end();
