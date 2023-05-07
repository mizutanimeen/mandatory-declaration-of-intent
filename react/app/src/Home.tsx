import React from 'react';
import { useNavigate } from "react-router-dom";
import axios from 'axios';
import { postRoomURL } from "./components/baseURL"
import './static/css/Home.css';
import { TextInput, Textarea, Button, Switch } from '@mantine/core';

export const Home: React.FC = () => {
  const [roomName, setRoomName] = React.useState('');
  const [roomDescription, setRoomDescription] = React.useState('');
  const [onPassword, setOnPassword] = React.useState(false);
  const [password, setPassword] = React.useState('');
  const navigate = useNavigate();

  const createRoom = () => {
    if (roomName == '' || roomDescription == '') {
      return;
    }
    if (roomName.length > 32) {
      alert(`名前は32文字以内で入力してください`)
      return
    }
    if (roomDescription.length > 255) {
      alert(`説明は255文字以内で入力してください`)
      return
    }

    axios.post(postRoomURL(), {
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
      <TextInput
        placeholder="部屋の名前"
        label="部屋の名前"
        withAsterisk
        value={roomName}
        onChange={(e) => setRoomName(e.target.value)}
      />
      <Textarea
        placeholder="部屋の説明"
        label="部屋の説明"
        withAsterisk
        value={roomDescription}
        onChange={(e) => setRoomDescription(e.target.value)}
      />
      <Switch
        label="パスワードを設定する"
        checked={onPassword}
        onChange={(event) => setOnPassword(event.target.checked)}
      />
      <TextInput
        disabled={!onPassword}
        placeholder="パスワード"
        label="パスワード"
        value={password}
        onChange={(e) => setPassword(e.target.value)}
      />
      <Button onClick={createRoom}>部屋を作成</Button >

    </>
  );
}