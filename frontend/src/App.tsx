/* eslint-disable @typescript-eslint/no-explicit-any */
import { useEffect, useState } from "react";
import "./App.css";
import { connect, sendMsg } from "./api";
import Header from "./components/Header/Header";
import ChatHistory from "./components/Chat/ChatHistory";

function App() {
  const [history, setHistory] = useState<any[]>([]);
  const [message, setMessage] = useState<string>("");

  useEffect(() => {
    connect((msg: any) => {
      setHistory((prevMessages) => [...prevMessages, msg]);
    });
  }, []);

  return (
    <div className="card">
      <Header />
      <ChatHistory chatHistory={history} />
      <input
        onChange={(e: any) => setMessage(e.target.value)}
        value={message}
      />
      <button onClick={() => sendMsg(message)}>Send Message</button>
    </div>
  );
}

export default App;
