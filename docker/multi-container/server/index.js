const keys = require("./keys");
const express = require("express");
const bodyParser = require("body-parser");
const cors = require("cors");
const { Pool } = require("pg");
const redis = require("redis");

// Express app setup
const app = express();
app.use(cors());
app.use(bodyParser.json());

// Postgres setup
const pgClient = new Pool({
  host: keys.pgHost,
  port: keys.pgPort,
  user: keys.pgUser,
  database: keys.pgDatabase,
  password: keys.pgPassword
});

pgClient.on("error", () => {
  console.log("Lost PG connection");
});

pgClient.query("CREATE TABLE IF NOT EXISTS values (number INT)").catch(err => {
  console.log(err);
});
