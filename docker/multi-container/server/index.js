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

pgClient.query("CREATE TABLE IF NOT EXISTS values (number INT);").catch(err => {
  console.log(err);
});

// Redis configuration
const redisClient = redis.createClient({
  host: keys.redisHost,
  port: keys.redisPort,
  retry_strategy: () => 1000
});
const redisPublisher = redisClient.duplicate();

// Express route handlers
app.get("/", (req, res) => {
  return res.send("hi");
});

app.get("/values/all", async (req, res) => {
  const values = await pgClient.query("SELECT * FROM values;");
  return res.send(values.rows);
});

app.get("/values/current", (req, res) => {
  redisClient.hgetall("values", (err, values) => {
    if (err) return err;
    return res.send(values);
  });
});

app.post("/values", async (req, res) => {
  const index = req.body["index"];
  if (parseInt(index) > 40) {
    return res.status(422).send("Index is too high");
  }
  redisClient.hset("values", index, "Nothing yet!");
  redisPublisher.publish("insert", index);
  pgClient.query("INSERT INTO values(number) VALUES $1", [index]);

  return res.send({ working: true });
});

app.listen(5000, err => console.log("Listening..."));
