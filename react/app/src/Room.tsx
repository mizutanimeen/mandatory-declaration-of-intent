import React, { useEffect } from "react";
import { Link, useParams } from "react-router-dom";
import axios from 'axios';
import { log } from "console";

export const Room: React.FC = () => {
    const [room, setRoom] = React.useState({} as Room);
    const { id } = useParams();
    const [gestUserName, setGestUserName] = React.useState("");
    const [text, setText] = React.useState("");

    type Room = {
        id: string;
        name: string;
        description: string;
    };

    useEffect(() => {
        const getRoomUrl = (process.env.REACT_APP_GO_URL ?? "") + (process.env.REACT_APP_GO_PATH ?? "") + '/rooms/' + id
        axios.get(getRoomUrl).then((response: any) => {
            setRoom({ id: response.data.roomid, name: response.data.name, description: response.data.description });
        });
    }, []);

    const inputGestUserName = (e: { target: { value: React.SetStateAction<string>; }; }) => {
        setGestUserName(e.target.value);
    }
    const inputText = (e: { target: { value: React.SetStateAction<string>; }; }) => {
        setText(e.target.value);
    }
    const createGestUser = (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (gestUserName == '' || text == '') {
            return;
        }
        const createGestUserUrl = (process.env.REACT_APP_GO_URL ?? "") + (process.env.REACT_APP_GO_PATH ?? "") + '/rooms/members/gest'
        axios.post(createGestUserUrl, {
            Roomid: id,
            Name: gestUserName,
            Description: text
        }, { withCredentials: true })
            .then((response) => {
                console.log(response);
            });
    };

    const getAllGestUser = () => {
        const getAllGestUserUrl = (process.env.REACT_APP_GO_URL ?? "") + (process.env.REACT_APP_GO_PATH ?? "") + '/rooms/' + (id ?? "") + '/members/gest'
        axios.get(getAllGestUserUrl).then((response: any) => {
            console.log(response);
        });
    }

    return (
        <>
            <div>
                <h2>{room.id}</h2>
                <h2>{room.name}</h2>
                <h2>{room.description}</h2>
            </div>

            <h1>意見追加</h1>
            <form onSubmit={(e) => { createGestUser(e) }}>
                <input type="text" onChange={(e) => { inputGestUserName(e) }} placeholder="名前" />
                <textarea onChange={(e) => { inputText(e) }} placeholder="意見" />
                <input type="submit" value="作成" />
            </form>

            <button onClick={getAllGestUser}>みる</button>
            <Link to="/">Home</Link>
        </>
    );
}
