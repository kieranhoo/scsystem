import React from "react";
import { View, Text, StyleSheet, Dimensions } from "react-native";
import { Header } from "@/components/header";

const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

export const Home = () => {
  return (
    <View style={styles.container}>
      <View style={styles.header_container}>
        <Header title="Home" />
      </View>
      <View style={styles.content_container}>

      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#FFFFFF",
  },
  header_container: {

  },
  content_container: {

    paddingHorizontal: 0.03 * screenWidth,
  }
});
