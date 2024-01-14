import {
  Poppins_100Thin,
  Poppins_100Thin_Italic,
  Poppins_200ExtraLight,
  Poppins_200ExtraLight_Italic,
  Poppins_300Light,
  Poppins_300Light_Italic,
  Poppins_400Regular,
  Poppins_400Regular_Italic,
  Poppins_500Medium,
  Poppins_500Medium_Italic,
  Poppins_600SemiBold,
  Poppins_600SemiBold_Italic,
  Poppins_700Bold,
  Poppins_700Bold_Italic,
  Poppins_800ExtraBold,
  Poppins_800ExtraBold_Italic,
  Poppins_900Black,
  Poppins_900Black_Italic,
} from "@expo-google-fonts/poppins";
import { useFonts } from "expo-font";
import * as Localization from "expo-localization";
import * as SplashScreen from 'expo-splash-screen';
import { StatusBar } from "expo-status-bar";
import { NativeBaseProvider } from "native-base";
import React, { useEffect, useCallback, useState } from "react";
import { View, Text } from "react-native";
import { SafeAreaView } from 'react-native-safe-area-context';
import { Provider } from "react-redux";

import { PersistGate } from "redux-persist/integration/react";
import { i18n, Language } from "@/localization";


import { store, persistor } from "@/store";


import { ApplicationNavigator } from "@/components/navigation";


i18n.locale = Localization.locale;
i18n.enableFallback = true;
i18n.defaultLocale = Language.ENGLISH;

// SplashScreen.preventAutoHideAsync();

export default function App() {
  const [appIsReady, setAppIsReady] = useState(false);

  const [fontsLoaded] = useFonts({
    Poppins_100Thin,
    Poppins_100Thin_Italic,
    Poppins_200ExtraLight,
    Poppins_200ExtraLight_Italic,
    Poppins_300Light,
    Poppins_300Light_Italic,
    Poppins_400Regular,
    Poppins_400Regular_Italic,
    Poppins_500Medium,
    Poppins_500Medium_Italic,
    Poppins_600SemiBold,
    Poppins_600SemiBold_Italic,
    Poppins_700Bold,
    Poppins_700Bold_Italic,
    Poppins_800ExtraBold,
    Poppins_800ExtraBold_Italic,
    Poppins_900Black,
    Poppins_900Black_Italic,
  });
  useEffect(() => {
    async function prepare() {
      try {
        await new Promise(resolve => setTimeout(resolve, 1000));
      } catch (e) {
        console.warn(e);
      } finally {
        setAppIsReady(true);
      }
    }
    prepare();
  }, []);

  if (!appIsReady || !fontsLoaded) {
    return null;
  } else {
    return (
      <NativeBaseProvider>
        <Provider store={store}>
          <PersistGate loading={null} persistor={persistor}>
            <SafeAreaView
              style={{ flex: 1 }}
              edges={["right", "top", "left"]}
            >
              <StatusBar style="dark" />
              < ApplicationNavigator />
            </SafeAreaView>
          </PersistGate>
        </Provider>
      </NativeBaseProvider>
    );
  }
}
