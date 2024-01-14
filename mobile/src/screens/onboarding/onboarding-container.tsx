
import { NativeStackScreenProps } from "@react-navigation/native-stack";
import React from "react";

import { Onboarding } from "./onboarding";
import { RootScreens } from "..";

import { RootStackParamList } from '@/components/navigation';


type OnboardingScreenNavigatorProps = NativeStackScreenProps<
  RootStackParamList,
  RootScreens.ONBOARDING
>;

export const OnboardingContainer = ({
  navigation,
}: OnboardingScreenNavigatorProps) => {
  const onNavigate = (screen: RootScreens) => {
    navigation.navigate(screen);
  };

  return <Onboarding onNavigate={onNavigate} />;
}

