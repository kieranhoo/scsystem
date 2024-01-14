
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import React, { useState, useEffect } from "react";

import { Profile } from "./profile";
import { RootScreens } from "..";

import { RootStackParamList } from '@/components/navigation';


type ProfileScreenNavigatorProps = NativeStackScreenProps<
    RootStackParamList,
    RootScreens.PROFILE
>;

export const ProfileContainer = () => {
    return <Profile/>;
}

