const express = require("express");
const path = require("path");
const { env } = require("process");

const app = express();
const PORT = env.PORT || 3000;

// Serve the PDF file
app.get("/ResumeAfwan", (req, res) => {
  const filePath = path.join(__dirname, "../Afwan_Shaikh_-_Resume.pdf"); // Replace with your PDF file name
  res.sendFile(filePath);
});

// Serve the HTML file to display the PDF
app.get("/", (req, res) => {
  const fileName = path.join(__dirname, "./index.html");
  res.sendfile(fileName);
});

// Start the server
app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
