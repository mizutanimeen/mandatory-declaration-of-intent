import React from 'react';
import { Link,useNavigate } from "react-router-dom";
import axios from 'axios';

function Home() {
  const [roomName, setRoomName] = React.useState('');
  const [roomDescription, setRoomDescription] = React.useState('');
  const navigate = useNavigate();

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

    axios.post('http://localhost:8123/mandatory-declaration-of-intent/api/v1/rooms', {
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
      <form onSubmit={(e) => {createRoom(e)}}>
        <input type="text" onChange={(e) => {roomNameChange(e)}} placeholder="名前" />
        <textarea  onChange={(e) => {roomDescriptionChange(e)}} placeholder="説明"  />
        <input type="submit" value="作成" />
      </form>
      <h2>{roomName}</h2>
      <h2>{roomDescription}</h2>
      <br></br>
      <Link to="/about">About</Link>
    </>
  );
}

export default Home;
