const express = require("express");
const io = require("socket.io-client");

// create instance server
const app = express();

// initialize web socket
var socket = io("http://127.0.0.1:5000", { path: "/v1/ws" });

// add listener
socket.on("connect", () => {
  console.log("your ws is connected");
});

socket.on("disconnect", () => {
  console.log("disconnect from ws");
});

// add path, so when user go this path send
// the message, to spesific number
app.get("/sendmessage", (req, res, next) => {
  let phoneNumber = req.param("phoneNumber");
  let message = req.param("message");

  // emit to the server go lang
  socket.emit(
    "chat message",
    JSON.stringify({
      phoneNumber,
      message
    })
  );
  socket.emit("test");
  socket.emit("ahah", "test");

  res.end();
});

app.listen(3030, () => {
  console.log("your server already up");
});
