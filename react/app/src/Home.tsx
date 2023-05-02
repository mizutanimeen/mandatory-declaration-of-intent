import React from 'react';
import { useNavigate } from "react-router-dom";
import axios from 'axios';

function Home() {
  const [roomName, setRoomName] = React.useState('');
  const [roomDescription, setRoomDescription] = React.useState('');
  const navigate = useNavigate();
  const createRoomURL = (process.env.REACT_APP_GO_URL ?? "") + (process.env.REACT_APP_GO_PATH ?? "") + '/rooms'

  const roomNameChange = (e: { target: { value: React.SetStateAction<string>; }; }) => {
    setRoomName(e.target.value);
  }

  const roomDescriptionChange = (e: { target: { value: React.SetStateAction<string>; }; }) => {
    setRoomDescription(e.target.value);
  }

  const createRoom = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (roomName == '' || roomDescription == '') {
      return;
    }

    axios.post(createRoomURL, {
      Name: roomName,
      Description: roomDescription
    })
      .then((response) => {
        navigate("/rooms/" + response.data);
      });
  };


  return (
    <>
      <h1>部屋作成</h1>
      <form onSubmit={(e) => { createRoom(e) }}>
        <input type="text" onChange={(e) => { roomNameChange(e) }} placeholder="名前" />
        <textarea onChange={(e) => { roomDescriptionChange(e) }} placeholder="部屋の詳細" />
        <input type="submit" value="作成" />
      </form>
    </>
  );
}

export default Home;
