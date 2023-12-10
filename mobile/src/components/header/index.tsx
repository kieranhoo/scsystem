import React from "react";
import { View, Text } from "native-base";
import { SafeAreaView, StyleSheet, Dimensions } from "react-native";
import { Colors } from "@/theme/variables";

interface HeaderProps {
    title: string;
}

const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

export const Header: React.FC<HeaderProps> = ({ title }) => {
    return (
        <View style={{ overflow: "hidden", paddingBottom: 1 }}>
            <View style={styles.container}>
                <Text style={styles.title}>{title}</Text>
            </View>
        </View>
    );
};

const styles = StyleSheet.create({
    container: {
        backgroundColor: Colors.WHITE,
        paddingHorizontal: 0.03 * screenWidth,
        paddingBottom: 0.02 * screenHeight,
        shadowColor: '#000',
        shadowOffset: { width: 1, height: 1 },
        shadowOpacity: 0.4,
        shadowRadius: 3,
        elevation: 5,
        justifyContent: 'center'
    },
    title: {
        paddingTop: 0.03 * screenHeight,
        fontFamily: "Poppins_600SemiBold",
        fontSize: 24
    }
});
