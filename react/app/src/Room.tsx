import React, { useEffect } from "react";
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

    const getUrl = (process.env.REACT_APP_GO_URL ?? "") + (process.env.REACT_APP_GO_PATH ?? "") +'/rooms/'+id

    console.log(getUrl);

    axios.get(getUrl).then((response: any)=>{
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
