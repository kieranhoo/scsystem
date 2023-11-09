import { i18n, LocalizationKey } from "@/Localization";
import React from "react";
import { View, Text, StyleSheet, Image } from "react-native";
import { StatusBar } from "expo-status-bar";
import { HStack, Spinner, Heading } from "native-base";
import { User } from "@/Services";

export interface IHomeProps {
  data: number | undefined;
  isLoading: boolean;
}

export const Onboarding = () => {
  return (
    <View style={styles.container}>
      <Image style={styles.logo} source={{ uri: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTyuA0nvDsTQ1oEs0UynwRL78ZVByg81UShUJJY9gkjZKauGaUV7sdPUqBqwT9QVymQ8Q&usqp=CAU" }} />
      <Image style={styles.illustration} source={require("../../../assets/onboarding/illustration2.png")} />
      <View>
        <Text style={styles.title}>
          Monitor Easily
        </Text>
        <Text>
          With history and dashboard, user can monitor all entry or exit activities of student easier
        </Text>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    alignItems: "center",
    // justifyContent: "center",
  },
  logo: {
    width: 250,
    height: 100,
    resizeMode: 'contain',
  },
  illustration: {
    resizeMode: 'contain',
  },
  title: {
    fontSize: 28,
    textAlign: "center",
    fontFamily: "Poppins_700Bold"
  }
});
