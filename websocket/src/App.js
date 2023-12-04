import React, { useEffect, useState } from 'react';

const WebSocketPage = () => {
  const [dataList, setDataList] = useState([]);

  useEffect(() => {
    const socket = new WebSocket('ws://localhost:8090/push-list');

    socket.onopen = () => {
      console.log('WebSocket connected');
    };

    socket.onmessage = (event) => {
      const newData = JSON.parse(event.data);
      setDataList((prevDataList) => [...prevDataList, newData]);
    };

    socket.onclose = () => {
      console.log('WebSocket disconnected');
    };

    return () => {
      socket.close();
    };
  }, []);

  const handleReject = (id) => {
    fetch('http://localhost:9090/reject', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        id: id,
        action: 'reject',
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log('Rejection response:', data);
      })
      .catch((error) => {
        console.error('Error rejecting:', error);
      });
  };

  const handleApprove = (id) => {
    fetch('http://localhost:9090/reject', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        id: id,
        action: 'approve',
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log('Approval response:', data);
      })
      .catch((error) => {
        console.error('Error approving:', error);
      });
  };

  return (
    <div
      style={{
        display: 'flex',
        justifyContent: 'center',
        border: '1px solid black',
        height: '100vh',
        overflow: 'auto',
      }}
    >
      <div
        style={{
          width: '100%',
          maxWidth: '600px',
          padding: '20px',
        }}
      >
        <h1>List Page</h1>
        {dataList.map((data) => (
          <div
            key={data.id}
            style={{
              margin: '10px',
              padding: '10px',
              border: '1px solid gray',
              position: 'relative',
            }}
          >
            <p>Data ID: {data.id}</p>
            <p>Data Content: {data.content}</p>
            <div
              style={{
                position: 'absolute',
                bottom: '10px',
                right: '10px',
              }}
            >
              <button onClick={() => handleReject(data.id)}>Reject</button>
              <button onClick={() => handleApprove(data.id)}>Approve</button>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default WebSocketPage;