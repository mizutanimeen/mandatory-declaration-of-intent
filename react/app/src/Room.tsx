import React, { useEffect, useRef } from "react";
import { Link, useParams } from "react-router-dom";
import axios from 'axios';
import { log } from "console";

export const Room: React.FC = () => {
    const [room, setRoom] = React.useState({} as Room);
    const { id } = useParams();
    const refGestUserName = useRef<HTMLInputElement>(null);
    const refGestUserDescription = useRef<HTMLTextAreaElement>(null);

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

    const createGestUser = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
        e.preventDefault()

        if (refGestUserName.current?.value == '' || refGestUserDescription.current?.value == '') {
            return;
        }
        const createGestUserUrl = (process.env.REACT_APP_GO_URL ?? "") + (process.env.REACT_APP_GO_PATH ?? "") + '/rooms/members/gest'
        axios.post(createGestUserUrl, {
            Roomid: id,
            Name: refGestUserName.current?.value,
            Description: refGestUserDescription.current?.value
        }, { withCredentials: true })
            .then((response) => {
                console.log(response);
            });
    };

    //ココでクッキー渡す必要ありそう
    const getAllGestUser = () => {
        const getAllGestUserUrl = (process.env.REACT_APP_GO_URL ?? "") + (process.env.REACT_APP_GO_PATH ?? "") + '/rooms/' + (id ?? "") + '/members/gest'
        axios.get(getAllGestUserUrl).then((response: any) => {
            console.log(response);
        });
    }

    return (
        <>
            <h3>{room.name}</h3>
            <div>{room.description}</div>

            <br></br>

            <div>意見追加</div>
            <input type="text" ref={refGestUserName} placeholder="名前" />
            <textarea ref={refGestUserDescription} placeholder="意見" />
            <button onClick={(e) => createGestUser(e)}>作成</button>

            <br></br>

            <button onClick={getAllGestUser}>ゲストユーザー一覧</button>

            <br></br>
            <Link to="/">Homeへ戻る</Link>
        </>
    );
}
