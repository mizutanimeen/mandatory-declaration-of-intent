import React from 'react';
import { Link } from "react-router-dom";
import axios from 'axios';

function About() {
  axios.get('http://localhost:8123/mandatory-declaration-of-intent/api/v1/rooms/a2bec7c4a1cc40da84e8e7a5db464921').then((response: any)=>{
    console.log(response);
  }).catch((error: any)=> {
    console.log(error);
  }).finally(()=> {
  });
  
  return (
    <>
      <Link to="/">Home</Link>
    </>
  );
}

export default About;
