import { NavigationContainer } from "@react-navigation/native";
import { createStackNavigator } from '@react-navigation/stack';
import React from "react";
import { StatusBar } from "react-native";

import { MainNavigator } from "./main";

import { RootScreens } from "@/screens";
import { OnboardingContainer } from "@/screens/onboarding";
import { ProfileContainer } from "@/screens/profile";
import { ScannerContainer } from "@/screens/scan/scanner-container";
import { WelcomeContainer } from "@/screens/welcome";

export type RootStackParamList = {
  [RootScreens.MAIN]: undefined;
  [RootScreens.WELCOME]: undefined;
  [RootScreens.ONBOARDING]: undefined;
  [RootScreens.SCANNER]: undefined;
  [RootScreens.ROOMS]: undefined;
  [RootScreens.NOTI]: undefined;
  [RootScreens.PROFILE]: undefined;
};

const RootStack = createStackNavigator<RootStackParamList>();

const ApplicationNavigator = () => {
  return (
    <NavigationContainer>
      <StatusBar />
      <RootStack.Navigator
        detachInactiveScreens
        screenOptions={{
          headerShown: false,
        }}
        initialRouteName={RootScreens.ONBOARDING}
      >
        <RootStack.Screen
          name={RootScreens.ONBOARDING}
          component={OnboardingContainer}
        />
        <RootStack.Screen
          name={RootScreens.WELCOME}
          component={WelcomeContainer}
        />
        <RootStack.Screen
          name={RootScreens.MAIN}
          component={MainNavigator}
          options={{ gestureEnabled: false }}

        />
        <RootStack.Screen
          name={RootScreens.SCANNER}
          component={ScannerContainer}
        />
        <RootStack.Screen
          name={RootScreens.PROFILE}
          component={ProfileContainer}
        />
      </RootStack.Navigator>
    </NavigationContainer>
  );
};

export { ApplicationNavigator };
