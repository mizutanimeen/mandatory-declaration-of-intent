import React, { useEffect, useRef, useState } from "react";
import { Link, useParams } from "react-router-dom";
import axios from 'axios';
import { GestUser } from "./components/models";
import { Text } from '@mantine/core';
import { DataTable } from 'mantine-datatable';

export const Room: React.FC = () => {
    const [room, setRoom] = React.useState({} as Room);
    const { id } = useParams();
    const [gestUserName, setGestUserName] = useState("");
    const [gestUserText, setGestUserText] = useState("");
    const [gestUsers, setGestUsers] = useState([] as GestUser[]);

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
        getAllGestUser();
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
        if (gestUserName.length > 32) {
            alert(`名前は32文字以内で入力してください`)
            return;
        }
        if (gestUserText.length > 2048) {
            alert(`意見は2048文字以内で入力してください`)
            return;
        }

        const createGestUserUrl = (process.env.REACT_APP_GO_URL ?? "") + (process.env.REACT_APP_GO_PATH ?? "") + '/rooms/members/gest'
        axios.post(createGestUserUrl, {
            Roomid: id,
            Name: gestUserName,
            Text: gestUserText
        }, { withCredentials: true })
            .then((response) => {
                setGestUserName("");
                setGestUserText("");
                getAllGestUser();
            });
    };

    const getAllGestUser = () => {
        const getAllGestUserUrl = (process.env.REACT_APP_GO_URL ?? "") + (process.env.REACT_APP_GO_PATH ?? "") + '/rooms/' + (id ?? "") + '/members/gest'
        axios.get(getAllGestUserUrl, { withCredentials: true }).then((response: any) => {
            if (response.data.length <= 0) {
                return
            }
            const data: GestUser[] = JSON.parse(response.request.response);
            setGestUsers(data);
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

            <GestUsersView gestUsers={gestUsers} />

            <Link to="/">Homeへ戻る</Link>
        </>
    );
}

const GestUsersView: React.FC<{ gestUsers: GestUser[] | undefined }> = ({ gestUsers }) => {

    const records = gestUsers?.map((gestUser) => {
        return {
            name: gestUser.name,
            text: gestUser.text,
        }
    });

    if ((gestUsers?.length ?? 0) <= 0) {
        return (
            <></>
        )
    }

    return (
        <>
            <DataTable
                withBorder
                borderRadius="sm"
                withColumnBorders
                striped
                highlightOnHover
                columns={[
                    {
                        accessor: 'name',
                        title: '名前',
                        width: 100,
                        ellipsis: false,
                    },
                    {
                        accessor: 'text',
                        title: '説明',
                        width: 1000,
                        ellipsis: false,
                    },
                ]}
                records={records}
                onRowClick={({ name, text }) =>
                    alert(`${name} \n ${text}`)
                }
            />
            <br></br>
        </>
    );
};
