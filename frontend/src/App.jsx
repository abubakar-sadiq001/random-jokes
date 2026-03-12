import { useState } from "react";
import "./index.css";

function App() {
  const [responseData, setResponseData] = useState([]);
  const [loading, setLoading] = useState(false);
  const [errMessage, setErrMessage] = useState("");

  async function getJokes() {
    try {
      setLoading(true);
      const res = await fetch(
        "https://random-jokes-y36x.onrender.com/api/jokes",
      );

      if (!res.ok) {
        throw new Error("❌ Server failed to retreive joke");
      }

      const jokes = await res.json();
      setResponseData(jokes);
      setLoading(false);
    } catch (err) {
      console.log(err);
      setErrMessage(err);
    } finally {
      setLoading(false);
    }
  }

  return (
    <div style={{ textAlign: "center" }}>
      <div
        style={{
          display: "flex",
          flexDirection: "column",
          width: "30%",
          margin: "0 auto",
          paddingTop: "150px",
          gap: "20px",
        }}
      >
        <h3>Click the below button to see random jokes</h3>

        <p>{responseData?.setup || errMessage?.message}</p>

        <textarea
          disabled
          style={{ height: "50px" }}
          value={responseData?.punchline}
        />

        <button
          onClick={getJokes}
          id="btn"
          disabled={loading}
          style={{
            cursor: loading ? "not-allowed" : "pointer",
          }}
        >
          New Joke
        </button>
      </div>
    </div>
  );
}

export default App;
