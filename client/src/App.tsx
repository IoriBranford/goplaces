import React from 'react'
import { useState } from 'react'
import './App.css'

function ExploreView() {
  const [place] = useState("Train_Day");
  const placeImage = '/' + place + '.png';
  return (
    <img
      width={window.innerWidth}
      height={window.innerHeight}
      className="placeImage"
      src={placeImage}
    />
  );
}

export default ExploreView
