export const ReceiveRe = () => {
  const socket = new WebSocket("ws://localhost:8080/wbsocket");
  socket.onopen = function (event) {
    console.log("Connected to server");
  };
  socket.onmessage = function (event) {
    console.log(event.data);
  };
  socket.onclose = function (event) {
    console.log("Disconnected from server");
  };
  socket.onerror = function (event) {
    console.log("Error from server");
  };
};

// let conn = socket()
export const SendReponseToBack = (methode, message, conn, idReceiver) => {
    conn.onopen = (event) => {
      console.log("Connected to server");
  
      try {
        const responseData = {
          type: methode,
          content: methode === "GET" ? "ok" : message,
          statut: "connected",
        };
  
        if (methode !== "GET") {
          responseData.Receiver = idReceiver; // En supposant que idReceiver est défini ailleurs
        }
  
        conn.send(JSON.stringify(responseData));
        conn.onmessage = (event) => {
            console.log(event.data);
            
        }
      } catch (error) {
        console.error("Error sending response:", error);
      }
    };
  };
  
