import React, { useEffect, useRef, useState } from "react";
import { Link, useParams } from "react-router-dom";
import axios from 'axios';
import { GestUser } from "./components/models";
import { DataTable } from 'mantine-datatable';
import { getRoomURL, postGestUserURL, getAllGestUserURL, getRoomPasswordCheckURL } from "./components/baseURL"
import { TextInput, Textarea, Button } from '@mantine/core';

export const Room: React.FC = () => {
    const [start, setStart] = useState(false);
    const [password, setPassword] = React.useState('');
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
        if (!id) { return }
        axios.get(getRoomURL(id)).then((response: any) => {
            setRoom({ id: response.data.roomid, name: response.data.name, description: response.data.description });
            checkPassword();
        });
        getAllGestUser();
    }, []);

    const createGestUser = () => {
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

        axios.post(postGestUserURL(), {
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
        if (!id) { return }

        axios.get(getAllGestUserURL(id), { withCredentials: true }).then((response: any) => {
            if (response.data.length <= 0) {
                return
            }
            const data: GestUser[] = JSON.parse(response.request.response);
            setGestUsers(data);
        });
    }

    const checkPassword = () => {
        if (!id) { return }

        axios.get(getRoomPasswordCheckURL(id), {
            params: {
                Password: password
            }
        }).then((response: any) => {
            if (response.data.password_ok != true) {
                alert(`パスワードが違います`)
                return
            }
            setStart(true);
        });
    }

    if (!start) {
        return (
            <>
                <h2>パスワードを入力してください</h2>
                <TextInput
                    placeholder="パスワード"
                    label="パスワード"
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                />
                <Button onClick={checkPassword}>送信</Button >
            </>
        );
    }

    return (
        <>
            <h3>{room.name}</h3>
            <div>{room.description}</div>

            <br></br>

            <div>意見追加</div>
            <TextInput
                placeholder="ゲストユーザー名"
                label="ゲストユーザー名"
                withAsterisk
                value={gestUserName}
                onChange={(e) => setGestUserName(e.target.value)}
            />
            <Textarea
                placeholder="意見"
                label="意見"
                withAsterisk
                value={gestUserText}
                onChange={(e) => setGestUserText(e.target.value)}
            />
            <Button onClick={createGestUser}>意見を投稿</Button >

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
