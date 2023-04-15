import React from 'react';
import { Link ,useParams } from "react-router-dom";
import axios from 'axios';

function Room() {

    const [room, setRoom] = React.useState({} as Room);
    const {id} = useParams();

    type Room = {
        id: string;
        name: string;
        description: string;
    };

    axios.get('http://localhost:8123/mandatory-declaration-of-intent/api/v1/rooms/'+id).then((response: any)=>{
        setRoom({id: response.data.roomid, name: response.data.name, description: response.data.description});
    });

    return (
    <>
        <div>
            <h2>{room.id}</h2>
            <h2>{room.name}</h2>
            <h2>{room.description}</h2>
        </div>
        <Link to="/">Home</Link>
    </>
    );
}

export default Room;
