import React, { useState } from "react";
import {
  Dimensions,
  View,
  Text,
  StyleSheet,
  TouchableOpacity,
  SafeAreaView,
} from "react-native";
import { Colors, FontSize } from "../../theme/variables";
import { Header } from "@/components/header";

const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

export const Profile = () => {

  return (
    <SafeAreaView style={styles.container}>
      <View>
        <Header title="Profile" />
      </View>
      <View style={styles.userCard}>
        <View style={styles.avatar}>
          <Text style={styles.avatarName}>NG</Text>
        </View>

        <View style={styles.userInfo}>
          <Text style={styles.userName}>
            Nguyễn Văn A
          </Text>
          <Text>Phone: 0987654321</Text>
          <Text>Organization: xxxxxxxxxxxxxxxxxxx</Text>
          <Text>Office: xxxxxxxxxxxxxxxxxxxx</Text>
        </View>

        <View style={{ flexDirection: "row" }}>
          <TouchableOpacity style={styles.primaryButton}>
            <Text style={styles.primaryButtonText}>
              Edit Profile
            </Text>
          </TouchableOpacity>

          <TouchableOpacity style={styles.secondaryButton}>
            <Text style={styles.secondaryButtonText}>
              Log Out
            </Text>
          </TouchableOpacity>
        </View>
      </View>
    </SafeAreaView>
  );
};

const styles = StyleSheet.create({
  container: {
    height: screenHeight,
    backgroundColor: Colors.WHITE,
  },
  userCard: {
    alignItems: "center",
        justifyContent: "center",
        backgroundColor: Colors.WHITE,
        borderRadius: 10,
        marginHorizontal: 20,
        marginVertical: 30,
        padding: 30,
        shadowOffset: {
          width: 0,
          height: 0,
        },
        shadowOpacity: 0.2,
        shadowRadius: 5,
  },
  avatar: {
    alignItems: "center",
    justifyContent: "center",
    width: 130,
    height: 130,
    borderRadius: 80,
    backgroundColor: Colors.SUCCESS,
  },
  avatarName: {
    color: Colors.WHITE, 
    fontSize: FontSize.LARGE
  },
  userInfo: {
    alignItems: "center",
    justifyContent: "center",
    marginTop: 10,
    marginBottom: 20,
  },
  userName: {
    fontSize: FontSize.REGULAR,
    fontWeight: "bold",
  },
  primaryButton: {
    width: 124,
    height: 36,
    alignItems: "center",
    justifyContent: "center",
    backgroundColor: Colors.WHITE,
    borderColor: Colors.PRIMARY,
    borderWidth: 1,
    borderRadius: 10,
    marginHorizontal: 10,
  },
  primaryButtonText: {
    fontSize: FontSize.SMALL,
    color: Colors.PRIMARY,
  },
  secondaryButton: {
    width: 124,
    height: 36,
    alignItems: "center",
    justifyContent: "center",
    backgroundColor: Colors.PRIMARY,
    borderRadius: 10,
    marginHorizontal: 10,
  },
  secondaryButtonText: {
    fontSize: FontSize.SMALL,
    color: Colors.WHITE,
  },
});
