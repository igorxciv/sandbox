import React, { useState, useEffect, useCallback } from "react";
import axios from "axios";

const SeenIndexes = ({ indexes }) => {
  return indexes.map(({ number }) => number).join(", ");
};

const CalculatedValues = ({ values }) => {
  const entries = [];
  for (const key in values) {
    entries.push(
      <div key={key}>
        For index {key} I calculated {values[key]}
      </div>
    );
  }
  return entries;
};

export const Fib = () => {
  const [seenIndexes, setSeenIndexes] = useState([]);
  const [values, setValues] = useState({});
  const [index, setIndex] = useState("");

  const fetchValues = async () => {
    const vals = await axios.get("/api/values/current");
    setValues(vals.data);
  };

  const fetchIndexes = async () => {
    const seen = await axios.get("/api/values/all");
    setSeenIndexes(seen);
  };

  const handleSubmit = async event => {
    event.preventDefault();

    await axios.post("/api/values", { index });
    setIndex("");
  };

  useEffect(() => {
    fetchValues();
    fetchIndexes();
  }, []);

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <label>Enter your index:</label>
        <input onChange={e => setIndex(e.target.value)} value={index} />
        <button>Submit</button>
      </form>

      <h3>Indexes I have seen</h3>
      <SeenIndexes indexes={seenIndexes} />

      <h3>Calculated values</h3>
      <CalculatedValues values={values} />
    </div>
  );
};
