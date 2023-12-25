
import React, { useState, useEffect } from "react";
import { Profile } from "./profile";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import { RootStackParamList } from '@/components/navigation';
import { RootScreens } from "..";

type ProfileScreenNavigatorProps = NativeStackScreenProps<
    RootStackParamList,
    RootScreens.PROFILE
>;

export const ProfileContainer = () => {
    return <Profile/>;
}

