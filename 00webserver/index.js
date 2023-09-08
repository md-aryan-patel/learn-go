const express = require("express");
const app = express();
const port = 8080;

app.use(express.json());
app.use(express.urlencoded({ extended: true }));

app.get("/", (req, res) => {
  res.status(200).send("Welcome to my live server");
});

app.get("/get", (req, res) => {
  res.status(200).json({ message: "Welcome to my live server" });
});

app.post("/post", (req, res) => {
  res.status(200).send(req.body);
});

app.post("/postform", (req, res) => {
  res.status(200).send(JSON.stringify(req.body));
});

// make the server listen to requests
app.listen(port, () => {
  console.log(`Server running at: http://localhost:${port}/`);
});
