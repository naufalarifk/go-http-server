// curl -X POST http://localhost:42069/coffee -H 'Content-Type: application/json' -d '{"type": "dark mode", "size": "medium"}'
//this is for testing purposes only, idk why i use js
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

makeRequest("/", body).then((response) => {
  console.log("response", response);
});
