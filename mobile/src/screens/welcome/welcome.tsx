import { StatusBar } from "expo-status-bar";
import { Button } from "native-base";
import React from "react";
import { View, Text, StyleSheet } from "react-native";

import { RootScreens } from "..";

import { i18n, LocalizationKey } from "@/localization";

export const Welcome = (props: {
  onNavigate: (string: RootScreens) => void;
}) => {
  return (
    <View style={styles.container}>
      <Text>{i18n.t(LocalizationKey.WELCOME)}</Text>
      <StatusBar style="auto" />
      <Button onPress={() => props.onNavigate(RootScreens.MAIN)}>
        {i18n.t(LocalizationKey.START)}
      </Button>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    alignItems: "center",
    justifyContent: "center",
  },
});
