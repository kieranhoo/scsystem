import React from "react";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { HomeContainer } from "@/screens/home";
import { ScannerContainer } from "@/screens/scan";

const Tab = createBottomTabNavigator();

// @refresh reset
export const MainNavigator = () => {
  return (
    <Tab.Navigator
    // screenOptions={{
    //   tabBarStyle: {
    //     height: 50,
    //   },
    //   tabBarItemStyle: {
    //     margin: 5,
    //   }
    // }}
    >
      <Tab.Screen
        name="Home"
        component={HomeContainer}
        options={{
          tabBarIconStyle: { display: "none" },
          tabBarLabelPosition: "beside-icon",
          headerShown: false,
        }}
      />
      <Tab.Screen
        name="Rooms"
        component={HomeContainer}
        options={{
          tabBarIconStyle: { display: "none" },
          tabBarLabelPosition: "beside-icon",
          headerShown: false,
        }}
      />
      <Tab.Screen
        name="Scan"
        component={ScannerContainer}
        options={{
          tabBarIconStyle: { display: "none" },
          tabBarLabelPosition: "beside-icon",
          headerShown: false,
        }}
      />
      <Tab.Screen
        name="Noti"
        component={HomeContainer}
        options={{
          tabBarIconStyle: { display: "none" },
          tabBarLabelPosition: "beside-icon",
          headerShown: false,
        }}
      />
      <Tab.Screen
        name="Profile"
        component={HomeContainer}
        options={{
          tabBarIconStyle: { display: "none" },
          tabBarLabelPosition: "beside-icon",
          headerShown: false,
        }}
      />
    </Tab.Navigator>
  );
};
