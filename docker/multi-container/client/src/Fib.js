import React, { useState, useEffect, useCallback } from "react";
import axios from "axios";

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

  useEffect(() => {
    fetchValues();
    fetchIndexes();
  }, []);

  return null;
};
