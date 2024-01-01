import React, { useState, useEffect } from "react";
import { createBottomTabNavigator } from "@react-navigation/bottom-tabs";
import { HomeContainer } from "@/screens/home";
import { ScannerContainer } from "@/screens/scan";
import Ionicons from "react-native-vector-icons/Ionicons";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";

import { RoomsContainer } from "@/screens/rooms";
import { ProfileContainer } from "@/screens/profile";
import { Colors } from "react-native/Libraries/NewAppScreen";

const Tab = createBottomTabNavigator();

interface CustomTabBarIconProps {
  focused: boolean;
  routeName: string;
  size: number;
}

const CustomTabBarIcon: React.FC<CustomTabBarIconProps> = ({
  focused,
  routeName,
  size,
}) => {
  const [iconName, setIconName] = useState("");

  useEffect(() => {
    if (routeName === "Home") {
      setIconName(focused ? "home" : "home-outline");
    } else if (routeName === "Scan") {
      setIconName(focused ? "scan-circle" : "scan-circle-outline");
    } else if (routeName === "Rooms") {
      setIconName(focused ? "home-city" : "home-city-outline");
    }
  }, [focused, routeName]);

  const getIconComponent = () => {
    if (routeName === "Home") {
      return (
        <Ionicons
          name={iconName}
          color={
            focused ? "rgba(27, 97, 181, 0.89)" : "rgba(108, 108, 108, 0.89)"
          }
          size={24}
        />
      );
    }
    if (routeName === "Scan") {
      return (
        <Ionicons
          name={iconName}
          color={
            focused ? "rgba(27, 97, 181, 0.89)" : "rgba(108, 108, 108, 0.89)"
          }
          size={24}
        />
      );
    }
    if (routeName === "Rooms") {
      return (
        <MaterialCommunityIcons
          name={iconName}
          color={
            focused ? "rgba(27, 97, 181, 0.89)" : "rgba(108, 108, 108, 0.89)"
          }
          size={24}
        />
      );
    }
    return null;
  };

  return getIconComponent();
};

// @refresh reset
export const MainNavigator = () => {
  return (
    <Tab.Navigator
      screenOptions={{
        tabBarLabelStyle: {
          fontSize: 13,
          fontFamily: "Poppins_400Regular",
        },
      }}
    >
      <Tab.Screen
        name="Home"
        component={HomeContainer}
        options={{
          tabBarIcon: ({ focused, size }) => (
            <CustomTabBarIcon focused={focused} routeName="Home" size={size} />
          ),
          // tabBarIconStyle: { display: "none" },
          // tabBarLabelPosition: "beside-icon",
          headerShown: false,
        }}
      />
      <Tab.Screen
        name="Scan"
        component={ScannerContainer}
        options={{
          tabBarIcon: ({ focused, size }) => (
            <CustomTabBarIcon focused={focused} routeName="Scan" size={size} />
          ),
          // tabBarIconStyle: { display: "none" },
          // tabBarLabelPosition: "beside-icon",
          headerShown: false,
        }}
      />
      <Tab.Screen
        name="Rooms"
        component={RoomsContainer}
        options={{
          tabBarIcon: ({ focused, size }) => (
            <CustomTabBarIcon focused={focused} routeName="Rooms" size={size} />
          ),
          // tabBarIconStyle: { display: "none" },
          // tabBarLabelPosition: "beside-icon",
          headerShown: false,
        }}
      />
      {/* <Tab.Screen
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
        component={ProfileContainer}
        options={{
          tabBarIconStyle: { display: "none" },
          tabBarLabelPosition: "beside-icon",
          headerShown: false,
        }}
      /> */}
    </Tab.Navigator>
  );
};
