import React, { useState } from "react";
import { View, Text, StyleSheet, Dimensions, TouchableOpacity } from "react-native";
import { Colors } from "@/theme/variables";
import DropDownPicker from "react-native-dropdown-picker";
import MaterialCommunityIcons from "react-native-vector-icons/MaterialCommunityIcons";

interface HeaderProps {}

const screenWidth = Dimensions.get("screen").width;
const screenHeight = Dimensions.get("screen").height;

export const RoomsHeader: React.FC<HeaderProps> = () => {
  const [open, setOpen] = useState(false);
  const [selectedValue, setSelectedValue] = useState(null);
  const [items, setItems] = useState([
    { label: "Option 1", value: "option1" },
    { label: "Option 2", value: "option2" },
    { label: "Option 3", value: "option3" },
  ]);

  return (
    <View style={styles.headerContainer}>
      <DropDownPicker
        open={open}
        items={items}
        value={selectedValue}
        setOpen={setOpen}
        setValue={setSelectedValue}
        setItems={setItems}
        containerStyle={styles.dropdownContainer}
        //   style={{ backgroundColor: "#fafafa" }}
        //   itemStyle={{
        //     justifyContent: "flex-start",
        //   }}
        //   dropDownStyle={{ backgroundColor: "#fafafa" }}
        //   onChangeItem={(item) => setSelectedValue(item.value)}
      />
      <TouchableOpacity style={styles.iconContainer}>
        <MaterialCommunityIcons
          name="account-box"
          size={50}
          color={"rgba(27, 97, 181, 0.89)"}
        />
      </TouchableOpacity>
    </View>
  );
};

const styles = StyleSheet.create({
  headerContainer: {
    flexDirection: "row",
    backgroundColor: Colors.WHITE,
    paddingBottom: 0.02 * screenHeight,
    shadowColor: "#000",
    shadowOffset: { width: 1, height: 1 },
    shadowOpacity: 0.4,
    shadowRadius: 3,
    elevation: 5,
    zIndex: 10,
  },
  dropdownContainer: {
    flex: 0.75,
    paddingLeft: 15,
    backgroundColor: Colors.WHITE,
    // paddingHorizontal: 0.03 * screenWidth,
    // paddingBottom: 0.02 * screenHeight,
    // shadowColor: "#000",
    // shadowOffset: { width: 1, height: 1 },
    // shadowOpacity: 0.4,
    // shadowRadius: 3,
    // elevation: 5,
    // justifyContent: 'center'
  },
  iconContainer: {
    flex: 0.25,
    alignItems: "center",
    justifyContent: "center",
  },
});
