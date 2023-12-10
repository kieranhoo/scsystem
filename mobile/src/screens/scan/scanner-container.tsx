
import React, { useState, useEffect } from "react";
import { Scanner } from "./scanner";
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import { RootStackParamList } from '@/components/navigation';
import { RootScreens } from "..";

type ScannerScreenNavigatorProps = NativeStackScreenProps<
    RootStackParamList,
    RootScreens.SCANNER
>;

export const ScannerContainer = () => {
    return <Scanner />;
}

