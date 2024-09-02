/* eslint-disable @typescript-eslint/no-explicit-any */
import "./ChatHistory.scss";

type PropType = {
  chatHistory: any[];
};

const ChatHistory = ({ chatHistory }: PropType) => {
  return (
    <div className="ChatHistory">
      <h2>Chat History</h2>
      {chatHistory.map((msg: any, index: number) => {
        return <p key={index}>{msg.data}</p>;
      })}
    </div>
  );
};

export default ChatHistory;
