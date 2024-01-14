
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import React, { useState, useEffect } from "react";

import { Scanner } from "./scanner";
import { RootScreens } from "..";

import { RootStackParamList } from '@/components/navigation';


type ScannerScreenNavigatorProps = NativeStackScreenProps<
    RootStackParamList,
    RootScreens.SCANNER
>;

export const ScannerContainer = () => {
    return <Scanner />;
}

