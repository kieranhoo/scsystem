import { NativeStackScreenProps } from "@react-navigation/native-stack";
import React from "react";

import { Welcome } from "./welcome";
import { RootScreens } from "..";

import { RootStackParamList } from "@/components/navigation";


type WelcomeScreenNavigatorProps = NativeStackScreenProps<
  RootStackParamList,
  RootScreens.WELCOME
>;

export const WelcomeContainer = ({
  navigation,
}: WelcomeScreenNavigatorProps) => {
  const onNavigate = (screen: RootScreens) => {
    navigation.navigate(screen);
  };

  return <Welcome onNavigate={onNavigate} />;
};
