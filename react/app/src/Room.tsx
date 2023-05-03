import React, { useEffect, useRef, useState } from "react";
import { Link, useParams } from "react-router-dom";
import axios from 'axios';

export const Room: React.FC = () => {
    const [room, setRoom] = React.useState({} as Room);
    const { id } = useParams();
    const [gestUserName, setGestUserName] = useState("");
    const [gestUserText, setGestUserText] = useState("");

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

    const changeUserName = (e: { preventDefault: () => void; target: { value: React.SetStateAction<string>; }; }) => {
        e.preventDefault();
        setGestUserName(e.target.value);
    }
    const changeUserText = (e: { preventDefault: () => void; target: { value: React.SetStateAction<string>; }; }) => {
        e.preventDefault();
        setGestUserText(e.target.value);
    }

    const createGestUser = (e: { preventDefault: () => void; }) => {
        e.preventDefault()

        if (gestUserName == '' || gestUserText == '') {
            return;
        }
        const createGestUserUrl = (process.env.REACT_APP_GO_URL ?? "") + (process.env.REACT_APP_GO_PATH ?? "") + '/rooms/members/gest'
        axios.post(createGestUserUrl, {
            Roomid: id,
            Name: gestUserName,
            Text: gestUserText
        }, { withCredentials: true })
            .then((response) => {
                console.log(response);
                setGestUserName("");
                setGestUserText("");
            });
    };

    //ココでクッキー渡す必要ありそう
    const getAllGestUser = () => {
        const getAllGestUserUrl = (process.env.REACT_APP_GO_URL ?? "") + (process.env.REACT_APP_GO_PATH ?? "") + '/rooms/' + (id ?? "") + '/members/gest'
        axios.get(getAllGestUserUrl, { withCredentials: true }).then((response: any) => {
            console.log(response);
        });
    }

    return (
        <>
            <h3>{room.name}</h3>
            <div>{room.description}</div>

            <br></br>

            <div>意見追加</div>
            <input type="text" value={gestUserName} onChange={changeUserName} placeholder="名前" />
            <textarea value={gestUserText} onChange={changeUserText} placeholder="意見" />
            <button onClick={createGestUser}>作成</button>

            <br></br>

            <button onClick={getAllGestUser}>ゲストユーザー一覧</button>

            <br></br>
            <Link to="/">Homeへ戻る</Link>
        </>
    );
}
