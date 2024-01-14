
import React, { useEffect, useState } from "react";

import { Home } from "./home";

import { rooms, useLazyGetUserQuery } from "@/services";

export const HomeContainer = () => {
  // const [generalStateOfAllRooms, setGeneralStateOfAllRooms] = useState(null);
  // setGeneralStateOfAllRooms(rooms.getgeneralstateallroom())

    // setTimeout(async () => {
    //   setGeneralStateOfAllRooms(await rooms.getgeneralstateallroom());
    //   console.log("1 check:", generalStateOfAllRooms);
    // }, 5000);

  // return <Home data={generalStateOfAllRooms} />;
  return <Home />;
};
